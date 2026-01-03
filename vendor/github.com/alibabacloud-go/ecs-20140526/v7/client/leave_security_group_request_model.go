// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *LeaveSecurityGroupRequest
	GetInstanceId() *string
	SetNetworkInterfaceId(v string) *LeaveSecurityGroupRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *LeaveSecurityGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *LeaveSecurityGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *LeaveSecurityGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *LeaveSecurityGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LeaveSecurityGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *LeaveSecurityGroupRequest
	GetSecurityGroupId() *string
}

type LeaveSecurityGroupRequest struct {
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
	// 	- If you want to remove an instance from a security group, you do not need to specify a region ID.
	//
	// 	- If you want to remove an ENI from a security group, you must specify the ID of the region in which the ENI resides.
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
}

func (s LeaveSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s LeaveSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LeaveSecurityGroupRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *LeaveSecurityGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *LeaveSecurityGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LeaveSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LeaveSecurityGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LeaveSecurityGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LeaveSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *LeaveSecurityGroupRequest) SetInstanceId(v string) *LeaveSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetNetworkInterfaceId(v string) *LeaveSecurityGroupRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetOwnerAccount(v string) *LeaveSecurityGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetOwnerId(v int64) *LeaveSecurityGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetRegionId(v string) *LeaveSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetResourceOwnerAccount(v string) *LeaveSecurityGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetResourceOwnerId(v int64) *LeaveSecurityGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) SetSecurityGroupId(v string) *LeaveSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *LeaveSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
