// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *JoinResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *JoinResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *JoinResourceGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *JoinResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *JoinResourceGroupRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *JoinResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *JoinResourceGroupRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *JoinResourceGroupRequest
	GetResourceType() *string
}

type JoinResourceGroupRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which you want to add the instance.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource. For example, if you set ResourceType to instance, set this parameter to the ID of the instance.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the ECS resource. Valid values:
	//
	// 	- instance: instance
	//
	// 	- disk: Elastic Block Storage (EBS) device
	//
	// 	- snapshot: snapshot
	//
	// 	- image: image
	//
	// 	- securitygroup: security group
	//
	// 	- ddh: dedicated host
	//
	// 	- ddhcluster: dedicated host cluster
	//
	// 	- eni: ENI
	//
	// 	- keypair: SSH key pair
	//
	// 	- launchtemplate: launch template
	//
	// 	- command: Cloud Assistant command
	//
	// 	- activation: activation code for a Cloud Assistant managed instance
	//
	// 	- managedinstance: Cloud Assistant managed instance
	//
	// The values are case-sensitive.
	//
	// example:
	//
	// securitygroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s JoinResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *JoinResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *JoinResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *JoinResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *JoinResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *JoinResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *JoinResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *JoinResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *JoinResourceGroupRequest) SetOwnerAccount(v string) *JoinResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *JoinResourceGroupRequest) SetOwnerId(v int64) *JoinResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetRegionId(v string) *JoinResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceGroupId(v string) *JoinResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceId(v string) *JoinResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceOwnerAccount(v string) *JoinResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceOwnerId(v int64) *JoinResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *JoinResourceGroupRequest) SetResourceType(v string) *JoinResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *JoinResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
