// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineTypes(v []*ListMachineTypesResponseBodyMachineTypes) *ListMachineTypesResponseBody
	GetMachineTypes() []*ListMachineTypesResponseBodyMachineTypes
	SetNextToken(v string) *ListMachineTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMachineTypesResponseBody
	GetRequestId() *string
}

type ListMachineTypesResponseBody struct {
	// Details of the machine types.
	MachineTypes []*ListMachineTypesResponseBodyMachineTypes `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty" type:"Repeated"`
	// The token to request the next page of results. Include this token in your next request to retrieve the next page.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F16BA4D8-FF50-53B6-A026-F443FE31006C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMachineTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBody) GetMachineTypes() []*ListMachineTypesResponseBodyMachineTypes {
	return s.MachineTypes
}

func (s *ListMachineTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMachineTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMachineTypesResponseBody) SetMachineTypes(v []*ListMachineTypesResponseBodyMachineTypes) *ListMachineTypesResponseBody {
	s.MachineTypes = v
	return s
}

func (s *ListMachineTypesResponseBody) SetNextToken(v string) *ListMachineTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMachineTypesResponseBody) SetRequestId(v string) *ListMachineTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMachineTypesResponseBody) Validate() error {
	if s.MachineTypes != nil {
		for _, item := range s.MachineTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMachineTypesResponseBodyMachineTypes struct {
	// The number of bonds.
	//
	// example:
	//
	// 2
	BondNum *int32 `json:"BondNum,omitempty" xml:"BondNum,omitempty"`
	// CPU information.
	//
	// example:
	//
	// 2x Intel Icelake 8369B 32C CPU
	CpuInfo       *string                                                `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	CpuInfoDetail *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail `json:"CpuInfoDetail,omitempty" xml:"CpuInfoDetail,omitempty" type:"Struct"`
	// Disk information.
	//
	// example:
	//
	// 2x 480GB SATA SSD
	DiskInfo              *string                                                        `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	DiskInfoDetail        *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail        `json:"DiskInfoDetail,omitempty" xml:"DiskInfoDetail,omitempty" type:"Struct"`
	FrontendNetworkDetail *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail `json:"FrontendNetworkDetail,omitempty" xml:"FrontendNetworkDetail,omitempty" type:"Struct"`
	// GPU information.
	//
	// example:
	//
	// 8x NVIDIA SXM4 80GB A100 GPU
	GpuInfo       *string                                                `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	GpuInfoDetail *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail `json:"GpuInfoDetail,omitempty" xml:"GpuInfoDetail,omitempty" type:"Struct"`
	// Memory information.
	//
	// example:
	//
	// 32x 64GB DDR4 3200 Memory
	MemoryInfo       *string                                                   `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	MemoryInfoDetail *ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail `json:"MemoryInfoDetail,omitempty" xml:"MemoryInfoDetail,omitempty" type:"Struct"`
	// The name of the machine type.
	//
	// example:
	//
	// efg1.nvga1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Network information.
	//
	// example:
	//
	// 2x 100Gbps DP NIC
	NetworkInfo *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 10
	NodeCount      *string                                                 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	RdmaInfoDetail *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail `json:"RdmaInfoDetail,omitempty" xml:"RdmaInfoDetail,omitempty" type:"Struct"`
	// The number of CPU cores.
	//
	// example:
	//
	// 48
	TotalCpuCore *int32 `json:"TotalCpuCore,omitempty" xml:"TotalCpuCore,omitempty"`
	// The type of the machine type.
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypes) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypes) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetBondNum() *int32 {
	return s.BondNum
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetCpuInfo() *string {
	return s.CpuInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetCpuInfoDetail() *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail {
	return s.CpuInfoDetail
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetDiskInfo() *string {
	return s.DiskInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetDiskInfoDetail() *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail {
	return s.DiskInfoDetail
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetFrontendNetworkDetail() *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail {
	return s.FrontendNetworkDetail
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetGpuInfo() *string {
	return s.GpuInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetGpuInfoDetail() *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail {
	return s.GpuInfoDetail
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetMemoryInfo() *string {
	return s.MemoryInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetMemoryInfoDetail() *ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail {
	return s.MemoryInfoDetail
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetName() *string {
	return s.Name
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetNetworkInfo() *string {
	return s.NetworkInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetNodeCount() *string {
	return s.NodeCount
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetRdmaInfoDetail() *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail {
	return s.RdmaInfoDetail
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetTotalCpuCore() *int32 {
	return s.TotalCpuCore
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetType() *string {
	return s.Type
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetBondNum(v int32) *ListMachineTypesResponseBodyMachineTypes {
	s.BondNum = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetCpuInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.CpuInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetCpuInfoDetail(v *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) *ListMachineTypesResponseBodyMachineTypes {
	s.CpuInfoDetail = v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetDiskInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.DiskInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetDiskInfoDetail(v *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) *ListMachineTypesResponseBodyMachineTypes {
	s.DiskInfoDetail = v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetFrontendNetworkDetail(v *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) *ListMachineTypesResponseBodyMachineTypes {
	s.FrontendNetworkDetail = v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetGpuInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.GpuInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetGpuInfoDetail(v *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) *ListMachineTypesResponseBodyMachineTypes {
	s.GpuInfoDetail = v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetMemoryInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.MemoryInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetMemoryInfoDetail(v *ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail) *ListMachineTypesResponseBodyMachineTypes {
	s.MemoryInfoDetail = v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetName(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.Name = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetNetworkInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.NetworkInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetNodeCount(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.NodeCount = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetRdmaInfoDetail(v *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) *ListMachineTypesResponseBodyMachineTypes {
	s.RdmaInfoDetail = v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetTotalCpuCore(v int32) *ListMachineTypesResponseBodyMachineTypes {
	s.TotalCpuCore = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetType(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.Type = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) Validate() error {
	if s.CpuInfoDetail != nil {
		if err := s.CpuInfoDetail.Validate(); err != nil {
			return err
		}
	}
	if s.DiskInfoDetail != nil {
		if err := s.DiskInfoDetail.Validate(); err != nil {
			return err
		}
	}
	if s.FrontendNetworkDetail != nil {
		if err := s.FrontendNetworkDetail.Validate(); err != nil {
			return err
		}
	}
	if s.GpuInfoDetail != nil {
		if err := s.GpuInfoDetail.Validate(); err != nil {
			return err
		}
	}
	if s.MemoryInfoDetail != nil {
		if err := s.MemoryInfoDetail.Validate(); err != nil {
			return err
		}
	}
	if s.RdmaInfoDetail != nil {
		if err := s.RdmaInfoDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMachineTypesResponseBodyMachineTypesCpuInfoDetail struct {
	// example:
	//
	// x86_64
	CpuArch *string `json:"CpuArch,omitempty" xml:"CpuArch,omitempty"`
	// example:
	//
	// 2
	CpuSockets *int32 `json:"CpuSockets,omitempty" xml:"CpuSockets,omitempty"`
	// example:
	//
	// 192
	VCpuCores *int32 `json:"VCpuCores,omitempty" xml:"VCpuCores,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) GetCpuArch() *string {
	return s.CpuArch
}

func (s *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) GetCpuSockets() *int32 {
	return s.CpuSockets
}

func (s *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) GetVCpuCores() *int32 {
	return s.VCpuCores
}

func (s *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) SetCpuArch(v string) *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail {
	s.CpuArch = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) SetCpuSockets(v int32) *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail {
	s.CpuSockets = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) SetVCpuCores(v int32) *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail {
	s.VCpuCores = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesCpuInfoDetail) Validate() error {
	return dara.Validate(s)
}

type ListMachineTypesResponseBodyMachineTypesDiskInfoDetail struct {
	// example:
	//
	// 4
	LocalDiskCount *int32 `json:"LocalDiskCount,omitempty" xml:"LocalDiskCount,omitempty"`
	// example:
	//
	// 3.84
	LocalDiskSizeInTB *float32 `json:"LocalDiskSizeInTB,omitempty" xml:"LocalDiskSizeInTB,omitempty"`
	// example:
	//
	// NVMe SSD
	LocalDiskType *string `json:"LocalDiskType,omitempty" xml:"LocalDiskType,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) GetLocalDiskCount() *int32 {
	return s.LocalDiskCount
}

func (s *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) GetLocalDiskSizeInTB() *float32 {
	return s.LocalDiskSizeInTB
}

func (s *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) GetLocalDiskType() *string {
	return s.LocalDiskType
}

func (s *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) SetLocalDiskCount(v int32) *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail {
	s.LocalDiskCount = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) SetLocalDiskSizeInTB(v float32) *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail {
	s.LocalDiskSizeInTB = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) SetLocalDiskType(v string) *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail {
	s.LocalDiskType = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesDiskInfoDetail) Validate() error {
	return dara.Validate(s)
}

type ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail struct {
	// example:
	//
	// vpc
	FrontendNetworkType *string `json:"FrontendNetworkType,omitempty" xml:"FrontendNetworkType,omitempty"`
	// example:
	//
	// true
	JumboFrameSupported *bool `json:"JumboFrameSupported,omitempty" xml:"JumboFrameSupported,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) GetFrontendNetworkType() *string {
	return s.FrontendNetworkType
}

func (s *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) GetJumboFrameSupported() *bool {
	return s.JumboFrameSupported
}

func (s *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) SetFrontendNetworkType(v string) *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail {
	s.FrontendNetworkType = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) SetJumboFrameSupported(v bool) *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail {
	s.JumboFrameSupported = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesFrontendNetworkDetail) Validate() error {
	return dara.Validate(s)
}

type ListMachineTypesResponseBodyMachineTypesGpuInfoDetail struct {
	// example:
	//
	// 8
	GpuCount *int32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// example:
	//
	// 144
	GpuMemoryInGB *int32 `json:"GpuMemoryInGB,omitempty" xml:"GpuMemoryInGB,omitempty"`
	// example:
	//
	// M890P-144G
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// example:
	//
	// T-HEAD
	GpuVendor *string `json:"GpuVendor,omitempty" xml:"GpuVendor,omitempty"`
	// example:
	//
	// 9216
	TotalGpuMemoryInGB *int32 `json:"TotalGpuMemoryInGB,omitempty" xml:"TotalGpuMemoryInGB,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) GetGpuCount() *int32 {
	return s.GpuCount
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) GetGpuMemoryInGB() *int32 {
	return s.GpuMemoryInGB
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) GetGpuName() *string {
	return s.GpuName
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) GetGpuVendor() *string {
	return s.GpuVendor
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) GetTotalGpuMemoryInGB() *int32 {
	return s.TotalGpuMemoryInGB
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) SetGpuCount(v int32) *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail {
	s.GpuCount = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) SetGpuMemoryInGB(v int32) *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail {
	s.GpuMemoryInGB = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) SetGpuName(v string) *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail {
	s.GpuName = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) SetGpuVendor(v string) *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail {
	s.GpuVendor = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) SetTotalGpuMemoryInGB(v int32) *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail {
	s.TotalGpuMemoryInGB = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesGpuInfoDetail) Validate() error {
	return dara.Validate(s)
}

type ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail struct {
	// example:
	//
	// 2048
	MemorySizeInGB *int32 `json:"MemorySizeInGB,omitempty" xml:"MemorySizeInGB,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail) GetMemorySizeInGB() *int32 {
	return s.MemorySizeInGB
}

func (s *ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail) SetMemorySizeInGB(v int32) *ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail {
	s.MemorySizeInGB = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesMemoryInfoDetail) Validate() error {
	return dara.Validate(s)
}

type ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail struct {
	// example:
	//
	// 400
	BackendRdmaNicBwInGbps *int32 `json:"BackendRdmaNicBwInGbps,omitempty" xml:"BackendRdmaNicBwInGbps,omitempty"`
	// example:
	//
	// 4
	BackendRdmaNicCount *int32 `json:"BackendRdmaNicCount,omitempty" xml:"BackendRdmaNicCount,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) GetBackendRdmaNicBwInGbps() *int32 {
	return s.BackendRdmaNicBwInGbps
}

func (s *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) GetBackendRdmaNicCount() *int32 {
	return s.BackendRdmaNicCount
}

func (s *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) SetBackendRdmaNicBwInGbps(v int32) *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail {
	s.BackendRdmaNicBwInGbps = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) SetBackendRdmaNicCount(v int32) *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail {
	s.BackendRdmaNicCount = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypesRdmaInfoDetail) Validate() error {
	return dara.Validate(s)
}
