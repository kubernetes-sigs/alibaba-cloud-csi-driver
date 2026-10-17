package agentidentity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestGetTokenRefreshMargin(t *testing.T) {
	t.Run("unset uses the default", func(t *testing.T) {
		assert.Equal(t, DefaultTokenRefreshMargin, GetTokenRefreshMargin())
		assert.Equal(t, 20*time.Minute, DefaultTokenRefreshMargin, "the documented default")
	})

	t.Run("a duration from the environment wins", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_TOKEN_REFRESH_MARGIN", "45m")
		assert.Equal(t, 45*time.Minute, GetTokenRefreshMargin())
	})

	t.Run("sub-minute durations are allowed", func(t *testing.T) {
		t.Setenv("AGENT_IDENTITY_TOKEN_REFRESH_MARGIN", "90s")
		assert.Equal(t, 90*time.Second, GetTokenRefreshMargin())
	})

	t.Run("a misconfigured value falls back instead of breaking mounts", func(t *testing.T) {
		for _, value := range []string{"20", "twenty minutes", "0", "-5m", " "} {
			t.Setenv("AGENT_IDENTITY_TOKEN_REFRESH_MARGIN", value)
			assert.Equal(t, DefaultTokenRefreshMargin, GetTokenRefreshMargin(), "value %q", value)
		}
	})
}

func TestActorIdentityDoesNotFallBackToWorkerWhenIncomplete(t *testing.T) {
	legacy := map[string]string{PodUIDKey: "legacy-actor"}
	require.Equal(t, "legacy-actor", ActorUID(legacy))
	current := map[string]string{ActorUIDKey: "actor-id", ActorNameKey: "actor-name", ActorNamespaceKey: "actor-space", PodUIDKey: "worker-id"}
	require.Equal(t, "actor-id", ActorUID(current))
	delete(current, ActorUIDKey)
	require.Empty(t, ActorUID(current))
}

func TestActorUIDCannotBecomePathOrMountOption(t *testing.T) {
	for _, uid := range []string{"../other", "..", "/root", "actor,option=value", "actor\x00id"} {
		values := map[string]string{ActorUIDKey: uid, ActorNameKey: "actor", ActorNamespaceKey: "space"}
		require.Error(t, ValidateActorIdentity(values))
	}
	require.NoError(t, ValidateActorIdentity(map[string]string{ActorUIDKey: "actor-id", ActorNameKey: "actor", ActorNamespaceKey: "space"}))
}
