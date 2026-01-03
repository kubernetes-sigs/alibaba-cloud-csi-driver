// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteSecurityGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSecurityGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSecurityGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSecurityGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSecurityGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *DeleteSecurityGroupRequest
	GetSecurityGroupId() *string
}

type DeleteSecurityGroupRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the security group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group ID. You can call the [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) operation to query the security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSecurityGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSecurityGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSecurityGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeleteSecurityGroupRequest) SetOwnerAccount(v string) *DeleteSecurityGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetOwnerId(v int64) *DeleteSecurityGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetRegionId(v string) *DeleteSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetResourceOwnerAccount(v string) *DeleteSecurityGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetResourceOwnerId(v int64) *DeleteSecurityGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
