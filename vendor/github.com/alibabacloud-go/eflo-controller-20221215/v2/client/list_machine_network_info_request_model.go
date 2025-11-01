// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineNetworkInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineHpnInfo(v []*ListMachineNetworkInfoRequestMachineHpnInfo) *ListMachineNetworkInfoRequest
	GetMachineHpnInfo() []*ListMachineNetworkInfoRequestMachineHpnInfo
}

type ListMachineNetworkInfoRequest struct {
	// hpn information of machine
	MachineHpnInfo []*ListMachineNetworkInfoRequestMachineHpnInfo `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty" type:"Repeated"`
}

func (s ListMachineNetworkInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequest) GetMachineHpnInfo() []*ListMachineNetworkInfoRequestMachineHpnInfo {
	return s.MachineHpnInfo
}

func (s *ListMachineNetworkInfoRequest) SetMachineHpnInfo(v []*ListMachineNetworkInfoRequestMachineHpnInfo) *ListMachineNetworkInfoRequest {
	s.MachineHpnInfo = v
	return s
}

func (s *ListMachineNetworkInfoRequest) Validate() error {
	if s.MachineHpnInfo != nil {
		for _, item := range s.MachineHpnInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMachineNetworkInfoRequestMachineHpnInfo struct {
	// hpn zone infomation
	//
	// example:
	//
	// C1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The type of machine.
	//
	// example:
	//
	// efg2.C48cNHmcn
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The ID of the region in which the application is located.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMachineNetworkInfoRequestMachineHpnInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoRequestMachineHpnInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) GetMachineType() *string {
	return s.MachineType
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetHpnZone(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.HpnZone = &v
	return s
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetMachineType(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.MachineType = &v
	return s
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetRegionId(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.RegionId = &v
	return s
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) Validate() error {
	return dara.Validate(s)
}
