// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCores(v int32) *DescribeRecommendInstanceTypeRequest
	GetCores() *int32
	SetInstanceChargeType(v string) *DescribeRecommendInstanceTypeRequest
	GetInstanceChargeType() *string
	SetInstanceFamilyLevel(v string) *DescribeRecommendInstanceTypeRequest
	GetInstanceFamilyLevel() *string
	SetInstanceType(v string) *DescribeRecommendInstanceTypeRequest
	GetInstanceType() *string
	SetInstanceTypeFamily(v []*string) *DescribeRecommendInstanceTypeRequest
	GetInstanceTypeFamily() []*string
	SetIoOptimized(v string) *DescribeRecommendInstanceTypeRequest
	GetIoOptimized() *string
	SetMaxPrice(v float32) *DescribeRecommendInstanceTypeRequest
	GetMaxPrice() *float32
	SetMemory(v float32) *DescribeRecommendInstanceTypeRequest
	GetMemory() *float32
	SetNetworkType(v string) *DescribeRecommendInstanceTypeRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeRecommendInstanceTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRecommendInstanceTypeRequest
	GetOwnerId() *int64
	SetPriorityStrategy(v string) *DescribeRecommendInstanceTypeRequest
	GetPriorityStrategy() *string
	SetRegionId(v string) *DescribeRecommendInstanceTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRecommendInstanceTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRecommendInstanceTypeRequest
	GetResourceOwnerId() *int64
	SetScene(v string) *DescribeRecommendInstanceTypeRequest
	GetScene() *string
	SetSpotStrategy(v string) *DescribeRecommendInstanceTypeRequest
	GetSpotStrategy() *string
	SetSystemDiskCategory(v string) *DescribeRecommendInstanceTypeRequest
	GetSystemDiskCategory() *string
	SetZoneId(v string) *DescribeRecommendInstanceTypeRequest
	GetZoneId() *string
	SetZoneMatchMode(v string) *DescribeRecommendInstanceTypeRequest
	GetZoneMatchMode() *string
}

type DescribeRecommendInstanceTypeRequest struct {
	// The number of vCPU cores of the instance type.
	//
	// >  If you specify both `Cores` and `Memory`, the system returns all instance types that match the values of the parameters.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The billing method of the ECS instance. For more information, see [Billing overview](https://help.aliyun.com/document_detail/25398.html). Valid values:
	//
	// 	- PrePaid: subscription.
	//
	// 	- PostPaid: pay-as-you-go
	//
	// Default value: PostPaid
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The level of the instance family. Valid values:
	//
	// 	- EntryLevel: entry level.
	//
	// 	- EnterpriseLevel: enterprise level.
	//
	// 	- CreditEntryLevel: credit-based entry level. For more information, see [Burstable instance families](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance type. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the most recent instance type list.
	//
	// >  If you specify `InstanceType`, you cannot specify `Cores` or `Memory`.
	//
	// example:
	//
	// ecs.hfg6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance families from which the alternative instance types are selected. You can specify up to 10 instance families.
	//
	// example:
	//
	// ecs.hfg6
	InstanceTypeFamily []*string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty" type:"Repeated"`
	// Specifies whether instances of the instance type are I/O optimized. You cannot specify IoOptimized if the instance type supports only non-I/O optimized instances. Valid values:
	//
	// 	- optimized: The instances are I/O optimized.
	//
	// 	- none: The instances are non-I/O optimized.
	//
	// Default value: optimized.
	//
	// If you query alternative instance types for retired instance types, this parameter is set to none by default.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The maximum hourly price for pay-as-you-go instances or spot instances.
	//
	// >  This parameter takes effect only when `SpotStrategy` is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 10.0
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The memory size of the instance type. Unit: GiB.
	//
	// >  If you specify both `Cores` and `Memory`, the system returns all instance types that match the values of the parameters.
	//
	// example:
	//
	// 8.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network type of ECS instances. Valid values:
	//
	// 	- classic
	//
	// 	- vpc
	//
	// Default value: vpc.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy for recommending instance types. Valid values:
	//
	// 	- InventoryFirst: recommends instance types in descending order of resource availability.
	//
	// 	- PriceFirst: recommends the most cost-effective instance types. Recommended instance types appear based on the hourly prices of vCPUs in ascending order.
	//
	// 	- NewProductFirst: recommends the latest instance types first.
	//
	// Default value: InventoryFirst.
	//
	// example:
	//
	// PriceFirst
	PriorityStrategy *string `json:"PriorityStrategy,omitempty" xml:"PriorityStrategy,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies the scenarios in which instance types are recommended. Valid values:
	//
	// 	- UPGRADE: instance type upgrade or downgrade
	//
	// 	- CREATE: instance creation
	//
	// Default value: CREATE.
	//
	// example:
	//
	// CREATE
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The bidding policy of the spot instance. Valid values:
	//
	// 	- NoSpot: The instance is created as a pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a spot instance that has a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is a spot instance for which the market price at the time of purchase is automatically used as the bid price. The market price can be up to the pay-as-you-go price.
	//
	// >  If you specify `SpotStrategy`, you must set `InstanceChargeType` to `PostPaid`.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The category of the system disk. Valid values:
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: Enterprise SSD (ESSD)
	//
	// 	- cloud: basic disk
	//
	// For non-I/O optimized instances, the default value is cloud.
	//
	// For I/O optimized instances, the default value is cloud_efficiency.
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The zone ID. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// We recommend that you set ZoneMatchMode to Include, which is the default value. This way, the system recommends instance types that are available in the zone specified by ZoneId based on the priority policy. The system also recommends instance types that are available in other zones within the same region.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// Specifies whether to recommend only instance types in the zone specified by ZoneId. Valid values:
	//
	// 	- Strict: recommends only instance types that are available in the zone specified by ZoneId.
	//
	// 	- Include: recommends instance types that are available in the zone specified by ZoneId and instance types that are available in other zones within the same region.
	//
	// If `ZoneId` is specified, the default value of this parameter is Strict, which indicates that only instance types in the zone specified by ZoneId are recommended.
	//
	// example:
	//
	// Strict
	ZoneMatchMode *string `json:"ZoneMatchMode,omitempty" xml:"ZoneMatchMode,omitempty"`
}

func (s DescribeRecommendInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeRecommendInstanceTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRecommendInstanceTypeRequest) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeRecommendInstanceTypeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRecommendInstanceTypeRequest) GetInstanceTypeFamily() []*string {
	return s.InstanceTypeFamily
}

func (s *DescribeRecommendInstanceTypeRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeRecommendInstanceTypeRequest) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *DescribeRecommendInstanceTypeRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeRecommendInstanceTypeRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeRecommendInstanceTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRecommendInstanceTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRecommendInstanceTypeRequest) GetPriorityStrategy() *string {
	return s.PriorityStrategy
}

func (s *DescribeRecommendInstanceTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRecommendInstanceTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRecommendInstanceTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRecommendInstanceTypeRequest) GetScene() *string {
	return s.Scene
}

func (s *DescribeRecommendInstanceTypeRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeRecommendInstanceTypeRequest) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeRecommendInstanceTypeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRecommendInstanceTypeRequest) GetZoneMatchMode() *string {
	return s.ZoneMatchMode
}

func (s *DescribeRecommendInstanceTypeRequest) SetCores(v int32) *DescribeRecommendInstanceTypeRequest {
	s.Cores = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetInstanceChargeType(v string) *DescribeRecommendInstanceTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetInstanceFamilyLevel(v string) *DescribeRecommendInstanceTypeRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetInstanceType(v string) *DescribeRecommendInstanceTypeRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetInstanceTypeFamily(v []*string) *DescribeRecommendInstanceTypeRequest {
	s.InstanceTypeFamily = v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetIoOptimized(v string) *DescribeRecommendInstanceTypeRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetMaxPrice(v float32) *DescribeRecommendInstanceTypeRequest {
	s.MaxPrice = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetMemory(v float32) *DescribeRecommendInstanceTypeRequest {
	s.Memory = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetNetworkType(v string) *DescribeRecommendInstanceTypeRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetOwnerAccount(v string) *DescribeRecommendInstanceTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetOwnerId(v int64) *DescribeRecommendInstanceTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetPriorityStrategy(v string) *DescribeRecommendInstanceTypeRequest {
	s.PriorityStrategy = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetRegionId(v string) *DescribeRecommendInstanceTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetResourceOwnerAccount(v string) *DescribeRecommendInstanceTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetResourceOwnerId(v int64) *DescribeRecommendInstanceTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetScene(v string) *DescribeRecommendInstanceTypeRequest {
	s.Scene = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetSpotStrategy(v string) *DescribeRecommendInstanceTypeRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetSystemDiskCategory(v string) *DescribeRecommendInstanceTypeRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetZoneId(v string) *DescribeRecommendInstanceTypeRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) SetZoneMatchMode(v string) *DescribeRecommendInstanceTypeRequest {
	s.ZoneMatchMode = &v
	return s
}

func (s *DescribeRecommendInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
