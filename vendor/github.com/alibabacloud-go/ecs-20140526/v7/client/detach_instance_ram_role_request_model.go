// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstanceRamRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DetachInstanceRamRoleRequest
	GetInstanceIds() *string
	SetOwnerId(v int64) *DetachInstanceRamRoleRequest
	GetOwnerId() *int64
	SetRamRoleName(v string) *DetachInstanceRamRoleRequest
	GetRamRoleName() *string
	SetRegionId(v string) *DetachInstanceRamRoleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachInstanceRamRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachInstanceRamRoleRequest
	GetResourceOwnerId() *int64
}

type DetachInstanceRamRoleRequest struct {
	// The IDs of ECS instances. You can specify 1 to 100 instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-bp67acfmxazb4p****", "i-bp67acfmxazb5p****", "i-bp67acfmxazb6p****"…]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the instance RAM role. You can call the [ListRoles](https://help.aliyun.com/document_detail/28713.html) operation of RAM to query the names of available instance RAM roles.
	//
	// example:
	//
	// RamRoleTest
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DetachInstanceRamRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceRamRoleRequest) GoString() string {
	return s.String()
}

func (s *DetachInstanceRamRoleRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DetachInstanceRamRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachInstanceRamRoleRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DetachInstanceRamRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachInstanceRamRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachInstanceRamRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachInstanceRamRoleRequest) SetInstanceIds(v string) *DetachInstanceRamRoleRequest {
	s.InstanceIds = &v
	return s
}

func (s *DetachInstanceRamRoleRequest) SetOwnerId(v int64) *DetachInstanceRamRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachInstanceRamRoleRequest) SetRamRoleName(v string) *DetachInstanceRamRoleRequest {
	s.RamRoleName = &v
	return s
}

func (s *DetachInstanceRamRoleRequest) SetRegionId(v string) *DetachInstanceRamRoleRequest {
	s.RegionId = &v
	return s
}

func (s *DetachInstanceRamRoleRequest) SetResourceOwnerAccount(v string) *DetachInstanceRamRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachInstanceRamRoleRequest) SetResourceOwnerId(v int64) *DetachInstanceRamRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachInstanceRamRoleRequest) Validate() error {
	return dara.Validate(s)
}
