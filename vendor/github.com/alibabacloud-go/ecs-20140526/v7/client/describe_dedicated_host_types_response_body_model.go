// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostTypes(v *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes) *DescribeDedicatedHostTypesResponseBody
	GetDedicatedHostTypes() *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes
	SetRequestId(v string) *DescribeDedicatedHostTypesResponseBody
	GetRequestId() *string
}

type DescribeDedicatedHostTypesResponseBody struct {
	// Details about the dedicated host types.
	DedicatedHostTypes *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes `json:"DedicatedHostTypes,omitempty" xml:"DedicatedHostTypes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5FE5FF06-3A33-4658-8495-6445FC54E327
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedHostTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostTypesResponseBody) GetDedicatedHostTypes() *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes {
	return s.DedicatedHostTypes
}

func (s *DescribeDedicatedHostTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedHostTypesResponseBody) SetDedicatedHostTypes(v *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes) *DescribeDedicatedHostTypesResponseBody {
	s.DedicatedHostTypes = v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBody) SetRequestId(v string) *DescribeDedicatedHostTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBody) Validate() error {
	if s.DedicatedHostTypes != nil {
		if err := s.DedicatedHostTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes struct {
	DedicatedHostType []*DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes) GetDedicatedHostType() []*DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	return s.DedicatedHostType
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes) SetDedicatedHostType(v []*DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes {
	s.DedicatedHostType = v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypes) Validate() error {
	if s.DedicatedHostType != nil {
		for _, item := range s.DedicatedHostType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType struct {
	// The number of cores per physical CPU.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The supported CPU overcommit ratio range.
	//
	// example:
	//
	// 1-5
	CpuOverCommitRatioRange *string `json:"CpuOverCommitRatioRange,omitempty" xml:"CpuOverCommitRatioRange,omitempty"`
	// The dedicated host type.
	//
	// example:
	//
	// ddh.sn1ne
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// The GPU model.
	//
	// example:
	//
	// gpu
	GPUSpec *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	// The number of local disks on a dedicated host.
	//
	// example:
	//
	// 0
	LocalStorageAmount *int32 `json:"LocalStorageAmount,omitempty" xml:"LocalStorageAmount,omitempty"`
	// The capacity of a local disk. Unit: GiB.
	//
	// example:
	//
	// 0
	LocalStorageCapacity *int64 `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty"`
	// The category of local disks.
	//
	// example:
	//
	// local
	LocalStorageCategory *string `json:"LocalStorageCategory,omitempty" xml:"LocalStorageCategory,omitempty"`
	// The memory size. Unit: GiB.
	//
	// example:
	//
	// 112.0
	MemorySize *float32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The number of physical GPUs.
	//
	// example:
	//
	// 2
	PhysicalGpus *int32 `json:"PhysicalGpus,omitempty" xml:"PhysicalGpus,omitempty"`
	// The number of physical CPUs.
	//
	// example:
	//
	// 2
	Sockets *int32 `json:"Sockets,omitempty" xml:"Sockets,omitempty"`
	// Indicates whether the CPU overcommit ratio settings are supported.
	//
	// example:
	//
	// true
	SupportCpuOverCommitRatio *bool `json:"SupportCpuOverCommitRatio,omitempty" xml:"SupportCpuOverCommitRatio,omitempty"`
	// The ECS instance families supported by the dedicated host type.
	SupportedInstanceTypeFamilies *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies `json:"SupportedInstanceTypeFamilies,omitempty" xml:"SupportedInstanceTypeFamilies,omitempty" type:"Struct"`
	// The ECS instance types supported by the dedicated host type.
	SupportedInstanceTypesList *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList `json:"SupportedInstanceTypesList,omitempty" xml:"SupportedInstanceTypesList,omitempty" type:"Struct"`
	// The total number of vCPUs.
	//
	// example:
	//
	// 56
	TotalVcpus *int32 `json:"TotalVcpus,omitempty" xml:"TotalVcpus,omitempty"`
	// The total number of vGPUs.
	//
	// example:
	//
	// 10
	TotalVgpus *int32 `json:"TotalVgpus,omitempty" xml:"TotalVgpus,omitempty"`
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetCpuOverCommitRatioRange() *string {
	return s.CpuOverCommitRatioRange
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetGPUSpec() *string {
	return s.GPUSpec
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetLocalStorageAmount() *int32 {
	return s.LocalStorageAmount
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetLocalStorageCapacity() *int64 {
	return s.LocalStorageCapacity
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetLocalStorageCategory() *string {
	return s.LocalStorageCategory
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetMemorySize() *float32 {
	return s.MemorySize
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetPhysicalGpus() *int32 {
	return s.PhysicalGpus
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetSockets() *int32 {
	return s.Sockets
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetSupportCpuOverCommitRatio() *bool {
	return s.SupportCpuOverCommitRatio
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetSupportedInstanceTypeFamilies() *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies {
	return s.SupportedInstanceTypeFamilies
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetSupportedInstanceTypesList() *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList {
	return s.SupportedInstanceTypesList
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetTotalVcpus() *int32 {
	return s.TotalVcpus
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) GetTotalVgpus() *int32 {
	return s.TotalVgpus
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetCores(v int32) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.Cores = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetCpuOverCommitRatioRange(v string) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.CpuOverCommitRatioRange = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetDedicatedHostType(v string) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.DedicatedHostType = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetGPUSpec(v string) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.GPUSpec = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetLocalStorageAmount(v int32) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.LocalStorageAmount = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetLocalStorageCapacity(v int64) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.LocalStorageCapacity = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetLocalStorageCategory(v string) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.LocalStorageCategory = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetMemorySize(v float32) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.MemorySize = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetPhysicalGpus(v int32) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.PhysicalGpus = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetSockets(v int32) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.Sockets = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetSupportCpuOverCommitRatio(v bool) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.SupportCpuOverCommitRatio = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetSupportedInstanceTypeFamilies(v *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.SupportedInstanceTypeFamilies = v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetSupportedInstanceTypesList(v *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.SupportedInstanceTypesList = v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetTotalVcpus(v int32) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.TotalVcpus = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) SetTotalVgpus(v int32) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType {
	s.TotalVgpus = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostType) Validate() error {
	if s.SupportedInstanceTypeFamilies != nil {
		if err := s.SupportedInstanceTypeFamilies.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedInstanceTypesList != nil {
		if err := s.SupportedInstanceTypesList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies struct {
	SupportedInstanceTypeFamily []*string `json:"SupportedInstanceTypeFamily,omitempty" xml:"SupportedInstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies) GetSupportedInstanceTypeFamily() []*string {
	return s.SupportedInstanceTypeFamily
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies) SetSupportedInstanceTypeFamily(v []*string) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies {
	s.SupportedInstanceTypeFamily = v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypeFamilies) Validate() error {
	return dara.Validate(s)
}

type DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList struct {
	SupportedInstanceTypesList []*string `json:"SupportedInstanceTypesList,omitempty" xml:"SupportedInstanceTypesList,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList) GetSupportedInstanceTypesList() []*string {
	return s.SupportedInstanceTypesList
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList) SetSupportedInstanceTypesList(v []*string) *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList {
	s.SupportedInstanceTypesList = v
	return s
}

func (s *DescribeDedicatedHostTypesResponseBodyDedicatedHostTypesDedicatedHostTypeSupportedInstanceTypesList) Validate() error {
	return dara.Validate(s)
}
