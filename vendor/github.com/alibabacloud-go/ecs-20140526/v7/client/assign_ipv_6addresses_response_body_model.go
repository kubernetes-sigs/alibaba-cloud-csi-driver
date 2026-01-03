// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignIpv6AddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6PrefixSets(v *AssignIpv6AddressesResponseBodyIpv6PrefixSets) *AssignIpv6AddressesResponseBody
	GetIpv6PrefixSets() *AssignIpv6AddressesResponseBodyIpv6PrefixSets
	SetIpv6Sets(v *AssignIpv6AddressesResponseBodyIpv6Sets) *AssignIpv6AddressesResponseBody
	GetIpv6Sets() *AssignIpv6AddressesResponseBodyIpv6Sets
	SetNetworkInterfaceId(v string) *AssignIpv6AddressesResponseBody
	GetNetworkInterfaceId() *string
	SetRequestId(v string) *AssignIpv6AddressesResponseBody
	GetRequestId() *string
}

type AssignIpv6AddressesResponseBody struct {
	// The IPv6 prefixes of the ENI.
	Ipv6PrefixSets *AssignIpv6AddressesResponseBodyIpv6PrefixSets `json:"Ipv6PrefixSets,omitempty" xml:"Ipv6PrefixSets,omitempty" type:"Struct"`
	// The IPv6 addresses assigned to the ENI.
	Ipv6Sets *AssignIpv6AddressesResponseBodyIpv6Sets `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	// The ENI ID.
	//
	// example:
	//
	// eni-bp1iqejowblx6h8j****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignIpv6AddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignIpv6AddressesResponseBody) GoString() string {
	return s.String()
}

func (s *AssignIpv6AddressesResponseBody) GetIpv6PrefixSets() *AssignIpv6AddressesResponseBodyIpv6PrefixSets {
	return s.Ipv6PrefixSets
}

func (s *AssignIpv6AddressesResponseBody) GetIpv6Sets() *AssignIpv6AddressesResponseBodyIpv6Sets {
	return s.Ipv6Sets
}

func (s *AssignIpv6AddressesResponseBody) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignIpv6AddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignIpv6AddressesResponseBody) SetIpv6PrefixSets(v *AssignIpv6AddressesResponseBodyIpv6PrefixSets) *AssignIpv6AddressesResponseBody {
	s.Ipv6PrefixSets = v
	return s
}

func (s *AssignIpv6AddressesResponseBody) SetIpv6Sets(v *AssignIpv6AddressesResponseBodyIpv6Sets) *AssignIpv6AddressesResponseBody {
	s.Ipv6Sets = v
	return s
}

func (s *AssignIpv6AddressesResponseBody) SetNetworkInterfaceId(v string) *AssignIpv6AddressesResponseBody {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignIpv6AddressesResponseBody) SetRequestId(v string) *AssignIpv6AddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignIpv6AddressesResponseBody) Validate() error {
	if s.Ipv6PrefixSets != nil {
		if err := s.Ipv6PrefixSets.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv6Sets != nil {
		if err := s.Ipv6Sets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignIpv6AddressesResponseBodyIpv6PrefixSets struct {
	Ipv6Prefix []*string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty" type:"Repeated"`
}

func (s AssignIpv6AddressesResponseBodyIpv6PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s AssignIpv6AddressesResponseBodyIpv6PrefixSets) GoString() string {
	return s.String()
}

func (s *AssignIpv6AddressesResponseBodyIpv6PrefixSets) GetIpv6Prefix() []*string {
	return s.Ipv6Prefix
}

func (s *AssignIpv6AddressesResponseBodyIpv6PrefixSets) SetIpv6Prefix(v []*string) *AssignIpv6AddressesResponseBodyIpv6PrefixSets {
	s.Ipv6Prefix = v
	return s
}

func (s *AssignIpv6AddressesResponseBodyIpv6PrefixSets) Validate() error {
	return dara.Validate(s)
}

type AssignIpv6AddressesResponseBodyIpv6Sets struct {
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
}

func (s AssignIpv6AddressesResponseBodyIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s AssignIpv6AddressesResponseBodyIpv6Sets) GoString() string {
	return s.String()
}

func (s *AssignIpv6AddressesResponseBodyIpv6Sets) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *AssignIpv6AddressesResponseBodyIpv6Sets) SetIpv6Address(v []*string) *AssignIpv6AddressesResponseBodyIpv6Sets {
	s.Ipv6Address = v
	return s
}

func (s *AssignIpv6AddressesResponseBodyIpv6Sets) Validate() error {
	return dara.Validate(s)
}
