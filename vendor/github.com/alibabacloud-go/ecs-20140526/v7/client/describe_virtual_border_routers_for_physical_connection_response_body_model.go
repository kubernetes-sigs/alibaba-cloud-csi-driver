// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersForPhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetTotalCount() *int32
	SetVirtualBorderRouterForPhysicalConnectionSet(v *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody
	GetVirtualBorderRouterForPhysicalConnectionSet() *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody struct {
	PageNumber                                  *int32                                                                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                                    *int32                                                                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                                   *string                                                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                                  *int32                                                                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VirtualBorderRouterForPhysicalConnectionSet *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet `json:"VirtualBorderRouterForPhysicalConnectionSet,omitempty" xml:"VirtualBorderRouterForPhysicalConnectionSet,omitempty" type:"Struct"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) GetVirtualBorderRouterForPhysicalConnectionSet() *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet {
	return s.VirtualBorderRouterForPhysicalConnectionSet
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetRequestId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetTotalCount(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) SetVirtualBorderRouterForPhysicalConnectionSet(v *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody {
	s.VirtualBorderRouterForPhysicalConnectionSet = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBody) Validate() error {
	if s.VirtualBorderRouterForPhysicalConnectionSet != nil {
		if err := s.VirtualBorderRouterForPhysicalConnectionSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet struct {
	VirtualBorderRouterForPhysicalConnectionType []*DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType `json:"VirtualBorderRouterForPhysicalConnectionType,omitempty" xml:"VirtualBorderRouterForPhysicalConnectionType,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) GetVirtualBorderRouterForPhysicalConnectionType() []*DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	return s.VirtualBorderRouterForPhysicalConnectionType
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) SetVirtualBorderRouterForPhysicalConnectionType(v []*DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet {
	s.VirtualBorderRouterForPhysicalConnectionType = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSet) Validate() error {
	if s.VirtualBorderRouterForPhysicalConnectionType != nil {
		for _, item := range s.VirtualBorderRouterForPhysicalConnectionType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType struct {
	ActivationTime  *string `json:"ActivationTime,omitempty" xml:"ActivationTime,omitempty"`
	CircuitCode     *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	RecoveryTime    *string `json:"RecoveryTime,omitempty" xml:"RecoveryTime,omitempty"`
	TerminationTime *string `json:"TerminationTime,omitempty" xml:"TerminationTime,omitempty"`
	VbrId           *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	VbrOwnerUid     *int64  `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	VlanId          *int32  `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetActivationTime() *string {
	return s.ActivationTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetRecoveryTime() *string {
	return s.RecoveryTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetTerminationTime() *string {
	return s.TerminationTime
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetVbrId() *string {
	return s.VbrId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetVbrOwnerUid() *int64 {
	return s.VbrOwnerUid
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) GetVlanId() *int32 {
	return s.VlanId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetActivationTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.ActivationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetCircuitCode(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.CircuitCode = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetCreationTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.CreationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetRecoveryTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.RecoveryTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetTerminationTime(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.TerminationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetVbrId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.VbrId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetVbrOwnerUid(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.VbrOwnerUid = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) SetVlanId(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType {
	s.VlanId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionResponseBodyVirtualBorderRouterForPhysicalConnectionSetVirtualBorderRouterForPhysicalConnectionType) Validate() error {
	return dara.Validate(s)
}
