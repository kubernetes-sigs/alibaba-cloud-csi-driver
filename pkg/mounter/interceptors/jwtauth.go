package interceptors

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"k8s.io/klog/v2"
)

var _ mounter.MountInterceptor = JWTAuthInterceptor

const (
	AuthTypeAgentIdentity = "agent-identity"

	AuthTypeJWTAuth = "jwtauth"

	// credentialBaseDir is the parent directory under which per-mount STS
	// credential directories are created.
	credentialBaseDir = "/var/run/nas/credentials"

	// Mount options consumed by the jwtauth interceptors. These are
	// infrastructure-only: they are removed from op.Options before the mount
	// handler runs (see jwtAuthInfraOptionKeys).
	optAuthType                = "authType"
	optSandboxId               = "sandboxId"
	optSandboxCredProviderName = "sandboxCredProviderName"
	optJWTAuthEndpoint         = "jwtauth_endpoint"
	optJWTAuthTokenFile        = "jwtauth_token_file"
	optJWTAuthCredProvider     = "jwtauth_cred_provider"
	optJWTAuthCAFile           = "jwtauth_ca_file"

	// optCredentialDir is the single option passed through to the FUSE
	// entrypoint. It points at the directory holding the rotated STS files.
	optCredentialDir = "credentialDir"
)

func isJWTAuth(authType string) bool {
	return authType == AuthTypeAgentIdentity || authType == AuthTypeJWTAuth
}

// jwtAuthInfraOptionKeys is the shared set of infrastructure-only jwtauth
// mount options stripped by all jwtauth interceptors before the mount handler
// runs.
var jwtAuthInfraOptionKeys = map[string]struct{}{
	optSandboxId:               {},
	optSandboxCredProviderName: {},
	optJWTAuthEndpoint:         {},
	optJWTAuthTokenFile:        {},
	optJWTAuthCredProvider:     {},
	optJWTAuthCAFile:           {},
}

// JWTAuthInterceptor provisions scoped STS credentials for FUSE mounts.
//
// It mirrors OssfsSecretInterceptor: before the mount it starts a background
// credential refresher (jwtauth.Refresher with a file-based sink) that
// exchanges the sandbox jwtauth token for an STS token and writes the
// credential files atomically to a per-mount directory. It rewrites
// op.Options so the entrypoint receives only credentialDir (plus authType),
// and binds the refresher lifetime to the mount process via
// OssfsMountResult.ExitChan.
//
// For any other authType (including the empty default) it is a no-op.
func JWTAuthInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	if op == nil {
		return handler(ctx, op)
	}
	idx := mounterutils.IndexMountOptions(op.Options)
	if !isJWTAuth(idx[optAuthType]) {
		return handler(ctx, op)
	}

	opts := resolveJWTAuthOpts(idx)
	if err := opts.Validate(); err != nil {
		return fmt.Errorf("jwtauth config error: %w", err)
	}

	volumeID := op.VolumeID
	if volumeID == "" {
		volumeID = idx[optSandboxId]
	}
	if volumeID == "" {
		// Never fall back to a shared directory: distinct mounts would collide
		// on the same credential directory and their refreshers would overwrite
		// each other's STS files. Fail rather than use a shared "default"
		// directory.
		return fmt.Errorf("jwtauth config error: neither volumeID nor sandboxId is set, cannot derive a unique credential directory")
	}
	sink := newJWTAuthFileSink(filepath.Join(credentialBaseDir, volumeID))

	refresher := jwtauth.NewRefresher(opts, sink)
	if err := refresher.Start(ctx); err != nil {
		return fmt.Errorf("start jwtauth credential refresher: %w", err)
	}
	jwtauth.DefaultManager.Add(op.Target, refresher)

	// Replace infrastructure-only options with the resolved credential dir so
	// the FUSE entrypoint only sees credentialDir + authType.
	op.Options = rewriteJWTAuthOptions(op.Options, sink.Dir())

	err := handler(ctx, op)
	if err != nil {
		// Mount failed: stop the refresher and remove the credential files
		// written by Start, so failed mounts do not leave STS files on disk.
		jwtauth.DefaultManager.StopRefresher(refresher)
		return err
	}

	if op.MountResult == nil {
		// Mount reported success but produced no result to hang cleanup on.
		// Stop the refresher and remove its credential directory to avoid
		// leaking the goroutine and the on-disk STS files.
		jwtauth.DefaultManager.StopRefresher(refresher)
		return nil
	}
	result, ok := op.MountResult.(server.OssfsMountResult)
	if !ok {
		klog.ErrorS(errors.New("failed to assert fuse mount result"),
			"stopping jwtauth refresher", "mountpoint", op.Target)
		jwtauth.DefaultManager.StopRefresher(refresher)
		return nil
	}

	// Bind the refresher lifetime to the mount process.
	go func() {
		<-result.ExitChan
		jwtauth.DefaultManager.StopRefresher(refresher)
	}()
	return nil
}

// resolveJWTAuthOpts extracts and defaults the jwtauth settings
// from the indexed mount options. Defaults are resolved here (formerly in the
// driver's ApplyOptionDefaults) so the interceptor is the single owner of
// jwtauth configuration.
func resolveJWTAuthOpts(idx map[string]string) jwtauth.Opts {
	opts := jwtauth.Opts{
		TokenFile:    idx[optJWTAuthTokenFile],
		Endpoint:     idx[optJWTAuthEndpoint],
		CredProvider: idx[optJWTAuthCredProvider],
		CAFile:       idx[optJWTAuthCAFile],
		SandboxId:    idx[optSandboxId],
	}
	if opts.Endpoint == "" {
		opts.Endpoint = agentidentity.GetEndpoint()
	}
	if opts.TokenFile == "" && opts.SandboxId != "" {
		opts.TokenFile = agentidentity.GetTokenFilePath(opts.SandboxId)
	}
	if opts.CredProvider == "" {
		opts.CredProvider = idx[optSandboxCredProviderName]
	}
	if opts.CAFile == "" {
		if caPath := agentidentity.GetCAFilePath(); caPath != "" && unix.Access(caPath, unix.R_OK) == nil {
			opts.CAFile = caPath
		}
	}
	return opts
}

// rewriteJWTAuthOptions strips infrastructure-only jwtauth options
// and appends the resolved credentialDir. authType is preserved so the
// entrypoint can branch on it.
func rewriteJWTAuthOptions(options []string, credentialDir string) []string {
	result := make([]string, 0, len(options)+1)
	for _, opt := range options {
		key, _, _ := strings.Cut(opt, "=")
		if _, infra := jwtAuthInfraOptionKeys[key]; infra {
			continue
		}
		result = append(result, opt)
	}
	result = append(result, optCredentialDir+"="+credentialDir)
	return result
}
