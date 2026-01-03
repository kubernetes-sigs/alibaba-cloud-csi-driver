// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeStorageSetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageSetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeStorageSetsResponseBody
	GetRequestId() *string
	SetStorageSets(v *DescribeStorageSetsResponseBodyStorageSets) *DescribeStorageSetsResponseBody
	GetStorageSets() *DescribeStorageSetsResponseBodyStorageSets
	SetTotalCount(v int32) *DescribeStorageSetsResponseBody
	GetTotalCount() *int32
}

type DescribeStorageSetsResponseBody struct {
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
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the storage sets. The value of this parameter is an array that consists of StorageSet data.
	StorageSets *DescribeStorageSetsResponseBodyStorageSets `json:"StorageSets,omitempty" xml:"StorageSets,omitempty" type:"Struct"`
	// The total number of storage sets.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStorageSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageSetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStorageSetsResponseBody) GetStorageSets() *DescribeStorageSetsResponseBodyStorageSets {
	return s.StorageSets
}

func (s *DescribeStorageSetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeStorageSetsResponseBody) SetPageNumber(v int32) *DescribeStorageSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageSetsResponseBody) SetPageSize(v int32) *DescribeStorageSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageSetsResponseBody) SetRequestId(v string) *DescribeStorageSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageSetsResponseBody) SetStorageSets(v *DescribeStorageSetsResponseBodyStorageSets) *DescribeStorageSetsResponseBody {
	s.StorageSets = v
	return s
}

func (s *DescribeStorageSetsResponseBody) SetTotalCount(v int32) *DescribeStorageSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStorageSetsResponseBody) Validate() error {
	if s.StorageSets != nil {
		if err := s.StorageSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStorageSetsResponseBodyStorageSets struct {
	StorageSet []*DescribeStorageSetsResponseBodyStorageSetsStorageSet `json:"StorageSet,omitempty" xml:"StorageSet,omitempty" type:"Repeated"`
}

func (s DescribeStorageSetsResponseBodyStorageSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsResponseBodyStorageSets) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsResponseBodyStorageSets) GetStorageSet() []*DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	return s.StorageSet
}

func (s *DescribeStorageSetsResponseBodyStorageSets) SetStorageSet(v []*DescribeStorageSetsResponseBodyStorageSetsStorageSet) *DescribeStorageSetsResponseBodyStorageSets {
	s.StorageSet = v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSets) Validate() error {
	if s.StorageSet != nil {
		for _, item := range s.StorageSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStorageSetsResponseBodyStorageSetsStorageSet struct {
	// The time when the storage set was created.
	//
	// example:
	//
	// 2019-06-01T00:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the storage set.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the region to which the storage set belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp1d6tsvznfghy7y****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The name of the storage set.
	//
	// example:
	//
	// testStorageSetName
	StorageSetName *string `json:"StorageSetName,omitempty" xml:"StorageSetName,omitempty"`
	// The maximum number of partitions supported by the storage set.
	//
	// example:
	//
	// 3
	StorageSetPartitionNumber *int32                                                    `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	Tags                      *DescribeStorageSetsResponseBodyStorageSetsStorageSetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the zone to which the storage set belongs.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeStorageSetsResponseBodyStorageSetsStorageSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsResponseBodyStorageSetsStorageSet) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetDescription() *string {
	return s.Description
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetStorageSetName() *string {
	return s.StorageSetName
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetTags() *DescribeStorageSetsResponseBodyStorageSetsStorageSetTags {
	return s.Tags
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetCreationTime(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetDescription(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.Description = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetRegionId(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetResourceGroupId(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetStorageSetId(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.StorageSetId = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetStorageSetName(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.StorageSetName = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetStorageSetPartitionNumber(v int32) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetTags(v *DescribeStorageSetsResponseBodyStorageSetsStorageSetTags) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.Tags = v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) SetZoneId(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSet {
	s.ZoneId = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSet) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStorageSetsResponseBodyStorageSetsStorageSetTags struct {
	Tag []*DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeStorageSetsResponseBodyStorageSetsStorageSetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsResponseBodyStorageSetsStorageSetTags) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTags) GetTag() []*DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag {
	return s.Tag
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTags) SetTag(v []*DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) *DescribeStorageSetsResponseBodyStorageSetsStorageSetTags {
	s.Tag = v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTags) Validate() error {
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

type DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) SetKey(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) SetValue(v string) *DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeStorageSetsResponseBodyStorageSetsStorageSetTagsTag) Validate() error {
	return dara.Validate(s)
}
