// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateNetworkInterfaceRequest
	GetBusinessType() *string
	SetClientToken(v string) *CreateNetworkInterfaceRequest
	GetClientToken() *string
	SetConnectionTrackingConfiguration(v *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) *CreateNetworkInterfaceRequest
	GetConnectionTrackingConfiguration() *CreateNetworkInterfaceRequestConnectionTrackingConfiguration
	SetDeleteOnRelease(v bool) *CreateNetworkInterfaceRequest
	GetDeleteOnRelease() *bool
	SetDescription(v string) *CreateNetworkInterfaceRequest
	GetDescription() *string
	SetEnhancedNetwork(v *CreateNetworkInterfaceRequestEnhancedNetwork) *CreateNetworkInterfaceRequest
	GetEnhancedNetwork() *CreateNetworkInterfaceRequestEnhancedNetwork
	SetInstanceType(v string) *CreateNetworkInterfaceRequest
	GetInstanceType() *string
	SetIpv4Prefix(v []*string) *CreateNetworkInterfaceRequest
	GetIpv4Prefix() []*string
	SetIpv4PrefixCount(v int32) *CreateNetworkInterfaceRequest
	GetIpv4PrefixCount() *int32
	SetIpv6Address(v []*string) *CreateNetworkInterfaceRequest
	GetIpv6Address() []*string
	SetIpv6AddressCount(v int32) *CreateNetworkInterfaceRequest
	GetIpv6AddressCount() *int32
	SetIpv6Prefix(v []*string) *CreateNetworkInterfaceRequest
	GetIpv6Prefix() []*string
	SetIpv6PrefixCount(v int32) *CreateNetworkInterfaceRequest
	GetIpv6PrefixCount() *int32
	SetNetworkInterfaceName(v string) *CreateNetworkInterfaceRequest
	GetNetworkInterfaceName() *string
	SetNetworkInterfaceTrafficConfig(v *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) *CreateNetworkInterfaceRequest
	GetNetworkInterfaceTrafficConfig() *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig
	SetNetworkInterfaceTrafficMode(v string) *CreateNetworkInterfaceRequest
	GetNetworkInterfaceTrafficMode() *string
	SetOwnerAccount(v string) *CreateNetworkInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNetworkInterfaceRequest
	GetOwnerId() *int64
	SetPrimaryIpAddress(v string) *CreateNetworkInterfaceRequest
	GetPrimaryIpAddress() *string
	SetPrivateIpAddress(v []*string) *CreateNetworkInterfaceRequest
	GetPrivateIpAddress() []*string
	SetQueueNumber(v int32) *CreateNetworkInterfaceRequest
	GetQueueNumber() *int32
	SetQueuePairNumber(v int32) *CreateNetworkInterfaceRequest
	GetQueuePairNumber() *int32
	SetRegionId(v string) *CreateNetworkInterfaceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateNetworkInterfaceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateNetworkInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNetworkInterfaceRequest
	GetResourceOwnerId() *int64
	SetRxQueueSize(v int32) *CreateNetworkInterfaceRequest
	GetRxQueueSize() *int32
	SetSecondaryPrivateIpAddressCount(v int32) *CreateNetworkInterfaceRequest
	GetSecondaryPrivateIpAddressCount() *int32
	SetSecurityGroupId(v string) *CreateNetworkInterfaceRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateNetworkInterfaceRequest
	GetSecurityGroupIds() []*string
	SetSourceDestCheck(v bool) *CreateNetworkInterfaceRequest
	GetSourceDestCheck() *bool
	SetTag(v []*CreateNetworkInterfaceRequestTag) *CreateNetworkInterfaceRequest
	GetTag() []*CreateNetworkInterfaceRequestTag
	SetTxQueueSize(v int32) *CreateNetworkInterfaceRequest
	GetTxQueueSize() *int32
	SetVSwitchId(v string) *CreateNetworkInterfaceRequest
	GetVSwitchId() *string
	SetVisible(v bool) *CreateNetworkInterfaceRequest
	GetVisible() *bool
}

type CreateNetworkInterfaceRequest struct {
	// > This parameter is no longer used.
	//
	// example:
	//
	// null
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connection tracking configurations of the ENI.
	//
	// Before you use this parameter, learn about how to manage connection timeout periods. For more information, see [Manage connection timeout periods](https://help.aliyun.com/document_detail/2865958.html).
	ConnectionTrackingConfiguration *CreateNetworkInterfaceRequestConnectionTrackingConfiguration `json:"ConnectionTrackingConfiguration,omitempty" xml:"ConnectionTrackingConfiguration,omitempty" type:"Struct"`
	// Specifies whether to release the ENI when the associated instance is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the ENI. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// >  This parameter is not publicly available.
	EnhancedNetwork *CreateNetworkInterfaceRequestEnhancedNetwork `json:"EnhancedNetwork,omitempty" xml:"EnhancedNetwork,omitempty" type:"Struct"`
	// The type of the ENI. Valid values:
	//
	// 	- Secondary: secondary ENI.
	//
	// 	- Trunk: trunk ENI. This value is in invitational preview.
	//
	// Default value: Secondary.
	//
	// example:
	//
	// null
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// IPv4 prefixes to assign to the ENI. Valid values of N: 1 to 10.
	//
	// >  To assign IPv4 prefixes to the ENI, you must specify the Ipv4Prefix.N or Ipv4PrefixCount parameter, but not both.
	Ipv4Prefix []*string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty" type:"Repeated"`
	// The number of IPv4 prefixes to assign to the ENI. Valid values: 1 to 10.
	//
	// >  To assign IPv4 prefixes to the ENI, you must specify the Ipv4Prefix.N or Ipv4PrefixCount parameter, but not both.
	//
	// example:
	//
	// hide
	Ipv4PrefixCount *int32 `json:"Ipv4PrefixCount,omitempty" xml:"Ipv4PrefixCount,omitempty"`
	// IPv6 addresses to assign to the ENI. Valid values of N: 1 to 10.
	//
	// Example: Ipv6Address.1=2001:db8:1234:1a00::\\*\\*\\*\\*
	//
	// >  To assign IPv6 addresses to the ENI, you must specify the `Ipv6Addresses.N` or `Ipv6AddressCount` parameter, but not both.
	//
	// example:
	//
	// 2001:db8:1234:1a00::****
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of IPv6 addresses to randomly generate for the ENI. Valid values: 1 to 10.
	//
	// >  To assign IPv6 addresses to the ENI, you must specify the `Ipv6Addresses.N` or `Ipv6AddressCount` parameter, but not both.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// IPv6 prefixes to assign to the ENI. Valid values of N: 1 to 10.
	//
	// >  To assign IPv6 prefixes to the ENI, you must specify the Ipv6Prefix.N or Ipv6PrefixCount parameter, but not both.
	Ipv6Prefix []*string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty" type:"Repeated"`
	// The number of IPv6 prefixes to assign to the ENI. Valid values: 1 to 10.
	//
	// >  To assign IPv6 prefixes to the ENI, you must specify the Ipv6Prefix.N or Ipv6PrefixCount parameter, but not both.
	//
	// example:
	//
	// hide
	Ipv6PrefixCount *int32 `json:"Ipv6PrefixCount,omitempty" xml:"Ipv6PrefixCount,omitempty"`
	// The name of the ENI. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// testNetworkInterfaceName
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication settings of the ENI.
	NetworkInterfaceTrafficConfig *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig `json:"NetworkInterfaceTrafficConfig,omitempty" xml:"NetworkInterfaceTrafficConfig,omitempty" type:"Struct"`
	// The communication mode of the ENI. Valid values:
	//
	// 	- Standard: uses the TCP communication mode.
	//
	// 	- HighPerformance: uses the remote direct memory access (RDMA) communication mode with Elastic RDMA Interface (ERI) enabled.
	//
	// >  ENIs in RDMA mode can be attached only to instances of the instance types that support ERIs. The number of ENIs in RDMA mode that are attached to an instance cannot exceed the maximum number of ENIs that the instance type supports. For more information, see [Overview of ECS instance families](https://help.aliyun.com/document_detail/25378.html) and [Configure eRDMA on an enterprise-level instance](https://help.aliyun.com/document_detail/336853.html).
	//
	// Default value: Standard.
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The primary private IP address of the ENI.
	//
	// The specified IP address must be an idle IP address within the CIDR block of the vSwitch. If you do not specify this parameter, a random idle IP address within the vSwitch CIDR block is assigned to the ENI.
	//
	// example:
	//
	// ``172.17.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// Secondary private IP addresses to assign to the ENI. The IP addresses must be idle IP addresses in the CIDR block of the vSwitch with which to associate the ENI. Valid values of N: 0 to 10.
	//
	// >  To assign secondary private IP addresses to the ENI, you can specify the `PrivateIpAddress.N` or `SecondaryPrivateIpAddressCount` parameter, but not both.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	// The number of queues supported by the ENI. Valid values: 1 to 2048.
	//
	// When you attach the ENI to an instance, make sure that the value of this parameter is less than the maximum number of queues per ENI that is allowed for the instance type. To view the maximum number of queues per ENI allowed for an instance type, you can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation and then check the return value of `MaximumQueueNumberPerEni`.
	//
	// This parameter is left empty by default. If you do not specify this parameter, the default number of queues per ENI for the instance type of an instance is used when you attach the ENI to the instance. To view the default number of queues per ENI for an instance type, you can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation and then check the return value of `SecondaryEniQueueNumber`.
	//
	// example:
	//
	// 1
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queue pairs (QPs) supported by the elastic RDMA interface (ERI).
	//
	// If you want to attach multiple ERIs to an instance, we recommend that you specify QueuePairNumber for each ERI based on the value of `QueuePairNumber` supported by the instance type and the number of ERIs that you want to use. Make sure that the total number of QPs of all ERIs does not exceed the maximum number of QPs supported by the instance type. For information about the maximum number of QPs supported by an instance type, see [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html).
	//
	// >  If you do not specify QueuePairNumber for an ERI, the maximum number of QPs supported by the instance type may be used as the number of QPs supported by the ERI. In this case, you cannot attach an additional ERI to the instance. However, you can attach other types of ENIs to the instance.
	//
	// example:
	//
	// 22
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The region in which to create the ENI. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which you want to assign the ENI. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the most recent resource group list.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The receive (Rx) queue depth of the ENI.
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- The Rx queue depth of an ENI must be the same as the Tx queue depth of the ENI. Valid values: powers of 2 in the range of 8192 to 16384.
	//
	// 	- A larger Rx queue depth yields higher inbound throughput but consumes more memory.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The number of private IP addresses to be assigned by ECS. Valid values: 1 to 49.
	//
	// example:
	//
	// 1
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The ID of the security group to which to assign the ENI. The security group and the ENI must belong to the same VPC.
	//
	// > You must specify `SecurityGroupId` or `SecurityGroupIds.N` but not both.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9i****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of security groups to which to assign the ENI. The security groups and the ENI must belong to the same VPC. The valid values of N are determined based on the maximum number of security groups to which an ENI can be assigned. For more information, see [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// >  You must specify `SecurityGroupId` or `SecurityGroupIds.N` but not both.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9i****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable the source and destination IP address check feature. We recommend that you enable the feature to improve network security. Valid value:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  This feature is available only in some regions. Before you use this method, read [Source and destination IP address check](https://help.aliyun.com/document_detail/2863210.html).
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The tags to add to the ENI.
	Tag []*CreateNetworkInterfaceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The transmit (Tx) queue depth of the ENI.
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- The Tx queue depth of an ENI must be the same as the Rx queue depth of the ENI. Valid values: powers of 2 in the range of 8192 to 16384.
	//
	// 	- A larger Tx queue depth yields higher outbound throughput but consumes more memory.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
	// The ID of the vSwitch to which to connect the ENI. Private IP addresses are assigned to the ENI from within the CIDR block of the vSwitch.
	//
	// >  A secondary ENI can be attached to only an instance that is in the same zone as the ENI. The instance and the ENI can be connected to different vSwitches.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws03****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// > This parameter is no longer used.
	//
	// example:
	//
	// null
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s CreateNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateNetworkInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkInterfaceRequest) GetConnectionTrackingConfiguration() *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	return s.ConnectionTrackingConfiguration
}

func (s *CreateNetworkInterfaceRequest) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *CreateNetworkInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkInterfaceRequest) GetEnhancedNetwork() *CreateNetworkInterfaceRequestEnhancedNetwork {
	return s.EnhancedNetwork
}

func (s *CreateNetworkInterfaceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateNetworkInterfaceRequest) GetIpv4Prefix() []*string {
	return s.Ipv4Prefix
}

func (s *CreateNetworkInterfaceRequest) GetIpv4PrefixCount() *int32 {
	return s.Ipv4PrefixCount
}

func (s *CreateNetworkInterfaceRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *CreateNetworkInterfaceRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateNetworkInterfaceRequest) GetIpv6Prefix() []*string {
	return s.Ipv6Prefix
}

func (s *CreateNetworkInterfaceRequest) GetIpv6PrefixCount() *int32 {
	return s.Ipv6PrefixCount
}

func (s *CreateNetworkInterfaceRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *CreateNetworkInterfaceRequest) GetNetworkInterfaceTrafficConfig() *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	return s.NetworkInterfaceTrafficConfig
}

func (s *CreateNetworkInterfaceRequest) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateNetworkInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNetworkInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNetworkInterfaceRequest) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *CreateNetworkInterfaceRequest) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *CreateNetworkInterfaceRequest) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *CreateNetworkInterfaceRequest) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *CreateNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkInterfaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateNetworkInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNetworkInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNetworkInterfaceRequest) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *CreateNetworkInterfaceRequest) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *CreateNetworkInterfaceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateNetworkInterfaceRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateNetworkInterfaceRequest) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *CreateNetworkInterfaceRequest) GetTag() []*CreateNetworkInterfaceRequestTag {
	return s.Tag
}

func (s *CreateNetworkInterfaceRequest) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *CreateNetworkInterfaceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNetworkInterfaceRequest) GetVisible() *bool {
	return s.Visible
}

func (s *CreateNetworkInterfaceRequest) SetBusinessType(v string) *CreateNetworkInterfaceRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetClientToken(v string) *CreateNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetConnectionTrackingConfiguration(v *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) *CreateNetworkInterfaceRequest {
	s.ConnectionTrackingConfiguration = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetDeleteOnRelease(v bool) *CreateNetworkInterfaceRequest {
	s.DeleteOnRelease = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetDescription(v string) *CreateNetworkInterfaceRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetEnhancedNetwork(v *CreateNetworkInterfaceRequestEnhancedNetwork) *CreateNetworkInterfaceRequest {
	s.EnhancedNetwork = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetInstanceType(v string) *CreateNetworkInterfaceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv4Prefix(v []*string) *CreateNetworkInterfaceRequest {
	s.Ipv4Prefix = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv4PrefixCount(v int32) *CreateNetworkInterfaceRequest {
	s.Ipv4PrefixCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6Address(v []*string) *CreateNetworkInterfaceRequest {
	s.Ipv6Address = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6AddressCount(v int32) *CreateNetworkInterfaceRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6Prefix(v []*string) *CreateNetworkInterfaceRequest {
	s.Ipv6Prefix = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetIpv6PrefixCount(v int32) *CreateNetworkInterfaceRequest {
	s.Ipv6PrefixCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetNetworkInterfaceName(v string) *CreateNetworkInterfaceRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetNetworkInterfaceTrafficConfig(v *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) *CreateNetworkInterfaceRequest {
	s.NetworkInterfaceTrafficConfig = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetNetworkInterfaceTrafficMode(v string) *CreateNetworkInterfaceRequest {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetOwnerAccount(v string) *CreateNetworkInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetOwnerId(v int64) *CreateNetworkInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetPrimaryIpAddress(v string) *CreateNetworkInterfaceRequest {
	s.PrimaryIpAddress = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetPrivateIpAddress(v []*string) *CreateNetworkInterfaceRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetQueueNumber(v int32) *CreateNetworkInterfaceRequest {
	s.QueueNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetQueuePairNumber(v int32) *CreateNetworkInterfaceRequest {
	s.QueuePairNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetRegionId(v string) *CreateNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetResourceGroupId(v string) *CreateNetworkInterfaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetResourceOwnerAccount(v string) *CreateNetworkInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetResourceOwnerId(v int64) *CreateNetworkInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetRxQueueSize(v int32) *CreateNetworkInterfaceRequest {
	s.RxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSecondaryPrivateIpAddressCount(v int32) *CreateNetworkInterfaceRequest {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSecurityGroupId(v string) *CreateNetworkInterfaceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSecurityGroupIds(v []*string) *CreateNetworkInterfaceRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSourceDestCheck(v bool) *CreateNetworkInterfaceRequest {
	s.SourceDestCheck = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetTag(v []*CreateNetworkInterfaceRequestTag) *CreateNetworkInterfaceRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetTxQueueSize(v int32) *CreateNetworkInterfaceRequest {
	s.TxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetVSwitchId(v string) *CreateNetworkInterfaceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetVisible(v bool) *CreateNetworkInterfaceRequest {
	s.Visible = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) Validate() error {
	if s.ConnectionTrackingConfiguration != nil {
		if err := s.ConnectionTrackingConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.EnhancedNetwork != nil {
		if err := s.EnhancedNetwork.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaceTrafficConfig != nil {
		if err := s.NetworkInterfaceTrafficConfig.Validate(); err != nil {
			return err
		}
	}
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

type CreateNetworkInterfaceRequestConnectionTrackingConfiguration struct {
	// The timeout period for TCP connections in the TIME_WAIT or CLOSE_WAIT state. Unit: seconds. Valid values: integers from 3 to 15.
	//
	// Default value: 3.
	//
	// >  If the associated Elastic Compute Service (ECS) instance is used with a Network Load Balancer (NLB) or Classic Load Balancer (CLB) instance, the default timeout period for TCP connections in the `TIME_WAIT` state is 15 seconds.
	//
	// example:
	//
	// 3
	TcpClosedAndTimeWaitTimeout *int32 `json:"TcpClosedAndTimeWaitTimeout,omitempty" xml:"TcpClosedAndTimeWaitTimeout,omitempty"`
	// The timeout period for TCP connections in the ESTABLISHED state. Unit: seconds. Valid values: 30, 60, 80, 100, 200, 300, 500, 700, and 910.
	//
	// Default value: 910.
	//
	// example:
	//
	// 910
	TcpEstablishedTimeout *int32 `json:"TcpEstablishedTimeout,omitempty" xml:"TcpEstablishedTimeout,omitempty"`
	// The timeout period for UDP flows. Unit: seconds. Valid values: 10, 20, 30, 60, 80, and 100.
	//
	// Default value: 30.
	//
	// >  If the associated ECS instance is used with an NLB or CLB instance, the default timeout period for UDP flows is 100 seconds.
	//
	// example:
	//
	// 30
	UdpTimeout *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s CreateNetworkInterfaceRequestConnectionTrackingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GetTcpClosedAndTimeWaitTimeout() *int32 {
	return s.TcpClosedAndTimeWaitTimeout
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GetTcpEstablishedTimeout() *int32 {
	return s.TcpEstablishedTimeout
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) SetTcpClosedAndTimeWaitTimeout(v int32) *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	s.TcpClosedAndTimeWaitTimeout = &v
	return s
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) SetTcpEstablishedTimeout(v int32) *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	s.TcpEstablishedTimeout = &v
	return s
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) SetUdpTimeout(v int32) *CreateNetworkInterfaceRequestConnectionTrackingConfiguration {
	s.UdpTimeout = &v
	return s
}

func (s *CreateNetworkInterfaceRequestConnectionTrackingConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceRequestEnhancedNetwork struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableRss *bool `json:"EnableRss,omitempty" xml:"EnableRss,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableSriov                     *bool  `json:"EnableSriov,omitempty" xml:"EnableSriov,omitempty"`
	VirtualFunctionQuantity         *int32 `json:"VirtualFunctionQuantity,omitempty" xml:"VirtualFunctionQuantity,omitempty"`
	VirtualFunctionTotalQueueNumber *int32 `json:"VirtualFunctionTotalQueueNumber,omitempty" xml:"VirtualFunctionTotalQueueNumber,omitempty"`
}

func (s CreateNetworkInterfaceRequestEnhancedNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestEnhancedNetwork) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetEnableRss() *bool {
	return s.EnableRss
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetEnableSriov() *bool {
	return s.EnableSriov
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetVirtualFunctionQuantity() *int32 {
	return s.VirtualFunctionQuantity
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) GetVirtualFunctionTotalQueueNumber() *int32 {
	return s.VirtualFunctionTotalQueueNumber
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetEnableRss(v bool) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.EnableRss = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetEnableSriov(v bool) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.EnableSriov = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetVirtualFunctionQuantity(v int32) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.VirtualFunctionQuantity = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) SetVirtualFunctionTotalQueueNumber(v int32) *CreateNetworkInterfaceRequestEnhancedNetwork {
	s.VirtualFunctionTotalQueueNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequestEnhancedNetwork) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig struct {
	// The communication mode of the ENI.
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of queues supported by the ENI.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of QPs supported by the ERI.
	//
	// example:
	//
	// 8
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The Rx queue depth of the ENI.
	//
	// >  This parameter is in invitational preview and is not publicly available. To use this parameter, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- This parameter is applicable only to 7th-generation or later ECS instance types.
	//
	// 	- This parameter is applicable to Linux images.
	//
	// 	- A larger Rx queue depth yields higher inbound throughput and reduces packet loss rates but consumes more memory.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The Tx queue depth of the ENI.
	//
	// >  This parameter is in invitational preview and is not publicly available. To use this parameter, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- This parameter is applicable only to 7th-generation or later ECS instance types.
	//
	// 	- This parameter is applicable to Linux images.
	//
	// 	- A larger Tx queue depth yields higher outbound throughput and reduces packet loss rates but consumes more memory.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
}

func (s CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetNetworkInterfaceTrafficMode(v string) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetQueueNumber(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.QueueNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetQueuePairNumber(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.QueuePairNumber = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetRxQueueSize(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.RxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) SetTxQueueSize(v int32) *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig {
	s.TxQueueSize = &v
	return s
}

func (s *CreateNetworkInterfaceRequestNetworkInterfaceTrafficConfig) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkInterfaceRequestTag struct {
	// The key of tag N to add to the ENI. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the ENI. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkInterfaceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNetworkInterfaceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNetworkInterfaceRequestTag) SetKey(v string) *CreateNetworkInterfaceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkInterfaceRequestTag) SetValue(v string) *CreateNetworkInterfaceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNetworkInterfaceRequestTag) Validate() error {
	return dara.Validate(s)
}
