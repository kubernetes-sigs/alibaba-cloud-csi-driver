package server

import "os"

// AgentIdentityCAFilePath is the default path for the agent identity CA certificate.
// See AgentIdentityConfig for file placement scenarios.
const AgentIdentityCAFilePath = "/etc/ssl/certs/agent-identity/ca.crt"

// GetAgentIdentityCAFilePath resolves the CA file path for agent identity authentication.
// It prefers the SSL_CERT_FILE environment variable; if unset or empty, it falls back to
// AgentIdentityCAFilePath.
func GetAgentIdentityCAFilePath() string {
	if p := os.Getenv("SSL_CERT_FILE"); p != "" {
		return p
	}
	return AgentIdentityCAFilePath
}
