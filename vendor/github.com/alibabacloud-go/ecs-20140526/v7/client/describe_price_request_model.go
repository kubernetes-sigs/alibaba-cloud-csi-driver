// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest
	GetDataDisk() []*DescribePriceRequestDataDisk
	SetSchedulerOptions(v *DescribePriceRequestSchedulerOptions) *DescribePriceRequest
	GetSchedulerOptions() *DescribePriceRequestSchedulerOptions
	SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest
	GetSystemDisk() *DescribePriceRequestSystemDisk
	SetAmount(v int32) *DescribePriceRequest
	GetAmount() *int32
	SetAssuranceTimes(v string) *DescribePriceRequest
	GetAssuranceTimes() *string
	SetCapacity(v int32) *DescribePriceRequest
	GetCapacity() *int32
	SetDedicatedHostType(v string) *DescribePriceRequest
	GetDedicatedHostType() *string
	SetImageId(v string) *DescribePriceRequest
	GetImageId() *string
	SetInstanceAmount(v int32) *DescribePriceRequest
	GetInstanceAmount() *int32
	SetInstanceCpuCoreCount(v int32) *DescribePriceRequest
	GetInstanceCpuCoreCount() *int32
	SetInstanceNetworkType(v string) *DescribePriceRequest
	GetInstanceNetworkType() *string
	SetInstanceType(v string) *DescribePriceRequest
	GetInstanceType() *string
	SetInstanceTypeList(v []*string) *DescribePriceRequest
	GetInstanceTypeList() []*string
	SetInternetChargeType(v string) *DescribePriceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *DescribePriceRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *DescribePriceRequest
	GetIoOptimized() *string
	SetIsp(v string) *DescribePriceRequest
	GetIsp() *string
	SetOfferingType(v string) *DescribePriceRequest
	GetOfferingType() *string
	SetOwnerAccount(v string) *DescribePriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePriceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribePriceRequest
	GetPeriod() *int32
	SetPlatform(v string) *DescribePriceRequest
	GetPlatform() *string
	SetPriceUnit(v string) *DescribePriceRequest
	GetPriceUnit() *string
	SetRecurrenceRules(v []*DescribePriceRequestRecurrenceRules) *DescribePriceRequest
	GetRecurrenceRules() []*DescribePriceRequestRecurrenceRules
	SetRegionId(v string) *DescribePriceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePriceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribePriceRequest
	GetResourceType() *string
	SetScope(v string) *DescribePriceRequest
	GetScope() *string
	SetSpotDuration(v int32) *DescribePriceRequest
	GetSpotDuration() *int32
	SetSpotStrategy(v string) *DescribePriceRequest
	GetSpotStrategy() *string
	SetStartTime(v string) *DescribePriceRequest
	GetStartTime() *string
	SetZoneId(v string) *DescribePriceRequest
	GetZoneId() *string
}

type DescribePriceRequest struct {
	DataDisk         []*DescribePriceRequestDataDisk       `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SchedulerOptions *DescribePriceRequestSchedulerOptions `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	SystemDisk       *DescribePriceRequestSystemDisk       `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The number of ECS instances. You can specify this parameter when you want to query the prices of multiple instances that have specific specifications. Valid values: 1 to 1000.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The total number of times that the elasticity assurance can be applied. Set the value to Unlimited. This value indicates that the elasticity assurance can be applied an unlimited number of times within its effective period.
	//
	// Default value: Unlimited.
	//
	// example:
	//
	// Unlimited
	AssuranceTimes *string `json:"AssuranceTimes,omitempty" xml:"AssuranceTimes,omitempty"`
	// The storage capacity. Unit: GiB.
	//
	// example:
	//
	// 1024
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The type of the dedicated host. You can call the [DescribeDedicatedHostTypes](https://help.aliyun.com/document_detail/134240.html) operation to query the most recent list of dedicated host types.
	//
	// example:
	//
	// ddh.c5
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// This parameter takes effect only when ResourceType is set to instance.
	//
	// The image ID. Images contain the runtime environments to load when instances start. You can call the [DescribeImages](https://help.aliyun.com/document_detail/25534.html) operation to query available images. If you do not specify this parameter, the system queries the prices of Linux images.
	//
	// example:
	//
	// centos_7_05_64_20G_alibase_20181212.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The total number of reserved instances for an instance type.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 100
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The total number of vCPUs supported by the elasticity assurance. When you call this API operation, the system calculates the number of instances that an elasticity assurance must support based on the specified value of InstanceType. The calculated value is rounded up to the nearest integer.
	//
	// > When you call this API operation to query the price of an elasticity assurance, you can only specify either InstanceCoreCpuCount or InstanceAmount.
	//
	// example:
	//
	// 1024
	InstanceCpuCoreCount *int32 `json:"InstanceCpuCoreCount,omitempty" xml:"InstanceCpuCoreCount,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- classic: classic network
	//
	// 	- vpc: Virtual Private Cloud (VPC)
	//
	// Default value: vpc.
	//
	// example:
	//
	// vpc
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The instance type. When `ResourceType` is set to `instance`, you must specify this parameter. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the most recent list of instance types.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance types. You can select only a single instance type when you configure an elasticity assurance in unlimited mode.
	//
	// example:
	//
	// ecs.g6.xlarge
	InstanceTypeList []*string `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic: pay-by-traffic
	//
	// Default value: PayByTraffic
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// 	- none: The instance is not I/O optimized.
	//
	// 	- optimized: The instance is I/O optimized.
	//
	// When the instance type specified by the InstanceType parameter belongs to [Generation I instance families](https://help.aliyun.com/document_detail/55263.html), the default value of this parameter is none.
	//
	// When the instance type specified by the InstanceType parameter does not belong to [Generation I instance families](https://help.aliyun.com/document_detail/55263.html), the default value of this parameter is optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The Internet service provider (ISP). Valid values:
	//
	// 	- cmcc: China Mobile
	//
	// 	- telecom: China Telecom
	//
	// 	- unicom: China Unicom
	//
	// 	- multiCarrier: multi-line ISP
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The payment option of the reserved instance. Valid values:
	//
	// 	- No Upfront
	//
	// 	- Partial Upfront
	//
	// 	- All Upfront
	//
	// example:
	//
	// All Upfront
	OfferingType *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing cycle of the ECS instance. Valid values:
	//
	// 	- Valid values when PriceUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// 	- Valid values when PriceUnit is set to Year: 1, 2, 3, 4, and 5.
	//
	// 	- Set the value to 1 when PriceUnit is set to Hour.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The operating system of the image that is used by the instance. Valid values:
	//
	// 	- Windows: Windows Server operating system
	//
	// 	- Linux: Linux and UNIX-like operating system
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The pricing unit of the ECS resource. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// 	- Hour (default)
	//
	// example:
	//
	// Year
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The assurance schedules of the time-segmented elasticity assurance.
	//
	// >  Time-segmented elasticity assurances are available only in specific regions and to specific users. To use time-segmented elasticity assurances, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	RecurrenceRules []*DescribePriceRequestRecurrenceRules `json:"RecurrenceRules,omitempty" xml:"RecurrenceRules,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- instance: queries the most recent prices of ECS instances. If you set this parameter to `instance`, specify `InstanceType`.
	//
	// 	- disk: queries the most recent prices of cloud disks. If you set this parameter to `disk`, specify `DataDisk.1.Category` and `DataDisk.1.Size`.
	//
	// 	- diskperformance: Queries the most recent prices of the provioned performance of the Enterprise SSD (ESSD) AutoPL disk. You must also specify `DataDisk.1.Category` and `DataDisk.1.ProvisionedIops`.
	//
	// 	- bandwidth: queries the most recent prices for network usage.
	//
	// 	- ddh: queries the most recent prices of dedicated hosts.
	//
	// 	- ElasticityAssurance: queries the most recent prices of elasticity assurances. If you set this parameter to `ElasticityAssurance`, specify `InstanceType`.
	//
	// 	- CapacityReservation: queries the most recent prices of capacity reservations. If you set this parameter to `CapacityReservation`, specify `InstanceType`.
	//
	// Default value: instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The scope of the reserved instance. Valid values:
	//
	// 	- Region: regional
	//
	// 	- Zone: zonal
	//
	// Default value: Region.
	//
	// example:
	//
	// Zone
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// 	- 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends an ECS system event to notify you 5 minutes before the instance is released. Spot instances are billed by second. We recommend that you specify a protection period based on your business requirements.
	//
	// >  This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit or SpotAsPriceGo.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The bidding policy for the pay-as-you-go instance. Valid values:
	//
	// 	- NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is created as a spot instance that has a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is created as a spot instance whose bid price is based on the market price at the time of purchase. The market price can be up to the pay-as-you-go price.
	//
	// Default value: NoSpot.
	//
	// >  This parameter takes effect only when `PriceUnit` is set to Hour and `Period` is set to 1. The default value of `PriceUnit` is `Hour` and the default value of `Period` is `1`. Therefore, you do not need to set `PriceUnit` or `Period` when you set SpotStrategy.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The time when the time-segmented assurance of the elasticity assurance takes effect. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2020-10-30T06:32:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The zone ID.
	//
	// > Prices of spot instances vary based on zones. When you query the price of a spot instance, specify ZoneId.
	//
	// example:
	//
	// cn-hagzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetDataDisk() []*DescribePriceRequestDataDisk {
	return s.DataDisk
}

func (s *DescribePriceRequest) GetSchedulerOptions() *DescribePriceRequestSchedulerOptions {
	return s.SchedulerOptions
}

func (s *DescribePriceRequest) GetSystemDisk() *DescribePriceRequestSystemDisk {
	return s.SystemDisk
}

func (s *DescribePriceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribePriceRequest) GetAssuranceTimes() *string {
	return s.AssuranceTimes
}

func (s *DescribePriceRequest) GetCapacity() *int32 {
	return s.Capacity
}

func (s *DescribePriceRequest) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *DescribePriceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribePriceRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *DescribePriceRequest) GetInstanceCpuCoreCount() *int32 {
	return s.InstanceCpuCoreCount
}

func (s *DescribePriceRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribePriceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePriceRequest) GetInstanceTypeList() []*string {
	return s.InstanceTypeList
}

func (s *DescribePriceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribePriceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribePriceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribePriceRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribePriceRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribePriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceRequest) GetPlatform() *string {
	return s.Platform
}

func (s *DescribePriceRequest) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *DescribePriceRequest) GetRecurrenceRules() []*DescribePriceRequestRecurrenceRules {
	return s.RecurrenceRules
}

func (s *DescribePriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePriceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePriceRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribePriceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribePriceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribePriceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePriceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePriceRequest) SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest {
	s.DataDisk = v
	return s
}

func (s *DescribePriceRequest) SetSchedulerOptions(v *DescribePriceRequestSchedulerOptions) *DescribePriceRequest {
	s.SchedulerOptions = v
	return s
}

func (s *DescribePriceRequest) SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribePriceRequest) SetAmount(v int32) *DescribePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequest) SetAssuranceTimes(v string) *DescribePriceRequest {
	s.AssuranceTimes = &v
	return s
}

func (s *DescribePriceRequest) SetCapacity(v int32) *DescribePriceRequest {
	s.Capacity = &v
	return s
}

func (s *DescribePriceRequest) SetDedicatedHostType(v string) *DescribePriceRequest {
	s.DedicatedHostType = &v
	return s
}

func (s *DescribePriceRequest) SetImageId(v string) *DescribePriceRequest {
	s.ImageId = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceAmount(v int32) *DescribePriceRequest {
	s.InstanceAmount = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceCpuCoreCount(v int32) *DescribePriceRequest {
	s.InstanceCpuCoreCount = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceNetworkType(v string) *DescribePriceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceTypeList(v []*string) *DescribePriceRequest {
	s.InstanceTypeList = v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetInternetMaxBandwidthOut(v int32) *DescribePriceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribePriceRequest) SetIoOptimized(v string) *DescribePriceRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribePriceRequest) SetIsp(v string) *DescribePriceRequest {
	s.Isp = &v
	return s
}

func (s *DescribePriceRequest) SetOfferingType(v string) *DescribePriceRequest {
	s.OfferingType = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerAccount(v string) *DescribePriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerId(v int64) *DescribePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetPlatform(v string) *DescribePriceRequest {
	s.Platform = &v
	return s
}

func (s *DescribePriceRequest) SetPriceUnit(v string) *DescribePriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribePriceRequest) SetRecurrenceRules(v []*DescribePriceRequestRecurrenceRules) *DescribePriceRequest {
	s.RecurrenceRules = v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerAccount(v string) *DescribePriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerId(v int64) *DescribePriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceType(v string) *DescribePriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePriceRequest) SetScope(v string) *DescribePriceRequest {
	s.Scope = &v
	return s
}

func (s *DescribePriceRequest) SetSpotDuration(v int32) *DescribePriceRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribePriceRequest) SetSpotStrategy(v string) *DescribePriceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribePriceRequest) SetStartTime(v string) *DescribePriceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePriceRequest) SetZoneId(v string) *DescribePriceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SchedulerOptions != nil {
		if err := s.SchedulerOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.RecurrenceRules != nil {
		for _, item := range s.RecurrenceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceRequestDataDisk struct {
	// The category of data disk N. Valid values:
	//
	// 	- cloud: basic disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- ephemeral_ssd: local SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// 	- cloud_auto: ESSD AutoPL disk.
	//
	// Valid values of N: 1 to 16.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of data disk N when the disk is an ESSD. This parameter takes effect only when `DataDisk.N.Category` is set to cloud_essd. Valid values:
	//
	// 	- PL0
	//
	// 	- PL1 (default)
	//
	// 	- PL2
	//
	// 	- PL3
	//
	// Valid values of N: 1 to 16.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of data disk N. Unit: GiB. Valid values:
	//
	// 	- Valid values if DataDisk.N.Category is set to cloud: 5 to 2000.
	//
	// 	- Valid values if DataDisk.N.Category is set to cloud_efficiency: 20 to 32768.
	//
	// 	- Valid values if DataDisk.N.Category is set to cloud_ssd: 20 to 32768.
	//
	// 	- Valid values if DataDisk.N.Category is set to cloud_auto: 1 to 32768.
	//
	// 	- Valid values if DataDisk.N.Category is set to cloud_essd: vary based on the `DataDisk.N.PerformanceLevel` value.
	//
	//     	- Valid values if DataDisk.N.PerformanceLevel is set to PL0: 1 to 32768.
	//
	//     	- Valid values if DataDisk.N.PerformanceLevel is set to PL1: 20 to 32768.
	//
	//     	- Valid values if DataDisk.N.PerformanceLevel is set to PL2: 461 to 32768.
	//
	//     	- Valid values if DataDisk.N.PerformanceLevel is set to PL3: 1261 to 32768.
	//
	// 	- Valid values if DataDisk.N.Category is set to ephemeral_ssd: 5 to 800.
	//
	// Valid values of N: 1 to 16.
	//
	// example:
	//
	// 2000
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk to use as data disk N. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// >  This parameter is available only if you set `DataDisk.N.Category` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
}

func (s DescribePriceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribePriceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribePriceRequestDataDisk) GetSize() *int64 {
	return s.Size
}

func (s *DescribePriceRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribePriceRequestDataDisk) SetCategory(v string) *DescribePriceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *DescribePriceRequestDataDisk) SetPerformanceLevel(v string) *DescribePriceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestDataDisk) SetSize(v int64) *DescribePriceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceRequestDataDisk) SetProvisionedIops(v int64) *DescribePriceRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribePriceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestSchedulerOptions struct {
	// This parameter takes effect only when ResourceType is set to instance.
	//
	// The ID of the dedicated host. You can call the [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) operation to query the dedicated host list.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId       *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	DeploymentSetStrategy *string `json:"DeploymentSetStrategy,omitempty" xml:"DeploymentSetStrategy,omitempty"`
}

func (s DescribePriceRequestSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestSchedulerOptions) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestSchedulerOptions) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribePriceRequestSchedulerOptions) GetDeploymentSetStrategy() *string {
	return s.DeploymentSetStrategy
}

func (s *DescribePriceRequestSchedulerOptions) SetDedicatedHostId(v string) *DescribePriceRequestSchedulerOptions {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribePriceRequestSchedulerOptions) SetDeploymentSetStrategy(v string) *DescribePriceRequestSchedulerOptions {
	s.DeploymentSetStrategy = &v
	return s
}

func (s *DescribePriceRequestSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestSystemDisk struct {
	// The category of the system disk. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- ephemeral_ssd: local SSD
	//
	// 	- cloud_essd: Enterprise SSD (ESSD)
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// Default value:
	//
	// 	- When InstanceType is set to a retired instance type and `IoOptimized` is set to `none`, the default value is `cloud`.
	//
	// 	- In other cases, the default value is `cloud_efficiency`.
	//
	// >  If you want to query the price of a system disk, you must also specify `ImageId`.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the system disk when the disk is an ESSD. This parameter is valid only when `SystemDiskCategory` is set to cloud_essd. Valid values:
	//
	// PL0, PL1 (default), PL2, PL3.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// 	- Basic disk (cloud): 20 to 500.
	//
	// 	- ESSD (cloud_essd): Valid values vary based on the SystemDisk.PerformanceLevel value.
	//
	//     	- Valid values when SystemDisk.PerformanceLevel is set to PL0: 1 to 2048.
	//
	//     	- Valid values when SystemDisk.PerformanceLevel is set to PL1: 20 to 2048.
	//
	//     	- Valid values when SystemDisk.PerformanceLevel is set to PL2: 461 to 2048.
	//
	//     	- Valid values when SystemDisk.PerformanceLevel is set to PL3: 1261 to 2048.
	//
	// 	- ESSD AutoPL disk (cloud_auto): 1 to 2048.
	//
	// 	- Other disk categories: 20 to 2048.
	//
	// Default value: 20 or the size of the image specified by ImageId, whichever is greater.
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribePriceRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribePriceRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribePriceRequestSystemDisk) SetCategory(v string) *DescribePriceRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribePriceRequestSystemDisk) SetPerformanceLevel(v string) *DescribePriceRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribePriceRequestSystemDisk) SetSize(v int32) *DescribePriceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestRecurrenceRules struct {
	// The end time of the assurance period for the capacity reservation of the time-segmented elasticity assurance. Specify an on-the-hour point in time.
	//
	// example:
	//
	// 10
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The type of the assurance schedule. Valid values:
	//
	// 	- Daily
	//
	// 	- Weekly
	//
	// 	- Monthly
	//
	// >  If you specify this parameter, you must specify `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The days of the week or month on which the capacity reservation of the time-segmented elasticity assurance takes effect or the interval, in number of days, at which the capacity reservation takes effect.
	//
	// 	- If you set `RecurrenceType` to `Daily`, you can specify only one value. Valid values: 1 to 31. The value specifies that the capacity reservation takes effect every few days.
	//
	// 	- If you set `RecurrenceType` to `Weekly`, you can specify multiple values. Separate the values with commas (,). Valid values: 0, 1, 2, 3, 4, 5, and 6, which specify Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday, respectively. Example: `1,2`, which specifies that the capacity reservation takes effect on Monday and Tuesday.
	//
	// 	- If you set `RecurrenceType` to `Monthly`, you can specify two values in the `A-B` format. Valid values of A and B: 1 to 31. B must be greater than or equal to A. Example: `1-5`, which specifies that the capacity reservation takes effect every day from the first day up to the fifth day of each month.
	//
	// >  If you specify this parameter, you must specify `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// 5
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The start time of the assurance period for the capacity reservation of the time-segmented elasticity assurance. Specify an on-the-hour point in time.
	//
	// >  You must specify both StartHour and EndHour. The EndHour value must be at least 4 hours later than the StartHour value.
	//
	// example:
	//
	// 4
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s DescribePriceRequestRecurrenceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestRecurrenceRules) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestRecurrenceRules) GetEndHour() *int32 {
	return s.EndHour
}

func (s *DescribePriceRequestRecurrenceRules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *DescribePriceRequestRecurrenceRules) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *DescribePriceRequestRecurrenceRules) GetStartHour() *int32 {
	return s.StartHour
}

func (s *DescribePriceRequestRecurrenceRules) SetEndHour(v int32) *DescribePriceRequestRecurrenceRules {
	s.EndHour = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) SetRecurrenceType(v string) *DescribePriceRequestRecurrenceRules {
	s.RecurrenceType = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) SetRecurrenceValue(v string) *DescribePriceRequestRecurrenceRules {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) SetStartHour(v int32) *DescribePriceRequestRecurrenceRules {
	s.StartHour = &v
	return s
}

func (s *DescribePriceRequestRecurrenceRules) Validate() error {
	return dara.Validate(s)
}
