// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignIpv6AddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6Address(v []*string) *UnassignIpv6AddressesRequest
	GetIpv6Address() []*string
	SetIpv6Prefix(v []*string) *UnassignIpv6AddressesRequest
	GetIpv6Prefix() []*string
	SetNetworkInterfaceId(v string) *UnassignIpv6AddressesRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *UnassignIpv6AddressesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassignIpv6AddressesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassignIpv6AddressesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassignIpv6AddressesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassignIpv6AddressesRequest
	GetResourceOwnerId() *int64
}

type UnassignIpv6AddressesRequest struct {
	// IPv6 address N to unassign. Valid values of N: 1 to 10.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// IPv6 prefix N to unassign. Valid values of N: 1 to 10.
	Ipv6Prefix []*string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty" type:"Repeated"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp14v2sdd3v8ht****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s UnassignIpv6AddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassignIpv6AddressesRequest) GoString() string {
	return s.String()
}

func (s *UnassignIpv6AddressesRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *UnassignIpv6AddressesRequest) GetIpv6Prefix() []*string {
	return s.Ipv6Prefix
}

func (s *UnassignIpv6AddressesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *UnassignIpv6AddressesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassignIpv6AddressesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassignIpv6AddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassignIpv6AddressesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassignIpv6AddressesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassignIpv6AddressesRequest) SetIpv6Address(v []*string) *UnassignIpv6AddressesRequest {
	s.Ipv6Address = v
	return s
}

func (s *UnassignIpv6AddressesRequest) SetIpv6Prefix(v []*string) *UnassignIpv6AddressesRequest {
	s.Ipv6Prefix = v
	return s
}

func (s *UnassignIpv6AddressesRequest) SetNetworkInterfaceId(v string) *UnassignIpv6AddressesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *UnassignIpv6AddressesRequest) SetOwnerAccount(v string) *UnassignIpv6AddressesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassignIpv6AddressesRequest) SetOwnerId(v int64) *UnassignIpv6AddressesRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassignIpv6AddressesRequest) SetRegionId(v string) *UnassignIpv6AddressesRequest {
	s.RegionId = &v
	return s
}

func (s *UnassignIpv6AddressesRequest) SetResourceOwnerAccount(v string) *UnassignIpv6AddressesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassignIpv6AddressesRequest) SetResourceOwnerId(v int64) *UnassignIpv6AddressesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassignIpv6AddressesRequest) Validate() error {
	return dara.Validate(s)
}
