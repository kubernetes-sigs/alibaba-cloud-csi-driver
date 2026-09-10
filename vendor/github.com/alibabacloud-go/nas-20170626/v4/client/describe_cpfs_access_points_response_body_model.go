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
	AccessPoints []*DescribeCpfsAccessPointsResponseBodyAccessPoints `json:"AccessPoints,omitempty" xml:"AccessPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A323836B-5BC6-45A6-8048-60675C23****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// acs:nas:cn-hangzhou:178321033379****:accesspoint/ap-ie15yd****
	ARN *string `json:"ARN,omitempty" xml:"ARN,omitempty"`
	// example:
	//
	// ap-ie15y*****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// example:
	//
	// 2026-03-28T06:32:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// bmcpfs-290r9c75fnb0il8d8v1
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 2025-03-28T06:32:14Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId      *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RootDirectory *DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty" type:"Struct"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeCpfsAccessPointsResponseBodyAccessPoints) Validate() error {
	if s.RootDirectory != nil {
		if err := s.RootDirectory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCpfsAccessPointsResponseBodyAccessPointsRootDirectory struct {
	// example:
	//
	// /path
	RootPath *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
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
