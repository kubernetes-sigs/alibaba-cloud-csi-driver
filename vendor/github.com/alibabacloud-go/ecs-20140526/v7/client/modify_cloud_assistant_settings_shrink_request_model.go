// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudAssistantSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentUpgradeConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetAgentUpgradeConfigShrink() *string
	SetOssDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetOssDeliveryConfigShrink() *string
	SetOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest
	GetResourceOwnerId() *int64
	SetSessionManagerConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetSessionManagerConfigShrink() *string
	SetSettingType(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetSettingType() *string
	SetSlsDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest
	GetSlsDeliveryConfigShrink() *string
}

type ModifyCloudAssistantSettingsShrinkRequest struct {
	// The configurations for upgrading Cloud Assistant Agent.
	AgentUpgradeConfigShrink *string `json:"AgentUpgradeConfig,omitempty" xml:"AgentUpgradeConfig,omitempty"`
	// The configurations for delivering records to Object Storage Service (OSS).
	OssDeliveryConfigShrink *string `json:"OssDeliveryConfig,omitempty" xml:"OssDeliveryConfig,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	SessionManagerConfigShrink *string `json:"SessionManagerConfig,omitempty" xml:"SessionManagerConfig,omitempty"`
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
	SlsDeliveryConfigShrink *string `json:"SlsDeliveryConfig,omitempty" xml:"SlsDeliveryConfig,omitempty"`
}

func (s ModifyCloudAssistantSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetAgentUpgradeConfigShrink() *string {
	return s.AgentUpgradeConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetOssDeliveryConfigShrink() *string {
	return s.OssDeliveryConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetSessionManagerConfigShrink() *string {
	return s.SessionManagerConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetSettingType() *string {
	return s.SettingType
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) GetSlsDeliveryConfigShrink() *string {
	return s.SlsDeliveryConfigShrink
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetAgentUpgradeConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.AgentUpgradeConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetOssDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.OssDeliveryConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetRegionId(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetResourceOwnerAccount(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetResourceOwnerId(v int64) *ModifyCloudAssistantSettingsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetSessionManagerConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.SessionManagerConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetSettingType(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.SettingType = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) SetSlsDeliveryConfigShrink(v string) *ModifyCloudAssistantSettingsShrinkRequest {
	s.SlsDeliveryConfigShrink = &v
	return s
}

func (s *ModifyCloudAssistantSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
