// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeNetworkInterfacesRequest
	GetInstanceId() *string
	SetIpv6Address(v []*string) *DescribeNetworkInterfacesRequest
	GetIpv6Address() []*string
	SetMaxResults(v int32) *DescribeNetworkInterfacesRequest
	GetMaxResults() *int32
	SetNetworkInterfaceId(v []*string) *DescribeNetworkInterfacesRequest
	GetNetworkInterfaceId() []*string
	SetNetworkInterfaceName(v string) *DescribeNetworkInterfacesRequest
	GetNetworkInterfaceName() *string
	SetNextToken(v string) *DescribeNetworkInterfacesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeNetworkInterfacesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNetworkInterfacesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeNetworkInterfacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkInterfacesRequest
	GetPageSize() *int32
	SetPrimaryIpAddress(v string) *DescribeNetworkInterfacesRequest
	GetPrimaryIpAddress() *string
	SetPrivateIpAddress(v []*string) *DescribeNetworkInterfacesRequest
	GetPrivateIpAddress() []*string
	SetRegionId(v string) *DescribeNetworkInterfacesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeNetworkInterfacesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeNetworkInterfacesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNetworkInterfacesRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *DescribeNetworkInterfacesRequest
	GetSecurityGroupId() *string
	SetServiceManaged(v bool) *DescribeNetworkInterfacesRequest
	GetServiceManaged() *bool
	SetStatus(v string) *DescribeNetworkInterfacesRequest
	GetStatus() *string
	SetTag(v []*DescribeNetworkInterfacesRequestTag) *DescribeNetworkInterfacesRequest
	GetTag() []*DescribeNetworkInterfacesRequestTag
	SetType(v string) *DescribeNetworkInterfacesRequest
	GetType() *string
	SetVSwitchId(v string) *DescribeNetworkInterfacesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeNetworkInterfacesRequest
	GetVpcId() *string
}

type DescribeNetworkInterfacesRequest struct {
	// The ID of the instance to which the ENI is attached.
	//
	// example:
	//
	// i-bp1e2l6djkndyuli****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// An array that consists of the IPv6 address of the ENI. You can specify multiple IPv6 addresses. Valid values of N: 1 to 100.
	//
	// example:
	//
	// 2408:4321:180:1701:94c7:bc38:3bfa:****
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page. Valid values: 10 to 500.
	//
	// Default values:
	//
	// 	- If this parameter is not specified or if this parameter is set to a value less than 10, the default value is 10.
	//
	// 	- If this parameter is set to a value greater than 500, the default value is 500.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// An array that consists of the IDs of the ENIs. You specify multiple ENI IDs. Valid values of N: 1 to 100.
	//
	// example:
	//
	// eni-bp125p95hhdhn3ot****
	NetworkInterfaceId []*string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" type:"Repeated"`
	// The name of the ENI.
	//
	// example:
	//
	// test-eni-name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The query token. Set the value to the `NextToken` value returned in the last call to this operation.
	//
	// For more information about how to check the responses returned by this operation, see the preceding "Description" section.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Deprecated
	//
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Deprecated
	//
	// The number of entries per page.
	//
	// Valid values: 1 to 1000.
	//
	// Default value: 10.
	//
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The primary private IPv4 address of the ENI.
	//
	// example:
	//
	// ``192.168.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// An array that consists of the secondary private IPv4 addresses of the ENI. You can specify multiple secondary private IPv4 addresses. Valid values of N: 1 to 100.
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the ENI belongs. If this parameter is specified to query resources, up to 1,000 resources that belong to the specified resource group can be returned.
	//
	// > Resources in the default resource group are displayed in the response regardless of how this parameter is set.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group to which the secondary ENI belongs.
	//
	// 	- To query the details of secondary ENIs based on the ID of a security group, specify this parameter.
	//
	// 	- To query the details of primary ENIs based on the ID of a security group, call the [DescribeInstances](https://help.aliyun.com/document_detail/25506.html) operation and specify the `SecurityGroupId` parameter.
	//
	// example:
	//
	// sg-bp144yr32sx6ndw****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Specifies whether the user of the ENI is an Alibaba Cloud service or a distributor.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The state of the ENI. Valid values:
	//
	// 	- Available: The ENI is available.
	//
	// 	- Attaching: The ENI is being attached to an instance.
	//
	// 	- InUse: The ENI is attached to an instance.
	//
	// 	- Detaching: The ENI is being detached from an instance.
	//
	// 	- Deleting: The ENI is being deleted.
	//
	// This parameter is empty by default, which indicates that ENIs in all states are queried.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags to use for query.
	Tag []*DescribeNetworkInterfacesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the ENI. Valid values:
	//
	// 	- Primary
	//
	// 	- Secondary
	//
	// This parameter is empty by default, which indicates that both primary and secondary ENIs are queried.
	//
	// example:
	//
	// Secondary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch with which the ENI is associated.
	//
	// example:
	//
	// vsw-bp16usj2p27htro3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the elastic network interface (ENI) belongs.
	//
	// example:
	//
	// vsw-bp16usj2p27htro3****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNetworkInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkInterfacesRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *DescribeNetworkInterfacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNetworkInterfacesRequest) GetNetworkInterfaceId() []*string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfacesRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *DescribeNetworkInterfacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNetworkInterfacesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNetworkInterfacesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkInterfacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkInterfacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkInterfacesRequest) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *DescribeNetworkInterfacesRequest) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *DescribeNetworkInterfacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkInterfacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNetworkInterfacesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNetworkInterfacesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNetworkInterfacesRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeNetworkInterfacesRequest) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeNetworkInterfacesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkInterfacesRequest) GetTag() []*DescribeNetworkInterfacesRequestTag {
	return s.Tag
}

func (s *DescribeNetworkInterfacesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkInterfacesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNetworkInterfacesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkInterfacesRequest) SetInstanceId(v string) *DescribeNetworkInterfacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetIpv6Address(v []*string) *DescribeNetworkInterfacesRequest {
	s.Ipv6Address = v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetMaxResults(v int32) *DescribeNetworkInterfacesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetNetworkInterfaceId(v []*string) *DescribeNetworkInterfacesRequest {
	s.NetworkInterfaceId = v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetNetworkInterfaceName(v string) *DescribeNetworkInterfacesRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetNextToken(v string) *DescribeNetworkInterfacesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetOwnerAccount(v string) *DescribeNetworkInterfacesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetOwnerId(v int64) *DescribeNetworkInterfacesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageNumber(v int32) *DescribeNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageSize(v int32) *DescribeNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPrimaryIpAddress(v string) *DescribeNetworkInterfacesRequest {
	s.PrimaryIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPrivateIpAddress(v []*string) *DescribeNetworkInterfacesRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetRegionId(v string) *DescribeNetworkInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetResourceGroupId(v string) *DescribeNetworkInterfacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetResourceOwnerAccount(v string) *DescribeNetworkInterfacesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetResourceOwnerId(v int64) *DescribeNetworkInterfacesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetSecurityGroupId(v string) *DescribeNetworkInterfacesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetServiceManaged(v bool) *DescribeNetworkInterfacesRequest {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetStatus(v string) *DescribeNetworkInterfacesRequest {
	s.Status = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetTag(v []*DescribeNetworkInterfacesRequestTag) *DescribeNetworkInterfacesRequest {
	s.Tag = v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetType(v string) *DescribeNetworkInterfacesRequest {
	s.Type = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetVSwitchId(v string) *DescribeNetworkInterfacesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetVpcId(v string) *DescribeNetworkInterfacesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) Validate() error {
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

type DescribeNetworkInterfacesRequestTag struct {
	// The key of tag N of the ENI. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the ENI. Valid values of N: 1 to 20.
	//
	// If a single tag is specified to query ENIs, up to 1,000 ENIs that have this tag can be returned. If multiple tags are specified to query ENIs, up to 1,000 ENIs that have all these tags can be returned. To query more than 1,000 resources that have specified tags, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkInterfacesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkInterfacesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworkInterfacesRequestTag) SetKey(v string) *DescribeNetworkInterfacesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeNetworkInterfacesRequestTag) SetValue(v string) *DescribeNetworkInterfacesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeNetworkInterfacesRequestTag) Validate() error {
	return dara.Validate(s)
}
