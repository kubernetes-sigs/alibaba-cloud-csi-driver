// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaunchTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateId(v string) *DeleteLaunchTemplateRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *DeleteLaunchTemplateRequest
	GetLaunchTemplateName() *string
	SetOwnerAccount(v string) *DeleteLaunchTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLaunchTemplateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLaunchTemplateRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteLaunchTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteLaunchTemplateRequest
	GetResourceOwnerId() *int64
}

type DeleteLaunchTemplateRequest struct {
	// The ID of the launch template. For more information, see [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html).
	//
	// You must specify `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template.
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the launch template.
	//
	// You must specify `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template.
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

func (s DeleteLaunchTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DeleteLaunchTemplateRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *DeleteLaunchTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLaunchTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLaunchTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLaunchTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteLaunchTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteLaunchTemplateRequest) SetLaunchTemplateId(v string) *DeleteLaunchTemplateRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *DeleteLaunchTemplateRequest) SetLaunchTemplateName(v string) *DeleteLaunchTemplateRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *DeleteLaunchTemplateRequest) SetOwnerAccount(v string) *DeleteLaunchTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLaunchTemplateRequest) SetOwnerId(v int64) *DeleteLaunchTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLaunchTemplateRequest) SetRegionId(v string) *DeleteLaunchTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLaunchTemplateRequest) SetResourceOwnerAccount(v string) *DeleteLaunchTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLaunchTemplateRequest) SetResourceOwnerId(v int64) *DeleteLaunchTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLaunchTemplateRequest) Validate() error {
	return dara.Validate(s)
}
