// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCpfsAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoints(v []*DescribeCpfsAccessPointsResponseBodyAccessPoints) *DescribeCpfsAccessPointsResponseBody
	GetAccessPoints() []*DescribeCpfsAccessPointsResponseBodyAccessPoints
	SetPageNumber(v int32) *DescribeCpfsAccessPointsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCpfsAccessPointsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCpfsAccessPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCpfsAccessPointsResponseBody
	GetTotalCount() *int32
}

type DescribeCpfsAccessPointsResponseBody struct {
	// The access point information.
	AccessPoints []*DescribeCpfsAccessPointsResponseBodyAccessPoints `json:"AccessPoints,omitempty" xml:"AccessPoints,omitempty" type:"Repeated"`
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of results per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A323836B-5BC6-45A6-8048-60675C23****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of access points.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCpfsAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointsResponseBody) GetAccessPoints() []*DescribeCpfsAccessPointsResponseBodyAccessPoints {
	return s.AccessPoints
}

func (s *DescribeCpfsAccessPointsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCpfsAccessPointsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCpfsAccessPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCpfsAccessPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCpfsAccessPointsResponseBody) SetAccessPoints(v []*DescribeCpfsAccessPointsResponseBodyAccessPoints) *DescribeCpfsAccessPointsResponseBody {
	s.AccessPoints = v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBody) SetPageNumber(v int32) *DescribeCpfsAccessPointsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBody) SetPageSize(v int32) *DescribeCpfsAccessPointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBody) SetRequestId(v string) *DescribeCpfsAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBody) SetTotalCount(v int32) *DescribeCpfsAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBody) Validate() error {
	if s.AccessPoints != nil {
		for _, item := range s.AccessPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCpfsAccessPointsResponseBodyAccessPoints struct {
	// The Alibaba Cloud Resource Name (ARN) of the access point.
	//
	// example:
	//
	// acs:nas:cn-hangzhou:178321033379****:accesspoint/ap-ie15yd****
	ARN *string `json:"ARN,omitempty" xml:"ARN,omitempty"`
	// The access point ID.
	//
	// example:
	//
	// ap-ie15y*****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The time when the access point was created. The time follows the ISO 8601 standard in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2026-03-28T06:32:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access point.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file system ID.
	//
	// example:
	//
	// bmcpfs-290r9c75fnb0il8d8v1
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The time when the access point was last modified. The time follows the ISO 8601 standard in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-03-28T06:32:14Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The root directory information.
	RootDirectory *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty" type:"Struct"`
	// The current status of the access point.
	//
	// Valid values:
	//
	// - Active: available.
	//
	// - Inactive: unavailable.
	//
	// - Pending: being created.
	//
	// - Deleting: being deleted.
	//
	// > You can mount the file system only when the status is Active.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of CPFS access point tags.
	Tags []*DescribeCpfsAccessPointsResponseBodyAccessPointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeCpfsAccessPointsResponseBodyAccessPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointsResponseBodyAccessPoints) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetARN() *string {
	return s.ARN
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetDescription() *string {
	return s.Description
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetRootDirectory() *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory {
	return s.RootDirectory
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetStatus() *string {
	return s.Status
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) GetTags() []*DescribeCpfsAccessPointsResponseBodyAccessPointsTags {
	return s.Tags
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetARN(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.ARN = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetAccessPointId(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.AccessPointId = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetCreateTime(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.CreateTime = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetDescription(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.Description = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetFileSystemId(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.FileSystemId = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetModifyTime(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.ModifyTime = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetRegionId(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.RegionId = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetRootDirectory(v *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.RootDirectory = v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetStatus(v string) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.Status = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) SetTags(v []*DescribeCpfsAccessPointsResponseBodyAccessPointsTags) *DescribeCpfsAccessPointsResponseBodyAccessPoints {
	s.Tags = v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) Validate() error {
	if s.RootDirectory != nil {
		if err := s.RootDirectory.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory struct {
	// The root directory.
	//
	// example:
	//
	// /path
	RootPath *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
	// The current status of the root directory.
	//
	// Valid values:
	//
	// - Unknown: the root path status is unknown.
	//
	// - NotExist: the root path does not exist. It may have been deleted by the user.
	//
	// - Ready: the root path status is normal.
	//
	// example:
	//
	// Ready
	RootPathStatus *string `json:"RootPathStatus,omitempty" xml:"RootPathStatus,omitempty"`
}

func (s DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) GetRootPath() *string {
	return s.RootPath
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) GetRootPathStatus() *string {
	return s.RootPathStatus
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) SetRootPath(v string) *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory {
	s.RootPath = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) SetRootPathStatus(v string) *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory {
	s.RootPathStatus = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory) Validate() error {
	return dara.Validate(s)
}

type DescribeCpfsAccessPointsResponseBodyAccessPointsTags struct {
	// The key of the CPFS access point tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the CPFS access point tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCpfsAccessPointsResponseBodyAccessPointsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointsResponseBodyAccessPointsTags) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsTags) SetKey(v string) *DescribeCpfsAccessPointsResponseBodyAccessPointsTags {
	s.Key = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsTags) SetValue(v string) *DescribeCpfsAccessPointsResponseBodyAccessPointsTags {
	s.Value = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponseBodyAccessPointsTags) Validate() error {
	return dara.Validate(s)
}
