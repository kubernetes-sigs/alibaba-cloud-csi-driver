// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRamRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeInstanceRamRoleRequest
	GetInstanceIds() *string
	SetOwnerId(v int64) *DescribeInstanceRamRoleRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstanceRamRoleRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceRamRoleRequest
	GetPageSize() *int32
	SetRamRoleName(v string) *DescribeInstanceRamRoleRequest
	GetRamRoleName() *string
	SetRegionId(v string) *DescribeInstanceRamRoleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceRamRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceRamRoleRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceRamRoleRequest struct {
	// The IDs of ECS instances. You can specify up to 50 instance IDs in a single request.
	//
	// >  You must specify at least one parameter from `InstanceIds` and `RamRoleName`.
	//
	// example:
	//
	// ["i-bp67acfmxazb1p****", "i-bp67acfmxazb2p****", "bp67acfmxazb3p****"…]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the instance RAM role. If you specify this parameter, all ECS instances to which the instance RAM role is attached are returned in the response. You can call the [ListRoles](https://help.aliyun.com/document_detail/28713.html) operation of RAM to query the names of available instance RAM roles.
	//
	// >  You must specify at least one parameter from `InstanceIds` and `RamRoleName`.
	//
	// example:
	//
	// EcsServiceRole-EcsDocGuideTest
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance RAM role. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
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

func (s DescribeInstanceRamRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRamRoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRamRoleRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstanceRamRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceRamRoleRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceRamRoleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceRamRoleRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeInstanceRamRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceRamRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceRamRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceRamRoleRequest) SetInstanceIds(v string) *DescribeInstanceRamRoleRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) SetOwnerId(v int64) *DescribeInstanceRamRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) SetPageNumber(v int32) *DescribeInstanceRamRoleRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) SetPageSize(v int32) *DescribeInstanceRamRoleRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) SetRamRoleName(v string) *DescribeInstanceRamRoleRequest {
	s.RamRoleName = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) SetRegionId(v string) *DescribeInstanceRamRoleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) SetResourceOwnerAccount(v string) *DescribeInstanceRamRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) SetResourceOwnerId(v int64) *DescribeInstanceRamRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceRamRoleRequest) Validate() error {
	return dara.Validate(s)
}
