// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssignedPrivateIpAddressesSet(v *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) *AssignPrivateIpAddressesResponseBody
	GetAssignedPrivateIpAddressesSet() *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet
	SetRequestId(v string) *AssignPrivateIpAddressesResponseBody
	GetRequestId() *string
}

type AssignPrivateIpAddressesResponseBody struct {
	// Details about the ENI and the secondary private IP addresses that are assigned to the ENI.
	AssignedPrivateIpAddressesSet *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet `json:"AssignedPrivateIpAddressesSet,omitempty" xml:"AssignedPrivateIpAddressesSet,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignPrivateIpAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesResponseBody) GetAssignedPrivateIpAddressesSet() *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet {
	return s.AssignedPrivateIpAddressesSet
}

func (s *AssignPrivateIpAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignPrivateIpAddressesResponseBody) SetAssignedPrivateIpAddressesSet(v *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) *AssignPrivateIpAddressesResponseBody {
	s.AssignedPrivateIpAddressesSet = v
	return s
}

func (s *AssignPrivateIpAddressesResponseBody) SetRequestId(v string) *AssignPrivateIpAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignPrivateIpAddressesResponseBody) Validate() error {
	if s.AssignedPrivateIpAddressesSet != nil {
		if err := s.AssignedPrivateIpAddressesSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet struct {
	// Details about the assigned IPv4 prefixes.
	Ipv4PrefixSet *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet `json:"Ipv4PrefixSet,omitempty" xml:"Ipv4PrefixSet,omitempty" type:"Struct"`
	// The ENI ID.
	//
	// example:
	//
	// eni-bp125p95hhdhn3ot****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The secondary private IP addresses that are assigned to the ENI.
	PrivateIpSet *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Struct"`
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) GetIpv4PrefixSet() *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet {
	return s.Ipv4PrefixSet
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) GetPrivateIpSet() *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet {
	return s.PrivateIpSet
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) SetIpv4PrefixSet(v *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet) *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet {
	s.Ipv4PrefixSet = v
	return s
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) SetNetworkInterfaceId(v string) *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) SetPrivateIpSet(v *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet) *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet {
	s.PrivateIpSet = v
	return s
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSet) Validate() error {
	if s.Ipv4PrefixSet != nil {
		if err := s.Ipv4PrefixSet.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateIpSet != nil {
		if err := s.PrivateIpSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet struct {
	Ipv4Prefixes []*string `json:"Ipv4Prefixes,omitempty" xml:"Ipv4Prefixes,omitempty" type:"Repeated"`
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet) GetIpv4Prefixes() []*string {
	return s.Ipv4Prefixes
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet) SetIpv4Prefixes(v []*string) *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet {
	s.Ipv4Prefixes = v
	return s
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetIpv4PrefixSet) Validate() error {
	return dara.Validate(s)
}

type AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet struct {
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet) SetPrivateIpAddress(v []*string) *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet {
	s.PrivateIpAddress = v
	return s
}

func (s *AssignPrivateIpAddressesResponseBodyAssignedPrivateIpAddressesSetPrivateIpSet) Validate() error {
	return dara.Validate(s)
}
