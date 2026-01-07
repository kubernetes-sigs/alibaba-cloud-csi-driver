// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSnapshotPackageResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotPackageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSnapshotPackageResponseBody
	GetRequestId() *string
	SetSnapshotPackages(v *DescribeSnapshotPackageResponseBodySnapshotPackages) *DescribeSnapshotPackageResponseBody
	GetSnapshotPackages() *DescribeSnapshotPackageResponseBodySnapshotPackages
	SetTotalCount(v int32) *DescribeSnapshotPackageResponseBody
	GetTotalCount() *int32
}

type DescribeSnapshotPackageResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the OSS storage plans.
	SnapshotPackages *DescribeSnapshotPackageResponseBodySnapshotPackages `json:"SnapshotPackages,omitempty" xml:"SnapshotPackages,omitempty" type:"Struct"`
	// The total number of OSS storage plans.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnapshotPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotPackageResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotPackageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotPackageResponseBody) GetSnapshotPackages() *DescribeSnapshotPackageResponseBodySnapshotPackages {
	return s.SnapshotPackages
}

func (s *DescribeSnapshotPackageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnapshotPackageResponseBody) SetPageNumber(v int32) *DescribeSnapshotPackageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBody) SetPageSize(v int32) *DescribeSnapshotPackageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBody) SetRequestId(v string) *DescribeSnapshotPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBody) SetSnapshotPackages(v *DescribeSnapshotPackageResponseBodySnapshotPackages) *DescribeSnapshotPackageResponseBody {
	s.SnapshotPackages = v
	return s
}

func (s *DescribeSnapshotPackageResponseBody) SetTotalCount(v int32) *DescribeSnapshotPackageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBody) Validate() error {
	if s.SnapshotPackages != nil {
		if err := s.SnapshotPackages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotPackageResponseBodySnapshotPackages struct {
	SnapshotPackage []*DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage `json:"SnapshotPackage,omitempty" xml:"SnapshotPackage,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotPackageResponseBodySnapshotPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotPackageResponseBodySnapshotPackages) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackages) GetSnapshotPackage() []*DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage {
	return s.SnapshotPackage
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackages) SetSnapshotPackage(v []*DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) *DescribeSnapshotPackageResponseBodySnapshotPackages {
	s.SnapshotPackage = v
	return s
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackages) Validate() error {
	if s.SnapshotPackage != nil {
		for _, item := range s.SnapshotPackage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage struct {
	// The name of the OSS storage plan.
	//
	// example:
	//
	// testDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the OSS storage plan expires. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-30T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum storage capacity offered by the OSS storage plan.
	//
	// example:
	//
	// 500
	InitCapacity *int64 `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	// The time when the OSS storage plan was purchased. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) GetInitCapacity() *int64 {
	return s.InitCapacity
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) SetDisplayName(v string) *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage {
	s.DisplayName = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) SetEndTime(v string) *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage {
	s.EndTime = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) SetInitCapacity(v int64) *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage {
	s.InitCapacity = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) SetStartTime(v string) *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeSnapshotPackageResponseBodySnapshotPackagesSnapshotPackage) Validate() error {
	return dara.Validate(s)
}
