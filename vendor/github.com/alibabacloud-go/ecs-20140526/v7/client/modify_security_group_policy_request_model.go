// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifySecurityGroupPolicyRequest
	GetClientToken() *string
	SetInnerAccessPolicy(v string) *ModifySecurityGroupPolicyRequest
	GetInnerAccessPolicy() *string
	SetOwnerAccount(v string) *ModifySecurityGroupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityGroupPolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySecurityGroupPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySecurityGroupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityGroupPolicyRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *ModifySecurityGroupPolicyRequest
	GetSecurityGroupId() *string
}

type ModifySecurityGroupPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The internal access control policy of the security group. Valid values:
	//
	// 	- Accept: the internal interconnectivity policy
	//
	// 	- Drop: the internal isolation policy
	//
	// >  The value of this parameter is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// Drop
	InnerAccessPolicy *string `json:"InnerAccessPolicy,omitempty" xml:"InnerAccessPolicy,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4ph****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s ModifySecurityGroupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifySecurityGroupPolicyRequest) GetInnerAccessPolicy() *string {
	return s.InnerAccessPolicy
}

func (s *ModifySecurityGroupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityGroupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityGroupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityGroupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityGroupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityGroupPolicyRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupPolicyRequest) SetClientToken(v string) *ModifySecurityGroupPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) SetInnerAccessPolicy(v string) *ModifySecurityGroupPolicyRequest {
	s.InnerAccessPolicy = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) SetOwnerAccount(v string) *ModifySecurityGroupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) SetOwnerId(v int64) *ModifySecurityGroupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) SetRegionId(v string) *ModifySecurityGroupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) SetResourceOwnerAccount(v string) *ModifySecurityGroupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) SetResourceOwnerId(v int64) *ModifySecurityGroupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) SetSecurityGroupId(v string) *ModifySecurityGroupPolicyRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
