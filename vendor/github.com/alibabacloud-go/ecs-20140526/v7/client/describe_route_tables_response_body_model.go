// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouteTablesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouteTablesResponseBody
	GetRequestId() *string
	SetRouteTables(v *DescribeRouteTablesResponseBodyRouteTables) *DescribeRouteTablesResponseBody
	GetRouteTables() *DescribeRouteTablesResponseBodyRouteTables
	SetTotalCount(v int32) *DescribeRouteTablesResponseBody
	GetTotalCount() *int32
}

type DescribeRouteTablesResponseBody struct {
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteTables *DescribeRouteTablesResponseBodyRouteTables `json:"RouteTables,omitempty" xml:"RouteTables,omitempty" type:"Struct"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteTablesResponseBody) GetRouteTables() *DescribeRouteTablesResponseBodyRouteTables {
	return s.RouteTables
}

func (s *DescribeRouteTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouteTablesResponseBody) SetPageNumber(v int32) *DescribeRouteTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetPageSize(v int32) *DescribeRouteTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetRequestId(v string) *DescribeRouteTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetRouteTables(v *DescribeRouteTablesResponseBodyRouteTables) *DescribeRouteTablesResponseBody {
	s.RouteTables = v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetTotalCount(v int32) *DescribeRouteTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) Validate() error {
	if s.RouteTables != nil {
		if err := s.RouteTables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTables struct {
	RouteTable []*DescribeRouteTablesResponseBodyRouteTablesRouteTable `json:"RouteTable,omitempty" xml:"RouteTable,omitempty" type:"Repeated"`
}

func (s DescribeRouteTablesResponseBodyRouteTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTables) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTables) GetRouteTable() []*DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	return s.RouteTable
}

func (s *DescribeRouteTablesResponseBodyRouteTables) SetRouteTable(v []*DescribeRouteTablesResponseBodyRouteTablesRouteTable) *DescribeRouteTablesResponseBodyRouteTables {
	s.RouteTable = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTables) Validate() error {
	if s.RouteTable != nil {
		for _, item := range s.RouteTable {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTable struct {
	CreationTime    *string                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ResourceGroupId *string                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RouteEntrys     *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys `json:"RouteEntrys,omitempty" xml:"RouteEntrys,omitempty" type:"Struct"`
	RouteTableId    *string                                                          `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	RouteTableType  *string                                                          `json:"RouteTableType,omitempty" xml:"RouteTableType,omitempty"`
	VRouterId       *string                                                          `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTable) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetRouteEntrys() *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys {
	return s.RouteEntrys
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetCreationTime(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.CreationTime = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetResourceGroupId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetRouteEntrys(v *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.RouteEntrys = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetRouteTableId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetRouteTableType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.RouteTableType = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetVRouterId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.VRouterId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) Validate() error {
	if s.RouteEntrys != nil {
		if err := s.RouteEntrys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys struct {
	RouteEntry []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry `json:"RouteEntry,omitempty" xml:"RouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) GetRouteEntry() []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	return s.RouteEntry
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) SetRouteEntry(v []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys {
	s.RouteEntry = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) Validate() error {
	if s.RouteEntry != nil {
		for _, item := range s.RouteEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry struct {
	DestinationCidrBlock *string                                                                            `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	InstanceId           *string                                                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NextHopType          *string                                                                            `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	NextHops             *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops `json:"NextHops,omitempty" xml:"NextHops,omitempty" type:"Struct"`
	RouteTableId         *string                                                                            `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	Status               *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                 *string                                                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetNextHops() *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops {
	return s.NextHops
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetType() *string {
	return s.Type
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetDestinationCidrBlock(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetInstanceId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetNextHopType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetNextHops(v *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.NextHops = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetRouteTableId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetStatus(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) Validate() error {
	if s.NextHops != nil {
		if err := s.NextHops.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops struct {
	NextHop []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop `json:"NextHop,omitempty" xml:"NextHop,omitempty" type:"Repeated"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) GetNextHop() []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	return s.NextHop
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) SetNextHop(v []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops {
	s.NextHop = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) Validate() error {
	if s.NextHop != nil {
		for _, item := range s.NextHop {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop struct {
	Enabled     *int32  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	NextHopId   *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	Weight      *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetEnabled() *int32 {
	return s.Enabled
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetEnabled(v int32) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.Enabled = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetNextHopId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetNextHopType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetWeight(v int32) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.Weight = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) Validate() error {
	return dara.Validate(s)
}
