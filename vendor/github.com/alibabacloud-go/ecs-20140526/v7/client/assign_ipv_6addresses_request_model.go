// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignIpv6AddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssignIpv6AddressesRequest
	GetClientToken() *string
	SetIpv6Address(v []*string) *AssignIpv6AddressesRequest
	GetIpv6Address() []*string
	SetIpv6AddressCount(v int32) *AssignIpv6AddressesRequest
	GetIpv6AddressCount() *int32
	SetIpv6Prefix(v []*string) *AssignIpv6AddressesRequest
	GetIpv6Prefix() []*string
	SetIpv6PrefixCount(v int32) *AssignIpv6AddressesRequest
	GetIpv6PrefixCount() *int32
	SetNetworkInterfaceId(v string) *AssignIpv6AddressesRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *AssignIpv6AddressesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssignIpv6AddressesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssignIpv6AddressesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssignIpv6AddressesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssignIpv6AddressesRequest
	GetResourceOwnerId() *int64
}

type AssignIpv6AddressesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.***	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IPv6 addresses to assign to the ENI. Valid values of N: 1 to 10.
	//
	// Example: Ipv6Address.1=2001:db8:1234:1a00::\\*\\*\\*\\*
	//
	// >  You must specify `Ipv6Addresses.N` or `Ipv6AddressCount`, but not both.
	//
	// example:
	//
	// 2001:db8:1234:1a00::****
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of IPv6 addresses to randomly generate for the ENI. Valid values: 1 to 10.
	//
	// >  You must specify `Ipv6Addresses.N` or `Ipv6AddressCount`, but not both.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The IPv6 prefixes to assign to the ENI. Valid values of N: 1 to 10.
	//
	// >  To assign IPv6 prefixes to the ENI, you must specify Ipv6Prefix.N or Ipv6PrefixCount, but not both.
	Ipv6Prefix []*string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty" type:"Repeated"`
	// The number of IPv6 prefixes to assign to the ENI. Valid values: 1 to 10.
	//
	// >  To assign IPv6 prefixes to the ENI, you must specify Ipv6Prefix.N or Ipv6PrefixCount, but not both.
	//
	// example:
	//
	// hide
	Ipv6PrefixCount *int32 `json:"Ipv6PrefixCount,omitempty" xml:"Ipv6PrefixCount,omitempty"`
	// The ENI ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp1iqejowblx6h8j****
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

func (s AssignIpv6AddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignIpv6AddressesRequest) GoString() string {
	return s.String()
}

func (s *AssignIpv6AddressesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssignIpv6AddressesRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *AssignIpv6AddressesRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *AssignIpv6AddressesRequest) GetIpv6Prefix() []*string {
	return s.Ipv6Prefix
}

func (s *AssignIpv6AddressesRequest) GetIpv6PrefixCount() *int32 {
	return s.Ipv6PrefixCount
}

func (s *AssignIpv6AddressesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignIpv6AddressesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssignIpv6AddressesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssignIpv6AddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssignIpv6AddressesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssignIpv6AddressesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssignIpv6AddressesRequest) SetClientToken(v string) *AssignIpv6AddressesRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetIpv6Address(v []*string) *AssignIpv6AddressesRequest {
	s.Ipv6Address = v
	return s
}

func (s *AssignIpv6AddressesRequest) SetIpv6AddressCount(v int32) *AssignIpv6AddressesRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetIpv6Prefix(v []*string) *AssignIpv6AddressesRequest {
	s.Ipv6Prefix = v
	return s
}

func (s *AssignIpv6AddressesRequest) SetIpv6PrefixCount(v int32) *AssignIpv6AddressesRequest {
	s.Ipv6PrefixCount = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetNetworkInterfaceId(v string) *AssignIpv6AddressesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetOwnerAccount(v string) *AssignIpv6AddressesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetOwnerId(v int64) *AssignIpv6AddressesRequest {
	s.OwnerId = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetRegionId(v string) *AssignIpv6AddressesRequest {
	s.RegionId = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetResourceOwnerAccount(v string) *AssignIpv6AddressesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssignIpv6AddressesRequest) SetResourceOwnerId(v int64) *AssignIpv6AddressesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssignIpv6AddressesRequest) Validate() error {
	return dara.Validate(s)
}
