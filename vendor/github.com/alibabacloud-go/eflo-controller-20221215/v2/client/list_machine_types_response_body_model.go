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
	// The instance types.
	MachineTypes []*ListMachineTypesResponseBodyMachineTypes `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
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
	// The CPU information.
	//
	// example:
	//
	// 2x Intel Icelake 8369B 32C CPU
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// The disk information.
	//
	// example:
	//
	// 2x 480GB SATA SSD
	DiskInfo *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	// The GPU information.
	//
	// example:
	//
	// 8x NVIDIA SXM4 80GB A100 GPU
	GpuInfo *string `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	// The storage information.
	//
	// example:
	//
	// 32x 64GB DDR4 3200 Memory
	MemoryInfo *string `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	// The name of the instance type.
	//
	// example:
	//
	// efg1.nvga1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network information.
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
	NodeCount *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 48
	TotalCpuCore *int32 `json:"TotalCpuCore,omitempty" xml:"TotalCpuCore,omitempty"`
	// The access type.
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

func (s *ListMachineTypesResponseBodyMachineTypes) GetDiskInfo() *string {
	return s.DiskInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetGpuInfo() *string {
	return s.GpuInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetMemoryInfo() *string {
	return s.MemoryInfo
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

func (s *ListMachineTypesResponseBodyMachineTypes) SetDiskInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.DiskInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetGpuInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.GpuInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetMemoryInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.MemoryInfo = &v
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

func (s *ListMachineTypesResponseBodyMachineTypes) SetTotalCpuCore(v int32) *ListMachineTypesResponseBodyMachineTypes {
	s.TotalCpuCore = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetType(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.Type = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) Validate() error {
	return dara.Validate(s)
}
