// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySnapshotGroupRequest
	GetDescription() *string
	SetName(v string) *ModifySnapshotGroupRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifySnapshotGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySnapshotGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySnapshotGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySnapshotGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySnapshotGroupRequest
	GetResourceOwnerId() *int64
	SetSnapshotGroupId(v string) *ModifySnapshotGroupRequest
	GetSnapshotGroupId() *string
}

type ModifySnapshotGroupRequest struct {
	// The new description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is new description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new name of the snapshot-consistent group. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with a `http://` or `https://`. The name can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
	//
	// example:
	//
	// testName02
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the snapshot-consistent group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the snapshot-consistent group. You can call the [DescribeSnapshotGroups](https://help.aliyun.com/document_detail/210940.html) operation to query the IDs of available snapshot-consistent groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// ssg-j6ciyh3k52qp7ovm****
	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" xml:"SnapshotGroupId,omitempty"`
}

func (s ModifySnapshotGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifySnapshotGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySnapshotGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifySnapshotGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySnapshotGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySnapshotGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySnapshotGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySnapshotGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySnapshotGroupRequest) GetSnapshotGroupId() *string {
	return s.SnapshotGroupId
}

func (s *ModifySnapshotGroupRequest) SetDescription(v string) *ModifySnapshotGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifySnapshotGroupRequest) SetName(v string) *ModifySnapshotGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifySnapshotGroupRequest) SetOwnerAccount(v string) *ModifySnapshotGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySnapshotGroupRequest) SetOwnerId(v int64) *ModifySnapshotGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySnapshotGroupRequest) SetRegionId(v string) *ModifySnapshotGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySnapshotGroupRequest) SetResourceOwnerAccount(v string) *ModifySnapshotGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySnapshotGroupRequest) SetResourceOwnerId(v int64) *ModifySnapshotGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySnapshotGroupRequest) SetSnapshotGroupId(v string) *ModifySnapshotGroupRequest {
	s.SnapshotGroupId = &v
	return s
}

func (s *ModifySnapshotGroupRequest) Validate() error {
	return dara.Validate(s)
}
