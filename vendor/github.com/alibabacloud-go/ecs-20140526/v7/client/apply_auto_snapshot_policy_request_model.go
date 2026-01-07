// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ApplyAutoSnapshotPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ApplyAutoSnapshotPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ApplyAutoSnapshotPolicyRequest
	GetResourceOwnerId() *int64
	SetAutoSnapshotPolicyId(v string) *ApplyAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyId() *string
	SetDiskIds(v string) *ApplyAutoSnapshotPolicyRequest
	GetDiskIds() *string
	SetRegionId(v string) *ApplyAutoSnapshotPolicyRequest
	GetRegionId() *string
}

type ApplyAutoSnapshotPolicyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-bp14yziiuvu3s6jn****
	AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty" xml:"autoSnapshotPolicyId,omitempty"`
	// The IDs of disks. The value is a JSON array that consists of disk IDs. Separate the disk IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["d-bp14k9cxvr5uzy54****", "d-bp1dtj8v7x6u08iw****", "d-bp1c0tyj9tfli2r8****"]
	DiskIds *string `json:"diskIds,omitempty" xml:"diskIds,omitempty"`
	// The region ID of the automatic snapshot policy and the disks. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ApplyAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ApplyAutoSnapshotPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ApplyAutoSnapshotPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ApplyAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ApplyAutoSnapshotPolicyRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *ApplyAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyAutoSnapshotPolicyRequest) SetOwnerId(v int64) *ApplyAutoSnapshotPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetResourceOwnerAccount(v string) *ApplyAutoSnapshotPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetResourceOwnerId(v int64) *ApplyAutoSnapshotPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetDiskIds(v string) *ApplyAutoSnapshotPolicyRequest {
	s.DiskIds = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetRegionId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
