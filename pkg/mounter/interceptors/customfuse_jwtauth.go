package interceptors

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

var _ mounter.MountInterceptor = CustomFuseJWTAuthInterceptor

const (
	// OptCredentialDir lets a volume pin the directory the credential files are
	// written to, for a client whose credential path is fixed in its own
	// configuration and therefore cannot be told where to look at mount time.
	// Otherwise each mount gets its own directory and the entrypoint reads the
	// path from the environment.
	OptCredentialDir = "credentialDir"

	// The refresh hook is found the way the entrypoint is: the image may ship
	// one, and a ConfigMap entry takes precedence over it. What puts a file at
	// /etc/fuse-config depends on where this runs — a fuse pod gets it from the
	// volume's own projected ConfigMap, an injected sandbox sidecar from
	// whatever the platform mounts there, which is the same for every volume.
	// Either way both paths are fixed, so a volume decides whether a hook runs
	// and what it contains, never which executable this runs.
	configMapRefreshHook = "/etc/fuse-config/refresh-hook.sh"
	imageRefreshHook     = "/refresh-hook.sh"

	// customFuseCredentialBaseDir is the parent of the default, per-mount
	// credential directory. It lives under /run so the credential never
	// survives a reboot of the node.
	customFuseCredentialBaseDir = "/run/customfuse-credentials"
)

// CustomFuseJWTAuthInterceptor provisions scoped STS credentials for customfuse
// mounts whose authType selects the agent-identity flow.
//
// The other drivers can each assume something about their client: that it
// accepts the credential as mount options, or that it performs the exchange
// itself. customfuse runs a client chosen by the volume, so it can assume
// neither. The exchange therefore happens here and the result is written to
// disk, leaving the entrypoint as the adapter between a fixed delivery and an
// arbitrary client (see fileCredentialSink for why that boundary sits there).
//
// The credential is on disk before the entrypoint starts, so the very first
// mount has it, and a jwtauth.Refresher keeps writing to the same directory for
// the lifetime of the mount. The refresher is stopped on unmount
// (Manager.StopByTarget) or driver Terminate (jwtauth.StopAll).
//
// For any other authType, including the empty default that passes Secret
// entries through as environment variables, this is a no-op.
func CustomFuseJWTAuthInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	if op == nil {
		return handler(ctx, op)
	}
	idx := mounterutils.IndexMountOptions(op.Options)
	if !jwtauth.IsAgentIdentity(idx[jwtauth.OptAuthType]) {
		return handler(ctx, op)
	}

	opts := jwtauth.ResolveOpts(idx)
	if err := opts.Validate(); err != nil {
		return fmt.Errorf("jwtauth config error: %w", err)
	}

	credDir := resolveCredentialDir(idx[OptCredentialDir], op.Target)

	// Strip the infrastructure options before the entrypoint sees them: they
	// configure this exchange and would otherwise reach the client as
	// environment variables. credentialDir replaces them, pointing the
	// entrypoint at the result instead of at the machinery.
	op.Options = append(stripCustomFuseInfraOptions(op.Options), OptCredentialDir+"="+credDir)

	sink := newFileCredentialSink(credDir, resolveRefreshHook(), customFuseHookEnv(op))

	// Fetch and write the credential before mounting: the entrypoint reads
	// these files as it starts, so a mount must not begin without them. Failing
	// the mount is the only correct outcome — there is no static credential to
	// fall back to, and a client left to start without one would fail later and
	// less clearly.
	refresher := jwtauth.NewRefresher(opts, sink)
	if err := refresher.Start(ctx); err != nil {
		sink.Cleanup()
		return fmt.Errorf("jwtauth provision credential: %w", err)
	}
	klog.V(4).InfoS("provisioned customfuse agent-identity credential", "target", op.Target, "dir", credDir)

	// The refresher is already running, so it has to be tracked before the
	// mount: a mount that fails or is skipped must not leave a refresh loop
	// writing to a directory nobody reads.
	jwtauth.DefaultManager.Add(op.Target, refresher)

	if err := handler(ctx, op); err != nil {
		jwtauth.DefaultManager.StopRefresher(refresher)
		return err
	}

	// A customfuse mount lives exactly as long as the entrypoint process, so
	// the process exiting is what ends the credential's usefulness — whether it
	// exited because of an unmount or because it crashed. Waiting on it covers
	// both, so this driver needs no unmount hook to stop the refresher.
	// StopRefresher also runs the sink's Cleanup, removing the files.
	result, ok := op.MountResult.(server.FuseMountResult)
	if !ok {
		klog.ErrorS(errors.New("mount result is not a FuseMountResult"),
			"credential refresher will only stop at driver terminate", "target", op.Target)
		return nil
	}
	go func() {
		<-result.ExitChan
		jwtauth.DefaultManager.StopRefresher(refresher)
		klog.V(4).InfoS("stopped customfuse credential refresher", "target", op.Target, "dir", credDir)
	}()
	return nil
}

// resolveCredentialDir decides where the credential files go. A volume-pinned
// directory is used as given, for clients configured with an absolute path;
// otherwise each mount gets its own directory keyed by the mount path, so two
// volumes on one node cannot read each other's credential.
func resolveCredentialDir(pinned, target string) string {
	if pinned != "" {
		return pinned
	}
	return filepath.Join(customFuseCredentialBaseDir, mounterutils.ComputeMountPathHash(target), "sts")
}

func resolveRefreshHook() string {
	return firstExecutableHook(configMapRefreshHook, imageRefreshHook)
}

// firstExecutableHook returns the first of paths that exists, or empty to run no
// hook. Whichever exists first decides the outcome even if it is not executable:
// a volume that projected a hook meant that one, so a wrong file mode must not
// silently hand execution to whatever the image happens to ship. A hook that
// cannot run is not fatal — the mount still works for a client that re-reads the
// credential files on its own.
func firstExecutableHook(paths ...string) string {
	for _, path := range paths {
		fi, err := os.Stat(path)
		if err != nil {
			continue
		}
		if fi.Mode()&0111 == 0 {
			klog.Warningf("credential refresh hook %s is not executable (mode %s), ignoring it", path, fi.Mode())
			return ""
		}
		return path
	}
	return ""
}

// stripCustomFuseInfraOptions drops the options that configure the credential
// exchange, plus any caller-supplied credential directory, which is re-added
// resolved. authType is deliberately kept: the entrypoint is a script, so an
// unknown key is harmless there, and knowing which auth flow ran lets one
// entrypoint serve both the agent-identity and the Secret-passthrough case.
func stripCustomFuseInfraOptions(options []string) []string {
	kept := make([]string, 0, len(options))
	for _, opt := range options {
		key, _, _ := strings.Cut(opt, "=")
		if _, infra := jwtauth.InfraOptionKeys[key]; infra {
			continue
		}
		if key == OptCredentialDir {
			continue
		}
		kept = append(kept, opt)
	}
	return kept
}

// customFuseHookEnv gives the refresh hook the volume's view of the mount: the
// mount point, the source, and the volume attributes, one key=value environment
// variable each — the shape the driver hands the entrypoint. A hook runs after
// the mount call has returned and has no way to ask, and a client that must be
// told about a rotation must be told which volume's.
//
// op.Secrets is left out. The hook is told where the credential is, never what it
// is, and a volume carrying both a Secret and agent-identity would otherwise hand
// a long-lived key to a process that lives as long as the mount. credentialDir is
// left to the sink, which owns the resolved value.
func customFuseHookEnv(op *mounter.MountOperation) []string {
	env := append(os.Environ(), "mountpoint="+op.Target)
	if op.Source != "" {
		env = append(env, "source="+op.Source)
	}
	for _, opt := range op.Options {
		if key, _, _ := strings.Cut(opt, "="); key != OptCredentialDir {
			env = append(env, opt)
		}
	}
	return env
}
