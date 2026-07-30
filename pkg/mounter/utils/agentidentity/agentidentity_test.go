package agentidentity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenDir(t *testing.T) {
	t.Run("env set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_TOKEN_DIR", "/var/opt/sandbox/agent-token")
		assert.Equal(t, "/var/opt/sandbox/agent-token", GetTokenDir())
	})

	t.Run("env not set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_TOKEN_DIR", "")
		assert.Equal(t, "", GetTokenDir())
	})
}

func TestGetTokenFilePath(t *testing.T) {
	t.Run("env set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_TOKEN_DIR", "/var/opt/sandbox/agent-token")
		assert.Equal(t, "/var/opt/sandbox/agent-token/sandbox-abc.token", GetTokenFilePath("sandbox-abc"))
	})

	t.Run("env not set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_TOKEN_DIR", "")
		assert.Equal(t, "sandbox-abc.token", GetTokenFilePath("sandbox-abc"))
	})
}

func TestGetEndpoint(t *testing.T) {
	t.Run("env set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_ENDPOINT", "https://custom-endpoint:9090")
		assert.Equal(t, "https://custom-endpoint:9090", GetEndpoint())
	})

	t.Run("env not set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_ENDPOINT", "")
		assert.Equal(t, "", GetEndpoint())
	})
}

func TestGetCAFilePath(t *testing.T) {
	t.Run("env set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", "/custom/path/ca.crt")
		assert.Equal(t, "/custom/path/ca.crt", GetCAFilePath())
	})

	t.Run("env not set", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_CERT_FILE", "")
		assert.Equal(t, "", GetCAFilePath())
	})
}
