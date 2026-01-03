// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudAssistantSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentUpgradeConfig(v *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) *ModifyCloudAssistantSettingsRequest
	GetAgentUpgradeConfig() *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig
	SetOssDeliveryConfig(v *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) *ModifyCloudAssistantSettingsRequest
	GetOssDeliveryConfig() *ModifyCloudAssistantSettingsRequestOssDeliveryConfig
	SetOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCloudAssistantSettingsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCloudAssistantSettingsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsRequest
	GetResourceOwnerId() *int64
	SetSessionManagerConfig(v *ModifyCloudAssistantSettingsRequestSessionManagerConfig) *ModifyCloudAssistantSettingsRequest
	GetSessionManagerConfig() *ModifyCloudAssistantSettingsRequestSessionManagerConfig
	SetSettingType(v string) *ModifyCloudAssistantSettingsRequest
	GetSettingType() *string
	SetSlsDeliveryConfig(v *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) *ModifyCloudAssistantSettingsRequest
	GetSlsDeliveryConfig() *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig
}

type ModifyCloudAssistantSettingsRequest struct {
	// The configurations for upgrading Cloud Assistant Agent.
	AgentUpgradeConfig *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig `json:"AgentUpgradeConfig,omitempty" xml:"AgentUpgradeConfig,omitempty" type:"Struct"`
	// The configurations for delivering records to Object Storage Service (OSS).
	OssDeliveryConfig *ModifyCloudAssistantSettingsRequestOssDeliveryConfig `json:"OssDeliveryConfig,omitempty" xml:"OssDeliveryConfig,omitempty" type:"Struct"`
	OwnerAccount      *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Cloud Assistant Session Manager configuration.
	SessionManagerConfig *ModifyCloudAssistantSettingsRequestSessionManagerConfig `json:"SessionManagerConfig,omitempty" xml:"SessionManagerConfig,omitempty" type:"Struct"`
	// The Cloud Assistant feature. Set SettingType to one of the following valid values:
	//
	// 	- SessionManagerDelivery: the Session Record Delivery configurations.
	//
	// 	- InvocationDelivery: the Operation Content and Result Delivery configurations.
	//
	// 	- AgentUpgradeConfig: the Cloud Assistant Agent Upgrade configurations.
	//
	// 	- SessionManagerConfig: Cloud Assistant the SessionManager configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// SessionManagerDelivery
	SettingType *string `json:"SettingType,omitempty" xml:"SettingType,omitempty"`
	// The configurations for delivering records to Simple Log Service.
	SlsDeliveryConfig *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig `json:"SlsDeliveryConfig,omitempty" xml:"SlsDeliveryConfig,omitempty" type:"Struct"`
}

func (s ModifyCloudAssistantSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequest) GetAgentUpgradeConfig() *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	return s.AgentUpgradeConfig
}

func (s *ModifyCloudAssistantSettingsRequest) GetOssDeliveryConfig() *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	return s.OssDeliveryConfig
}

func (s *ModifyCloudAssistantSettingsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCloudAssistantSettingsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCloudAssistantSettingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudAssistantSettingsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCloudAssistantSettingsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCloudAssistantSettingsRequest) GetSessionManagerConfig() *ModifyCloudAssistantSettingsRequestSessionManagerConfig {
	return s.SessionManagerConfig
}

func (s *ModifyCloudAssistantSettingsRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *ModifyCloudAssistantSettingsRequest) GetSlsDeliveryConfig() *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	return s.SlsDeliveryConfig
}

func (s *ModifyCloudAssistantSettingsRequest) SetAgentUpgradeConfig(v *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) *ModifyCloudAssistantSettingsRequest {
	s.AgentUpgradeConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetOssDeliveryConfig(v *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) *ModifyCloudAssistantSettingsRequest {
	s.OssDeliveryConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetOwnerId(v int64) *ModifyCloudAssistantSettingsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetRegionId(v string) *ModifyCloudAssistantSettingsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetSessionManagerConfig(v *ModifyCloudAssistantSettingsRequestSessionManagerConfig) *ModifyCloudAssistantSettingsRequest {
	s.SessionManagerConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetSettingType(v string) *ModifyCloudAssistantSettingsRequest {
	s.SettingType = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) SetSlsDeliveryConfig(v *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) *ModifyCloudAssistantSettingsRequest {
	s.SlsDeliveryConfig = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequest) Validate() error {
	if s.AgentUpgradeConfig != nil {
		if err := s.AgentUpgradeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OssDeliveryConfig != nil {
		if err := s.OssDeliveryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SessionManagerConfig != nil {
		if err := s.SessionManagerConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SlsDeliveryConfig != nil {
		if err := s.SlsDeliveryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCloudAssistantSettingsRequestAgentUpgradeConfig struct {
	// The time windows during which Cloud Assistant Agent can be upgraded. The time windows can be accurate to minutes. The Coordinated Universal Time (UTC) time zone is used by default.
	//
	// Make sure that the upgrade windows specified by this parameter are not shorter than 1 hour.
	//
	// Specify each upgrade window in the following format: \\<Start time in the HH:mm format>-\\<End time in the HH:mm format>.
	//
	// For example, [ "02:00-03:00", "05:00-06:00" ] specifies that Cloud Assistant Agent can be upgraded from 2:00:00 to 3:00:00 and from 5:00:00 to 6:00:00 every day in the UTC time zone.
	AllowedUpgradeWindow []*string `json:"AllowedUpgradeWindow,omitempty" xml:"AllowedUpgradeWindow,omitempty" type:"Repeated"`
	// Specifies whether to enable custom upgrade for Cloud Assistant Agent. If you set this parameter to false, an upgrade attempt is performed for Cloud Assistant Agent every 30 minutes.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time zone of the time windows. Default value: UTC. You can specify a time zone in the following forms:
	//
	// 	- The time zone name. Examples: Asia/Shanghai and America/Los_Angeles.
	//
	// 	- The time offset from GMT. Examples: GMT+8:00 (UTC+8) and GMT-7:00 (UTC-7). You cannot add leading zeros to the hour value.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetAllowedUpgradeWindow() []*string {
	return s.AllowedUpgradeWindow
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetAllowedUpgradeWindow(v []*string) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.AllowedUpgradeWindow = v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetEnabled(v bool) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) SetTimeZone(v string) *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig {
	s.TimeZone = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestAgentUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudAssistantSettingsRequestOssDeliveryConfig struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// example-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// Specifies whether to deliver records to OSS. Default value: false.
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
	// 	- KMS: server-side encryption by using Key Management Service managed keys (SSE-KMS).
	//
	// example:
	//
	// Inherit
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The prefix of the OSS bucket directory. The prefix must meet the following requirements:
	//
	// 	- The prefix can be up to 254 characters in length.
	//
	// 	- The prefix cannot start with a forward slash (/) or a backslash (\\\\).
	//
	// Note: If you do not need a directory prefix, specify a pair of double quotation marks ("") for this parameter to clear the directory prefix that you specified.
	//
	// example:
	//
	// sessionmanager/audit
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestOssDeliveryConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEncryptionAlgorithm() *string {
	return s.EncryptionAlgorithm
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetBucketName(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.BucketName = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEnabled(v bool) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEncryptionAlgorithm(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.EncryptionAlgorithm = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEncryptionKeyId(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.EncryptionKeyId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetEncryptionType(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.EncryptionType = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) SetPrefix(v string) *ModifyCloudAssistantSettingsRequestOssDeliveryConfig {
	s.Prefix = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestOssDeliveryConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudAssistantSettingsRequestSessionManagerConfig struct {
	// Specify whether to enable Cloud Assistant Session Manager. Valid values:
	//
	// 	- true: Enables the feature.
	//
	// 	- false: Disables the feature.
	//
	// Notes:
	//
	// 	- The feature applies to all regions.
	//
	// example:
	//
	// true
	SessionManagerEnabled *bool `json:"SessionManagerEnabled,omitempty" xml:"SessionManagerEnabled,omitempty"`
}

func (s ModifyCloudAssistantSettingsRequestSessionManagerConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestSessionManagerConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestSessionManagerConfig) GetSessionManagerEnabled() *bool {
	return s.SessionManagerEnabled
}

func (s *ModifyCloudAssistantSettingsRequestSessionManagerConfig) SetSessionManagerEnabled(v bool) *ModifyCloudAssistantSettingsRequestSessionManagerConfig {
	s.SessionManagerEnabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSessionManagerConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudAssistantSettingsRequestSlsDeliveryConfig struct {
	// Specifies whether to deliver records to Simple Log Service. Default value: false.
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

func (s ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) GetProjectName() *string {
	return s.ProjectName
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) SetEnabled(v bool) *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) SetLogstoreName(v string) *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	s.LogstoreName = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) SetProjectName(v string) *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig {
	s.ProjectName = &v
	return s
}

func (s *ModifyCloudAssistantSettingsRequestSlsDeliveryConfig) Validate() error {
	return dara.Validate(s)
}
