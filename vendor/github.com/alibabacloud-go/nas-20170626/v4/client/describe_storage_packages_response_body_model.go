// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackages(v *DescribeStoragePackagesResponseBodyPackages) *DescribeStoragePackagesResponseBody
	GetPackages() *DescribeStoragePackagesResponseBodyPackages
	SetPageNumber(v int32) *DescribeStoragePackagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStoragePackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeStoragePackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeStoragePackagesResponseBody
	GetTotalCount() *int32
}

type DescribeStoragePackagesResponseBody struct {
	// The list of storage plans.
	Packages *DescribeStoragePackagesResponseBodyPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of storage plans returned per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 035B3A3A-E514-4B41-B906-5D906CFB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of storage plans.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStoragePackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBody) GetPackages() *DescribeStoragePackagesResponseBodyPackages {
	return s.Packages
}

func (s *DescribeStoragePackagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStoragePackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStoragePackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStoragePackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeStoragePackagesResponseBody) SetPackages(v *DescribeStoragePackagesResponseBodyPackages) *DescribeStoragePackagesResponseBody {
	s.Packages = v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPageNumber(v int32) *DescribeStoragePackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetPageSize(v int32) *DescribeStoragePackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetRequestId(v string) *DescribeStoragePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) SetTotalCount(v int32) *DescribeStoragePackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStoragePackagesResponseBody) Validate() error {
	if s.Packages != nil {
		if err := s.Packages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStoragePackagesResponseBodyPackages struct {
	Package []*DescribeStoragePackagesResponseBodyPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeStoragePackagesResponseBodyPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePackagesResponseBodyPackages) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBodyPackages) GetPackage() []*DescribeStoragePackagesResponseBodyPackagesPackage {
	return s.Package
}

func (s *DescribeStoragePackagesResponseBodyPackages) SetPackage(v []*DescribeStoragePackagesResponseBodyPackagesPackage) *DescribeStoragePackagesResponseBodyPackages {
	s.Package = v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackages) Validate() error {
	if s.Package != nil {
		for _, item := range s.Package {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStoragePackagesResponseBodyPackagesPackage struct {
	// The end time of the validity period for the storage plan.
	//
	// example:
	//
	// 2020-01-05T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the file system that is bound to the storage plan.
	//
	// example:
	//
	// 109c****66
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the storage plan.
	//
	// example:
	//
	// naspackage-@string(\\"*****\\", *)-@string(\\"*****\\", *)
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The capacity of the storage plan.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 10
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time of the validity period for the storage plan.
	//
	// example:
	//
	// 2019-12-05T01:40:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the storage plan.
	//
	// Valid values:
	//
	// 	- free: The storage plan is not bound to a file system. You can bind the storage plan to a file system of the same storage type.
	//
	// 	- bound: The storage plan is bound to a file system.
	//
	// example:
	//
	// free
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the storage plan.
	//
	// Valid values:
	//
	// 	- Performance
	//
	// 	- Capacity
	//
	// example:
	//
	// Capacity
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePackagesResponseBodyPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) GetSize() *int64 {
	return s.Size
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) GetStatus() *string {
	return s.Status
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetExpiredTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetFileSystemId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.FileSystemId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetPackageId(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetSize(v int64) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStartTime(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStatus(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.Status = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) SetStorageType(v string) *DescribeStoragePackagesResponseBodyPackagesPackage {
	s.StorageType = &v
	return s
}

func (s *DescribeStoragePackagesResponseBodyPackagesPackage) Validate() error {
	return dara.Validate(s)
}
