// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaunchTemplateVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteVersion(v []*int64) *DeleteLaunchTemplateVersionRequest
	GetDeleteVersion() []*int64
	SetLaunchTemplateId(v string) *DeleteLaunchTemplateVersionRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *DeleteLaunchTemplateVersionRequest
	GetLaunchTemplateName() *string
	SetOwnerAccount(v string) *DeleteLaunchTemplateVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLaunchTemplateVersionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLaunchTemplateVersionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteLaunchTemplateVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteLaunchTemplateVersionRequest
	GetResourceOwnerId() *int64
}

type DeleteLaunchTemplateVersionRequest struct {
	// The version numbers of the launch template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	DeleteVersion []*int64 `json:"DeleteVersion,omitempty" xml:"DeleteVersion,omitempty" type:"Repeated"`
	// The ID of the launch template. For more information, call the [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html) operation.
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the launch template.
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

func (s DeleteLaunchTemplateVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateVersionRequest) GetDeleteVersion() []*int64 {
	return s.DeleteVersion
}

func (s *DeleteLaunchTemplateVersionRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DeleteLaunchTemplateVersionRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *DeleteLaunchTemplateVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLaunchTemplateVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLaunchTemplateVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLaunchTemplateVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteLaunchTemplateVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteLaunchTemplateVersionRequest) SetDeleteVersion(v []*int64) *DeleteLaunchTemplateVersionRequest {
	s.DeleteVersion = v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) SetLaunchTemplateId(v string) *DeleteLaunchTemplateVersionRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) SetLaunchTemplateName(v string) *DeleteLaunchTemplateVersionRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) SetOwnerAccount(v string) *DeleteLaunchTemplateVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) SetOwnerId(v int64) *DeleteLaunchTemplateVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) SetRegionId(v string) *DeleteLaunchTemplateVersionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) SetResourceOwnerAccount(v string) *DeleteLaunchTemplateVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) SetResourceOwnerId(v int64) *DeleteLaunchTemplateVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLaunchTemplateVersionRequest) Validate() error {
	return dara.Validate(s)
}
