package jwtauth

import (
	"golang.org/x/sys/unix"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
)

// This file owns the mount-option contract of the agent-identity credential
// provider: which options select the flow, which of them are infrastructure
// only, and how the settings are defaulted from the mount-proxy environment.
// It lives here rather than next to a driver because every storage that
// authorizes through the credential provider shares this contract; only the
// delivery of the resulting credential is storage specific (see CredentialSink).

const (
	// AuthTypeAgentIdentity is the canonical authType selecting this flow.
	AuthTypeAgentIdentity = "agent-identity"

	// AuthTypeJWTAuth is a backward-compatible alias of AuthTypeAgentIdentity.
	AuthTypeJWTAuth = "jwtauth"
)

// Mount options carrying the credential provider configuration. All but
// OptAuthType are infrastructure only: they configure the credential exchange
// and must be stripped before the mount handler runs, see InfraOptionKeys.
const (
	OptAuthType                = "authType"
	OptSandboxId               = "sandboxId"
	OptSandboxCredProviderName = "sandboxCredProviderName"
	OptEndpoint                = "jwtauth_endpoint"
	OptTokenFile               = "jwtauth_token_file"
	OptCredProvider            = "jwtauth_cred_provider"
	OptCAFile                  = "jwtauth_ca_file"
)

// InfraOptionKeys is the set of infrastructure-only options that every
// consumer must strip from the mount options before handing them to the
// mount client, which would otherwise reject or leak them.
var InfraOptionKeys = map[string]struct{}{
	OptSandboxId:               {},
	OptSandboxCredProviderName: {},
	OptEndpoint:                {},
	OptTokenFile:               {},
	OptCredProvider:            {},
	OptCAFile:                  {},
}

// IsAgentIdentity reports whether the given authType selects the
// agent-identity credential flow, accepting the legacy jwtauth alias.
func IsAgentIdentity(authType string) bool {
	return authType == AuthTypeAgentIdentity || authType == AuthTypeJWTAuth
}

// ResolveOpts extracts the credential provider settings from indexed mount
// options, falling back to the shared mount-proxy environment for anything the
// caller did not pin explicitly. A CA file discovered from the environment is
// only used when it is actually readable, since it is optional.
func ResolveOpts(idx map[string]string) Opts {
	opts := Opts{
		TokenFile:    idx[OptTokenFile],
		Endpoint:     idx[OptEndpoint],
		CredProvider: idx[OptCredProvider],
		CAFile:       idx[OptCAFile],
		SandboxId:    idx[OptSandboxId],
	}
	if opts.Endpoint == "" {
		opts.Endpoint = agentidentity.GetEndpoint()
	}
	if opts.TokenFile == "" && opts.SandboxId != "" {
		opts.TokenFile = agentidentity.GetTokenFilePath(opts.SandboxId)
	}
	if opts.CredProvider == "" {
		opts.CredProvider = idx[OptSandboxCredProviderName]
	}
	if opts.CAFile == "" {
		if caPath := agentidentity.GetCAFilePath(); caPath != "" && unix.Access(caPath, unix.R_OK) == nil {
			opts.CAFile = caPath
		}
	}
	return opts
}
