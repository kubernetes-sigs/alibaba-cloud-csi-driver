// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkInterfaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionTrackingConfiguration(v *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) *ModifyNetworkInterfaceAttributeRequest
	GetConnectionTrackingConfiguration() *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration
	SetDeleteOnRelease(v bool) *ModifyNetworkInterfaceAttributeRequest
	GetDeleteOnRelease() *bool
	SetDescription(v string) *ModifyNetworkInterfaceAttributeRequest
	GetDescription() *string
	SetEnhancedNetwork(v *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) *ModifyNetworkInterfaceAttributeRequest
	GetEnhancedNetwork() *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork
	SetNetworkInterfaceId(v string) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceId() *string
	SetNetworkInterfaceName(v string) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceName() *string
	SetNetworkInterfaceTrafficConfig(v *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceTrafficConfig() *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig
	SetOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest
	GetOwnerId() *int64
	SetQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequest
	GetQueueNumber() *int32
	SetRegionId(v string) *ModifyNetworkInterfaceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest
	GetResourceOwnerId() *int64
	SetRxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest
	GetRxQueueSize() *int32
	SetSecurityGroupId(v []*string) *ModifyNetworkInterfaceAttributeRequest
	GetSecurityGroupId() []*string
	SetSourceDestCheck(v bool) *ModifyNetworkInterfaceAttributeRequest
	GetSourceDestCheck() *bool
	SetTxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest
	GetTxQueueSize() *int32
}

type ModifyNetworkInterfaceAttributeRequest struct {
	// The connection tracking configuration of the ENI.
	ConnectionTrackingConfiguration *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration `json:"ConnectionTrackingConfiguration,omitempty" xml:"ConnectionTrackingConfiguration,omitempty" type:"Struct"`
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
	// The description of the ENI. The description must be 2 to 255 characters in length and cannot start with [http:// or https://](http://https://。).
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is not publicly available.
	EnhancedNetwork *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork `json:"EnhancedNetwork,omitempty" xml:"EnhancedNetwork,omitempty" type:"Struct"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp67acfmxazb4p****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the ENI. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// eniTestName
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication parameters of the ENI.
	NetworkInterfaceTrafficConfig *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig `json:"NetworkInterfaceTrafficConfig,omitempty" xml:"NetworkInterfaceTrafficConfig,omitempty" type:"Struct"`
	OwnerAccount                  *string                                                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                       *int64                                                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of queues supported by the ENI. Valid values: 1 to 2048.
	//
	// 	- You can change the number of queues supported by an ENI only when the ENI is in the `Available` state or the ENI is attached (`InUse`) to an instance that is in the `Stopped` state.
	//
	// 	- The number of queues supported by the ENI cannot exceed the maximum number of queues that the instance type allows for each ENI. The total number of queues on all ENIs on an instance cannot exceed the queue quota that the instance type supports. To query the maximum number of queues per ENI and the queue quota for an instance type, you can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation and check the `MaximumQueueNumberPerEni` and `TotalEniQueueQuantity` values in the response.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
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
	// The receive (Rx) queue depth of the ENI.
	//
	// Take note of the following items:
	//
	// 	- The Rx queue depth of an ENI must be the same as the transmit (Tx) queue depth of the ENI. Valid values: powers of 2 in the range of 8192 to 16384.
	//
	// 	- A larger Rx queue depth yields higher inbound throughput but consumes more memory.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The IDs of the security groups to which to add the secondary ENI. The secondary ENI is added to the specified security groups and removed from the original security groups.
	//
	// 	- The valid values of N vary based on the maximum number of security groups to which an ENI can be added. For more information, see the [Security group limits](~~25412#SecurityGroupQuota~~) section of the "Limits and quotas" topic.
	//
	// 	- The new security groups take effect after a short delay.
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	// Source and destination IP address check We recommend that you enable the feature to improve network security. Valid value:
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
	// The Tx queue depth of the ENI.
	//
	// Take note of the following items:
	//
	// 	- The Tx queue depth of an ENI must be the same as the Rx queue depth of the ENI. Valid values: powers of 2 in the range of 8192 to 16384.
	//
	// 	- A larger Tx queue depth yields higher outbound throughput but consumes more memory.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetConnectionTrackingConfiguration() *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	return s.ConnectionTrackingConfiguration
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetEnhancedNetwork() *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	return s.EnhancedNetwork
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceTrafficConfig() *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	return s.NetworkInterfaceTrafficConfig
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetConnectionTrackingConfiguration(v *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) *ModifyNetworkInterfaceAttributeRequest {
	s.ConnectionTrackingConfiguration = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetDeleteOnRelease(v bool) *ModifyNetworkInterfaceAttributeRequest {
	s.DeleteOnRelease = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetDescription(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetEnhancedNetwork(v *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) *ModifyNetworkInterfaceAttributeRequest {
	s.EnhancedNetwork = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceId(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceName(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceTrafficConfig(v *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceTrafficConfig = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequest {
	s.QueueNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetRegionId(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetResourceOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetRxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest {
	s.RxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetSecurityGroupId(v []*string) *ModifyNetworkInterfaceAttributeRequest {
	s.SecurityGroupId = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetSourceDestCheck(v bool) *ModifyNetworkInterfaceAttributeRequest {
	s.SourceDestCheck = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetTxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest {
	s.TxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) Validate() error {
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
	return nil
}

type ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration struct {
	// The timeout period for TCP connections in the TIME_WAIT or CLOSE_WAIT state. Unit: seconds. Valid values: integers from 3 to 15.
	//
	// Default value: 3.
	//
	// >  If the associated ECS instance is used together with a Network Load Balancer (NLB) or Classic Load Balancer (CLB) instance, the default timeout period for TCP connections in the `TIME_WAIT` state is 15 seconds.
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
	// >  If the associated ECS instance is used together with an NLB or CLB instance, the default timeout period for UDP flows is 100 seconds.
	//
	// example:
	//
	// 30
	UdpTimeout *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GetTcpClosedAndTimeWaitTimeout() *int32 {
	return s.TcpClosedAndTimeWaitTimeout
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GetTcpEstablishedTimeout() *int32 {
	return s.TcpEstablishedTimeout
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) SetTcpClosedAndTimeWaitTimeout(v int32) *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	s.TcpClosedAndTimeWaitTimeout = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) SetTcpEstablishedTimeout(v int32) *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	s.TcpEstablishedTimeout = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) SetUdpTimeout(v int32) *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	s.UdpTimeout = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) Validate() error {
	return dara.Validate(s)
}

type ModifyNetworkInterfaceAttributeRequestEnhancedNetwork struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// false
	EnableRss *bool `json:"EnableRss,omitempty" xml:"EnableRss,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableSriov                     *bool  `json:"EnableSriov,omitempty" xml:"EnableSriov,omitempty"`
	VirtualFunctionQuantity         *int32 `json:"VirtualFunctionQuantity,omitempty" xml:"VirtualFunctionQuantity,omitempty"`
	VirtualFunctionTotalQueueNumber *int32 `json:"VirtualFunctionTotalQueueNumber,omitempty" xml:"VirtualFunctionTotalQueueNumber,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetEnableRss() *bool {
	return s.EnableRss
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetEnableSriov() *bool {
	return s.EnableSriov
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetVirtualFunctionQuantity() *int32 {
	return s.VirtualFunctionQuantity
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetVirtualFunctionTotalQueueNumber() *int32 {
	return s.VirtualFunctionTotalQueueNumber
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetEnableRss(v bool) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.EnableRss = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetEnableSriov(v bool) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.EnableSriov = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetVirtualFunctionQuantity(v int32) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.VirtualFunctionQuantity = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetVirtualFunctionTotalQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.VirtualFunctionTotalQueueNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) Validate() error {
	return dara.Validate(s)
}

type ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig struct {
	// The communication mode of the ENI. Valid values:
	//
	// 	- Standard: uses the TCP communication mode.
	//
	// 	- HighPerformance: uses the remote direct memory access (RDMA) communication mode with Elastic RDMA Interface (ERI) enabled.
	//
	// When the ENI is in the InUse state, take note of the following items:
	//
	// 	- The total number of ERIs attached to the instance cannot exceed the ERI quota for the instance type. To query the ERI quota for an instance type, call the DescribeInstanceTypes operation and check the EriQuantity value in the response.
	//
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of queues supported by the ENI. When the ENI is in the InUse state, take note of the following items:
	//
	// 	- The value of this parameter cannot exceed the maximum number of queues allowed per ENI for the instance type.
	//
	// 	- The total number of queues for all ENIs on the instance cannot exceed the queue quota for the instance type. To query the maximum number of queues per ENI and the queue quota for an instance type, call the DescribeInstanceTypes operation and check the MaximumQueueNumberPerEnig and TotalEniQueueQuantity values in the response.
	//
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queues supported by the ERI. When the ERI is in the InUse state, take note of the following items:
	//
	// 	- The value of this parameter cannot exceed the maximum number of queues allowed per ERI for the instance type. To query the maximum number of queues allowed per ERI for an instance type, call the DescribeInstanceTypes operation and check the QueuePairNumber value in the response.
	//
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 8
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The receive (Rx) queue depth of the ENI.
	//
	// Take note of the following items:
	//
	// 	- The Rx queue depth of an ENI must be the same as the transmit (Tx) queue depth of the ENI. Valid values: powers of 2 in the range of 8192 to 16384.
	//
	// 	- A larger Rx queue depth yields higher inbound throughput but consumes more memory.
	//
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The Tx queue depth of the ENI.
	//
	// Take note of the following items:
	//
	// 	- The Tx queue depth of an ENI must be the same as the Rx queue depth of the ENI. Valid values: powers of 2 in the range of 8192 to 16384.
	//
	// 	- A larger Tx queue depth yields higher outbound throughput but consumes more memory.
	//
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetNetworkInterfaceTrafficMode(v string) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.QueueNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetQueuePairNumber(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.QueuePairNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetRxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.RxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetTxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.TxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) Validate() error {
	return dara.Validate(s)
}
