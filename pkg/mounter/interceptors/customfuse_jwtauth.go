package interceptors

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"k8s.io/klog/v2"
)

var _ mounter.MountInterceptor = CustomFuseJWTAuthInterceptor

var (
	configMapRefreshHook = "/etc/fuse-config/refresh-hook.sh"
	imageRefreshHook     = "/refresh-hook.sh"
)

// CustomFuseJWTAuthInterceptor provisions scoped STS credentials for customfuse
// mounts whose authType selects the agent-identity flow.
//
// The initial credential is passed to the entrypoint via stdin as four lines
// (AccessKeyId, AccessKeySecret, SecurityToken, Expiration). Subsequent
// rotations run the refresh hook with the new credential on its stdin in the
// same format.
//
// For any other authType, including the empty default that passes Secret
// entries through as environment variables, this is a no-op.
func CustomFuseJWTAuthInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	if op == nil {
		return handler(ctx, op)
	}
	opts, _, active, err := resolveAgentIdentityOpts(op.Options)
	if err != nil {
		return err
	}
	if !active {
		return handler(ctx, op)
	}

	// Strip infrastructure options before the entrypoint sees them.
	op.Options = stripCustomFuseInfraOptions(op.Options)

	hook := resolveRefreshHook(configMapRefreshHook, imageRefreshHook)
	if hook == "" {
		return fmt.Errorf("agent-identity requires a credential refresh hook; provide one via credentialRefreshHookKey or ship /refresh-hook.sh in the image")
	}
	sink := newStdinCredentialSink(hook, customFuseHookEnv(op))

	refresher, err := jwtauth.NewRefresher(opts, sink)
	if err != nil {
		return fmt.Errorf("jwtauth create refresher: %w", err)
	}

	cred, err := refresher.FetchCredentials(ctx)
	if err != nil {
		return fmt.Errorf("jwtauth fetch initial credential: %w", err)
	}

	// Pass the initial credential to the entrypoint via stdin. The driver
	// itself does not put a credential value in the filesystem, argv or
	// environment; what the entrypoint does after reading stdin is the
	// entrypoint's responsibility (see the example scripts for the suggested
	// patterns). The entrypoint must read all four lines before handing off,
	// otherwise the unread lines stay in the pipe and are inherited by the
	// process it exec's into.
	op.Stdin = bytes.NewReader(formatCredential(cred))
	klog.V(4).InfoS("provisioned customfuse agent-identity credential via stdin", "target", op.Target)

	stopRefresher := jwtauth.DefaultManager.Start(op.Target, refresher, cred)

	if err := handler(ctx, op); err != nil {
		stopRefresher()
		return err
	}

	result, ok := op.MountResult.(server.FuseMountResult)
	if !ok {
		klog.ErrorS(errors.New("mount result is not a FuseMountResult"),
			"credential refresher will only stop at driver terminate", "target", op.Target)
		return nil
	}
	go func() {
		<-result.ExitChan
		stopRefresher()
		klog.V(4).InfoS("stopped customfuse credential refresher", "target", op.Target)
	}()
	return nil
}

func resolveRefreshHook(candidates ...string) string {
	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

// stripCustomFuseInfraOptions drops the options that configure the credential
// exchange. authType is deliberately kept.
func stripCustomFuseInfraOptions(options []string) []string {
	kept := make([]string, 0, len(options))
	for _, opt := range options {
		key, _, _ := strings.Cut(opt, "=")
		if _, infra := jwtauth.InfraOptionKeys[key]; infra {
			continue
		}
		kept = append(kept, opt)
	}
	return kept
}

// customFuseHookEnv builds the environment for the refresh hook.
func customFuseHookEnv(op *mounter.MountOperation) []string {
	inherited := os.Environ()
	taken := make(map[string]struct{}, len(inherited))
	for _, kv := range inherited {
		if name, _, found := strings.Cut(kv, "="); found {
			taken[name] = struct{}{}
		}
	}

	env := inherited
	env = append(env, "mountpoint="+op.Target)
	taken["mountpoint"] = struct{}{}
	if op.Source != "" {
		env = append(env, "source="+op.Source)
		taken["source"] = struct{}{}
	}
	for _, opt := range op.Options {
		key, _, _ := strings.Cut(opt, "=")
		if _, ok := taken[key]; ok {
			continue
		}
		taken[key] = struct{}{}
		env = append(env, opt)
	}
	return env
}
