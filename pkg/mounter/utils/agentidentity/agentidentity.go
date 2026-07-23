package agentidentity

import (
	"fmt"
	"os"
	"path"
)

// GetTokenDir returns the base directory for agent identity token files.
// It reads the value from the AGENT_IDENTITY_TOKEN_DIR environment variable.
func GetTokenDir() string {
	return os.Getenv("AGENT_IDENTITY_TOKEN_DIR")
}

// GetTokenFilePath returns the agent identity token file path for the given sandbox.
func GetTokenFilePath(sandboxId string) string {
	return path.Join(GetTokenDir(), fmt.Sprintf("%s.token", sandboxId))
}

// GetEndpoint returns the agent identity credential provider endpoint.
// It reads the value from the AGENT_IDENTITY_ENDPOINT environment variable.
func GetEndpoint() string {
	return os.Getenv("AGENT_IDENTITY_ENDPOINT")
}

// GetCAFilePath returns the CA file path for agent identity authentication.
// It reads the value from the AGENT_IDENTITY_CERT_FILE environment variable.
// Returns empty string if not configured (CA is optional).
func GetCAFilePath() string {
	return os.Getenv("AGENT_IDENTITY_CERT_FILE")
}
