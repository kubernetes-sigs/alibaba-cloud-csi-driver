// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeCloudAssistantSettingsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCloudAssistantSettingsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCloudAssistantSettingsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCloudAssistantSettingsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCloudAssistantSettingsRequest
	GetResourceOwnerId() *int64
	SetSettingType(v []*string) *DescribeCloudAssistantSettingsRequest
	GetSettingType() []*string
}

type DescribeCloudAssistantSettingsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The Cloud Assistant configurations.
	//
	// This parameter is required.
	SettingType []*string `json:"SettingType,omitempty" xml:"SettingType,omitempty" type:"Repeated"`
}

func (s DescribeCloudAssistantSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCloudAssistantSettingsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCloudAssistantSettingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudAssistantSettingsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCloudAssistantSettingsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCloudAssistantSettingsRequest) GetSettingType() []*string {
	return s.SettingType
}

func (s *DescribeCloudAssistantSettingsRequest) SetOwnerAccount(v string) *DescribeCloudAssistantSettingsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCloudAssistantSettingsRequest) SetOwnerId(v int64) *DescribeCloudAssistantSettingsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCloudAssistantSettingsRequest) SetRegionId(v string) *DescribeCloudAssistantSettingsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudAssistantSettingsRequest) SetResourceOwnerAccount(v string) *DescribeCloudAssistantSettingsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCloudAssistantSettingsRequest) SetResourceOwnerId(v int64) *DescribeCloudAssistantSettingsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCloudAssistantSettingsRequest) SetSettingType(v []*string) *DescribeCloudAssistantSettingsRequest {
	s.SettingType = v
	return s
}

func (s *DescribeCloudAssistantSettingsRequest) Validate() error {
	return dara.Validate(s)
}
