// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineNetworkInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineNetworkInfo(v []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo) *ListMachineNetworkInfoResponseBody
	GetMachineNetworkInfo() []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo
	SetRequestId(v string) *ListMachineNetworkInfoResponseBody
	GetRequestId() *string
}

type ListMachineNetworkInfoResponseBody struct {
	// machine network infomation
	MachineNetworkInfo []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo `json:"MachineNetworkInfo,omitempty" xml:"MachineNetworkInfo,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMachineNetworkInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBody) GetMachineNetworkInfo() []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	return s.MachineNetworkInfo
}

func (s *ListMachineNetworkInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMachineNetworkInfoResponseBody) SetMachineNetworkInfo(v []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo) *ListMachineNetworkInfoResponseBody {
	s.MachineNetworkInfo = v
	return s
}

func (s *ListMachineNetworkInfoResponseBody) SetRequestId(v string) *ListMachineNetworkInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBody) Validate() error {
	if s.MachineNetworkInfo != nil {
		for _, item := range s.MachineNetworkInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMachineNetworkInfoResponseBodyMachineNetworkInfo struct {
	// Network of cluster
	//
	// example:
	//
	// vpc/acl
	ClusterNet *string `json:"ClusterNet,omitempty" xml:"ClusterNet,omitempty"`
	// Specifies whether to enable the Jumbo Frames feature for the instance. Valid values:
	//
	// 	- true: The Jumbo Frame feature is enabled for the instance.
	//
	// 	- false: The Jumbo Frame feature is disabled for the instance.
	//
	// Take note of the following items:
	//
	// 	- The instance must be in the Running (`Running`) or Stopped (`Stopped`) state.
	//
	// 	- The instance must reside in a VPC.
	//
	// 	- After the Jumbo Frames feature is enabled, the MTU value of the instance is set to 8500. After the Jumbo Frames feature is disabled, the MTU value of the instance is set to 1500. You can enable the Jumbo Frames feature only for specific instance types. For more information, see [Jumbo Frames](https://help.aliyun.com/document_detail/200512.html).
	//
	// example:
	//
	// true
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// HPN zone
	//
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Specifies whether dpu machine.
	//
	// example:
	//
	// true
	IsDpuMode *bool `json:"IsDpuMode,omitempty" xml:"IsDpuMode,omitempty"`
	// The type of machine.
	//
	// example:
	//
	// efg1.nvga8n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Network architecture
	//
	// example:
	//
	// XX-7.0
	NetArch *string `json:"NetArch,omitempty" xml:"NetArch,omitempty"`
	// The ID of the region in which the application is located.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMachineNetworkInfoResponseBodyMachineNetworkInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetClusterNet() *string {
	return s.ClusterNet
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetIsDpuMode() *bool {
	return s.IsDpuMode
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetMachineType() *string {
	return s.MachineType
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetNetArch() *string {
	return s.NetArch
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetClusterNet(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.ClusterNet = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetEnableJumboFrame(v bool) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.EnableJumboFrame = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetHpnZone(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.HpnZone = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetIsDpuMode(v bool) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.IsDpuMode = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetMachineType(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.MachineType = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetNetArch(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.NetArch = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetRegionId(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.RegionId = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) Validate() error {
	return dara.Validate(s)
}
