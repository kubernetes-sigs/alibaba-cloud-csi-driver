package agentidentity

import (
	"fmt"
	"os"
	"path"
	"time"

	"k8s.io/klog/v2"
)

// DefaultTokenRefreshMargin is how long before a credential expires it is
// renewed, when the environment does not say otherwise.
const DefaultTokenRefreshMargin = 20 * time.Minute

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

// GetTokenRefreshMargin returns how long before expiry a credential is renewed.
// It reads the value from the AGENT_IDENTITY_TOKEN_REFRESH_MARGIN environment
// variable, which takes a Go duration such as "20m" or "90s".
//
// A missing, unparseable or non-positive value falls back to
// DefaultTokenRefreshMargin: a misconfigured margin must not stop volumes from
// mounting, so it is reported and ignored. Note that a margin larger than the
// credential lifetime leaves no room to wait, and the renewal then runs at the
// minimum interval the refresher enforces.
func GetTokenRefreshMargin() time.Duration {
	value := os.Getenv("AGENT_IDENTITY_TOKEN_REFRESH_MARGIN")
	if value == "" {
		return DefaultTokenRefreshMargin
	}
	margin, err := time.ParseDuration(value)
	if err != nil || margin <= 0 {
		klog.Warningf("AGENT_IDENTITY_TOKEN_REFRESH_MARGIN %q is not a duration, using %v: %v", value, DefaultTokenRefreshMargin, err)
		return DefaultTokenRefreshMargin
	}
	return margin
}
