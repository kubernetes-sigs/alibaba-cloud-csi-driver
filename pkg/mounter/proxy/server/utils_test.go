package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAgentIdentityCAFilePath(t *testing.T) {
	t.Run("SSL_CERT_FILE set, returns env value", func(t *testing.T) {
		t.Setenv("SSL_CERT_FILE", "/custom/truststore/ca-bundle.pem")
		assert.Equal(t, "/custom/truststore/ca-bundle.pem", GetAgentIdentityCAFilePath())
	})

	t.Run("SSL_CERT_FILE empty, falls back to default", func(t *testing.T) {
		t.Setenv("SSL_CERT_FILE", "")
		assert.Equal(t, AgentIdentityCAFilePath, GetAgentIdentityCAFilePath())
	})

	t.Run("SSL_CERT_FILE not set, falls back to default", func(t *testing.T) {
		assert.Equal(t, AgentIdentityCAFilePath, GetAgentIdentityCAFilePath())
	})
}
