// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotAdviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCores(v int32) *DescribeSpotAdviceRequest
	GetCores() *int32
	SetGpuAmount(v int32) *DescribeSpotAdviceRequest
	GetGpuAmount() *int32
	SetGpuSpec(v string) *DescribeSpotAdviceRequest
	GetGpuSpec() *string
	SetInstanceFamilyLevel(v string) *DescribeSpotAdviceRequest
	GetInstanceFamilyLevel() *string
	SetInstanceTypeFamily(v string) *DescribeSpotAdviceRequest
	GetInstanceTypeFamily() *string
	SetInstanceTypes(v []*string) *DescribeSpotAdviceRequest
	GetInstanceTypes() []*string
	SetMemory(v float32) *DescribeSpotAdviceRequest
	GetMemory() *float32
	SetMinCores(v int32) *DescribeSpotAdviceRequest
	GetMinCores() *int32
	SetMinMemory(v float32) *DescribeSpotAdviceRequest
	GetMinMemory() *float32
	SetOwnerAccount(v string) *DescribeSpotAdviceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSpotAdviceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSpotAdviceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSpotAdviceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSpotAdviceRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeSpotAdviceRequest
	GetZoneId() *string
}

type DescribeSpotAdviceRequest struct {
	// The number of vCPUs of the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The number of GPUs that a GPU-accelerated instance has. For information about the valid values, see [GPU-accelerated compute optimized instance types](https://help.aliyun.com/document_detail/108496.html).
	//
	// example:
	//
	// 2
	GpuAmount *int32 `json:"GpuAmount,omitempty" xml:"GpuAmount,omitempty"`
	// The GPU type. Valid values:
	//
	// 	- NVIDIA P4
	//
	// 	- NVIDIA T4
	//
	// 	- NVIDIA P100
	//
	// 	- NVIDIA V100
	//
	// This parameter is left empty by default, which indicates that all GPU types are queried. For more information, see [GPU-accelerated compute-optimized and vGPU-accelerated instance families](https://help.aliyun.com/document_detail/108496.html).
	//
	// example:
	//
	// NVIDIA T4
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The level of the instance family. Valid values:
	//
	// 	- EntryLevel.
	//
	// 	- EnterpriseLevel.
	//
	// 	- CreditEntryLevel. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// This parameter is left empty by default, which indicates that instance families at all levels are queried.
	//
	// example:
	//
	// EntryLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.c5
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The instance types. You can specify up to 10 instance types.
	//
	// example:
	//
	// ecs.c5.large
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The memory size of the instance type. Unit: GiB. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 8.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The minimum number of vCPUs of the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 2
	MinCores *int32 `json:"MinCores,omitempty" xml:"MinCores,omitempty"`
	// The minimum memory size of the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 8.0
	MinMemory    *float32 `json:"MinMemory,omitempty" xml:"MinMemory,omitempty"`
	OwnerAccount *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The zone ID.
	//
	// This parameter is left empty by default, which indicates that all zones in the specified region are queried.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSpotAdviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotAdviceRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpotAdviceRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeSpotAdviceRequest) GetGpuAmount() *int32 {
	return s.GpuAmount
}

func (s *DescribeSpotAdviceRequest) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeSpotAdviceRequest) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeSpotAdviceRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeSpotAdviceRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeSpotAdviceRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeSpotAdviceRequest) GetMinCores() *int32 {
	return s.MinCores
}

func (s *DescribeSpotAdviceRequest) GetMinMemory() *float32 {
	return s.MinMemory
}

func (s *DescribeSpotAdviceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSpotAdviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSpotAdviceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSpotAdviceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSpotAdviceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSpotAdviceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSpotAdviceRequest) SetCores(v int32) *DescribeSpotAdviceRequest {
	s.Cores = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetGpuAmount(v int32) *DescribeSpotAdviceRequest {
	s.GpuAmount = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetGpuSpec(v string) *DescribeSpotAdviceRequest {
	s.GpuSpec = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetInstanceFamilyLevel(v string) *DescribeSpotAdviceRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetInstanceTypeFamily(v string) *DescribeSpotAdviceRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetInstanceTypes(v []*string) *DescribeSpotAdviceRequest {
	s.InstanceTypes = v
	return s
}

func (s *DescribeSpotAdviceRequest) SetMemory(v float32) *DescribeSpotAdviceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetMinCores(v int32) *DescribeSpotAdviceRequest {
	s.MinCores = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetMinMemory(v float32) *DescribeSpotAdviceRequest {
	s.MinMemory = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetOwnerAccount(v string) *DescribeSpotAdviceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetOwnerId(v int64) *DescribeSpotAdviceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetRegionId(v string) *DescribeSpotAdviceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetResourceOwnerAccount(v string) *DescribeSpotAdviceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetResourceOwnerId(v int64) *DescribeSpotAdviceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSpotAdviceRequest) SetZoneId(v string) *DescribeSpotAdviceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeSpotAdviceRequest) Validate() error {
	return dara.Validate(s)
}
