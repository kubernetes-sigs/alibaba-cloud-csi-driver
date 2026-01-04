package utils

type AuthConfig struct {
	AuthType string
	RoleName string

	AccountID string
	ClusterID string

	AccessKey     string
	AccessSecret  string
	SecurityToken string
}
