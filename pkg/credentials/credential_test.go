package credentials

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvCredential(t *testing.T) {
	t.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "test-accessKeyId")
	t.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "test-accessKeySecret")

	cred, err := NewCredential()
	assert.NoError(t, err)

	m, err := cred.GetCredential()
	assert.NoError(t, err)

	assert.Equal(t, "test-accessKeyId", *m.AccessKeyId)
	assert.Equal(t, "test-accessKeySecret", *m.AccessKeySecret)
	assert.Equal(t, "access_key", *m.Type)
}

func TestDeprecatedEnvCredential(t *testing.T) {
	t.Setenv("ACCESS_KEY_ID", "test-accessKeyId")
	t.Setenv("ACCESS_KEY_SECRET", "test-accessKeySecret")
	// To enable cleanup, also set the new env var
	t.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "")
	t.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "")

	cred, err := NewCredential()
	assert.NoError(t, err)

	m, err := cred.GetCredential()
	assert.NoError(t, err)

	assert.Equal(t, "test-accessKeyId", *m.AccessKeyId)
	assert.Equal(t, "test-accessKeySecret", *m.AccessKeySecret)
	assert.Equal(t, "access_key", *m.Type)
}

func TestFallbackEcsRamRole(t *testing.T) {
	t.Setenv(TokenPathEnvName, "/no/token/path")

	cred, err := NewCredential()
	assert.NoError(t, err)

	assert.Equal(t, "ecs_ram_role", *cred.GetType())
}

func TestInvalidManagedToken(t *testing.T) {
	tokenPath := t.TempDir() + "/token-config"
	t.Setenv(TokenPathEnvName, tokenPath)
	os.WriteFile(tokenPath, []byte("invalid token"), 0644)

	_, err := NewCredential()
	assert.Error(t, err)
}

var testConfig = []byte(`[default]
type = access_key
access_key_id = foo
access_key_secret = bar`)

func TestConfigFile(t *testing.T) {
	configPath := t.TempDir() + "/config"
	t.Setenv("ALIBABA_CLOUD_CREDENTIALS_FILE", configPath)
	os.WriteFile(configPath, testConfig, 0600)

	cred, err := NewCredential()
	assert.NoError(t, err)

	m, err := cred.GetCredential()
	assert.NoError(t, err)

	assert.Equal(t, "foo", *m.AccessKeyId)
	assert.Equal(t, "bar", *m.AccessKeySecret)
	assert.Empty(t, m.SecurityToken)
	assert.Equal(t, "access_key", *m.Type)
}
