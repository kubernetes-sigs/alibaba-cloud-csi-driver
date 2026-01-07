// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVSwitchesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVSwitchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVSwitchesResponseBody
	GetTotalCount() *int32
	SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody
	GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches
}

type DescribeVSwitchesResponseBody struct {
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitches  *DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVSwitchesResponseBody) GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches {
	return s.VSwitches
}

func (s *DescribeVSwitchesResponseBody) SetPageNumber(v int32) *DescribeVSwitchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageSize(v int32) *DescribeVSwitchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalCount(v int32) *DescribeVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeVSwitchesResponseBody) Validate() error {
	if s.VSwitches != nil {
		if err := s.VSwitches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitches struct {
	VSwitch []*DescribeVSwitchesResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetVSwitch() []*DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	return s.VSwitch
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitch(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitch) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) Validate() error {
	if s.VSwitch != nil {
		for _, item := range s.VSwitch {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitch struct {
	AvailableIpAddressCount *int64  `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	CidrBlock               *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsDefault               *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ResourceGroupId         *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName             *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetAvailableIpAddressCount(v int64) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCreationTime(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CreationTime = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetDescription(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetResourceGroupId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVpcId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetZoneId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) Validate() error {
	return dara.Validate(s)
}
