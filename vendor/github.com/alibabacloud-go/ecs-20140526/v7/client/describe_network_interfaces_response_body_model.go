// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) *DescribeNetworkInterfacesResponseBody
	GetNetworkInterfaceSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets
	SetNextToken(v string) *DescribeNetworkInterfacesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeNetworkInterfacesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkInterfacesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNetworkInterfacesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNetworkInterfacesResponseBody
	GetTotalCount() *int32
}

type DescribeNetworkInterfacesResponseBody struct {
	// Details of the ENIs.
	NetworkInterfaceSets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets `json:"NetworkInterfaceSets,omitempty" xml:"NetworkInterfaceSets,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the returned page.
	//
	// > This parameter will be removed in the future. We recommend that you use the NextToken and MaxResults parameters for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// > This parameter will be removed in the future. We recommend that you use the NextToken and MaxResults parameters for a paged query.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of ENIs.
	//
	// > If you specify the `MaxResults` and `NextToken` parameters to perform a paged query, the value of the `TotalCount` response parameter is invalid.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBody) GetNetworkInterfaceSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets {
	return s.NetworkInterfaceSets
}

func (s *DescribeNetworkInterfacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNetworkInterfacesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkInterfacesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkInterfacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNetworkInterfacesResponseBody) SetNetworkInterfaceSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) *DescribeNetworkInterfacesResponseBody {
	s.NetworkInterfaceSets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetNextToken(v string) *DescribeNetworkInterfacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetPageNumber(v int32) *DescribeNetworkInterfacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetPageSize(v int32) *DescribeNetworkInterfacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetRequestId(v string) *DescribeNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) SetTotalCount(v int32) *DescribeNetworkInterfacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBody) Validate() error {
	if s.NetworkInterfaceSets != nil {
		if err := s.NetworkInterfaceSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets struct {
	NetworkInterfaceSet []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet `json:"NetworkInterfaceSet,omitempty" xml:"NetworkInterfaceSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) GetNetworkInterfaceSet() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	return s.NetworkInterfaceSet
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) SetNetworkInterfaceSet(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets {
	s.NetworkInterfaceSet = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSets) Validate() error {
	if s.NetworkInterfaceSet != nil {
		for _, item := range s.NetworkInterfaceSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet struct {
	// The EIPs that are associated with the secondary private IP addresses of the ENI.
	AssociatedPublicIp *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp `json:"AssociatedPublicIp,omitempty" xml:"AssociatedPublicIp,omitempty" type:"Struct"`
	// >  This parameter is in invitational preview and is not publicly available.
	Attachment *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment `json:"Attachment,omitempty" xml:"Attachment,omitempty" type:"Struct"`
	// The time when the security group was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-12-25T12:31:31Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether to retain the ENI when the associated instance is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the ENI.
	//
	// example:
	//
	// DescriptionTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance to which the ENI is attached.
	//
	// >  If the ENI is managed by other Alibaba Cloud services, no instance ID is returned.
	//
	// example:
	//
	// i-bp1e2l6djkndyuli****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IPv4 prefixes of the ENI.
	Ipv4PrefixSets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets `json:"Ipv4PrefixSets,omitempty" xml:"Ipv4PrefixSets,omitempty" type:"Struct"`
	// The IPv6 prefixes of the ENI.
	Ipv6PrefixSets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets `json:"Ipv6PrefixSets,omitempty" xml:"Ipv6PrefixSets,omitempty" type:"Struct"`
	// The IPv6 addresses of the ENI.
	Ipv6Sets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	// The MAC address of the ENI.
	//
	// example:
	//
	// 00:16:3e:12:**:**
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-bp125p95hhdhn3ot****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the ENI.
	//
	// example:
	//
	// my-eni-name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication mode of the ENI. Valid values:
	//
	// 	- Standard: The TCP communication mode is used.
	//
	// 	- HighPerformance: The Elastic RDMA Interface (ERI) is enabled and the remote direct memory access (RDMA) communication mode is used.
	//
	// >  This parameter can have a value of HighPerformance only when the ENI is attached to a c7re RDMA-enhanced instance that resides in Beijing Zone K.
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The ID of the account to which the ENI belongs.
	//
	// example:
	//
	// 123456****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The primary private IP address of the ENI.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// Details about the private IP addresses of the ENI.
	PrivateIpSets *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Struct"`
	// The number of queues supported by the ENI.
	//
	// 	- If the ENI is a secondary ENI in the InUse state and the number of queues supported by the ENI has never been modified, the default number of queues per secondary ENI that the instance type supports is returned.
	//
	// 	- If the ENI is a secondary ENI and the number of queues supported by the ENI has been modified, the new number of queues is returned.
	//
	// 	- If the ENI is a secondary ENI in the Available state and the number of queues supported by the ENI has never been modified, an empty value is returned.
	//
	// 	- If the ENI is a primary ENI, the default number of queues per primary ENI that the instance type supports is returned.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 0
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The ID of the resource group to which the ENI belongs.
	//
	// example:
	//
	// rg-2ze88m67qx5z****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security groups to which the ENI belongs.
	SecurityGroupIds *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
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
	// Indicates whether the source and destination IP address check feature is enabled. To improve network security, enable this feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  Before you use this parameter, read [Source and destination IP address check](https://help.aliyun.com/document_detail/2863210.html).
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
	Tags *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the ENI.
	//
	// example:
	//
	// Secondary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
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

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetAssociatedPublicIp() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp {
	return s.AssociatedPublicIp
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetAttachment() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment {
	return s.Attachment
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetIpv4PrefixSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets {
	return s.Ipv4PrefixSets
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetIpv6PrefixSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets {
	return s.Ipv6PrefixSets
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetIpv6Sets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets {
	return s.Ipv6Sets
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetPrivateIpSets() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets {
	return s.PrivateIpSets
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetSecurityGroupIds() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetServiceID() *int64 {
	return s.ServiceID
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetTags() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags {
	return s.Tags
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetAssociatedPublicIp(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.AssociatedPublicIp = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetAttachment(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Attachment = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetCreationTime(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetDeleteOnRelease(v bool) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.DeleteOnRelease = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetDescription(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Description = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetInstanceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetIpv4PrefixSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Ipv4PrefixSets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetIpv6PrefixSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Ipv6PrefixSets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetIpv6Sets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Ipv6Sets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetMacAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.MacAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkInterfaceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkInterfaceName(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkInterfaceName = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetNetworkInterfaceTrafficMode(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetOwnerId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetPrivateIpAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetPrivateIpSets(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.PrivateIpSets = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetQueueNumber(v int32) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.QueueNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetQueuePairNumber(v int32) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.QueuePairNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetResourceGroupId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetSecurityGroupIds(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetServiceID(v int64) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.ServiceID = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetServiceManaged(v bool) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetSourceDestCheck(v bool) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.SourceDestCheck = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetStatus(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Status = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetTags(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Tags = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetType(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.Type = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetVSwitchId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetVpcId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) SetZoneId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet {
	s.ZoneId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSet) Validate() error {
	if s.AssociatedPublicIp != nil {
		if err := s.AssociatedPublicIp.Validate(); err != nil {
			return err
		}
	}
	if s.Attachment != nil {
		if err := s.Attachment.Validate(); err != nil {
			return err
		}
	}
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

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The EIP.
	//
	// example:
	//
	// ``116.62.**.**``
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) SetAllocationId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp {
	s.AllocationId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) SetPublicIpAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAssociatedPublicIp) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 0
	DeviceIndex *int32 `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The index of the network card.
	//
	// 	- If the ENI is in the Available state or if no network card index was specified when the ENI was attached, this parameter is empty.
	//
	// 	- If the ENI is in the InUse state and a network card index was specified when the ENI was attached, the specified network card index is returned as the value of this parameter.
	//
	// example:
	//
	// 0
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	TrunkNetworkInterfaceId *string `json:"TrunkNetworkInterfaceId,omitempty" xml:"TrunkNetworkInterfaceId,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) GetNetworkCardIndex() *int32 {
	return s.NetworkCardIndex
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) GetTrunkNetworkInterfaceId() *string {
	return s.TrunkNetworkInterfaceId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) SetDeviceIndex(v int32) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment {
	s.DeviceIndex = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) SetInstanceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) SetNetworkCardIndex(v int32) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment {
	s.NetworkCardIndex = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) SetTrunkNetworkInterfaceId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment {
	s.TrunkNetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetAttachment) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets struct {
	Ipv4PrefixSet []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet `json:"Ipv4PrefixSet,omitempty" xml:"Ipv4PrefixSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets) GetIpv4PrefixSet() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet {
	return s.Ipv4PrefixSet
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets) SetIpv4PrefixSet(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets {
	s.Ipv4PrefixSet = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSets) Validate() error {
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

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet struct {
	// The IPv4 prefix of the ENI.
	//
	// example:
	//
	// hide
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet) SetIpv4Prefix(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet {
	s.Ipv4Prefix = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv4PrefixSetsIpv4PrefixSet) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets struct {
	Ipv6PrefixSet []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet `json:"Ipv6PrefixSet,omitempty" xml:"Ipv6PrefixSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets) GetIpv6PrefixSet() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet {
	return s.Ipv6PrefixSet
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets) SetIpv6PrefixSet(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets {
	s.Ipv6PrefixSet = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSets) Validate() error {
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

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet struct {
	// The IPv6 prefix of the ENI.
	//
	// example:
	//
	// hide
	Ipv6Prefix *string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet) GetIpv6Prefix() *string {
	return s.Ipv6Prefix
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet) SetIpv6Prefix(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet {
	s.Ipv6Prefix = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6PrefixSetsIpv6PrefixSet) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets struct {
	Ipv6Set []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set `json:"Ipv6Set,omitempty" xml:"Ipv6Set,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) GetIpv6Set() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set {
	return s.Ipv6Set
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) SetIpv6Set(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets {
	s.Ipv6Set = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6Sets) Validate() error {
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

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set struct {
	// The IPv6 address of the ENI.
	//
	// example:
	//
	// 2408:4321:180:1701:94c7:bc38:3bfa:****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) SetIpv6Address(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetIpv6SetsIpv6Set) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets struct {
	PrivateIpSet []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) GetPrivateIpSet() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	return s.PrivateIpSet
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) SetPrivateIpSet(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets {
	s.PrivateIpSet = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSets) Validate() error {
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

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet struct {
	// The elastic IP address (EIP) that is associated with the private IP address.
	AssociatedPublicIp *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp `json:"AssociatedPublicIp,omitempty" xml:"AssociatedPublicIp,omitempty" type:"Struct"`
	// Indicates whether the private IP address is the primary private IP address. Valid values:
	//
	// 	- true: The IP address is the primary private IP address.
	//
	// 	- false: The IP address is a secondary private IP address.
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// The private domain name of the ECS instance.
	//
	// >  A private domain name can be returned in a specific format only when `HostnameType` is set to `IP` or `InstanceId`.
	//
	// example:
	//
	// DnsTestName
	PrivateDnsName *string `json:"PrivateDnsName,omitempty" xml:"PrivateDnsName,omitempty"`
	// The private IP address of the ENI.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GetAssociatedPublicIp() *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp {
	return s.AssociatedPublicIp
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GetPrivateDnsName() *string {
	return s.PrivateDnsName
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetAssociatedPublicIp(v *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.AssociatedPublicIp = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrimary(v bool) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.Primary = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrivateDnsName(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.PrivateDnsName = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) SetPrivateIpAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSet) Validate() error {
	if s.AssociatedPublicIp != nil {
		if err := s.AssociatedPublicIp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The EIP.
	//
	// example:
	//
	// ``116.62.**.**``
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) SetAllocationId(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp {
	s.AllocationId = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) SetPublicIpAddress(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetPrivateIpSetsPrivateIpSetAssociatedPublicIp) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetSecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags struct {
	Tag []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags) GetTag() []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag {
	return s.Tag
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags) SetTag(v []*DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags {
	s.Tag = v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTags) Validate() error {
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

type DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag struct {
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

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) SetTagKey(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) SetTagValue(v string) *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeNetworkInterfacesResponseBodyNetworkInterfaceSetsNetworkInterfaceSetTagsTag) Validate() error {
	return dara.Validate(s)
}
