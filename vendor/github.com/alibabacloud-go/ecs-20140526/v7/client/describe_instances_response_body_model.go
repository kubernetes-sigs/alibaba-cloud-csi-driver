// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() *DescribeInstancesResponseBodyInstances
	SetNextToken(v string) *DescribeInstancesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesResponseBody struct {
	// Details about the queried instances.
	Instances *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances queried.
	//
	// >  If you specify the `MaxResults` and `NextToken` request parameters to perform a paged query, the value of the `TotalCount` response parameter is invalid.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() *DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetNextToken(v string) *DescribeInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	Instance []*DescribeInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetInstance() []*DescribeInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeInstancesResponseBodyInstances) SetInstance(v []*DescribeInstancesResponseBodyInstancesInstance) *DescribeInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstance struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// if can be null:
	// true
	AdditionalInfo *DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo `json:"AdditionalInfo,omitempty" xml:"AdditionalInfo,omitempty" type:"Struct"`
	// The automatic release time of the pay-as-you-go instance.
	//
	// example:
	//
	// 2017-12-10T04:04Z
	AutoReleaseTime *string                                                     `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	ClockOptions    *DescribeInstancesResponseBodyInstancesInstanceClockOptions `json:"ClockOptions,omitempty" xml:"ClockOptions,omitempty" type:"Struct"`
	// The ID of the cluster to which the instance belongs.
	//
	// >  This parameter will be removed in the future. We recommend that you use other parameters to ensure future compatibility.
	//
	// example:
	//
	// c-bp67acfmxazb4p****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 8
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Details about the CPU options.
	CpuOptions *DescribeInstancesResponseBodyInstancesInstanceCpuOptions `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2017-12-10T04:04Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// 	- Standard: the standard mode. For more information, see the "Standard mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// 	- Unlimited: the unlimited mode. For more information, see the "Unlimited mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The information about the dedicated host. The value is an array that consists of DedicatedHostClusterId, DedicatedHostId, and DedicatedHostName.
	DedicatedHostAttribute *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute `json:"DedicatedHostAttribute,omitempty" xml:"DedicatedHostAttribute,omitempty" type:"Struct"`
	// The attributes of the instance on the dedicated host.
	DedicatedInstanceAttribute *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute `json:"DedicatedInstanceAttribute,omitempty" xml:"DedicatedInstanceAttribute,omitempty" type:"Struct"`
	// Indicates whether release protection is enabled for the instance. This parameter determines whether you can use the ECS console or call the DeleteInstance operation to release the instance. Valid values:
	//
	// 	- true: Release protection is enabled for the instance.
	//
	// 	- false: Release protection is disabled for the instance.
	//
	// >  This parameter is applicable only to pay-as-you-go instances. The release protection feature can protect instances against manual releases, but not against automatic releases.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The number of the deployment set group to which the instance belongs in a deployment set.
	//
	// example:
	//
	// 1
	DeploymentSetGroupNo *int32 `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the deployment set to which the instance belongs.
	//
	// example:
	//
	// ds-bp67acfmxazb4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether data disks can be attached to the instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DeviceAvailable *bool `json:"DeviceAvailable,omitempty" xml:"DeviceAvailable,omitempty"`
	// Details about the capacity reservation associated with the instance.
	EcsCapacityReservationAttr *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr `json:"EcsCapacityReservationAttr,omitempty" xml:"EcsCapacityReservationAttr,omitempty" type:"Struct"`
	// Details about the EIP associated with the instance.
	EipAddress *DescribeInstancesResponseBodyInstancesInstanceEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Struct"`
	EnableNVS  *bool                                                     `json:"EnableNVS,omitempty" xml:"EnableNVS,omitempty"`
	// The expiration time of the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2017-12-10T04:04Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The number of GPUs for the instance type.
	//
	// example:
	//
	// 4
	GPUAmount *int32 `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	// The category of GPUs for the instance type.
	//
	// example:
	//
	// NVIDIA V100
	GPUSpec *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	HibernationOptions *DescribeInstancesResponseBodyInstancesInstanceHibernationOptions `json:"HibernationOptions,omitempty" xml:"HibernationOptions,omitempty" type:"Struct"`
	// The hostname of the instance.
	//
	// example:
	//
	// testHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the HPC cluster to which the instance belongs.
	//
	// example:
	//
	// hpc-bp67acfmxazb4p****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The image ID of the instance.
	//
	// example:
	//
	// m-bp67acfmxazb4p****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image options.
	ImageOptions *DescribeInstancesResponseBodyInstancesInstanceImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The internal IP addresses of the instance located in the classic network.
	InnerIpAddress *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// InstanceNameTest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- classic
	//
	// 	- vpc
	//
	// example:
	//
	// vpc
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance family of the instance.
	//
	// example:
	//
	// ecs.g5
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic: pay-by-traffic
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Indicates whether the instance is an I/O optimized instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IoOptimized *bool `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The number of local disks attached to the instance.
	//
	// example:
	//
	// 2
	LocalStorageAmount *int32 `json:"LocalStorageAmount,omitempty" xml:"LocalStorageAmount,omitempty"`
	// The capacity of local disks attached to the instance. Unit: GiB.
	//
	// example:
	//
	// 1000
	LocalStorageCapacity *int64 `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty"`
	// The memory size. Unit: MiB.
	//
	// example:
	//
	// 16384
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Details about the metadata options.
	MetadataOptions *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions `json:"MetadataOptions,omitempty" xml:"MetadataOptions,omitempty" type:"Struct"`
	// The ENIs attached to the instance.
	NetworkInterfaces *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Struct"`
	// The name of the operating system of the instance.
	//
	// example:
	//
	// CentOS 7.4 64-bit
	OSName *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	// The English name of the operating system of the instance.
	//
	// example:
	//
	// CentOS  7.4 64 bit
	OSNameEn *string `json:"OSNameEn,omitempty" xml:"OSNameEn,omitempty"`
	// The type of the operating system of the instance. Valid values:
	//
	// 	- windows: Windows operating systems
	//
	// 	- linux: Linux operating systems
	//
	// example:
	//
	// linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The reasons why the instance was locked.
	OperationLocks *DescribeInstancesResponseBodyInstancesInstanceOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The private domain name options of the instance.
	//
	// For information about the resolution of ECS private domain names, see [ECS private DNS resolution](https://help.aliyun.com/document_detail/2844797.html).
	//
	// >  This parameter is returned only when the `AdditionalAttributes` parameter contains `PRIVATE_DNS_OPTIONS` in the request.
	PrivateDnsNameOptions *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions `json:"PrivateDnsNameOptions,omitempty" xml:"PrivateDnsNameOptions,omitempty" type:"Struct"`
	// The public IP addresses of the instance.
	PublicIpAddress *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Struct"`
	// The RDMA IP addresses of the instance in the HPC cluster.
	RdmaIpAddress *DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress `json:"RdmaIpAddress,omitempty" xml:"RdmaIpAddress,omitempty" type:"Struct"`
	// Indicates whether the instance can be recycled.
	//
	// example:
	//
	// false
	Recyclable *bool `json:"Recyclable,omitempty" xml:"Recyclable,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// >  The parameter is removed.
	//
	// example:
	//
	// month
	SaleCycle *string `json:"SaleCycle,omitempty" xml:"SaleCycle,omitempty"`
	// The IDs of the security groups to which the instance belongs.
	SecurityGroupIds *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	// The serial number of the instance.
	//
	// example:
	//
	// 51d1353b-22bf-4567-a176-8b3e12e4****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The protection period of the spot instance. Unit: hours. Valid values:
	//
	// 	- 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends an ECS system event to notify you 5 minutes before the instance is released. Spot instances are billed by second. We recommend that you specify a protection period based on your business requirements.
	//
	// >  This parameter is returned when SpotStrategy is set to SpotWithPriceLimit or SpotAsPriceGo.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the spot instance when the system initiates a spot instance interruption operation. Valid values:
	//
	// 	- Terminate: releases the spot instance.
	//
	// 	- Stop: stops the instance in economical mode.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the instance. The value can be accurate to three decimal places. This parameter is valid when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.98
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the pay-as-you-go instance. Valid values:
	//
	// 	- NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a spot instance with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is a spot instance for which the market price is automatically used as the bid price. The market price can be up to the pay-as-you-go price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The time when the instance was last started. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2017-12-10T04:04Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the instance continues to be billed after it is stopped. Valid values:
	//
	// 	- KeepCharging: The instance is stopped in standard mode. Billing for the instance continues after the instance is stopped, and resources are retained for the instance.
	//
	// 	- StopCharging: The instance is stopped in economical mode. Billing for some resources of the instance stops after the instance is stopped. When the instance is stopped, its resources such as vCPUs, memory, and public IP addresses are released. The instance may be unable to restart if some required resources are out of stock in the current region.
	//
	// 	- Not-applicable: The instance does not support economical mode.
	//
	// example:
	//
	// KeepCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
	// The tags of the instance.
	Tags *DescribeInstancesResponseBodyInstancesInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The virtual LAN (VLAN) ID of the instance.
	//
	// >  This parameter will be removed in the future. We recommend that you use other parameters to ensure future compatibility.
	//
	// example:
	//
	// 10
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The VPC attributes of the instance.
	VpcAttributes *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes `json:"VpcAttributes,omitempty" xml:"VpcAttributes,omitempty" type:"Struct"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetAdditionalInfo() *DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo {
	return s.AdditionalInfo
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetClockOptions() *DescribeInstancesResponseBodyInstancesInstanceClockOptions {
	return s.ClockOptions
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCpuOptions() *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	return s.CpuOptions
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDedicatedHostAttribute() *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute {
	return s.DedicatedHostAttribute
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDedicatedInstanceAttribute() *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute {
	return s.DedicatedInstanceAttribute
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDeviceAvailable() *bool {
	return s.DeviceAvailable
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetEcsCapacityReservationAttr() *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr {
	return s.EcsCapacityReservationAttr
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetEipAddress() *DescribeInstancesResponseBodyInstancesInstanceEipAddress {
	return s.EipAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetEnableNVS() *bool {
	return s.EnableNVS
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetGPUAmount() *int32 {
	return s.GPUAmount
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetGPUSpec() *string {
	return s.GPUSpec
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetHibernationOptions() *DescribeInstancesResponseBodyInstancesInstanceHibernationOptions {
	return s.HibernationOptions
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetHostName() *string {
	return s.HostName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetISP() *string {
	return s.ISP
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetImageOptions() *DescribeInstancesResponseBodyInstancesInstanceImageOptions {
	return s.ImageOptions
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInnerIpAddress() *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress {
	return s.InnerIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetIoOptimized() *bool {
	return s.IoOptimized
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetLocalStorageAmount() *int32 {
	return s.LocalStorageAmount
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetLocalStorageCapacity() *int64 {
	return s.LocalStorageCapacity
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMetadataOptions() *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions {
	return s.MetadataOptions
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetNetworkInterfaces() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetOSName() *string {
	return s.OSName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetOSNameEn() *string {
	return s.OSNameEn
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetOSType() *string {
	return s.OSType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetOperationLocks() *DescribeInstancesResponseBodyInstancesInstanceOperationLocks {
	return s.OperationLocks
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetPrivateDnsNameOptions() *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions {
	return s.PrivateDnsNameOptions
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetPublicIpAddress() *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress {
	return s.PublicIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetRdmaIpAddress() *DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress {
	return s.RdmaIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetRecyclable() *bool {
	return s.Recyclable
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSaleCycle() *string {
	return s.SaleCycle
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSecurityGroupIds() *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetTags() *DescribeInstancesResponseBodyInstancesInstanceTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetVlanId() *string {
	return s.VlanId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetVpcAttributes() *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes {
	return s.VpcAttributes
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetAdditionalInfo(v *DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo) *DescribeInstancesResponseBodyInstancesInstance {
	s.AdditionalInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetAutoReleaseTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.AutoReleaseTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClockOptions(v *DescribeInstancesResponseBodyInstancesInstanceClockOptions) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClockOptions = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCpuOptions(v *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) *DescribeInstancesResponseBodyInstancesInstance {
	s.CpuOptions = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreationTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreditSpecification(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreditSpecification = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDedicatedHostAttribute(v *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) *DescribeInstancesResponseBodyInstancesInstance {
	s.DedicatedHostAttribute = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDedicatedInstanceAttribute(v *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) *DescribeInstancesResponseBodyInstancesInstance {
	s.DedicatedInstanceAttribute = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDeletionProtection(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDeploymentSetGroupNo(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDeploymentSetId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDescription(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDeviceAvailable(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.DeviceAvailable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEcsCapacityReservationAttr(v *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) *DescribeInstancesResponseBodyInstancesInstance {
	s.EcsCapacityReservationAttr = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEipAddress(v *DescribeInstancesResponseBodyInstancesInstanceEipAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.EipAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEnableNVS(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.EnableNVS = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpiredTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetGPUAmount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.GPUAmount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetGPUSpec(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.GPUSpec = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetHibernationOptions(v *DescribeInstancesResponseBodyInstancesInstanceHibernationOptions) *DescribeInstancesResponseBodyInstancesInstance {
	s.HibernationOptions = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetHostName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.HostName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetHpcClusterId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.HpcClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetISP(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ISP = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetImageId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetImageOptions(v *DescribeInstancesResponseBodyInstancesInstanceImageOptions) *DescribeInstancesResponseBodyInstancesInstance {
	s.ImageOptions = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInnerIpAddress(v *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.InnerIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceChargeType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceNetworkType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceTypeFamily(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInternetChargeType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInternetMaxBandwidthIn(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInternetMaxBandwidthOut(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetIoOptimized(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.IoOptimized = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetKeyPairName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.KeyPairName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetLocalStorageAmount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.LocalStorageAmount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetLocalStorageCapacity(v int64) *DescribeInstancesResponseBodyInstancesInstance {
	s.LocalStorageCapacity = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMemory(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Memory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMetadataOptions(v *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) *DescribeInstancesResponseBodyInstancesInstance {
	s.MetadataOptions = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetNetworkInterfaces(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) *DescribeInstancesResponseBodyInstancesInstance {
	s.NetworkInterfaces = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetOSName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.OSName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetOSNameEn(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.OSNameEn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetOSType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.OSType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetOperationLocks(v *DescribeInstancesResponseBodyInstancesInstanceOperationLocks) *DescribeInstancesResponseBodyInstancesInstance {
	s.OperationLocks = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPrivateDnsNameOptions(v *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) *DescribeInstancesResponseBodyInstancesInstance {
	s.PrivateDnsNameOptions = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPublicIpAddress(v *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.PublicIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetRdmaIpAddress(v *DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress) *DescribeInstancesResponseBodyInstancesInstance {
	s.RdmaIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetRecyclable(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.Recyclable = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSaleCycle(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.SaleCycle = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSecurityGroupIds(v *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) *DescribeInstancesResponseBodyInstancesInstance {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSerialNumber(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.SerialNumber = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSpotDuration(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.SpotDuration = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSpotInterruptionBehavior(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSpotPriceLimit(v float32) *DescribeInstancesResponseBodyInstancesInstance {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetSpotStrategy(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStartTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStoppedMode(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.StoppedMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetTags(v *DescribeInstancesResponseBodyInstancesInstanceTags) *DescribeInstancesResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVlanId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.VlanId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVpcAttributes(v *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) *DescribeInstancesResponseBodyInstancesInstance {
	s.VpcAttributes = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) Validate() error {
	if s.AdditionalInfo != nil {
		if err := s.AdditionalInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ClockOptions != nil {
		if err := s.ClockOptions.Validate(); err != nil {
			return err
		}
	}
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
	if s.DedicatedHostAttribute != nil {
		if err := s.DedicatedHostAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.DedicatedInstanceAttribute != nil {
		if err := s.DedicatedInstanceAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.EcsCapacityReservationAttr != nil {
		if err := s.EcsCapacityReservationAttr.Validate(); err != nil {
			return err
		}
	}
	if s.EipAddress != nil {
		if err := s.EipAddress.Validate(); err != nil {
			return err
		}
	}
	if s.HibernationOptions != nil {
		if err := s.HibernationOptions.Validate(); err != nil {
			return err
		}
	}
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.InnerIpAddress != nil {
		if err := s.InnerIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.MetadataOptions != nil {
		if err := s.MetadataOptions.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaces != nil {
		if err := s.NetworkInterfaces.Validate(); err != nil {
			return err
		}
	}
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateDnsNameOptions != nil {
		if err := s.PrivateDnsNameOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PublicIpAddress != nil {
		if err := s.PublicIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.RdmaIpAddress != nil {
		if err := s.RdmaIpAddress.Validate(); err != nil {
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
	if s.VpcAttributes != nil {
		if err := s.VpcAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// true
	EnableHighDensityMode *bool `json:"EnableHighDensityMode,omitempty" xml:"EnableHighDensityMode,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo) GetEnableHighDensityMode() *bool {
	return s.EnableHighDensityMode
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo) SetEnableHighDensityMode(v bool) *DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo {
	s.EnableHighDensityMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceAdditionalInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceClockOptions struct {
	PtpStatus *string `json:"PtpStatus,omitempty" xml:"PtpStatus,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceClockOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceClockOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceClockOptions) GetPtpStatus() *string {
	return s.PtpStatus
}

func (s *DescribeInstancesResponseBodyInstancesInstanceClockOptions) SetPtpStatus(v string) *DescribeInstancesResponseBodyInstancesInstanceClockOptions {
	s.PtpStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceClockOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceCpuOptions struct {
	// The number of physical CPU cores.
	//
	// example:
	//
	// 2
	CoreCount   *int32 `json:"CoreCount,omitempty" xml:"CoreCount,omitempty"`
	EnableVISST *bool  `json:"EnableVISST,omitempty" xml:"EnableVISST,omitempty"`
	EnableVRDT  *bool  `json:"EnableVRDT,omitempty" xml:"EnableVRDT,omitempty"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// 2
	Numa *string `json:"Numa,omitempty" xml:"Numa,omitempty"`
	// The number of threads per CPU core.
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
	// The CPU topology type of the instance. Valid values:
	//
	// 	- ContinuousCoreToHTMapping: Hyper-Threading (HT) continuous CPU topology
	//
	// 	- DiscreteCoreToHTMapping: HT discrete CPU topology
	//
	// example:
	//
	// ContinuousCoreToHTMapping
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	TurboMode    *string `json:"TurboMode,omitempty" xml:"TurboMode,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GetCoreCount() *int32 {
	return s.CoreCount
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GetEnableVISST() *bool {
	return s.EnableVISST
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GetEnableVRDT() *bool {
	return s.EnableVRDT
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GetNuma() *string {
	return s.Numa
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GetTopologyType() *string {
	return s.TopologyType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) GetTurboMode() *string {
	return s.TurboMode
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) SetCoreCount(v int32) *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	s.CoreCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) SetEnableVISST(v bool) *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	s.EnableVISST = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) SetEnableVRDT(v bool) *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	s.EnableVRDT = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) SetNuma(v string) *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	s.Numa = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) SetThreadsPerCore(v int32) *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) SetTopologyType(v string) *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	s.TopologyType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) SetTurboMode(v string) *DescribeInstancesResponseBodyInstancesInstanceCpuOptions {
	s.TurboMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceCpuOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute struct {
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-bp67acfmxazb4h****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The name of the dedicated host.
	//
	// example:
	//
	// testDedicatedHostName
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) SetDedicatedHostClusterId(v string) *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) SetDedicatedHostId(v string) *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) SetDedicatedHostName(v string) *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedHostAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute struct {
	// Indicates whether the instance on the dedicated host is associated with the dedicated host. Valid values:
	//
	// 	- default: The instance is not associated with the dedicated host. When the instance is restarted from economical mode, the instance may be automatically deployed on another dedicated host in the automatic deployment resource pool.
	//
	// 	- host: The instance is associated with the dedicated host. When the instance is restarted from economical mode, the instance is still deployed on the original dedicated host.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// Indicates whether the instance is hosted on a dedicated host. Valid values:
	//
	// 	- default: The instance is not hosted on a dedicated host.
	//
	// 	- host: The instance is hosted on a dedicated host.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) GetAffinity() *string {
	return s.Affinity
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) GetTenancy() *string {
	return s.Tenancy
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) SetAffinity(v string) *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute {
	s.Affinity = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) SetTenancy(v string) *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute {
	s.Tenancy = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceDedicatedInstanceAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr struct {
	// The ID of the capacity reservation.
	//
	// example:
	//
	// cr-bp67acfmxazb4p****
	CapacityReservationId *string `json:"CapacityReservationId,omitempty" xml:"CapacityReservationId,omitempty"`
	// The preference of the capacity reservation.
	//
	// example:
	//
	// cr-bp67acfmxazb4p****
	CapacityReservationPreference *string `json:"CapacityReservationPreference,omitempty" xml:"CapacityReservationPreference,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) GetCapacityReservationId() *string {
	return s.CapacityReservationId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) GetCapacityReservationPreference() *string {
	return s.CapacityReservationPreference
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) SetCapacityReservationId(v string) *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr {
	s.CapacityReservationId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) SetCapacityReservationPreference(v string) *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr {
	s.CapacityReservationPreference = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEcsCapacityReservationAttr) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceEipAddress struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-2ze88m67qx5z****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The maximum public bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- PayByBandwidth
	//
	// 	- PayByTraffic
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The EIP.
	//
	// example:
	//
	// ``42.112.**.**``
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// Indicates whether the EIP can be disassociated.
	//
	// example:
	//
	// true
	IsSupportUnassociate *bool `json:"IsSupportUnassociate,omitempty" xml:"IsSupportUnassociate,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceEipAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) GetIsSupportUnassociate() *bool {
	return s.IsSupportUnassociate
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) SetAllocationId(v string) *DescribeInstancesResponseBodyInstancesInstanceEipAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) SetBandwidth(v int32) *DescribeInstancesResponseBodyInstancesInstanceEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) SetInternetChargeType(v string) *DescribeInstancesResponseBodyInstancesInstanceEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) SetIpAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceEipAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) SetIsSupportUnassociate(v bool) *DescribeInstancesResponseBodyInstancesInstanceEipAddress {
	s.IsSupportUnassociate = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceEipAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceHibernationOptions struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceHibernationOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceHibernationOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceHibernationOptions) GetConfigured() *bool {
	return s.Configured
}

func (s *DescribeInstancesResponseBodyInstancesInstanceHibernationOptions) SetConfigured(v bool) *DescribeInstancesResponseBodyInstancesInstanceHibernationOptions {
	s.Configured = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceHibernationOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceImageOptions struct {
	// Indicates whether the operating system supports access to disks over the NVMe protocol. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is returned only if you specify CURRENT_OS_NVME_SUPPORTED in AdditionalAttributes in the request.
	//
	// example:
	//
	// true
	CurrentOSNVMeSupported *bool `json:"CurrentOSNVMeSupported,omitempty" xml:"CurrentOSNVMeSupported,omitempty"`
	// Indicates whether the instance that uses the image supports logons of the ecs-user user. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceImageOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceImageOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceImageOptions) GetCurrentOSNVMeSupported() *bool {
	return s.CurrentOSNVMeSupported
}

func (s *DescribeInstancesResponseBodyInstancesInstanceImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *DescribeInstancesResponseBodyInstancesInstanceImageOptions) SetCurrentOSNVMeSupported(v bool) *DescribeInstancesResponseBodyInstancesInstanceImageOptions {
	s.CurrentOSNVMeSupported = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceImageOptions) SetLoginAsNonRoot(v bool) *DescribeInstancesResponseBodyInstancesInstanceImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceImageOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceInnerIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceMetadataOptions struct {
	// Indicates whether the access channel is enabled for instance metadata. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 0
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Indicates whether the security hardening mode (IMDSv2) is forcefully used to access instance metadata. Valid values:
	//
	// 	- optional: The security hardening mode (IMDSv2) is not forcefully used.
	//
	// 	- required: The security hardening mode (IMDSv2) is forcefully used.
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) SetHttpEndpoint(v string) *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions {
	s.HttpEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) SetHttpPutResponseHopLimit(v int32) *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) SetHttpTokens(v string) *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions {
	s.HttpTokens = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceMetadataOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces struct {
	NetworkInterface []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) GetNetworkInterface() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	return s.NetworkInterface
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) SetNetworkInterface(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces {
	s.NetworkInterface = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfaces) Validate() error {
	if s.NetworkInterface != nil {
		for _, item := range s.NetworkInterface {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface struct {
	// The IPv4 prefixes of the ENI. This parameter has a value only when `AdditionalAttributes.N` is set to `NETWORK_PRIMARY_ENI_IP`.
	Ipv4PrefixSets *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets `json:"Ipv4PrefixSets,omitempty" xml:"Ipv4PrefixSets,omitempty" type:"Struct"`
	// The IPv6 prefixes of the ENI. This parameter has a value only when `AdditionalAttributes.N` is set to `NETWORK_PRIMARY_ENI_IP`.
	Ipv6PrefixSets *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets `json:"Ipv6PrefixSets,omitempty" xml:"Ipv6PrefixSets,omitempty" type:"Struct"`
	// The IPv6 addresses of the ENI. This parameter has a value only when `AdditionalAttributes.N` is set to `NETWORK_PRIMARY_ENI_IP`.
	Ipv6Sets *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets `json:"Ipv6Sets,omitempty" xml:"Ipv6Sets,omitempty" type:"Struct"`
	// The MAC address of the ENI.
	//
	// example:
	//
	// 00:16:3e:32:b4:**
	MacAddress *string `json:"MacAddress,omitempty" xml:"MacAddress,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-2zeh9atclduxvf1z****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The primary private IP address of the ENI.
	//
	// example:
	//
	// ``172.17.**.**``*
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The private IP addresses of the ENI.
	PrivateIpSets *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Struct"`
	// The type of the ENI. Valid values:
	//
	// 	- Primary
	//
	// 	- Secondary
	//
	// example:
	//
	// Primary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetIpv4PrefixSets() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets {
	return s.Ipv4PrefixSets
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetIpv6PrefixSets() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets {
	return s.Ipv6PrefixSets
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetIpv6Sets() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets {
	return s.Ipv6Sets
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetMacAddress() *string {
	return s.MacAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetPrivateIpSets() *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets {
	return s.PrivateIpSets
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) GetType() *string {
	return s.Type
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetIpv4PrefixSets(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.Ipv4PrefixSets = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetIpv6PrefixSets(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.Ipv6PrefixSets = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetIpv6Sets(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.Ipv6Sets = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetMacAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.MacAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetNetworkInterfaceId(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetPrimaryIpAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetPrivateIpSets(v *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.PrivateIpSets = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) SetType(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface {
	s.Type = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterface) Validate() error {
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
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets struct {
	Ipv4PrefixSet []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet `json:"Ipv4PrefixSet,omitempty" xml:"Ipv4PrefixSet,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets) GetIpv4PrefixSet() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet {
	return s.Ipv4PrefixSet
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets) SetIpv4PrefixSet(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets {
	s.Ipv4PrefixSet = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSets) Validate() error {
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

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet struct {
	// The IPv4 prefix of the ENI.
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet) SetIpv4Prefix(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet {
	s.Ipv4Prefix = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv4PrefixSetsIpv4PrefixSet) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets struct {
	Ipv6PrefixSet []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet `json:"Ipv6PrefixSet,omitempty" xml:"Ipv6PrefixSet,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets) GetIpv6PrefixSet() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet {
	return s.Ipv6PrefixSet
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets) SetIpv6PrefixSet(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets {
	s.Ipv6PrefixSet = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSets) Validate() error {
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

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet struct {
	// The IPv6 prefix of the ENI.
	//
	// example:
	//
	// 2001:1111:*:*::/64
	Ipv6Prefix *string `json:"Ipv6Prefix,omitempty" xml:"Ipv6Prefix,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet) GetIpv6Prefix() *string {
	return s.Ipv6Prefix
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet) SetIpv6Prefix(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet {
	s.Ipv6Prefix = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6PrefixSetsIpv6PrefixSet) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets struct {
	Ipv6Set []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set `json:"Ipv6Set,omitempty" xml:"Ipv6Set,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets) GetIpv6Set() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set {
	return s.Ipv6Set
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets) SetIpv6Set(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets {
	s.Ipv6Set = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6Sets) Validate() error {
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

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set struct {
	// The IPv6 address of the ENI.
	//
	// example:
	//
	// 2408:4321:180:1701:94c7:bc38:3bfa:***
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set) SetIpv6Address(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfaceIpv6SetsIpv6Set) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets struct {
	PrivateIpSet []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet `json:"PrivateIpSet,omitempty" xml:"PrivateIpSet,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets) GetPrivateIpSet() []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet {
	return s.PrivateIpSet
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets) SetPrivateIpSet(v []*DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets {
	s.PrivateIpSet = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSets) Validate() error {
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

type DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet struct {
	// Indicates whether the IP address is the primary private IP address. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// The private domain name of the instance.
	//
	// >  This parameter has a value in a specific format only if `HostnameType` is set to `IpBased` or `InstanceIdBased`.
	//
	// example:
	//
	// dnsTestName
	PrivateDnsName *string `json:"PrivateDnsName,omitempty" xml:"PrivateDnsName,omitempty"`
	// The private IP address of the ENI.
	//
	// example:
	//
	// ``172.17.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) GetPrimary() *bool {
	return s.Primary
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) GetPrivateDnsName() *string {
	return s.PrivateDnsName
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) SetPrimary(v bool) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet {
	s.Primary = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) SetPrivateDnsName(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet {
	s.PrivateDnsName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) SetPrivateIpAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceNetworkInterfacesNetworkInterfacePrivateIpSetsPrivateIpSet) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceOperationLocks struct {
	LockReason []*DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocks) GetLockReason() []*DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocks) SetLockReason(v []*DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) *DescribeInstancesResponseBodyInstancesInstanceOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocks) Validate() error {
	if s.LockReason != nil {
		for _, item := range s.LockReason {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason struct {
	// The message returned when the instance was locked.
	//
	// example:
	//
	// The specified instance is locked due to financial reason.
	LockMsg *string `json:"LockMsg,omitempty" xml:"LockMsg,omitempty"`
	// The reason why the instance was locked. Valid values:
	//
	// 	- financial: The instance was locked due to overdue payments.
	//
	// 	- security: The instance was locked due to security reasons.
	//
	// 	- recycling: The spot instance was locked and pending release.
	//
	// 	- dedicatedhostfinancial: The instance was locked due to overdue payments for the dedicated host.
	//
	// 	- refunded: The instance was locked because a refund was made for the instance.
	//
	// example:
	//
	// Recycling
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) GetLockMsg() *string {
	return s.LockMsg
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) SetLockMsg(v string) *DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason {
	s.LockMsg = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) SetLockReason(v string) *DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions struct {
	// Indicates whether DNS Resolution from the Instance ID-based Hostname to the Instance Primary Private IPv6 Address (AAAA Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsAAAARecord *bool `json:"EnableInstanceIdDnsAAAARecord,omitempty" xml:"EnableInstanceIdDnsAAAARecord,omitempty"`
	// Indicates whether DNS Resolution from the Instance ID-based Hostname to the Instance Primary Private IPv4 Address (A Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsARecord *bool `json:"EnableInstanceIdDnsARecord,omitempty" xml:"EnableInstanceIdDnsARecord,omitempty"`
	// Indicates whether DNS Resolution from the IP Address-based Hostname to the Instance Primary Private IPv4 Address (A Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableIpDnsARecord *bool `json:"EnableIpDnsARecord,omitempty" xml:"EnableIpDnsARecord,omitempty"`
	// Indicates whether Reverse DNS Resolution from the Instance Primary Private IPv4 Address to the IP Address-based Hostname (PTR Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableIpDnsPtrRecord *bool `json:"EnableIpDnsPtrRecord,omitempty" xml:"EnableIpDnsPtrRecord,omitempty"`
	// The type of hostname. Valid values:
	//
	// 	- Custom: custom hostname
	//
	// 	- IpBased: IP address-based hostname
	//
	// 	- InstanceIdBased: instance ID-based hostname
	//
	// Default value: Custom.
	//
	// example:
	//
	// Custom
	HostnameType *string `json:"HostnameType,omitempty" xml:"HostnameType,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) GetEnableInstanceIdDnsAAAARecord() *bool {
	return s.EnableInstanceIdDnsAAAARecord
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) GetEnableInstanceIdDnsARecord() *bool {
	return s.EnableInstanceIdDnsARecord
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) GetEnableIpDnsARecord() *bool {
	return s.EnableIpDnsARecord
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) GetEnableIpDnsPtrRecord() *bool {
	return s.EnableIpDnsPtrRecord
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) GetHostnameType() *string {
	return s.HostnameType
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) SetEnableInstanceIdDnsAAAARecord(v bool) *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions {
	s.EnableInstanceIdDnsAAAARecord = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) SetEnableInstanceIdDnsARecord(v bool) *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions {
	s.EnableInstanceIdDnsARecord = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) SetEnableIpDnsARecord(v bool) *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions {
	s.EnableIpDnsARecord = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) SetEnableIpDnsPtrRecord(v bool) *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions {
	s.EnableIpDnsPtrRecord = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) SetHostnameType(v string) *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions {
	s.HostnameType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePrivateDnsNameOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstancePublicIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstancePublicIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceRdmaIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceSecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceTags struct {
	Tag []*DescribeInstancesResponseBodyInstancesInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) GetTag() []*DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	return s.Tag
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) SetTag(v []*DescribeInstancesResponseBodyInstancesInstanceTagsTag) *DescribeInstancesResponseBodyInstancesInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) Validate() error {
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

type DescribeInstancesResponseBodyInstancesInstanceTagsTag struct {
	// The tag key of the instance.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the instance.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetTagKey(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetTagValue(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesInstanceVpcAttributes struct {
	// The NAT IP address of the instance. The NAT IP address is used by ECS instances in different VPCs for communication.
	//
	// example:
	//
	// ``172.17.**.**``
	NatIpAddress *string `json:"NatIpAddress,omitempty" xml:"NatIpAddress,omitempty"`
	// The private IP addresses of the instance.
	PrivateIpAddress *DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Struct"`
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vsw-2zeh0r1pabwtg6wcs****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2zeuphj08tt7q3brd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) GetNatIpAddress() *string {
	return s.NatIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) GetPrivateIpAddress() *DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress {
	return s.PrivateIpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) SetNatIpAddress(v string) *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes {
	s.NatIpAddress = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) SetPrivateIpAddress(v *DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress) *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes {
	s.PrivateIpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributes) Validate() error {
	if s.PrivateIpAddress != nil {
		if err := s.PrivateIpAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress) SetIpAddress(v []*string) *DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceVpcAttributesPrivateIpAddress) Validate() error {
	return dara.Validate(s)
}
