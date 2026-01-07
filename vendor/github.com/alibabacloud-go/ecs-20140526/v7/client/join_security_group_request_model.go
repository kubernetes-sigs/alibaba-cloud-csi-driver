// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *JoinSecurityGroupRequest
	GetInstanceId() *string
	SetNetworkInterfaceId(v string) *JoinSecurityGroupRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *JoinSecurityGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *JoinSecurityGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *JoinSecurityGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *JoinSecurityGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *JoinSecurityGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *JoinSecurityGroupRequest
	GetSecurityGroupId() *string
}

type JoinSecurityGroupRequest struct {
	// The instance ID.
	//
	// > If you configure this parameter, you cannot configure `NetworkInterfaceId`.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ENI ID.
	//
	// > If you configure this parameter, you cannot configure `InstanceId`.
	//
	// example:
	//
	// eni-bp13kd656hxambfe****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// 	- If you want to add an instance to a security group, you do not need to specify a region ID.
	//
	// 	- If you want to add an ENI to a security group, you must specify the region ID of the ENI.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group. You can call the [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) operation to query the most recent security group list.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s JoinSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *JoinSecurityGroupRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *JoinSecurityGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *JoinSecurityGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *JoinSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *JoinSecurityGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *JoinSecurityGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *JoinSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *JoinSecurityGroupRequest) SetInstanceId(v string) *JoinSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetNetworkInterfaceId(v string) *JoinSecurityGroupRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetOwnerAccount(v string) *JoinSecurityGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetOwnerId(v int64) *JoinSecurityGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetRegionId(v string) *JoinSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetResourceOwnerAccount(v string) *JoinSecurityGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetResourceOwnerId(v int64) *JoinSecurityGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *JoinSecurityGroupRequest) SetSecurityGroupId(v string) *JoinSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *JoinSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
