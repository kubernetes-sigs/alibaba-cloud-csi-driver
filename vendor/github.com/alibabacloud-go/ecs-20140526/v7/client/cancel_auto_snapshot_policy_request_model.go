// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CancelAutoSnapshotPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelAutoSnapshotPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelAutoSnapshotPolicyRequest
	GetResourceOwnerId() *int64
	SetAutoSnapshotPolicyId(v string) *CancelAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyId() *string
	SetDiskIds(v string) *CancelAutoSnapshotPolicyRequest
	GetDiskIds() *string
	SetRegionId(v string) *CancelAutoSnapshotPolicyRequest
	GetRegionId() *string
}

type CancelAutoSnapshotPolicyRequest struct {
	// RAM用户的虚拟账号ID。
	//
	// example:
	//
	// 155780923770
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 资源主账号的账号名称。
	//
	// example:
	//
	// ECSforCloud
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// 资源主账号的ID，亦即UID。
	//
	// example:
	//
	// 155780923770
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty" xml:"autoSnapshotPolicyId,omitempty"`
	// The IDs of the disks for which you want to disable the automatic snapshot policy. To disable the automatic snapshot policy for multiple disks, you can set this parameter to a JSON array that consists of multiple disk IDs, such as ["dxxxxxxxxx", "dyyyyyyyyy", … "dzzzzzzzzz"]. Separate the disk IDs with commas (,).
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

func (s CancelAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelAutoSnapshotPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelAutoSnapshotPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CancelAutoSnapshotPolicyRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *CancelAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelAutoSnapshotPolicyRequest) SetOwnerId(v int64) *CancelAutoSnapshotPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetResourceOwnerAccount(v string) *CancelAutoSnapshotPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetResourceOwnerId(v int64) *CancelAutoSnapshotPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *CancelAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetDiskIds(v string) *CancelAutoSnapshotPolicyRequest {
	s.DiskIds = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) SetRegionId(v string) *CancelAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
