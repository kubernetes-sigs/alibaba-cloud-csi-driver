// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLaunchTemplateDefaultVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultVersionNumber(v int64) *ModifyLaunchTemplateDefaultVersionRequest
	GetDefaultVersionNumber() *int64
	SetLaunchTemplateId(v string) *ModifyLaunchTemplateDefaultVersionRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *ModifyLaunchTemplateDefaultVersionRequest
	GetLaunchTemplateName() *string
	SetOwnerAccount(v string) *ModifyLaunchTemplateDefaultVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLaunchTemplateDefaultVersionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLaunchTemplateDefaultVersionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLaunchTemplateDefaultVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyLaunchTemplateDefaultVersionRequest
	GetResourceOwnerId() *int64
}

type ModifyLaunchTemplateDefaultVersionRequest struct {
	// The default version number of the instance launch template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	DefaultVersionNumber *int64 `json:"DefaultVersionNumber,omitempty" xml:"DefaultVersionNumber,omitempty"`
	// The ID of the launch template. You must specify the LaunchTemplateId or LaunchTemplateName parameter to determine an instance launch template.
	//
	// example:
	//
	// lt-s-bp177juajht6****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the instance launch template. You must specify the LaunchTemplateId or LaunchTemplateName parameter to determine an instance launch template.
	//
	// example:
	//
	// testLaunchTemplateName
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the launch template. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyLaunchTemplateDefaultVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLaunchTemplateDefaultVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetDefaultVersionNumber() *int64 {
	return s.DefaultVersionNumber
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetDefaultVersionNumber(v int64) *ModifyLaunchTemplateDefaultVersionRequest {
	s.DefaultVersionNumber = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetLaunchTemplateId(v string) *ModifyLaunchTemplateDefaultVersionRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetLaunchTemplateName(v string) *ModifyLaunchTemplateDefaultVersionRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetOwnerAccount(v string) *ModifyLaunchTemplateDefaultVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetOwnerId(v int64) *ModifyLaunchTemplateDefaultVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetRegionId(v string) *ModifyLaunchTemplateDefaultVersionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetResourceOwnerAccount(v string) *ModifyLaunchTemplateDefaultVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) SetResourceOwnerId(v int64) *ModifyLaunchTemplateDefaultVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionRequest) Validate() error {
	return dara.Validate(s)
}
