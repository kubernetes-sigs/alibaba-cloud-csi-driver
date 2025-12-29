// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageSharePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddAccount(v []*string) *ModifyImageSharePermissionRequest
	GetAddAccount() []*string
	SetDryRun(v bool) *ModifyImageSharePermissionRequest
	GetDryRun() *bool
	SetImageId(v string) *ModifyImageSharePermissionRequest
	GetImageId() *string
	SetIsPublic(v bool) *ModifyImageSharePermissionRequest
	GetIsPublic() *bool
	SetLaunchPermission(v string) *ModifyImageSharePermissionRequest
	GetLaunchPermission() *string
	SetOwnerAccount(v string) *ModifyImageSharePermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyImageSharePermissionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyImageSharePermissionRequest
	GetRegionId() *string
	SetRemoveAccount(v []*string) *ModifyImageSharePermissionRequest
	GetRemoveAccount() []*string
	SetResourceOwnerAccount(v string) *ModifyImageSharePermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyImageSharePermissionRequest
	GetResourceOwnerId() *int64
}

type ModifyImageSharePermissionRequest struct {
	// The IDs of Alibaba Cloud accounts to which you want to share the custom image. Valid values of N: 1 to 10. If you specify more than 10 Alibaba Cloud account IDs, the system processes only the first 10 account IDs. The excess account IDs are ignored.
	//
	// example:
	//
	// 1234567890
	AddAccount []*string `json:"AddAccount,omitempty" xml:"AddAccount,omitempty" type:"Repeated"`
	DryRun     *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the shared custom image.
	//
	// >  You can share images encrypted by using CMKs but cannot share images encrypted by using service keys. When you share an image encrypted by using a service key, an error is reported.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp18ygjuqnwhechc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Specifies whether to publish or unpublish a community image. Valid values:
	//
	// 	- true: publishes the custom image as a community image.
	//
	// 	- false: unpublishes a community image. The unpublish operation takes effect only on community images.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	LaunchPermission *string `json:"LaunchPermission,omitempty" xml:"LaunchPermission,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the custom image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of Alibaba Cloud accounts from which you want to unshare the custom image. Valid values of N: 1 to 10. If you specify more than 10 Alibaba Cloud account IDs, the system processes only the first 10 account IDs. The excess account IDs are ignored.
	//
	// example:
	//
	// 1234567890
	RemoveAccount        []*string `json:"RemoveAccount,omitempty" xml:"RemoveAccount,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyImageSharePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionRequest) GetAddAccount() []*string {
	return s.AddAccount
}

func (s *ModifyImageSharePermissionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyImageSharePermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageSharePermissionRequest) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *ModifyImageSharePermissionRequest) GetLaunchPermission() *string {
	return s.LaunchPermission
}

func (s *ModifyImageSharePermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyImageSharePermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyImageSharePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImageSharePermissionRequest) GetRemoveAccount() []*string {
	return s.RemoveAccount
}

func (s *ModifyImageSharePermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyImageSharePermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyImageSharePermissionRequest) SetAddAccount(v []*string) *ModifyImageSharePermissionRequest {
	s.AddAccount = v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetDryRun(v bool) *ModifyImageSharePermissionRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetImageId(v string) *ModifyImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetIsPublic(v bool) *ModifyImageSharePermissionRequest {
	s.IsPublic = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetLaunchPermission(v string) *ModifyImageSharePermissionRequest {
	s.LaunchPermission = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetOwnerAccount(v string) *ModifyImageSharePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetOwnerId(v int64) *ModifyImageSharePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetRegionId(v string) *ModifyImageSharePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetRemoveAccount(v []*string) *ModifyImageSharePermissionRequest {
	s.RemoveAccount = v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetResourceOwnerAccount(v string) *ModifyImageSharePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) SetResourceOwnerId(v int64) *ModifyImageSharePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyImageSharePermissionRequest) Validate() error {
	return dara.Validate(s)
}
