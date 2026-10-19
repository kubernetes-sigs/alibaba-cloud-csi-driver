package interceptors

import (
	"context"
	"fmt"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	"k8s.io/klog/v2"
)

var _ mounter.MountInterceptor = AlinasJWTAuthInterceptor

const (
	// alinas mount options that carry the STS credential. The alinas client
	// accepts the STS triple directly as mount options:
	//   mount -t alinas -o tls,vers=3,ram,access_key_id=<AK>,access_key_secret=<SK>,security_token=<TOKEN> ...
	optAlinasAccessKeyID     = "access_key_id"
	optAlinasAccessKeySecret = "access_key_secret"
	optAlinasSecurityToken   = "security_token"
	optAlinasTLS             = "tls"
	optAlinasRAM             = "ram"
)

// AlinasJWTAuthInterceptor provisions scoped STS credentials for alinas mounts.
//
// The alinas client consumes the STS credential directly as mount options, so
// this interceptor exchanges the sandbox jwtauth token for an STS token in
// memory and injects the resolved credential into op.SensitiveOptions (never
// op.Options, so the secret is masked in mount logs and error messages). After
// a successful mount it starts a jwtauth.Refresher with the alinas cert-refresh
// sink that pushes each rotated credential to the live mount via
// alinas-tls-cert-refresh; the refresher is stopped on unmount
// (Manager.StopByTarget) or driver Terminate (jwtauth.StopAll). Nothing is
// written to disk.
//
// For any other authType (including the empty default) it is a no-op.
func AlinasJWTAuthInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
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
	refresher, err := jwtauth.NewRefresher(opts, newAlinasCertRefreshSink(op.Target))
	if err != nil {
		return err
	}

	cred, err := refresher.FetchCredentials(ctx)
	if err != nil {
		// STS acquisition failure is fatal: alinas requires cloud credentials
		// and silently falling back to static AK/SK would broaden the effective
		// permission scope. Fail the mount instead.
		return fmt.Errorf("jwtauth fetch STS token: %w", err)
	}

	options, sensitive := splitAlinasSTSOptions(flattenMountOptions(op.Options), cred)
	op.Options = options
	op.SensitiveOptions = append(op.SensitiveOptions, sensitive...)
	klog.V(4).InfoS("injected jwtauth STS credential into alinas sensitive mount options", "target", op.Target)

	if err := handler(ctx, op); err != nil {
		return err
	}

	jwtauth.DefaultManager.Start(op.Target, refresher, cred)
	return nil
}

// splitAlinasSTSOptions strips the infrastructure-only jwtauth options and
// any pre-existing static credential options, ensures tls and ram are present,
// and returns the resolved STS triple separately so it can be passed as
// sensitive mount options (masked in logs). Input must already be flattened.
func splitAlinasSTSOptions(flatOptions []string, cred *jwtauth.STSToken) (options, sensitiveOptions []string) {
	options = make([]string, 0, len(flatOptions)+2)
	hasTLS := false
	hasRAM := false
	for _, opt := range flatOptions {
		key, _, _ := strings.Cut(opt, "=")
		if _, infra := jwtauth.InfraOptionKeys[key]; infra {
			continue
		}
		switch key {
		case jwtauth.OptAuthType:
			// authType is an agent-identity marker consumed by this interceptor.
			// Unlike the OSS FUSE entrypoint, the alinas mount goes straight to
			// mount.nfs, which rejects unknown options ("an incorrect mount
			// option was specified"). Strip it so it never reaches mount.nfs.
			continue
		case optAlinasAccessKeyID, optAlinasAccessKeySecret, optAlinasSecurityToken:
			continue
		case optAlinasTLS:
			hasTLS = true
		case optAlinasRAM:
			hasRAM = true
		}
		options = append(options, opt)
	}
	if !hasTLS {
		options = append(options, optAlinasTLS)
	}
	if !hasRAM {
		options = append(options, optAlinasRAM)
	}
	sensitiveOptions = []string{
		optAlinasAccessKeyID + "=" + cred.AccessKeyID,
		optAlinasAccessKeySecret + "=" + cred.AccessKeySecret,
		optAlinasSecurityToken + "=" + cred.SecurityToken,
	}
	return options, sensitiveOptions
}
