// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePhysicalConnectionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePhysicalConnectionsResponseBody
	GetPageSize() *int32
	SetPhysicalConnectionSet(v *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) *DescribePhysicalConnectionsResponseBody
	GetPhysicalConnectionSet() *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet
	SetRequestId(v string) *DescribePhysicalConnectionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePhysicalConnectionsResponseBody
	GetTotalCount() *int32
}

type DescribePhysicalConnectionsResponseBody struct {
	PageNumber            *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhysicalConnectionSet *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet `json:"PhysicalConnectionSet,omitempty" xml:"PhysicalConnectionSet,omitempty" type:"Struct"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePhysicalConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePhysicalConnectionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePhysicalConnectionsResponseBody) GetPhysicalConnectionSet() *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet {
	return s.PhysicalConnectionSet
}

func (s *DescribePhysicalConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhysicalConnectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePhysicalConnectionsResponseBody) SetPageNumber(v int32) *DescribePhysicalConnectionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetPageSize(v int32) *DescribePhysicalConnectionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetPhysicalConnectionSet(v *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) *DescribePhysicalConnectionsResponseBody {
	s.PhysicalConnectionSet = v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetRequestId(v string) *DescribePhysicalConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) SetTotalCount(v int32) *DescribePhysicalConnectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBody) Validate() error {
	if s.PhysicalConnectionSet != nil {
		if err := s.PhysicalConnectionSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet struct {
	PhysicalConnectionType []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType `json:"PhysicalConnectionType,omitempty" xml:"PhysicalConnectionType,omitempty" type:"Repeated"`
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) GetPhysicalConnectionType() []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	return s.PhysicalConnectionType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) SetPhysicalConnectionType(v []*DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet {
	s.PhysicalConnectionType = v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSet) Validate() error {
	if s.PhysicalConnectionType != nil {
		for _, item := range s.PhysicalConnectionType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType struct {
	AccessPointId                 *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AdLocation                    *string `json:"AdLocation,omitempty" xml:"AdLocation,omitempty"`
	Bandwidth                     *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BusinessStatus                *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CircuitCode                   *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	CreationTime                  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnabledTime                   *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	LineOperator                  *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	Name                          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PeerLocation                  *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	PhysicalConnectionId          *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	PortNumber                    *string `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	PortType                      *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	Spec                          *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status                        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetAdLocation() *string {
	return s.AdLocation
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetDescription() *string {
	return s.Description
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetEnabledTime() *string {
	return s.EnabledTime
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetLineOperator() *string {
	return s.LineOperator
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetName() *string {
	return s.Name
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPortNumber() *string {
	return s.PortNumber
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetPortType() *string {
	return s.PortType
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetSpec() *string {
	return s.Spec
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetStatus() *string {
	return s.Status
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) GetType() *string {
	return s.Type
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetAccessPointId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.AccessPointId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetAdLocation(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.AdLocation = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetBandwidth(v int64) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Bandwidth = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetBusinessStatus(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.BusinessStatus = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetCircuitCode(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.CircuitCode = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetCreationTime(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.CreationTime = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetDescription(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Description = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetEnabledTime(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.EnabledTime = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetLineOperator(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.LineOperator = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetName(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Name = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPeerLocation(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PeerLocation = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPhysicalConnectionId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPortNumber(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PortNumber = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetPortType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.PortType = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetRedundantPhysicalConnectionId(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetSpec(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Spec = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetStatus(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Status = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) SetType(v string) *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType {
	s.Type = &v
	return s
}

func (s *DescribePhysicalConnectionsResponseBodyPhysicalConnectionSetPhysicalConnectionType) Validate() error {
	return dara.Validate(s)
}
