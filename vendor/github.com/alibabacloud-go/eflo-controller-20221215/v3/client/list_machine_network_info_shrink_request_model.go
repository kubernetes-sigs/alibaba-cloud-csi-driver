// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineNetworkInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineHpnInfoShrink(v string) *ListMachineNetworkInfoShrinkRequest
	GetMachineHpnInfoShrink() *string
}

type ListMachineNetworkInfoShrinkRequest struct {
	// hpn information of machine
	MachineHpnInfoShrink *string `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty"`
}

func (s ListMachineNetworkInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoShrinkRequest) GetMachineHpnInfoShrink() *string {
	return s.MachineHpnInfoShrink
}

func (s *ListMachineNetworkInfoShrinkRequest) SetMachineHpnInfoShrink(v string) *ListMachineNetworkInfoShrinkRequest {
	s.MachineHpnInfoShrink = &v
	return s
}

func (s *ListMachineNetworkInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
