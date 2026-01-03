// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisks(v *DescribeStorageSetDetailsResponseBodyDisks) *DescribeStorageSetDetailsResponseBody
	GetDisks() *DescribeStorageSetDetailsResponseBodyDisks
	SetPageNumber(v int32) *DescribeStorageSetDetailsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageSetDetailsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeStorageSetDetailsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeStorageSetDetailsResponseBody
	GetTotalCount() *int32
}

type DescribeStorageSetDetailsResponseBody struct {
	// Details about the disks or Shared Block Storage devices in the storage set.
	Disks *DescribeStorageSetDetailsResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
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
	// 49765E79-0D5D-4451-B3AE-580A20831846
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of storage sets.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStorageSetDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetDetailsResponseBody) GetDisks() *DescribeStorageSetDetailsResponseBodyDisks {
	return s.Disks
}

func (s *DescribeStorageSetDetailsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageSetDetailsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageSetDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStorageSetDetailsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeStorageSetDetailsResponseBody) SetDisks(v *DescribeStorageSetDetailsResponseBodyDisks) *DescribeStorageSetDetailsResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeStorageSetDetailsResponseBody) SetPageNumber(v int32) *DescribeStorageSetDetailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBody) SetPageSize(v int32) *DescribeStorageSetDetailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBody) SetRequestId(v string) *DescribeStorageSetDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBody) SetTotalCount(v int32) *DescribeStorageSetDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBody) Validate() error {
	if s.Disks != nil {
		if err := s.Disks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStorageSetDetailsResponseBodyDisks struct {
	Disk []*DescribeStorageSetDetailsResponseBodyDisksDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s DescribeStorageSetDetailsResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetDetailsResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetDetailsResponseBodyDisks) GetDisk() []*DescribeStorageSetDetailsResponseBodyDisksDisk {
	return s.Disk
}

func (s *DescribeStorageSetDetailsResponseBodyDisks) SetDisk(v []*DescribeStorageSetDetailsResponseBodyDisksDisk) *DescribeStorageSetDetailsResponseBodyDisks {
	s.Disk = v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisks) Validate() error {
	if s.Disk != nil {
		for _, item := range s.Disk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStorageSetDetailsResponseBodyDisksDisk struct {
	// The category of the disk or Shared Block Storage device.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk or Shared Block Storage device was created.
	//
	// example:
	//
	// 2019-06-01T00:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the disk or Shared Block Storage device.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk or Shared Block Storage device.
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The region to which the disk or Shared Block Storage device belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The number of partitions in the storage set.
	//
	// example:
	//
	// 3
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The zone to which the disk or Shared Block Storage device belongs.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeStorageSetDetailsResponseBodyDisksDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetDetailsResponseBodyDisksDisk) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetCategory(v string) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.Category = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetCreationTime(v string) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.CreationTime = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetDiskId(v string) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetDiskName(v string) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetRegionId(v string) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetStorageSetId(v string) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.StorageSetId = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetStorageSetPartitionNumber(v int32) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) SetZoneId(v string) *DescribeStorageSetDetailsResponseBodyDisksDisk {
	s.ZoneId = &v
	return s
}

func (s *DescribeStorageSetDetailsResponseBodyDisksDisk) Validate() error {
	return dara.Validate(s)
}
