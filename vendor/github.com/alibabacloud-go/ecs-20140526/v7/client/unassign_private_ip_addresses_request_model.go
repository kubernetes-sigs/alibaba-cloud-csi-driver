// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignPrivateIpAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpv4Prefix(v []*string) *UnassignPrivateIpAddressesRequest
	GetIpv4Prefix() []*string
	SetNetworkInterfaceId(v string) *UnassignPrivateIpAddressesRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *UnassignPrivateIpAddressesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassignPrivateIpAddressesRequest
	GetOwnerId() *int64
	SetPrivateIpAddress(v []*string) *UnassignPrivateIpAddressesRequest
	GetPrivateIpAddress() []*string
	SetRegionId(v string) *UnassignPrivateIpAddressesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassignPrivateIpAddressesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassignPrivateIpAddressesRequest
	GetResourceOwnerId() *int64
}

type UnassignPrivateIpAddressesRequest struct {
	// The IPv4 prefixes to unassign.
	Ipv4Prefix []*string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty" type:"Repeated"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp67acfmxazb4ph****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The secondary private IP addresses to unassign.
	//
	// example:
	//
	// ``192.168.**.**``
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	// The region ID of the ENI. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s UnassignPrivateIpAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassignPrivateIpAddressesRequest) GoString() string {
	return s.String()
}

func (s *UnassignPrivateIpAddressesRequest) GetIpv4Prefix() []*string {
	return s.Ipv4Prefix
}

func (s *UnassignPrivateIpAddressesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *UnassignPrivateIpAddressesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassignPrivateIpAddressesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassignPrivateIpAddressesRequest) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *UnassignPrivateIpAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassignPrivateIpAddressesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassignPrivateIpAddressesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassignPrivateIpAddressesRequest) SetIpv4Prefix(v []*string) *UnassignPrivateIpAddressesRequest {
	s.Ipv4Prefix = v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetNetworkInterfaceId(v string) *UnassignPrivateIpAddressesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetOwnerAccount(v string) *UnassignPrivateIpAddressesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetOwnerId(v int64) *UnassignPrivateIpAddressesRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetPrivateIpAddress(v []*string) *UnassignPrivateIpAddressesRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetRegionId(v string) *UnassignPrivateIpAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetResourceOwnerAccount(v string) *UnassignPrivateIpAddressesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) SetResourceOwnerId(v int64) *UnassignPrivateIpAddressesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassignPrivateIpAddressesRequest) Validate() error {
	return dara.Validate(s)
}
