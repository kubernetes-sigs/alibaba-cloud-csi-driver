// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentUpgradeConfig(v *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) *DescribeCloudAssistantSettingsResponseBody
	GetAgentUpgradeConfig() *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig
	SetOssDeliveryConfigs(v *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs) *DescribeCloudAssistantSettingsResponseBody
	GetOssDeliveryConfigs() *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs
	SetRequestId(v string) *DescribeCloudAssistantSettingsResponseBody
	GetRequestId() *string
	SetSessionManagerConfig(v *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig) *DescribeCloudAssistantSettingsResponseBody
	GetSessionManagerConfig() *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig
	SetSlsDeliveryConfigs(v *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs) *DescribeCloudAssistantSettingsResponseBody
	GetSlsDeliveryConfigs() *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs
}

type DescribeCloudAssistantSettingsResponseBody struct {
	// The configurations for upgrading Cloud Assistant Agent.
	AgentUpgradeConfig *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig `json:"AgentUpgradeConfig,omitempty" xml:"AgentUpgradeConfig,omitempty" type:"Struct"`
	// The configurations for delivering items to Object Storage Service (OSS).
	OssDeliveryConfigs *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs `json:"OssDeliveryConfigs,omitempty" xml:"OssDeliveryConfigs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Cloud Assistant Session Manager configuration.
	SessionManagerConfig *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig `json:"SessionManagerConfig,omitempty" xml:"SessionManagerConfig,omitempty" type:"Struct"`
	// The configurations for delivering items to Simple Log Service.
	SlsDeliveryConfigs *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs `json:"SlsDeliveryConfigs,omitempty" xml:"SlsDeliveryConfigs,omitempty" type:"Struct"`
}

func (s DescribeCloudAssistantSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBody) GetAgentUpgradeConfig() *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig {
	return s.AgentUpgradeConfig
}

func (s *DescribeCloudAssistantSettingsResponseBody) GetOssDeliveryConfigs() *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs {
	return s.OssDeliveryConfigs
}

func (s *DescribeCloudAssistantSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudAssistantSettingsResponseBody) GetSessionManagerConfig() *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig {
	return s.SessionManagerConfig
}

func (s *DescribeCloudAssistantSettingsResponseBody) GetSlsDeliveryConfigs() *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs {
	return s.SlsDeliveryConfigs
}

func (s *DescribeCloudAssistantSettingsResponseBody) SetAgentUpgradeConfig(v *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) *DescribeCloudAssistantSettingsResponseBody {
	s.AgentUpgradeConfig = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBody) SetOssDeliveryConfigs(v *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs) *DescribeCloudAssistantSettingsResponseBody {
	s.OssDeliveryConfigs = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBody) SetRequestId(v string) *DescribeCloudAssistantSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBody) SetSessionManagerConfig(v *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig) *DescribeCloudAssistantSettingsResponseBody {
	s.SessionManagerConfig = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBody) SetSlsDeliveryConfigs(v *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs) *DescribeCloudAssistantSettingsResponseBody {
	s.SlsDeliveryConfigs = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBody) Validate() error {
	if s.AgentUpgradeConfig != nil {
		if err := s.AgentUpgradeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OssDeliveryConfigs != nil {
		if err := s.OssDeliveryConfigs.Validate(); err != nil {
			return err
		}
	}
	if s.SessionManagerConfig != nil {
		if err := s.SessionManagerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SlsDeliveryConfigs != nil {
		if err := s.SlsDeliveryConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig struct {
	// The time windows during which Cloud Assistant Agent can be upgraded.
	AllowedUpgradeWindows *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows `json:"AllowedUpgradeWindows,omitempty" xml:"AllowedUpgradeWindows,omitempty" type:"Struct"`
	// Indicates whether custom upgrade is enabled for Cloud Assistant Agent. If the value is false or empty, an upgrade attempt is performed for Cloud Assistant Agent every 30 minutes.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time zone of the time windows.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) GetAllowedUpgradeWindows() *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows {
	return s.AllowedUpgradeWindows
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) SetAllowedUpgradeWindows(v *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows) *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig {
	s.AllowedUpgradeWindows = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) SetEnabled(v bool) *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig {
	s.Enabled = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) SetTimeZone(v string) *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig {
	s.TimeZone = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfig) Validate() error {
	if s.AllowedUpgradeWindows != nil {
		if err := s.AllowedUpgradeWindows.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows struct {
	AllowedUpgradeWindow []*string `json:"AllowedUpgradeWindow,omitempty" xml:"AllowedUpgradeWindow,omitempty" type:"Repeated"`
}

func (s DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows) GetAllowedUpgradeWindow() []*string {
	return s.AllowedUpgradeWindow
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows) SetAllowedUpgradeWindow(v []*string) *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows {
	s.AllowedUpgradeWindow = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyAgentUpgradeConfigAllowedUpgradeWindows) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs struct {
	OssDeliveryConfig []*DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig `json:"OssDeliveryConfig,omitempty" xml:"OssDeliveryConfig,omitempty" type:"Repeated"`
}

func (s DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs) GetOssDeliveryConfig() []*DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	return s.OssDeliveryConfig
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs) SetOssDeliveryConfig(v []*DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs {
	s.OssDeliveryConfig = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigs) Validate() error {
	if s.OssDeliveryConfig != nil {
		for _, item := range s.OssDeliveryConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The type of items to be delivered. Valid values:
	//
	// 	- SessionManager: session records.
	//
	// 	- Invocation: task execution records.
	//
	// example:
	//
	// SessionManager
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// Indicates whether to deliver the specified items to OSS.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The OSS encryption algorithm. Valid values:
	//
	// 	- AES256
	//
	// 	- SM4
	//
	// example:
	//
	// AES256
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" xml:"EncryptionAlgorithm,omitempty"`
	// The ID of the customer master key (CMK) when EncryptionType is set to KMS.
	//
	// example:
	//
	// a807****7a70e
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The OSS encryption method. Valid values:
	//
	// 	- Inherit: the encryption method used by the specified bucket.
	//
	// 	- OssManaged: server-side encryption by using OSS-managed keys (SSE-OSS).
	//
	// 	- KMS: server-side encryption with Key Management Service (SSE-KMS).
	//
	// example:
	//
	// Inherit
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The prefix of the OSS bucket directory.
	//
	// example:
	//
	// sessionmanager/audit
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GetEncryptionAlgorithm() *string {
	return s.EncryptionAlgorithm
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) SetBucketName(v string) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	s.BucketName = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) SetDeliveryType(v string) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	s.DeliveryType = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) SetEnabled(v bool) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	s.Enabled = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) SetEncryptionAlgorithm(v string) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	s.EncryptionAlgorithm = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) SetEncryptionKeyId(v string) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	s.EncryptionKeyId = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) SetEncryptionType(v string) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	s.EncryptionType = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) SetPrefix(v string) *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig {
	s.Prefix = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodyOssDeliveryConfigsOssDeliveryConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudAssistantSettingsResponseBodySessionManagerConfig struct {
	// Specify whether to enable Cloud Assistant Session Manager. Valid values:
	//
	// 	- true: Enables the feature.
	//
	// 	- false: Disables the feature.
	//
	// Note:
	//
	// 	- The feature applies to all regions.
	//
	// example:
	//
	// true
	SessionManagerEnabled *bool `json:"SessionManagerEnabled,omitempty" xml:"SessionManagerEnabled,omitempty"`
}

func (s DescribeCloudAssistantSettingsResponseBodySessionManagerConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBodySessionManagerConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig) GetSessionManagerEnabled() *bool {
	return s.SessionManagerEnabled
}

func (s *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig) SetSessionManagerEnabled(v bool) *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig {
	s.SessionManagerEnabled = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodySessionManagerConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs struct {
	SlsDeliveryConfig []*DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig `json:"SlsDeliveryConfig,omitempty" xml:"SlsDeliveryConfig,omitempty" type:"Repeated"`
}

func (s DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs) GetSlsDeliveryConfig() []*DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig {
	return s.SlsDeliveryConfig
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs) SetSlsDeliveryConfig(v []*DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs {
	s.SlsDeliveryConfig = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigs) Validate() error {
	if s.SlsDeliveryConfig != nil {
		for _, item := range s.SlsDeliveryConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig struct {
	// The type of items to be delivered. Valid values:
	//
	// 	- SessionManager: session records.
	//
	// 	- Invocation: task execution records.
	//
	// example:
	//
	// SessionManager
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// Indicates whether to deliver the specified items to Simple Log Service.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// example-logstore
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// example-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) SetDeliveryType(v string) *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig {
	s.DeliveryType = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) SetEnabled(v bool) *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig {
	s.Enabled = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) SetLogstoreName(v string) *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig {
	s.LogstoreName = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) SetProjectName(v string) *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig {
	s.ProjectName = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponseBodySlsDeliveryConfigsSlsDeliveryConfig) Validate() error {
	return dara.Validate(s)
}
