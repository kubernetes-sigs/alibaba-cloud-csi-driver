// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySecurityGroupAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifySecurityGroupAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityGroupAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySecurityGroupAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySecurityGroupAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityGroupAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *ModifySecurityGroupAttributeRequest
	GetSecurityGroupId() *string
	SetSecurityGroupName(v string) *ModifySecurityGroupAttributeRequest
	GetSecurityGroupName() *string
}

type ModifySecurityGroupAttributeRequest struct {
	// The new description of the security group. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// By default, the parameter is empty, which indicates that the description remains unchanged.
	//
	// example:
	//
	// TestDescription
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The new name of the security group. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// By default, the parameter is empty, which indicates that the name remains unchanged.
	//
	// example:
	//
	// SecurityGroupTestName
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s ModifySecurityGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityGroupAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityGroupAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityGroupAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityGroupAttributeRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupAttributeRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *ModifySecurityGroupAttributeRequest) SetDescription(v string) *ModifySecurityGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetOwnerAccount(v string) *ModifySecurityGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetOwnerId(v int64) *ModifySecurityGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetRegionId(v string) *ModifySecurityGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetResourceOwnerAccount(v string) *ModifySecurityGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetResourceOwnerId(v int64) *ModifySecurityGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetSecurityGroupId(v string) *ModifySecurityGroupAttributeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) SetSecurityGroupName(v string) *ModifySecurityGroupAttributeRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *ModifySecurityGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
