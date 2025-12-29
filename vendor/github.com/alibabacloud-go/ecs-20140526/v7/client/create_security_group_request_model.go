// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSecurityGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateSecurityGroupRequest
	GetDescription() *string
	SetOwnerAccount(v string) *CreateSecurityGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSecurityGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateSecurityGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSecurityGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateSecurityGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSecurityGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupName(v string) *CreateSecurityGroupRequest
	GetSecurityGroupName() *string
	SetSecurityGroupType(v string) *CreateSecurityGroupRequest
	GetSecurityGroupType() *string
	SetServiceManaged(v bool) *CreateSecurityGroupRequest
	GetServiceManaged() *bool
	SetTag(v []*CreateSecurityGroupRequestTag) *CreateSecurityGroupRequest
	GetTag() []*CreateSecurityGroupRequestTag
	SetVpcId(v string) *CreateSecurityGroupRequest
	GetVpcId() *string
}

type CreateSecurityGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.*	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the security group. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// testDescription
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the security group belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the security group. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// testSecurityGroupName
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// The type of the security group. Valid values:
	//
	// 	- normal: basic security group
	//
	// 	- enterprise: advanced security group For more information, see [Advanced security groups](https://help.aliyun.com/document_detail/120621.html).
	//
	// Default value: normal.
	//
	// example:
	//
	// enterprise
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The tags to add to the security group. You can add up to 20 tags.
	Tag []*CreateSecurityGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC in which you want to create the security group.
	//
	// > The VpcId parameter is required only if you want to create security groups of the VPC type. In regions that support the classic network, you can create security groups of the classic network type without the need to specify the VpcId parameter.
	//
	// example:
	//
	// vpc-bp1opxu1zkhn00gzv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSecurityGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSecurityGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSecurityGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSecurityGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSecurityGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSecurityGroupRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *CreateSecurityGroupRequest) GetSecurityGroupType() *string {
	return s.SecurityGroupType
}

func (s *CreateSecurityGroupRequest) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *CreateSecurityGroupRequest) GetTag() []*CreateSecurityGroupRequestTag {
	return s.Tag
}

func (s *CreateSecurityGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateSecurityGroupRequest) SetClientToken(v string) *CreateSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetDescription(v string) *CreateSecurityGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetOwnerAccount(v string) *CreateSecurityGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetOwnerId(v int64) *CreateSecurityGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetRegionId(v string) *CreateSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetResourceGroupId(v string) *CreateSecurityGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetResourceOwnerAccount(v string) *CreateSecurityGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetResourceOwnerId(v int64) *CreateSecurityGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetSecurityGroupName(v string) *CreateSecurityGroupRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetSecurityGroupType(v string) *CreateSecurityGroupRequest {
	s.SecurityGroupType = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetServiceManaged(v bool) *CreateSecurityGroupRequest {
	s.ServiceManaged = &v
	return s
}

func (s *CreateSecurityGroupRequest) SetTag(v []*CreateSecurityGroupRequestTag) *CreateSecurityGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateSecurityGroupRequest) SetVpcId(v string) *CreateSecurityGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSecurityGroupRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSecurityGroupRequestTag struct {
	// The key of the tag to add to the security group.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the security group.
	//
	// The tag value can be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSecurityGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateSecurityGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateSecurityGroupRequestTag) SetKey(v string) *CreateSecurityGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSecurityGroupRequestTag) SetValue(v string) *CreateSecurityGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateSecurityGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
