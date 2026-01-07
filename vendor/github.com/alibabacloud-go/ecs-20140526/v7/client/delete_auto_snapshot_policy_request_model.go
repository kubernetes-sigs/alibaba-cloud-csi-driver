// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteAutoSnapshotPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAutoSnapshotPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAutoSnapshotPolicyRequest
	GetResourceOwnerId() *int64
	SetAutoSnapshotPolicyId(v string) *DeleteAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyId() *string
	SetRegionId(v string) *DeleteAutoSnapshotPolicyRequest
	GetRegionId() *string
}

type DeleteAutoSnapshotPolicyRequest struct {
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
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the automatic snapshot policy. You can call the [DescribeAutoSnapshotPolicyEx](https://help.aliyun.com/document_detail/25530.html) operation to query the IDs of available automatic snapshot policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-bp14yziiuvu3s6jn****
	AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty" xml:"autoSnapshotPolicyId,omitempty"`
	// The ID of the region to which the automatic snapshot policy belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoSnapshotPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAutoSnapshotPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAutoSnapshotPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DeleteAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAutoSnapshotPolicyRequest) SetOwnerId(v int64) *DeleteAutoSnapshotPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) SetResourceOwnerAccount(v string) *DeleteAutoSnapshotPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) SetResourceOwnerId(v int64) *DeleteAutoSnapshotPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) SetRegionId(v string) *DeleteAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
