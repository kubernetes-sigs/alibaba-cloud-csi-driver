// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstanceRamRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *AttachInstanceRamRoleRequest
	GetInstanceIds() *string
	SetOwnerId(v int64) *AttachInstanceRamRoleRequest
	GetOwnerId() *int64
	SetPolicy(v string) *AttachInstanceRamRoleRequest
	GetPolicy() *string
	SetRamRoleName(v string) *AttachInstanceRamRoleRequest
	GetRamRoleName() *string
	SetRegionId(v string) *AttachInstanceRamRoleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachInstanceRamRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachInstanceRamRoleRequest
	GetResourceOwnerId() *int64
}

type AttachInstanceRamRoleRequest struct {
	// The IDs of ECS instances. You can specify 1 to 100 ECS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// [“i-bp14ss25xca5ex1u****”, “i-bp154z5o1qjalfse****”, “i-bp10ws62o04ubhvi****”…]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The additional policy. When you attach an instance RAM role to instances, you can specify an additional policy to further limit the permissions of the role. For more information, see [Policy overview](https://help.aliyun.com/document_detail/93732.html). The value of this parameter must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// {"Statement": [{"Action": ["*"],"Effect": "Allow","Resource": ["*"]}],"Version":"1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The name of the instance RAM role. You can call the [ListRoles](https://help.aliyun.com/document_detail/28713.html) operation provided by RAM to query the instance RAM roles that you created.
	//
	// This parameter is required.
	//
	// example:
	//
	// testRamRoleName
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

func (s AttachInstanceRamRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceRamRoleRequest) GoString() string {
	return s.String()
}

func (s *AttachInstanceRamRoleRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *AttachInstanceRamRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachInstanceRamRoleRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AttachInstanceRamRoleRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *AttachInstanceRamRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachInstanceRamRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachInstanceRamRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachInstanceRamRoleRequest) SetInstanceIds(v string) *AttachInstanceRamRoleRequest {
	s.InstanceIds = &v
	return s
}

func (s *AttachInstanceRamRoleRequest) SetOwnerId(v int64) *AttachInstanceRamRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachInstanceRamRoleRequest) SetPolicy(v string) *AttachInstanceRamRoleRequest {
	s.Policy = &v
	return s
}

func (s *AttachInstanceRamRoleRequest) SetRamRoleName(v string) *AttachInstanceRamRoleRequest {
	s.RamRoleName = &v
	return s
}

func (s *AttachInstanceRamRoleRequest) SetRegionId(v string) *AttachInstanceRamRoleRequest {
	s.RegionId = &v
	return s
}

func (s *AttachInstanceRamRoleRequest) SetResourceOwnerAccount(v string) *AttachInstanceRamRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachInstanceRamRoleRequest) SetResourceOwnerId(v int64) *AttachInstanceRamRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachInstanceRamRoleRequest) Validate() error {
	return dara.Validate(s)
}
