package jwtauth

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsAgentIdentity(t *testing.T) {
	cases := map[string]bool{
		"agent-identity": true,
		"jwtauth":        true,
		"rrsa":           false,
		"sts":            false,
		"":               false,
	}
	for authType, want := range cases {
		assert.Equal(t, want, IsAgentIdentity(authType), "IsAgentIdentity(%q)", authType)
	}
}

func TestResolveOptsDefaults(t *testing.T) {
	t.Setenv("AGENT_IDENTITY_ENDPOINT", "https://cred:8443/")
	t.Setenv("AGENT_IDENTITY_TOKEN_DIR", "/var/run/agent-token")
	idx := map[string]string{
		OptAuthType:                "jwtauth",
		OptSandboxId:               "sb-123",
		OptSandboxCredProviderName: "my-cred",
	}
	opts := ResolveOpts(idx)
	assert.Equal(t, "sb-123", opts.SandboxId)
	assert.Equal(t, "my-cred", opts.CredProvider)
	assert.Equal(t, agentidentity.GetTokenFilePath("sb-123"), opts.TokenFile)
	assert.Equal(t, agentidentity.GetEndpoint(), opts.Endpoint)
}

func TestResolveOptsExplicitOverrides(t *testing.T) {
	idx := map[string]string{
		OptSandboxId:    "sb-1",
		OptEndpoint:     "https://custom:9443/",
		OptTokenFile:    "/custom/token",
		OptCredProvider: "explicit-cred",
	}
	opts := ResolveOpts(idx)
	assert.Equal(t, "https://custom:9443/", opts.Endpoint)
	assert.Equal(t, "/custom/token", opts.TokenFile)
	assert.Equal(t, "explicit-cred", opts.CredProvider)
}

func TestResolveOptsCAFile(t *testing.T) {
	t.Run("readable CA file from env is used", func(t *testing.T) {
		caPath := filepath.Join(t.TempDir(), "ca.crt")
		require.NoError(t, os.WriteFile(caPath, []byte("pem"), 0600))
		t.Setenv("AGENT_IDENTITY_CERT_FILE", caPath)
		opts := ResolveOpts(map[string]string{OptSandboxId: "sb-1"})
		assert.Equal(t, caPath, opts.CAFile)
	})

	t.Run("unreadable CA file from env is ignored", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", filepath.Join(t.TempDir(), "missing.crt"))
		opts := ResolveOpts(map[string]string{OptSandboxId: "sb-1"})
		assert.Empty(t, opts.CAFile)
	})

	t.Run("explicit CA file option wins over env", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", "/env/ca.crt")
		opts := ResolveOpts(map[string]string{
			OptSandboxId: "sb-1",
			OptCAFile:    "/explicit/ca.crt",
		})
		assert.Equal(t, "/explicit/ca.crt", opts.CAFile)
	})
}

func TestInfraOptionKeysExcludesAuthType(t *testing.T) {
	// authType must survive stripping: it is the marker consumers branch on,
	// unlike the settings that only configure the credential exchange.
	_, isInfra := InfraOptionKeys[OptAuthType]
	assert.False(t, isInfra, "authType must not be treated as infrastructure only")
	for _, key := range []string{
		OptSandboxId, OptSandboxCredProviderName, OptEndpoint,
		OptTokenFile, OptCredProvider, OptCAFile,
	} {
		_, ok := InfraOptionKeys[key]
		assert.True(t, ok, "%s must be stripped before the mount", key)
	}
}
