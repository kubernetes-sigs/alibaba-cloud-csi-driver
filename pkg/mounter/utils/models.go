package utils

type MountRequest struct {
	Source      string            `json:"source,omitempty"`
	Target      string            `json:"target,omitempty"`
	Fstype      string            `json:"fstype,omitempty"`
	Args        []string          `json:"args,omitempty"`
	Options     []string          `json:"options,omitempty"`
	MountFlags  []string          `json:"mountFlags,omitempty"`
	Secrets     map[string]string `json:"secrets,omitempty"`
	MetricsPath string            `json:"metricsPath,omitempty"`
	VolumeID    string            `json:"volumeID,omitempty"`
	AuthConfig  *AuthConfig       `json:"authConfig,omitempty"`

	MountResult any
}

type AuthConfig struct {
	AuthType string `json:"authType,omitempty"`
	RoleName string `json:"roleName,omitempty"`

	AccountID string `json:"accountID,omitempty"`
	ClusterID string `json:"clusterID,omitempty"`

	AccessKey     string `json:"accessKey,omitempty"`
	AccessSecret  string `json:"accessSecret,omitempty"`
	SecurityToken string `json:"securityToken,omitempty"`
}
