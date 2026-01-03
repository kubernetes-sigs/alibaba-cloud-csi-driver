// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteSnapshotGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSnapshotGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSnapshotGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSnapshotGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSnapshotGroupRequest
	GetResourceOwnerId() *int64
	SetSnapshotGroupId(v string) *DeleteSnapshotGroupRequest
	GetSnapshotGroupId() *string
}

type DeleteSnapshotGroupRequest struct {
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
	// The ID of the snapshot-consistent group. You can call the [DescribeSnapshotGroups](https://help.aliyun.com/document_detail/210940.html) operation to query the IDs of one or more snapshot-consistent groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// ssg-j6c9lpuyxo2uxxny****
	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" xml:"SnapshotGroupId,omitempty"`
}

func (s DeleteSnapshotGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSnapshotGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSnapshotGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnapshotGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSnapshotGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSnapshotGroupRequest) GetSnapshotGroupId() *string {
	return s.SnapshotGroupId
}

func (s *DeleteSnapshotGroupRequest) SetOwnerAccount(v string) *DeleteSnapshotGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSnapshotGroupRequest) SetOwnerId(v int64) *DeleteSnapshotGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSnapshotGroupRequest) SetRegionId(v string) *DeleteSnapshotGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotGroupRequest) SetResourceOwnerAccount(v string) *DeleteSnapshotGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSnapshotGroupRequest) SetResourceOwnerId(v int64) *DeleteSnapshotGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSnapshotGroupRequest) SetSnapshotGroupId(v string) *DeleteSnapshotGroupRequest {
	s.SnapshotGroupId = &v
	return s
}

func (s *DeleteSnapshotGroupRequest) Validate() error {
	return dara.Validate(s)
}
