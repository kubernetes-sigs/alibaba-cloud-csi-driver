// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalAttributes(v []*string) *DescribeInstanceTypesRequest
	GetAdditionalAttributes() []*string
	SetCpuArchitecture(v string) *DescribeInstanceTypesRequest
	GetCpuArchitecture() *string
	SetCpuArchitectures(v []*string) *DescribeInstanceTypesRequest
	GetCpuArchitectures() []*string
	SetGPUSpec(v string) *DescribeInstanceTypesRequest
	GetGPUSpec() *string
	SetGpuSpecs(v []*string) *DescribeInstanceTypesRequest
	GetGpuSpecs() []*string
	SetInstanceCategories(v []*string) *DescribeInstanceTypesRequest
	GetInstanceCategories() []*string
	SetInstanceCategory(v string) *DescribeInstanceTypesRequest
	GetInstanceCategory() *string
	SetInstanceFamilyLevel(v string) *DescribeInstanceTypesRequest
	GetInstanceFamilyLevel() *string
	SetInstanceTypeFamilies(v []*string) *DescribeInstanceTypesRequest
	GetInstanceTypeFamilies() []*string
	SetInstanceTypeFamily(v string) *DescribeInstanceTypesRequest
	GetInstanceTypeFamily() *string
	SetInstanceTypes(v []*string) *DescribeInstanceTypesRequest
	GetInstanceTypes() []*string
	SetLocalStorageCategories(v []*string) *DescribeInstanceTypesRequest
	GetLocalStorageCategories() []*string
	SetLocalStorageCategory(v string) *DescribeInstanceTypesRequest
	GetLocalStorageCategory() *string
	SetMaxResults(v int64) *DescribeInstanceTypesRequest
	GetMaxResults() *int64
	SetMaximumCpuCoreCount(v int32) *DescribeInstanceTypesRequest
	GetMaximumCpuCoreCount() *int32
	SetMaximumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest
	GetMaximumCpuSpeedFrequency() *float32
	SetMaximumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest
	GetMaximumCpuTurboFrequency() *float32
	SetMaximumGPUAmount(v int32) *DescribeInstanceTypesRequest
	GetMaximumGPUAmount() *int32
	SetMaximumMemorySize(v float32) *DescribeInstanceTypesRequest
	GetMaximumMemorySize() *float32
	SetMinimumBaselineCredit(v int32) *DescribeInstanceTypesRequest
	GetMinimumBaselineCredit() *int32
	SetMinimumCpuCoreCount(v int32) *DescribeInstanceTypesRequest
	GetMinimumCpuCoreCount() *int32
	SetMinimumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest
	GetMinimumCpuSpeedFrequency() *float32
	SetMinimumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest
	GetMinimumCpuTurboFrequency() *float32
	SetMinimumDiskQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumDiskQuantity() *int32
	SetMinimumEniIpv6AddressQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEniIpv6AddressQuantity() *int32
	SetMinimumEniPrivateIpAddressQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEniPrivateIpAddressQuantity() *int32
	SetMinimumEniQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEniQuantity() *int32
	SetMinimumEriQuantity(v int32) *DescribeInstanceTypesRequest
	GetMinimumEriQuantity() *int32
	SetMinimumGPUAmount(v int32) *DescribeInstanceTypesRequest
	GetMinimumGPUAmount() *int32
	SetMinimumInitialCredit(v int32) *DescribeInstanceTypesRequest
	GetMinimumInitialCredit() *int32
	SetMinimumInstanceBandwidthRx(v int32) *DescribeInstanceTypesRequest
	GetMinimumInstanceBandwidthRx() *int32
	SetMinimumInstanceBandwidthTx(v int32) *DescribeInstanceTypesRequest
	GetMinimumInstanceBandwidthTx() *int32
	SetMinimumInstancePpsRx(v int64) *DescribeInstanceTypesRequest
	GetMinimumInstancePpsRx() *int64
	SetMinimumInstancePpsTx(v int64) *DescribeInstanceTypesRequest
	GetMinimumInstancePpsTx() *int64
	SetMinimumLocalStorageAmount(v int32) *DescribeInstanceTypesRequest
	GetMinimumLocalStorageAmount() *int32
	SetMinimumLocalStorageCapacity(v int64) *DescribeInstanceTypesRequest
	GetMinimumLocalStorageCapacity() *int64
	SetMinimumMemorySize(v float32) *DescribeInstanceTypesRequest
	GetMinimumMemorySize() *float32
	SetMinimumPrimaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest
	GetMinimumPrimaryEniQueueNumber() *int32
	SetMinimumQueuePairNumber(v int32) *DescribeInstanceTypesRequest
	GetMinimumQueuePairNumber() *int32
	SetMinimumSecondaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest
	GetMinimumSecondaryEniQueueNumber() *int32
	SetNextToken(v string) *DescribeInstanceTypesRequest
	GetNextToken() *string
	SetNvmeSupport(v string) *DescribeInstanceTypesRequest
	GetNvmeSupport() *string
	SetOwnerAccount(v string) *DescribeInstanceTypesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceTypesRequest
	GetOwnerId() *int64
	SetPhysicalProcessorModel(v string) *DescribeInstanceTypesRequest
	GetPhysicalProcessorModel() *string
	SetPhysicalProcessorModels(v []*string) *DescribeInstanceTypesRequest
	GetPhysicalProcessorModels() []*string
	SetResourceOwnerAccount(v string) *DescribeInstanceTypesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceTypesRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceTypesRequest struct {
	AdditionalAttributes []*string `json:"AdditionalAttributes,omitempty" xml:"AdditionalAttributes,omitempty" type:"Repeated"`
	// The CPU architecture. Valid values:
	//
	// 	- X86
	//
	// 	- ARM
	//
	// example:
	//
	// X86
	CpuArchitecture *string `json:"CpuArchitecture,omitempty" xml:"CpuArchitecture,omitempty"`
	// The CPU architectures of instance types. You can specify 1 or 2 CPU architectures.
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The GPU model.
	//
	// >  Fuzzy match is supported. For example, if an instance type provides NVIDIA V100 GPUs and you set this parameter to NVIDIA, information about the instance type is queried.
	//
	// example:
	//
	// NVIDIA V100
	GPUSpec *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	// The GPU models of instance types. You can specify 1 to 10 GPU models.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The categories of instance types. You can specify 1 to 10 categories of instance types.
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The category of the instance type. Valid values:
	//
	// 	- General-purpose: general-purpose instance type
	//
	// 	- Compute-optimized: compute-optimized instance type
	//
	// 	- Memory-optimized: memory-optimized instance type
	//
	// 	- Big data: big data instance type
	//
	// 	- Local SSDs: instance type with local SSDs
	//
	// 	- High Clock Speed: instance type with high clock speeds
	//
	// 	- Enhanced: enhanced instance type
	//
	// 	- Shared: shared instance type
	//
	// 	- Compute-optimized with GPU: GPU-accelerated compute-optimized instance type
	//
	// 	- Visual Compute-optimized: visual compute-optimized instance type
	//
	// 	- Heterogeneous Service: heterogeneous service instance type
	//
	// 	- Compute-optimized with FPGA: FPGA-accelerated compute-optimized instance type
	//
	// 	- Compute-optimized with NPU: NPU-accelerated compute-optimized instance type
	//
	// 	- ECS Bare Metal: ECS Bare Metal Instance type
	//
	// 	- Super Computing Cluster: Super Computing Cluster (SCC) instance type
	//
	// 	- High Performance Compute: high-performance computing instance type
	//
	// example:
	//
	// Big data
	InstanceCategory *string `json:"InstanceCategory,omitempty" xml:"InstanceCategory,omitempty"`
	// The level of the instance family. Valid values:
	//
	// 	- EntryLevel: entry level (shared)
	//
	// 	- EnterpriseLevel: enterprise level
	//
	// 	- CreditEntryLevel: credit-based entry level
	//
	// example:
	//
	// EntryLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families. You can specify 1 to 10 instance families.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The instance family to which the instance type belongs. For information about the valid values of this parameter, see [DescribeInstanceTypeFamilies](https://help.aliyun.com/document_detail/25621.html).
	//
	// For more information about instance families, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g6
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The instance types. You can specify 1 to 10 instance types. If this parameter is empty, information about all instance types is queried.
	//
	// example:
	//
	// ecs.g6.large
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The categories of local disks used by instance types. You can specify 1 or 2 categories of local disks.
	LocalStorageCategories []*string `json:"LocalStorageCategories,omitempty" xml:"LocalStorageCategories,omitempty" type:"Repeated"`
	// The category of local disks. For more information, see [Local disks](~~63138#section_n2w_8yc_5u1~~). Valid values:
	//
	// 	- local_hdd_pro: local Serial Advanced Technology Attachment (SATA) HDDs, which are attached to d1ne or d1 instances.
	//
	// 	- local_ssd_pro: local Non-Volatile Memory Express (NVMe) SSDs, which are attached to i2, i2g, i1, ga1, or gn5 instances.
	//
	// Valid values:
	//
	// 	- local_hdd_pro
	//
	// 	- local_ssd_pro
	//
	// example:
	//
	// local_ssd_pro
	LocalStorageCategory *string `json:"LocalStorageCategory,omitempty" xml:"LocalStorageCategory,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 1600.
	//
	// Default value: 1600.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The maximum number of vCPUs. The value must be a positive integer.
	//
	// >  If an instance type has more vCPUs than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 10
	MaximumCpuCoreCount *int32 `json:"MaximumCpuCoreCount,omitempty" xml:"MaximumCpuCoreCount,omitempty"`
	// The maximum clock speed.
	//
	// >  If an instance type uses processors that have a higher clock speed than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 3.2
	MaximumCpuSpeedFrequency *float32 `json:"MaximumCpuSpeedFrequency,omitempty" xml:"MaximumCpuSpeedFrequency,omitempty"`
	// The maximum turbo frequency.
	//
	// >  If an instance type uses processors that deliver a higher turbo frequency than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 4.1
	MaximumCpuTurboFrequency *float32 `json:"MaximumCpuTurboFrequency,omitempty" xml:"MaximumCpuTurboFrequency,omitempty"`
	// The maximum number of GPUs. The value must be a positive integer.
	//
	// >  If an instance type provides more GPUs than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 10
	MaximumGPUAmount *int32 `json:"MaximumGPUAmount,omitempty" xml:"MaximumGPUAmount,omitempty"`
	// The maximum memory size. Unit: GiB.
	//
	// >  If the memory size of an instance type is larger than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 60
	MaximumMemorySize *float32 `json:"MaximumMemorySize,omitempty" xml:"MaximumMemorySize,omitempty"`
	// The minimum baseline CPU performance (overall baseline performance of all vCPUs) of a t5 or t6 burstable instance.
	//
	// >  If a t5 or t6 instance type provides baseline CPU performance lower than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 12
	MinimumBaselineCredit *int32 `json:"MinimumBaselineCredit,omitempty" xml:"MinimumBaselineCredit,omitempty"`
	// The minimum number of vCPUs. The value must be a positive integer.
	//
	// >  If an instance type has fewer vCPUs than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum clock speed.
	//
	// >  If an instance type uses processors that have a lower clock speed than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 2.5
	MinimumCpuSpeedFrequency *float32 `json:"MinimumCpuSpeedFrequency,omitempty" xml:"MinimumCpuSpeedFrequency,omitempty"`
	// The minimum turbo frequency.
	//
	// >  If an instance type uses processors that deliver a lower turbo frequency than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 3.2
	MinimumCpuTurboFrequency *float32 `json:"MinimumCpuTurboFrequency,omitempty" xml:"MinimumCpuTurboFrequency,omitempty"`
	// The minimum number of cloud disks per instance.
	//
	// >  If an instance type supports fewer cloud disks than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 4
	MinimumDiskQuantity *int32 `json:"MinimumDiskQuantity,omitempty" xml:"MinimumDiskQuantity,omitempty"`
	// The minimum number of IPv6 addresses per ENI.
	//
	// >  If an instance type supports fewer IPv6 addresses per ENI than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 2
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum number of IPv4 addresses per ENI.
	//
	// >  If an instance type supports fewer IPv4 addresses per ENI than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum number of elastic network interfaces (ENIs) per instance.
	//
	// >  If an instance type supports fewer ENIs than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 4
	MinimumEniQuantity *int32 `json:"MinimumEniQuantity,omitempty" xml:"MinimumEniQuantity,omitempty"`
	// The minimum number of ERIs per instance.
	//
	// >  If an instance type supports fewer ERIs than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 0
	MinimumEriQuantity *int32 `json:"MinimumEriQuantity,omitempty" xml:"MinimumEriQuantity,omitempty"`
	// The minimum number of GPUs. The value must be a positive integer.
	//
	// >  If an instance type provides fewer GPUs than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 3
	MinimumGPUAmount *int32 `json:"MinimumGPUAmount,omitempty" xml:"MinimumGPUAmount,omitempty"`
	// The minimum initial CPU credits of a t5 or t6 burstable instance.
	//
	// >  If a t5 or t6 instance type provides less initial vCPU credits than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 12
	MinimumInitialCredit *int32 `json:"MinimumInitialCredit,omitempty" xml:"MinimumInitialCredit,omitempty"`
	// The minimum inbound internal bandwidth. Unit: Kbit/s.
	//
	// >  If an instance type provides an inbound internal bandwidth that is lower than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 12288
	MinimumInstanceBandwidthRx *int32 `json:"MinimumInstanceBandwidthRx,omitempty" xml:"MinimumInstanceBandwidthRx,omitempty"`
	// The minimum outbound internal bandwidth. Unit: Kbit/s.
	//
	// >  If an instance type provides an outbound internal bandwidth that is lower than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 12288
	MinimumInstanceBandwidthTx *int32 `json:"MinimumInstanceBandwidthTx,omitempty" xml:"MinimumInstanceBandwidthTx,omitempty"`
	// The minimum inbound packet forwarding rate over the internal network. Unit: pps.
	//
	// >  If an instance type provides an inbound packet forwarding rate over the internal network that is lower than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 15
	MinimumInstancePpsRx *int64 `json:"MinimumInstancePpsRx,omitempty" xml:"MinimumInstancePpsRx,omitempty"`
	// The minimum outbound packet forwarding rate over the internal network. Unit: pps.
	//
	// >  If an instance type provides an outbound packet forwarding rate over the internal network that is lower than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 15
	MinimumInstancePpsTx *int64 `json:"MinimumInstancePpsTx,omitempty" xml:"MinimumInstancePpsTx,omitempty"`
	// The minimum number of local disks per instance.
	//
	// >  If an instance type supports fewer local disks than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 4
	MinimumLocalStorageAmount *int32 `json:"MinimumLocalStorageAmount,omitempty" xml:"MinimumLocalStorageAmount,omitempty"`
	// The capacity of each local disk attached per instance. Unit: GiB.
	//
	// example:
	//
	// 40
	MinimumLocalStorageCapacity *int64 `json:"MinimumLocalStorageCapacity,omitempty" xml:"MinimumLocalStorageCapacity,omitempty"`
	// The minimum memory size. Unit: GiB.
	//
	// >  If the memory size of an instance type is smaller than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 20
	MinimumMemorySize *float32 `json:"MinimumMemorySize,omitempty" xml:"MinimumMemorySize,omitempty"`
	// The minimum default number of queues per primary network interface controller (NIC).
	//
	// >  If an instance type supports fewer queues per primary NIC than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 8
	MinimumPrimaryEniQueueNumber *int32 `json:"MinimumPrimaryEniQueueNumber,omitempty" xml:"MinimumPrimaryEniQueueNumber,omitempty"`
	// The minimum number of queue pair (QP) queues per elastic RDMA interface (ERI).
	//
	// >  If an instance type supports fewer QP queues per ERI than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 8
	MinimumQueuePairNumber *int32 `json:"MinimumQueuePairNumber,omitempty" xml:"MinimumQueuePairNumber,omitempty"`
	// The minimum default number of queues per secondary NIC.
	//
	// >  If an instance type supports fewer queues per secondary NIC than the specified value, information about the instance type is not queried.
	//
	// example:
	//
	// 4
	MinimumSecondaryEniQueueNumber *int32 `json:"MinimumSecondaryEniQueueNumber,omitempty" xml:"MinimumSecondaryEniQueueNumber,omitempty"`
	// The query token. Set the value to the NextToken value returned in the previous call to the DescribeInstanceTypes operation. You do not need to specify this parameter for the first request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether cloud disks can be attached by using the NVMe protocol. Valid values:
	//
	// 	- required: Cloud disks can be attached by using the NVMe protocol.
	//
	// 	- unsupported: Cloud disks cannot be attached by using the NVMe protocol.
	//
	// example:
	//
	// required
	NvmeSupport  *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The CPU model.
	//
	// >  Fuzzy match is supported. For example, if an instance type uses Intel Xeon (Ice Lake) Platinum 8369B processors and you set this parameter to Intel, information about the instance type is queried.
	//
	// example:
	//
	// Intel Xeon(Ice Lake) Platinum 8369B
	PhysicalProcessorModel *string `json:"PhysicalProcessorModel,omitempty" xml:"PhysicalProcessorModel,omitempty"`
	// The CPU models of instance types. You can specify 1 to 10 CPU models.
	PhysicalProcessorModels []*string `json:"PhysicalProcessorModels,omitempty" xml:"PhysicalProcessorModels,omitempty" type:"Repeated"`
	ResourceOwnerAccount    *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesRequest) GetAdditionalAttributes() []*string {
	return s.AdditionalAttributes
}

func (s *DescribeInstanceTypesRequest) GetCpuArchitecture() *string {
	return s.CpuArchitecture
}

func (s *DescribeInstanceTypesRequest) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *DescribeInstanceTypesRequest) GetGPUSpec() *string {
	return s.GPUSpec
}

func (s *DescribeInstanceTypesRequest) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *DescribeInstanceTypesRequest) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *DescribeInstanceTypesRequest) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *DescribeInstanceTypesRequest) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeInstanceTypesRequest) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *DescribeInstanceTypesRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeInstanceTypesRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeInstanceTypesRequest) GetLocalStorageCategories() []*string {
	return s.LocalStorageCategories
}

func (s *DescribeInstanceTypesRequest) GetLocalStorageCategory() *string {
	return s.LocalStorageCategory
}

func (s *DescribeInstanceTypesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeInstanceTypesRequest) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *DescribeInstanceTypesRequest) GetMaximumCpuSpeedFrequency() *float32 {
	return s.MaximumCpuSpeedFrequency
}

func (s *DescribeInstanceTypesRequest) GetMaximumCpuTurboFrequency() *float32 {
	return s.MaximumCpuTurboFrequency
}

func (s *DescribeInstanceTypesRequest) GetMaximumGPUAmount() *int32 {
	return s.MaximumGPUAmount
}

func (s *DescribeInstanceTypesRequest) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *DescribeInstanceTypesRequest) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *DescribeInstanceTypesRequest) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *DescribeInstanceTypesRequest) GetMinimumCpuSpeedFrequency() *float32 {
	return s.MinimumCpuSpeedFrequency
}

func (s *DescribeInstanceTypesRequest) GetMinimumCpuTurboFrequency() *float32 {
	return s.MinimumCpuTurboFrequency
}

func (s *DescribeInstanceTypesRequest) GetMinimumDiskQuantity() *int32 {
	return s.MinimumDiskQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumEriQuantity() *int32 {
	return s.MinimumEriQuantity
}

func (s *DescribeInstanceTypesRequest) GetMinimumGPUAmount() *int32 {
	return s.MinimumGPUAmount
}

func (s *DescribeInstanceTypesRequest) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstanceBandwidthRx() *int32 {
	return s.MinimumInstanceBandwidthRx
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstanceBandwidthTx() *int32 {
	return s.MinimumInstanceBandwidthTx
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstancePpsRx() *int64 {
	return s.MinimumInstancePpsRx
}

func (s *DescribeInstanceTypesRequest) GetMinimumInstancePpsTx() *int64 {
	return s.MinimumInstancePpsTx
}

func (s *DescribeInstanceTypesRequest) GetMinimumLocalStorageAmount() *int32 {
	return s.MinimumLocalStorageAmount
}

func (s *DescribeInstanceTypesRequest) GetMinimumLocalStorageCapacity() *int64 {
	return s.MinimumLocalStorageCapacity
}

func (s *DescribeInstanceTypesRequest) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *DescribeInstanceTypesRequest) GetMinimumPrimaryEniQueueNumber() *int32 {
	return s.MinimumPrimaryEniQueueNumber
}

func (s *DescribeInstanceTypesRequest) GetMinimumQueuePairNumber() *int32 {
	return s.MinimumQueuePairNumber
}

func (s *DescribeInstanceTypesRequest) GetMinimumSecondaryEniQueueNumber() *int32 {
	return s.MinimumSecondaryEniQueueNumber
}

func (s *DescribeInstanceTypesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceTypesRequest) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *DescribeInstanceTypesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceTypesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceTypesRequest) GetPhysicalProcessorModel() *string {
	return s.PhysicalProcessorModel
}

func (s *DescribeInstanceTypesRequest) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *DescribeInstanceTypesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceTypesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceTypesRequest) SetAdditionalAttributes(v []*string) *DescribeInstanceTypesRequest {
	s.AdditionalAttributes = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetCpuArchitecture(v string) *DescribeInstanceTypesRequest {
	s.CpuArchitecture = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetCpuArchitectures(v []*string) *DescribeInstanceTypesRequest {
	s.CpuArchitectures = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetGPUSpec(v string) *DescribeInstanceTypesRequest {
	s.GPUSpec = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetGpuSpecs(v []*string) *DescribeInstanceTypesRequest {
	s.GpuSpecs = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceCategories(v []*string) *DescribeInstanceTypesRequest {
	s.InstanceCategories = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceCategory(v string) *DescribeInstanceTypesRequest {
	s.InstanceCategory = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceFamilyLevel(v string) *DescribeInstanceTypesRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceTypeFamilies(v []*string) *DescribeInstanceTypesRequest {
	s.InstanceTypeFamilies = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceTypeFamily(v string) *DescribeInstanceTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetInstanceTypes(v []*string) *DescribeInstanceTypesRequest {
	s.InstanceTypes = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetLocalStorageCategories(v []*string) *DescribeInstanceTypesRequest {
	s.LocalStorageCategories = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetLocalStorageCategory(v string) *DescribeInstanceTypesRequest {
	s.LocalStorageCategory = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaxResults(v int64) *DescribeInstanceTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumCpuCoreCount(v int32) *DescribeInstanceTypesRequest {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MaximumCpuSpeedFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MaximumCpuTurboFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumGPUAmount(v int32) *DescribeInstanceTypesRequest {
	s.MaximumGPUAmount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMaximumMemorySize(v float32) *DescribeInstanceTypesRequest {
	s.MaximumMemorySize = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumBaselineCredit(v int32) *DescribeInstanceTypesRequest {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumCpuCoreCount(v int32) *DescribeInstanceTypesRequest {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumCpuSpeedFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MinimumCpuSpeedFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumCpuTurboFrequency(v float32) *DescribeInstanceTypesRequest {
	s.MinimumCpuTurboFrequency = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumDiskQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumDiskQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEniIpv6AddressQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEniPrivateIpAddressQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEniQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEniQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumEriQuantity(v int32) *DescribeInstanceTypesRequest {
	s.MinimumEriQuantity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumGPUAmount(v int32) *DescribeInstanceTypesRequest {
	s.MinimumGPUAmount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInitialCredit(v int32) *DescribeInstanceTypesRequest {
	s.MinimumInitialCredit = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstanceBandwidthRx(v int32) *DescribeInstanceTypesRequest {
	s.MinimumInstanceBandwidthRx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstanceBandwidthTx(v int32) *DescribeInstanceTypesRequest {
	s.MinimumInstanceBandwidthTx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstancePpsRx(v int64) *DescribeInstanceTypesRequest {
	s.MinimumInstancePpsRx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumInstancePpsTx(v int64) *DescribeInstanceTypesRequest {
	s.MinimumInstancePpsTx = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumLocalStorageAmount(v int32) *DescribeInstanceTypesRequest {
	s.MinimumLocalStorageAmount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumLocalStorageCapacity(v int64) *DescribeInstanceTypesRequest {
	s.MinimumLocalStorageCapacity = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumMemorySize(v float32) *DescribeInstanceTypesRequest {
	s.MinimumMemorySize = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumPrimaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest {
	s.MinimumPrimaryEniQueueNumber = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumQueuePairNumber(v int32) *DescribeInstanceTypesRequest {
	s.MinimumQueuePairNumber = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetMinimumSecondaryEniQueueNumber(v int32) *DescribeInstanceTypesRequest {
	s.MinimumSecondaryEniQueueNumber = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetNextToken(v string) *DescribeInstanceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetNvmeSupport(v string) *DescribeInstanceTypesRequest {
	s.NvmeSupport = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetOwnerAccount(v string) *DescribeInstanceTypesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetOwnerId(v int64) *DescribeInstanceTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetPhysicalProcessorModel(v string) *DescribeInstanceTypesRequest {
	s.PhysicalProcessorModel = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetPhysicalProcessorModels(v []*string) *DescribeInstanceTypesRequest {
	s.PhysicalProcessorModels = v
	return s
}

func (s *DescribeInstanceTypesRequest) SetResourceOwnerAccount(v string) *DescribeInstanceTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceTypesRequest) SetResourceOwnerId(v int64) *DescribeInstanceTypesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceTypesRequest) Validate() error {
	return dara.Validate(s)
}
