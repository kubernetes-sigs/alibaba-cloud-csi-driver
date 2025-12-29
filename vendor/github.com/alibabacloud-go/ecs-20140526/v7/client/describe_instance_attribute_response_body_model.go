// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeInstanceAttributeResponseBody
	GetClusterId() *string
	SetCpu(v int32) *DescribeInstanceAttributeResponseBody
	GetCpu() *int32
	SetCreationTime(v string) *DescribeInstanceAttributeResponseBody
	GetCreationTime() *string
	SetCreditSpecification(v string) *DescribeInstanceAttributeResponseBody
	GetCreditSpecification() *string
	SetDedicatedHostAttribute(v *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) *DescribeInstanceAttributeResponseBody
	GetDedicatedHostAttribute() *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute
	SetDescription(v string) *DescribeInstanceAttributeResponseBody
	GetDescription() *string
	SetEipAddress(v *DescribeInstanceAttributeResponseBodyEipAddress) *DescribeInstanceAttributeResponseBody
	GetEipAddress() *DescribeInstanceAttributeResponseBodyEipAddress
	SetEnableJumboFrame(v bool) *DescribeInstanceAttributeResponseBody
	GetEnableJumboFrame() *bool
	SetEnableNetworkEncryption(v bool) *DescribeInstanceAttributeResponseBody
	GetEnableNetworkEncryption() *bool
	SetExpiredTime(v string) *DescribeInstanceAttributeResponseBody
	GetExpiredTime() *string
	SetHostName(v string) *DescribeInstanceAttributeResponseBody
	GetHostName() *string
	SetImageId(v string) *DescribeInstanceAttributeResponseBody
	GetImageId() *string
	SetInnerIpAddress(v *DescribeInstanceAttributeResponseBodyInnerIpAddress) *DescribeInstanceAttributeResponseBody
	GetInnerIpAddress() *DescribeInstanceAttributeResponseBodyInnerIpAddress
	SetInstanceChargeType(v string) *DescribeInstanceAttributeResponseBody
	GetInstanceChargeType() *string
	SetInstanceId(v string) *DescribeInstanceAttributeResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeInstanceAttributeResponseBody
	GetInstanceName() *string
	SetInstanceNetworkType(v string) *DescribeInstanceAttributeResponseBody
	GetInstanceNetworkType() *string
	SetInstanceType(v string) *DescribeInstanceAttributeResponseBody
	GetInstanceType() *string
	SetInternetChargeType(v string) *DescribeInstanceAttributeResponseBody
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *DescribeInstanceAttributeResponseBody
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *DescribeInstanceAttributeResponseBody
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *DescribeInstanceAttributeResponseBody
	GetIoOptimized() *string
	SetMemory(v int32) *DescribeInstanceAttributeResponseBody
	GetMemory() *int32
	SetNetworkOptions(v *DescribeInstanceAttributeResponseBodyNetworkOptions) *DescribeInstanceAttributeResponseBody
	GetNetworkOptions() *DescribeInstanceAttributeResponseBodyNetworkOptions
	SetOperationLocks(v *DescribeInstanceAttributeResponseBodyOperationLocks) *DescribeInstanceAttributeResponseBody
	GetOperationLocks() *DescribeInstanceAttributeResponseBodyOperationLocks
	SetPublicIpAddress(v *DescribeInstanceAttributeResponseBodyPublicIpAddress) *DescribeInstanceAttributeResponseBody
	GetPublicIpAddress() *DescribeInstanceAttributeResponseBodyPublicIpAddress
	SetRegionId(v string) *DescribeInstanceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeInstanceAttributeResponseBody
	GetRequestId() *string
	SetSecurityGroupIds(v *DescribeInstanceAttributeResponseBodySecurityGroupIds) *DescribeInstanceAttributeResponseBody
	GetSecurityGroupIds() *DescribeInstanceAttributeResponseBodySecurityGroupIds
	SetSerialNumber(v string) *DescribeInstanceAttributeResponseBody
	GetSerialNumber() *string
	SetStatus(v string) *DescribeInstanceAttributeResponseBody
	GetStatus() *string
	SetStoppedMode(v string) *DescribeInstanceAttributeResponseBody
	GetStoppedMode() *string
	SetVlanId(v string) *DescribeInstanceAttributeResponseBody
	GetVlanId() *string
	SetVpcAttributes(v *DescribeInstanceAttributeResponseBodyVpcAttributes) *DescribeInstanceAttributeResponseBody
	GetVpcAttributes() *DescribeInstanceAttributeResponseBodyVpcAttributes
	SetZoneId(v string) *DescribeInstanceAttributeResponseBody
	GetZoneId() *string
}

type DescribeInstanceAttributeResponseBody struct {
	// The ID of the cluster to which the instance belongs.
	//
	// > This parameter will be removed in the future. To ensure future compatibility, we recommend that you use other parameters.
	//
	// example:
	//
	// cls-bp67acfmxazb4p****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 8
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2017-12-10T04:04Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// 	- Standard: the standard mode. For more information, see the [Performance modes](~~59977#section-svb-w9d-dju~~) section of the "Overview of burstable instances" topic.
	//
	// 	- Unlimited: the unlimited mode. For more information, see the [Performance modes](~~59977#section-svb-w9d-dju~~) section of the "Overview of burstable instances" topic.
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// Details about the dedicated host. It is an array that consists of the DedicatedHostClusterId, DedicatedHostId, and DedicatedHostName parameters.
	DedicatedHostAttribute *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute `json:"DedicatedHostAttribute,omitempty" xml:"DedicatedHostAttribute,omitempty" type:"Struct"`
	// The description of the instance.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The elastic IP address (EIP) associated with the instance.
	EipAddress *DescribeInstanceAttributeResponseBodyEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Struct"`
	// Indicates whether the Jumbo Frame feature is enabled for the instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// For more information, see [MTUs](https://help.aliyun.com/document_detail/200512.html).
	//
	// example:
	//
	// false
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// example:
	//
	// True
	EnableNetworkEncryption *bool `json:"EnableNetworkEncryption,omitempty" xml:"EnableNetworkEncryption,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2017-12-10T04:04Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The hostname of the instance.
	//
	// example:
	//
	// testHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the image that the instance is running.
	//
	// example:
	//
	// m-bp1h46wfpjsjastc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The internal IP address of the instance located in the classic network.
	InnerIpAddress *DescribeInstanceAttributeResponseBodyInnerIpAddress `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID
	//
	// example:
	//
	// i-uf6f5trc95ug8t33****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// testInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- classic: classic network
	//
	// 	- vpc: VPC
	//
	// example:
	//
	// vpc
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth
	//
	// 	- PayByTraffic
	//
	// >  When the **pay-by-traffic*	- billing method is used for network usage, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios in which demands exceed resource supplies, the maximum bandwidths may not be reached. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
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
	// Indicates whether the ECS instance is I/O optimized. Valid values:
	//
	// 	- optimized: The ECS instance is I/O optimized.
	//
	// 	- none: The ECS instance is not I/O optimized.
	//
	// example:
	//
	// true
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The memory size of the instance. Unit: MiB.
	//
	// example:
	//
	// 16384
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Details about network options.
	//
	// > This parameter is in invitational preview and is not publicly available.
	NetworkOptions *DescribeInstanceAttributeResponseBodyNetworkOptions `json:"NetworkOptions,omitempty" xml:"NetworkOptions,omitempty" type:"Struct"`
	// The reason why the instance was locked. Valid values:
	//
	// 	- financial: The dedicated host was locked due to overdue payments.
	//
	// 	- security: The instance was locked due to security reasons.
	//
	// 	- recycling: The spot instance was locked and pending release.
	//
	// 	- dedicatedhostfinancial: The instance was locked due to overdue payments for the dedicated host.
	//
	// 	- refunded: The instance was locked because a refund was made for the instance.
	OperationLocks *DescribeInstanceAttributeResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The public IP address of the instance.
	PublicIpAddress *DescribeInstanceAttributeResponseBodyPublicIpAddress `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Struct"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the security groups to which the instance belongs.
	SecurityGroupIds *DescribeInstanceAttributeResponseBodySecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	// The serial number of the instance.
	//
	// example:
	//
	// 51d1353b-22bf-4567-a176-8b3e12e4****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- Pending: The instance is being created.
	//
	// 	- Running: The instance is running.
	//
	// 	- Starting: The instance is being started.
	//
	// 	- Stopping: The instance is being stopped.
	//
	// 	- Stopped: The instance is stopped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the system implements billing after the instance is stopped. Valid values:
	//
	// 	- KeepCharging: The instance is stopped in standard mode. The billing of the instance continues after the instance is stopped, and resources are retained for the instance.
	//
	// 	- StopCharging: The instance is stopped in economical mode. The billing of some resources of the instance stops after the instance is stopped. When the instance is stopped, its resources such as vCPUs, memory, and public IP address are released. The instance may be unable to start again if some required resources are out of stock in the current region.
	//
	// 	- Not-applicable: The instance does not support economical mode.
	//
	// example:
	//
	// KeepCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
	// The virtual LAN (VLAN) ID of the instance.
	//
	// > This parameter will be removed in the future. To ensure future compatibility, we recommend that you use other parameters.
	//
	// example:
	//
	// 10
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The VPC attributes of the instance.
	VpcAttributes *DescribeInstanceAttributeResponseBodyVpcAttributes `json:"VpcAttributes,omitempty" xml:"VpcAttributes,omitempty" type:"Struct"`
	// The ID of the zone in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstanceAttributeResponseBody) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstanceAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInstanceAttributeResponseBody) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *DescribeInstanceAttributeResponseBody) GetDedicatedHostAttribute() *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute {
	return s.DedicatedHostAttribute
}

func (s *DescribeInstanceAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceAttributeResponseBody) GetEipAddress() *DescribeInstanceAttributeResponseBodyEipAddress {
	return s.EipAddress
}

func (s *DescribeInstanceAttributeResponseBody) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *DescribeInstanceAttributeResponseBody) GetEnableNetworkEncryption() *bool {
	return s.EnableNetworkEncryption
}

func (s *DescribeInstanceAttributeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeInstanceAttributeResponseBody) GetHostName() *string {
	return s.HostName
}

func (s *DescribeInstanceAttributeResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstanceAttributeResponseBody) GetInnerIpAddress() *DescribeInstanceAttributeResponseBodyInnerIpAddress {
	return s.InnerIpAddress
}

func (s *DescribeInstanceAttributeResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeInstanceAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAttributeResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceAttributeResponseBody) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeInstanceAttributeResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceAttributeResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeInstanceAttributeResponseBody) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *DescribeInstanceAttributeResponseBody) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeInstanceAttributeResponseBody) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeInstanceAttributeResponseBody) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeInstanceAttributeResponseBody) GetNetworkOptions() *DescribeInstanceAttributeResponseBodyNetworkOptions {
	return s.NetworkOptions
}

func (s *DescribeInstanceAttributeResponseBody) GetOperationLocks() *DescribeInstanceAttributeResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *DescribeInstanceAttributeResponseBody) GetPublicIpAddress() *DescribeInstanceAttributeResponseBodyPublicIpAddress {
	return s.PublicIpAddress
}

func (s *DescribeInstanceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAttributeResponseBody) GetSecurityGroupIds() *DescribeInstanceAttributeResponseBodySecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeInstanceAttributeResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeInstanceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceAttributeResponseBody) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *DescribeInstanceAttributeResponseBody) GetVlanId() *string {
	return s.VlanId
}

func (s *DescribeInstanceAttributeResponseBody) GetVpcAttributes() *DescribeInstanceAttributeResponseBodyVpcAttributes {
	return s.VpcAttributes
}

func (s *DescribeInstanceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstanceAttributeResponseBody) SetClusterId(v string) *DescribeInstanceAttributeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetCpu(v int32) *DescribeInstanceAttributeResponseBody {
	s.Cpu = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetCreationTime(v string) *DescribeInstanceAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetCreditSpecification(v string) *DescribeInstanceAttributeResponseBody {
	s.CreditSpecification = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetDedicatedHostAttribute(v *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) *DescribeInstanceAttributeResponseBody {
	s.DedicatedHostAttribute = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetDescription(v string) *DescribeInstanceAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetEipAddress(v *DescribeInstanceAttributeResponseBodyEipAddress) *DescribeInstanceAttributeResponseBody {
	s.EipAddress = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetEnableJumboFrame(v bool) *DescribeInstanceAttributeResponseBody {
	s.EnableJumboFrame = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetEnableNetworkEncryption(v bool) *DescribeInstanceAttributeResponseBody {
	s.EnableNetworkEncryption = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetExpiredTime(v string) *DescribeInstanceAttributeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetHostName(v string) *DescribeInstanceAttributeResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetImageId(v string) *DescribeInstanceAttributeResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInnerIpAddress(v *DescribeInstanceAttributeResponseBodyInnerIpAddress) *DescribeInstanceAttributeResponseBody {
	s.InnerIpAddress = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceChargeType(v string) *DescribeInstanceAttributeResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceId(v string) *DescribeInstanceAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceName(v string) *DescribeInstanceAttributeResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceNetworkType(v string) *DescribeInstanceAttributeResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceType(v string) *DescribeInstanceAttributeResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInternetChargeType(v string) *DescribeInstanceAttributeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInternetMaxBandwidthIn(v int32) *DescribeInstanceAttributeResponseBody {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInternetMaxBandwidthOut(v int32) *DescribeInstanceAttributeResponseBody {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetIoOptimized(v string) *DescribeInstanceAttributeResponseBody {
	s.IoOptimized = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetMemory(v int32) *DescribeInstanceAttributeResponseBody {
	s.Memory = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetNetworkOptions(v *DescribeInstanceAttributeResponseBodyNetworkOptions) *DescribeInstanceAttributeResponseBody {
	s.NetworkOptions = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetOperationLocks(v *DescribeInstanceAttributeResponseBodyOperationLocks) *DescribeInstanceAttributeResponseBody {
	s.OperationLocks = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetPublicIpAddress(v *DescribeInstanceAttributeResponseBodyPublicIpAddress) *DescribeInstanceAttributeResponseBody {
	s.PublicIpAddress = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRegionId(v string) *DescribeInstanceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetSecurityGroupIds(v *DescribeInstanceAttributeResponseBodySecurityGroupIds) *DescribeInstanceAttributeResponseBody {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetSerialNumber(v string) *DescribeInstanceAttributeResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetStatus(v string) *DescribeInstanceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetStoppedMode(v string) *DescribeInstanceAttributeResponseBody {
	s.StoppedMode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetVlanId(v string) *DescribeInstanceAttributeResponseBody {
	s.VlanId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetVpcAttributes(v *DescribeInstanceAttributeResponseBodyVpcAttributes) *DescribeInstanceAttributeResponseBody {
	s.VpcAttributes = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetZoneId(v string) *DescribeInstanceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) Validate() error {
	if s.DedicatedHostAttribute != nil {
		if err := s.DedicatedHostAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.EipAddress != nil {
		if err := s.EipAddress.Validate(); err != nil {
			return err
		}
	}
	if s.InnerIpAddress != nil {
		if err := s.InnerIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkOptions != nil {
		if err := s.NetworkOptions.Validate(); err != nil {
			return err
		}
	}
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
			return err
		}
	}
	if s.PublicIpAddress != nil {
		if err := s.PublicIpAddress.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
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

type DescribeInstanceAttributeResponseBodyDedicatedHostAttribute struct {
	// The ID of the dedicated host.
	//
	// example:
	//
	// dh-2ze7qrzz6lvbfhr0****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The name of the dedicated host.
	//
	// example:
	//
	// ecs-autoui-create-ddh-temp
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) SetDedicatedHostId(v string) *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) SetDedicatedHostName(v string) *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyDedicatedHostAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodyEipAddress struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-wz9uilio26dfscamm****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The maximum public bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 8
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth
	//
	// 	- PayByTraffic
	//
	// >  When the **pay-by-traffic*	- billing method is used for network usage, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios in which demands exceed resource supplies, the maximum bandwidths may not be reached. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The ID of the elastic IP address (EIP).
	//
	// example:
	//
	// ``30.21.**.**``
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyEipAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) SetAllocationId(v string) *DescribeInstanceAttributeResponseBodyEipAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) SetBandwidth(v int32) *DescribeInstanceAttributeResponseBodyEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) SetInternetChargeType(v string) *DescribeInstanceAttributeResponseBodyEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) SetIpAddress(v string) *DescribeInstanceAttributeResponseBodyEipAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyEipAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodyInnerIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyInnerIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInnerIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInnerIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstanceAttributeResponseBodyInnerIpAddress) SetIpAddress(v []*string) *DescribeInstanceAttributeResponseBodyInnerIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInnerIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodyNetworkOptions struct {
	// The bandwidth weight.
	//
	// The supported values vary with instance types. You can query the bandwidth weights supported by the current instance type by using the DescribeInstanceTypes.
	//
	// Valid values:
	//
	// 	- Vpc-L1.
	//
	// 	- Vpc-L2.
	//
	// 	- Ebs-L1.
	//
	// 	- Ebs-L2.
	//
	// 	- Default.
	//
	// example:
	//
	// Vpc-L1
	BandwidthWeighting *string `json:"BandwidthWeighting,omitempty" xml:"BandwidthWeighting,omitempty"`
	// example:
	//
	// false
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// example:
	//
	// False
	EnableNetworkEncryption *bool `json:"EnableNetworkEncryption,omitempty" xml:"EnableNetworkEncryption,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyNetworkOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyNetworkOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyNetworkOptions) GetBandwidthWeighting() *string {
	return s.BandwidthWeighting
}

func (s *DescribeInstanceAttributeResponseBodyNetworkOptions) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *DescribeInstanceAttributeResponseBodyNetworkOptions) GetEnableNetworkEncryption() *bool {
	return s.EnableNetworkEncryption
}

func (s *DescribeInstanceAttributeResponseBodyNetworkOptions) SetBandwidthWeighting(v string) *DescribeInstanceAttributeResponseBodyNetworkOptions {
	s.BandwidthWeighting = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyNetworkOptions) SetEnableJumboFrame(v bool) *DescribeInstanceAttributeResponseBodyNetworkOptions {
	s.EnableJumboFrame = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyNetworkOptions) SetEnableNetworkEncryption(v bool) *DescribeInstanceAttributeResponseBodyNetworkOptions {
	s.EnableNetworkEncryption = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyNetworkOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodyOperationLocks struct {
	LockReason []*DescribeInstanceAttributeResponseBodyOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyOperationLocks) GetLockReason() []*DescribeInstanceAttributeResponseBodyOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeInstanceAttributeResponseBodyOperationLocks) SetLockReason(v []*DescribeInstanceAttributeResponseBodyOperationLocksLockReason) *DescribeInstanceAttributeResponseBodyOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyOperationLocks) Validate() error {
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

type DescribeInstanceAttributeResponseBodyOperationLocksLockReason struct {
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
	// 	- refunded: The instance was locked because a refund is made for the instance.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeInstanceAttributeResponseBodyOperationLocksLockReason) SetLockReason(v string) *DescribeInstanceAttributeResponseBodyOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodyPublicIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyPublicIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyPublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyPublicIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstanceAttributeResponseBodyPublicIpAddress) SetIpAddress(v []*string) *DescribeInstanceAttributeResponseBodyPublicIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyPublicIpAddress) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodySecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodySecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodySecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodySecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeInstanceAttributeResponseBodySecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeInstanceAttributeResponseBodySecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodySecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAttributeResponseBodyVpcAttributes struct {
	// The NAT IP address of the instance. It is used by ECS instances in different VPCs for communication.
	//
	// example:
	//
	// ``172.17.**.**``
	NatIpAddress *string `json:"NatIpAddress,omitempty" xml:"NatIpAddress,omitempty"`
	// The private IP address of the instance.
	PrivateIpAddress *DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Struct"`
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vsw-uf6ixacqz8osrwnqb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-wz9e4e9pmbcnj6ki6****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyVpcAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyVpcAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) GetNatIpAddress() *string {
	return s.NatIpAddress
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) GetPrivateIpAddress() *DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress {
	return s.PrivateIpAddress
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) SetNatIpAddress(v string) *DescribeInstanceAttributeResponseBodyVpcAttributes {
	s.NatIpAddress = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) SetPrivateIpAddress(v *DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) *DescribeInstanceAttributeResponseBodyVpcAttributes {
	s.PrivateIpAddress = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) SetVSwitchId(v string) *DescribeInstanceAttributeResponseBodyVpcAttributes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyVpcAttributes {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributes) Validate() error {
	if s.PrivateIpAddress != nil {
		if err := s.PrivateIpAddress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress struct {
	IpAddress []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) SetIpAddress(v []*string) *DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress {
	s.IpAddress = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyVpcAttributesPrivateIpAddress) Validate() error {
	return dara.Validate(s)
}
