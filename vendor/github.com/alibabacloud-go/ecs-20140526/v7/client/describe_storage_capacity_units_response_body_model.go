// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageCapacityUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeStorageCapacityUnitsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageCapacityUnitsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeStorageCapacityUnitsResponseBody
	GetRequestId() *string
	SetStorageCapacityUnits(v *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits) *DescribeStorageCapacityUnitsResponseBody
	GetStorageCapacityUnits() *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits
	SetTotalCount(v int32) *DescribeStorageCapacityUnitsResponseBody
	GetTotalCount() *int32
}

type DescribeStorageCapacityUnitsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageCapacityUnits *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits `json:"StorageCapacityUnits,omitempty" xml:"StorageCapacityUnits,omitempty" type:"Struct"`
	// The total number of SCUs.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStorageCapacityUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageCapacityUnitsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageCapacityUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStorageCapacityUnitsResponseBody) GetStorageCapacityUnits() *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits {
	return s.StorageCapacityUnits
}

func (s *DescribeStorageCapacityUnitsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeStorageCapacityUnitsResponseBody) SetPageNumber(v int32) *DescribeStorageCapacityUnitsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBody) SetPageSize(v int32) *DescribeStorageCapacityUnitsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBody) SetRequestId(v string) *DescribeStorageCapacityUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBody) SetStorageCapacityUnits(v *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits) *DescribeStorageCapacityUnitsResponseBody {
	s.StorageCapacityUnits = v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBody) SetTotalCount(v int32) *DescribeStorageCapacityUnitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBody) Validate() error {
	if s.StorageCapacityUnits != nil {
		if err := s.StorageCapacityUnits.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits struct {
	StorageCapacityUnit []*DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit `json:"StorageCapacityUnit,omitempty" xml:"StorageCapacityUnit,omitempty" type:"Repeated"`
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits) GetStorageCapacityUnit() []*DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	return s.StorageCapacityUnit
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits) SetStorageCapacityUnit(v []*DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits {
	s.StorageCapacityUnit = v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnits) Validate() error {
	if s.StorageCapacityUnit != nil {
		for _, item := range s.StorageCapacityUnit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit struct {
	AllocationStatus      *string                                                                              `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	Capacity              *int32                                                                               `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreationTime          *string                                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description           *string                                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime           *string                                                                              `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Name                  *string                                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId              *string                                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime             *string                                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageCapacityUnitId *string                                                                              `json:"StorageCapacityUnitId,omitempty" xml:"StorageCapacityUnitId,omitempty"`
	Tags                  *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetAllocationStatus() *string {
	return s.AllocationStatus
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetCapacity() *int32 {
	return s.Capacity
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetDescription() *string {
	return s.Description
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetName() *string {
	return s.Name
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetStatus() *string {
	return s.Status
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetStorageCapacityUnitId() *string {
	return s.StorageCapacityUnitId
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) GetTags() *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags {
	return s.Tags
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetAllocationStatus(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetCapacity(v int32) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.Capacity = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetCreationTime(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.CreationTime = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetDescription(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.Description = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetExpiredTime(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetName(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.Name = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetRegionId(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetStartTime(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.StartTime = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetStatus(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.Status = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetStorageCapacityUnitId(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.StorageCapacityUnitId = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) SetTags(v *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit {
	s.Tags = v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnit) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags struct {
	Tag []*DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags) GetTag() []*DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag {
	return s.Tag
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags) SetTag(v []*DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags {
	s.Tag = v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) SetTagKey(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) SetTagValue(v string) *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponseBodyStorageCapacityUnitsStorageCapacityUnitTagsTag) Validate() error {
	return dara.Validate(s)
}
