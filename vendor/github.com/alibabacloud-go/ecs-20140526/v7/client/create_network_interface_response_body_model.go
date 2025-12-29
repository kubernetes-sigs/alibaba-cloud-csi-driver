// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNetworkInterfaceResponseBody
	GetDescription() *string
	SetIpv4PrefixSets(v *CreateNetworkInterfaceResponseBodyIpv4PrefixSets) *CreateNetworkInterfaceResponseBody
	GetIpv4PrefixSets() *CreateNetworkInterfaceResponseBodyIpv4PrefixSets
	SetIpv6PrefixSets(v *CreateNetworkInterfaceResponseBodyIpv6PrefixSets) *CreateNetworkInterfaceResponseBody
	GetIpv6PrefixSets() *CreateNetworkInterfaceResponseBodyIpv6PrefixSets
	SetIpv6Sets(v *CreateNetworkInterfaceResponseBodyIpv6Sets) *CreateNetworkInterfaceResponseBody
	GetIpv6Sets() *CreateNetworkInterfaceResponseBodyIpv6Sets
	SetMacAddress(v string) *CreateNetworkInterfaceResponseBody
	GetMacAddress() *string
	SetNetworkInterfaceId(v string) *CreateNetworkInterfaceResponseBody
	GetNetworkInterfaceId() *string
	SetNetworkInterfaceName(v string) *CreateNetworkInterfaceResponseBody
	GetNetworkInterfaceName() *string
	SetOwnerId(v string) *CreateNetworkInterfaceResponseBody
	GetOwnerId() *string
	SetPrivateIpAddress(v string) *CreateNetworkInterfaceResponseBody
	GetPrivateIpAddress() *string
	SetPrivateIpSets(v *CreateNetworkInterfaceResponseBodyPrivateIpSets) *CreateNetworkInterfaceResponseBody
	GetPrivateIpSets() *CreateNetworkInterfaceResponseBodyPrivateIpSets
	SetRequestId(v string) *CreateNetworkInterfaceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateNetworkInterfaceResponseBody
	GetResourceGroupId() *string
	SetSecurityGroupIds(v *CreateNetworkInterfaceResponseBodySecurityGroupIds) *CreateNetworkInterfaceResponseBody
	GetSecurityGroupIds() *CreateNetworkInterfaceResponseBodySecurityGroupIds
	SetServiceID(v int64) *CreateNetworkInterfaceResponseBody
	GetServiceID() *int64
	SetServiceManaged(v bool) *CreateNetworkInterfaceResponseBody
	GetServiceManaged() *bool
	SetSourceDestCheck(v bool) *CreateNetworkInterfaceResponseBody
	GetSourceDestCheck() *bool
	SetStatus(v string) *CreateNetworkInterfaceResponseBody
	GetStatus() *string
	SetTags(v *CreateNetworkInterfaceResponseBodyTags) *CreateNetworkInterfaceResponseBody
	GetTags() *CreateNetworkInterfaceResponseBodyTags
	SetType(v string) *CreateNetworkInterfaceResponseBody
	GetType() *string
	SetVSwitchId(v string) *CreateNetworkInterfaceResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *CreateNetworkInterfaceResponseBody
	GetVpcId() *string
	SetZoneId(v string) *CreateNetworkInterfaceResponseBody
	GetZoneId() *string
}

type CreateNetworkInterfaceResponseBody struct {
	// The description of the ENI.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IPv4 prefixes assigned to the ENI.
	Ipv4PrefixSets *CreateNetworkInterfaceResponseBodyIpv4PrefixSets `json:"Ipv4PrefixSets,omitempty" xml:"Ipv4PrefixSets,omitempty" type:"Struct"`
	// The IPv6 prefixes assigned to the ENI.
	Ipv6PrefixSets *CreateNetworkInterfaceResponseBodyIpv6PrefixSets `json:"Ipv6PrefixSets,omitempty" xml:"Ipv6PrefixSets,omitempty" type:"Struct"`
	// The IPv6 addresses assigned to the ENI.
	Ipv6Sets *CreateNetworkInterfaceResponseBodyIpv6Sets `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	// The media access control (MAC) address of the ENI.
	//
	// example:
	//
	// 00:16:3e:12:**:**
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-bp14v2sdd3v8htln****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the ENI.
	//
	// example:
	//
	// my-eni-name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The ID of the account to which the ENI belongs.
	//
	// example:
	//
	// 123456****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private IP address of the ENI.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The private IP addresses.
	PrivateIpSets *CreateNetworkInterfaceResponseBodyPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the ENI belongs.
	//
	// example:
	//
	// rg-2ze88m67qx5z****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the security groups to which the ENI belongs.
	SecurityGroupIds *CreateNetworkInterfaceResponseBodySecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	// The ID of the distributor to which the ENI belongs.
	//
	// example:
	//
	// 12345678910
	ServiceID *int64 `json:"ServiceID,omitempty" xml:"ServiceID,omitempty"`
	// Indicates whether the user of the ENI is an Alibaba Cloud service or a distributor.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The state of the ENI.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the ENI.
	Tags *CreateNetworkInterfaceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the ENI.
	//
	// example:
	//
	// Secondary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch to which the ENI is connected.
	//
	// example:
	//
	// vsw-bp16usj2p27htro3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the ENI belongs.
	//
	// example:
	//
	// vpc-bp1j7w3gc1cexjqd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the ENI.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkInterfaceResponseBody) GetIpv4PrefixSets() *CreateNetworkInterfaceResponseBodyIpv4PrefixSets {
	return s.Ipv4PrefixSets
}

func (s *CreateNetworkInterfaceResponseBody) GetIpv6PrefixSets() *CreateNetworkInterfaceResponseBodyIpv6PrefixSets {
	return s.Ipv6PrefixSets
}

func (s *CreateNetworkInterfaceResponseBody) GetIpv6Sets() *CreateNetworkInterfaceResponseBodyIpv6Sets {
	return s.Ipv6Sets
}

func (s *CreateNetworkInterfaceResponseBody) GetMacAddress() *string {
	return s.MacAddress
}

func (s *CreateNetworkInterfaceResponseBody) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *CreateNetworkInterfaceResponseBody) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *CreateNetworkInterfaceResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateNetworkInterfaceResponseBody) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateNetworkInterfaceResponseBody) GetPrivateIpSets() *CreateNetworkInterfaceResponseBodyPrivateIpSets {
	return s.PrivateIpSets
}

func (s *CreateNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkInterfaceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateNetworkInterfaceResponseBody) GetSecurityGroupIds() *CreateNetworkInterfaceResponseBodySecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *CreateNetworkInterfaceResponseBody) GetServiceID() *int64 {
	return s.ServiceID
}

func (s *CreateNetworkInterfaceResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *CreateNetworkInterfaceResponseBody) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *CreateNetworkInterfaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateNetworkInterfaceResponseBody) GetTags() *CreateNetworkInterfaceResponseBodyTags {
	return s.Tags
}

func (s *CreateNetworkInterfaceResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateNetworkInterfaceResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNetworkInterfaceResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkInterfaceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateNetworkInterfaceResponseBody) SetDescription(v string) *CreateNetworkInterfaceResponseBody {
	s.Description = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetIpv4PrefixSets(v *CreateNetworkInterfaceResponseBodyIpv4PrefixSets) *CreateNetworkInterfaceResponseBody {
	s.Ipv4PrefixSets = v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetIpv6PrefixSets(v *CreateNetworkInterfaceResponseBodyIpv6PrefixSets) *CreateNetworkInterfaceResponseBody {
	s.Ipv6PrefixSets = v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetIpv6Sets(v *CreateNetworkInterfaceResponseBodyIpv6Sets) *CreateNetworkInterfaceResponseBody {
	s.Ipv6Sets = v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetMacAddress(v string) *CreateNetworkInterfaceResponseBody {
	s.MacAddress = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetNetworkInterfaceId(v string) *CreateNetworkInterfaceResponseBody {
	s.NetworkInterfaceId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetNetworkInterfaceName(v string) *CreateNetworkInterfaceResponseBody {
	s.NetworkInterfaceName = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetOwnerId(v string) *CreateNetworkInterfaceResponseBody {
	s.OwnerId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetPrivateIpAddress(v string) *CreateNetworkInterfaceResponseBody {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetPrivateIpSets(v *CreateNetworkInterfaceResponseBodyPrivateIpSets) *CreateNetworkInterfaceResponseBody {
	s.PrivateIpSets = v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetRequestId(v string) *CreateNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetResourceGroupId(v string) *CreateNetworkInterfaceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetSecurityGroupIds(v *CreateNetworkInterfaceResponseBodySecurityGroupIds) *CreateNetworkInterfaceResponseBody {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetServiceID(v int64) *CreateNetworkInterfaceResponseBody {
	s.ServiceID = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetServiceManaged(v bool) *CreateNetworkInterfaceResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetSourceDestCheck(v bool) *CreateNetworkInterfaceResponseBody {
	s.SourceDestCheck = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetStatus(v string) *CreateNetworkInterfaceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetTags(v *CreateNetworkInterfaceResponseBodyTags) *CreateNetworkInterfaceResponseBody {
	s.Tags = v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetType(v string) *CreateNetworkInterfaceResponseBody {
	s.Type = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetVSwitchId(v string) *CreateNetworkInterfaceResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetVpcId(v string) *CreateNetworkInterfaceResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) SetZoneId(v string) *CreateNetworkInterfaceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBody) Validate() error {
	if s.Ipv4PrefixSets != nil {
		if err := s.Ipv4PrefixSets.Validate(); err != nil {
			return err
		}
	}
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
	if s.PrivateIpSets != nil {
		if err := s.PrivateIpSets.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNetworkInterfaceResponseBodyIpv4PrefixSets struct {
	Ipv4PrefixSet []*CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet `json:"Ipv4PrefixSet,omitempty" xml:"Ipv4PrefixSet,omitempty" type:"Repeated"`
}

func (s CreateNetworkInterfaceResponseBodyIpv4PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyIpv4PrefixSets) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyIpv4PrefixSets) GetIpv4PrefixSet() []*CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet {
	return s.Ipv4PrefixSet
}

func (s *CreateNetworkInterfaceResponseBodyIpv4PrefixSets) SetIpv4PrefixSet(v []*CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet) *CreateNetworkInterfaceResponseBodyIpv4PrefixSets {
	s.Ipv4PrefixSet = v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyIpv4PrefixSets) Validate() error {
	if s.Ipv4PrefixSet != nil {
		for _, item := range s.Ipv4PrefixSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet struct {
	// The IPv4 prefix assigned to the ENI.
	//
	// example:
	//
	// hide
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
}

func (s CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet) SetIpv4Prefix(v string) *CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet {
	s.Ipv4Prefix = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyIpv4PrefixSetsIpv4PrefixSet) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceResponseBodyIpv6PrefixSets struct {
	Ipv6PrefixSet []*CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet `json:"Ipv6PrefixSet,omitempty" xml:"Ipv6PrefixSet,omitempty" type:"Repeated"`
}

func (s CreateNetworkInterfaceResponseBodyIpv6PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyIpv6PrefixSets) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyIpv6PrefixSets) GetIpv6PrefixSet() []*CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet {
	return s.Ipv6PrefixSet
}

func (s *CreateNetworkInterfaceResponseBodyIpv6PrefixSets) SetIpv6PrefixSet(v []*CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet) *CreateNetworkInterfaceResponseBodyIpv6PrefixSets {
	s.Ipv6PrefixSet = v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyIpv6PrefixSets) Validate() error {
	if s.Ipv6PrefixSet != nil {
		for _, item := range s.Ipv6PrefixSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet struct {
	// The IPv6 prefix assigned to the ENI.
	//
	// example:
	//
	// hide
	Ipv6Prefix *string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty"`
}

func (s CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet) GetIpv6Prefix() *string {
	return s.Ipv6Prefix
}

func (s *CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet) SetIpv6Prefix(v string) *CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet {
	s.Ipv6Prefix = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyIpv6PrefixSetsIpv6PrefixSet) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceResponseBodyIpv6Sets struct {
	Ipv6Set []*CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set `json:"Ipv6Set,omitempty" xml:"Ipv6Set,omitempty" type:"Repeated"`
}

func (s CreateNetworkInterfaceResponseBodyIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyIpv6Sets) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyIpv6Sets) GetIpv6Set() []*CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set {
	return s.Ipv6Set
}

func (s *CreateNetworkInterfaceResponseBodyIpv6Sets) SetIpv6Set(v []*CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set) *CreateNetworkInterfaceResponseBodyIpv6Sets {
	s.Ipv6Set = v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyIpv6Sets) Validate() error {
	if s.Ipv6Set != nil {
		for _, item := range s.Ipv6Set {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set struct {
	// The IPv6 address assigned to the ENI.
	//
	// example:
	//
	// 2001:db8:1234:1a00::****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set) SetIpv6Address(v string) *CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set {
	s.Ipv6Address = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyIpv6SetsIpv6Set) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceResponseBodyPrivateIpSets struct {
	PrivateIpSet []*CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Repeated"`
}

func (s CreateNetworkInterfaceResponseBodyPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyPrivateIpSets) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSets) GetPrivateIpSet() []*CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet {
	return s.PrivateIpSet
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSets) SetPrivateIpSet(v []*CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) *CreateNetworkInterfaceResponseBodyPrivateIpSets {
	s.PrivateIpSet = v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSets) Validate() error {
	if s.PrivateIpSet != nil {
		for _, item := range s.PrivateIpSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet struct {
	// Indicates whether the private IP address is the primary private IP address.
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) GetPrimary() *bool {
	return s.Primary
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) SetPrimary(v bool) *CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet {
	s.Primary = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) SetPrivateIpAddress(v string) *CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyPrivateIpSetsPrivateIpSet) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceResponseBodySecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s CreateNetworkInterfaceResponseBodySecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodySecurityGroupIds) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodySecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *CreateNetworkInterfaceResponseBodySecurityGroupIds) SetSecurityGroupId(v []*string) *CreateNetworkInterfaceResponseBodySecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *CreateNetworkInterfaceResponseBodySecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceResponseBodyTags struct {
	Tag []*CreateNetworkInterfaceResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateNetworkInterfaceResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyTags) GetTag() []*CreateNetworkInterfaceResponseBodyTagsTag {
	return s.Tag
}

func (s *CreateNetworkInterfaceResponseBodyTags) SetTag(v []*CreateNetworkInterfaceResponseBodyTagsTag) *CreateNetworkInterfaceResponseBodyTags {
	s.Tag = v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyTags) Validate() error {
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

type CreateNetworkInterfaceResponseBodyTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateNetworkInterfaceResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateNetworkInterfaceResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateNetworkInterfaceResponseBodyTagsTag) SetTagKey(v string) *CreateNetworkInterfaceResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyTagsTag) SetTagValue(v string) *CreateNetworkInterfaceResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *CreateNetworkInterfaceResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
