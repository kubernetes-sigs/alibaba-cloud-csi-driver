// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageShareGroupPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddGroup(v []*string) *ModifyImageShareGroupPermissionRequest
	GetAddGroup() []*string
	SetRemoveGroup(v []*string) *ModifyImageShareGroupPermissionRequest
	GetRemoveGroup() []*string
	SetImageId(v string) *ModifyImageShareGroupPermissionRequest
	GetImageId() *string
	SetOwnerAccount(v string) *ModifyImageShareGroupPermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyImageShareGroupPermissionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyImageShareGroupPermissionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyImageShareGroupPermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyImageShareGroupPermissionRequest
	GetResourceOwnerId() *int64
}

type ModifyImageShareGroupPermissionRequest struct {
	AddGroup    []*string `json:"AddGroup,omitempty" xml:"AddGroup,omitempty" type:"Repeated"`
	RemoveGroup []*string `json:"RemoveGroup,omitempty" xml:"RemoveGroup,omitempty" type:"Repeated"`
	// This parameter is required.
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyImageShareGroupPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageShareGroupPermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageShareGroupPermissionRequest) GetAddGroup() []*string {
	return s.AddGroup
}

func (s *ModifyImageShareGroupPermissionRequest) GetRemoveGroup() []*string {
	return s.RemoveGroup
}

func (s *ModifyImageShareGroupPermissionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageShareGroupPermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyImageShareGroupPermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyImageShareGroupPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImageShareGroupPermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyImageShareGroupPermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyImageShareGroupPermissionRequest) SetAddGroup(v []*string) *ModifyImageShareGroupPermissionRequest {
	s.AddGroup = v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) SetRemoveGroup(v []*string) *ModifyImageShareGroupPermissionRequest {
	s.RemoveGroup = v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) SetImageId(v string) *ModifyImageShareGroupPermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) SetOwnerAccount(v string) *ModifyImageShareGroupPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) SetOwnerId(v int64) *ModifyImageShareGroupPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) SetRegionId(v string) *ModifyImageShareGroupPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) SetResourceOwnerAccount(v string) *ModifyImageShareGroupPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) SetResourceOwnerId(v int64) *ModifyImageShareGroupPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyImageShareGroupPermissionRequest) Validate() error {
	return dara.Validate(s)
}
