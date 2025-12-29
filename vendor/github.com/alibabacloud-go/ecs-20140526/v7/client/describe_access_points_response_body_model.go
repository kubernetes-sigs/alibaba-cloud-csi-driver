// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointSet(v *DescribeAccessPointsResponseBodyAccessPointSet) *DescribeAccessPointsResponseBody
	GetAccessPointSet() *DescribeAccessPointsResponseBodyAccessPointSet
	SetPageNumber(v int32) *DescribeAccessPointsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessPointsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccessPointsResponseBody
	GetTotalCount() *int32
}

type DescribeAccessPointsResponseBody struct {
	AccessPointSet *DescribeAccessPointsResponseBodyAccessPointSet `json:"AccessPointSet,omitempty" xml:"AccessPointSet,omitempty" type:"Struct"`
	PageNumber     *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBody) GetAccessPointSet() *DescribeAccessPointsResponseBodyAccessPointSet {
	return s.AccessPointSet
}

func (s *DescribeAccessPointsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessPointsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccessPointsResponseBody) SetAccessPointSet(v *DescribeAccessPointsResponseBodyAccessPointSet) *DescribeAccessPointsResponseBody {
	s.AccessPointSet = v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetPageNumber(v int32) *DescribeAccessPointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetPageSize(v int32) *DescribeAccessPointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetRequestId(v string) *DescribeAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetTotalCount(v int32) *DescribeAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) Validate() error {
	if s.AccessPointSet != nil {
		if err := s.AccessPointSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessPointsResponseBodyAccessPointSet struct {
	AccessPointType []*DescribeAccessPointsResponseBodyAccessPointSetAccessPointType `json:"AccessPointType,omitempty" xml:"AccessPointType,omitempty" type:"Repeated"`
}

func (s DescribeAccessPointsResponseBodyAccessPointSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPointSet) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPointSet) GetAccessPointType() []*DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	return s.AccessPointType
}

func (s *DescribeAccessPointsResponseBodyAccessPointSet) SetAccessPointType(v []*DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) *DescribeAccessPointsResponseBodyAccessPointSet {
	s.AccessPointType = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSet) Validate() error {
	if s.AccessPointType != nil {
		for _, item := range s.AccessPointType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessPointsResponseBodyAccessPointSetAccessPointType struct {
	AccessPointId    *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AttachedRegionNo *string `json:"AttachedRegionNo,omitempty" xml:"AttachedRegionNo,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HostOperator     *string `json:"HostOperator,omitempty" xml:"HostOperator,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetAttachedRegionNo() *string {
	return s.AttachedRegionNo
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetDescription() *string {
	return s.Description
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetHostOperator() *string {
	return s.HostOperator
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetLocation() *string {
	return s.Location
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetName() *string {
	return s.Name
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetStatus() *string {
	return s.Status
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) GetType() *string {
	return s.Type
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetAccessPointId(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.AccessPointId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetAttachedRegionNo(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.AttachedRegionNo = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetDescription(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.Description = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetHostOperator(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.HostOperator = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetLocation(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.Location = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetName(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.Name = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetStatus(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.Status = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) SetType(v string) *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType {
	s.Type = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointSetAccessPointType) Validate() error {
	return dara.Validate(s)
}
