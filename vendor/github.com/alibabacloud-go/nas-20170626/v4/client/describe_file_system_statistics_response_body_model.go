// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemStatistics(v *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) *DescribeFileSystemStatisticsResponseBody
	GetFileSystemStatistics() *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics
	SetFileSystems(v *DescribeFileSystemStatisticsResponseBodyFileSystems) *DescribeFileSystemStatisticsResponseBody
	GetFileSystems() *DescribeFileSystemStatisticsResponseBodyFileSystems
	SetPageNumber(v int32) *DescribeFileSystemStatisticsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFileSystemStatisticsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFileSystemStatisticsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBody
	GetTotalCount() *int32
}

type DescribeFileSystemStatisticsResponseBody struct {
	FileSystemStatistics *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics `json:"FileSystemStatistics,omitempty" xml:"FileSystemStatistics,omitempty" type:"Struct"`
	FileSystems          *DescribeFileSystemStatisticsResponseBodyFileSystems          `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
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
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9F088138-FD73-4B68-95CC-DFAD4D85****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of file system entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBody) GetFileSystemStatistics() *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics {
	return s.FileSystemStatistics
}

func (s *DescribeFileSystemStatisticsResponseBody) GetFileSystems() *DescribeFileSystemStatisticsResponseBodyFileSystems {
	return s.FileSystems
}

func (s *DescribeFileSystemStatisticsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFileSystemStatisticsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileSystemStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFileSystemStatisticsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystemStatistics(v *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystemStatistics = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetFileSystems(v *DescribeFileSystemStatisticsResponseBodyFileSystems) *DescribeFileSystemStatisticsResponseBody {
	s.FileSystems = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageNumber(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetPageSize(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetRequestId(v string) *DescribeFileSystemStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBody) Validate() error {
	if s.FileSystemStatistics != nil {
		if err := s.FileSystemStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.FileSystems != nil {
		if err := s.FileSystems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFileSystemStatisticsResponseBodyFileSystemStatistics struct {
	FileSystemStatistic []*DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic `json:"FileSystemStatistic,omitempty" xml:"FileSystemStatistic,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) GetFileSystemStatistic() []*DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	return s.FileSystemStatistic
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) SetFileSystemStatistic(v []*DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics {
	s.FileSystemStatistic = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics) Validate() error {
	if s.FileSystemStatistic != nil {
		for _, item := range s.FileSystemStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic struct {
	ExpiredCount   *int32  `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	ExpiringCount  *int32  `json:"ExpiringCount,omitempty" xml:"ExpiringCount,omitempty"`
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredSize    *int64  `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	TotalCount     *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GetExpiredCount() *int32 {
	return s.ExpiredCount
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GetExpiringCount() *int32 {
	return s.ExpiringCount
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GetMeteredSize() *int64 {
	return s.MeteredSize
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetExpiredCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.ExpiredCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetExpiringCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.ExpiringCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetFileSystemType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetMeteredSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) SetTotalCount(v int32) *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic {
	s.TotalCount = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemStatisticsFileSystemStatistic) Validate() error {
	return dara.Validate(s)
}

type DescribeFileSystemStatisticsResponseBodyFileSystems struct {
	FileSystem []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem `json:"FileSystem,omitempty" xml:"FileSystem,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystems) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystems) GetFileSystem() []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	return s.FileSystem
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystems) SetFileSystem(v []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) *DescribeFileSystemStatisticsResponseBodyFileSystems {
	s.FileSystem = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystems) Validate() error {
	if s.FileSystem != nil {
		for _, item := range s.FileSystem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem struct {
	Capacity       *int64                                                                 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType     *string                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime     *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime    *string                                                                `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FileSystemId   *string                                                                `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemType *string                                                                `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	MeteredIASize  *int64                                                                 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	MeteredSize    *int64                                                                 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	Packages       *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	ProtocolType   *string                                                                `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RegionId       *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status         *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType    *string                                                                `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	ZoneId         *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetDescription() *string {
	return s.Description
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetMeteredIASize() *int64 {
	return s.MeteredIASize
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetMeteredSize() *int64 {
	return s.MeteredSize
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetPackages() *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages {
	return s.Packages
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetStatus() *string {
	return s.Status
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCapacity(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Capacity = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetChargeType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ChargeType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetCreateTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetDescription(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Description = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetExpiredTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetFileSystemType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.FileSystemType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredIASize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredIASize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetMeteredSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.MeteredSize = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetPackages(v *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Packages = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetProtocolType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ProtocolType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetRegionId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.RegionId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStatus(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.Status = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetStorageType(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.StorageType = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) SetZoneId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem {
	s.ZoneId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystem) Validate() error {
	if s.Packages != nil {
		if err := s.Packages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages struct {
	Package []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage `json:"Package,omitempty" xml:"Package,omitempty" type:"Repeated"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) GetPackage() []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	return s.Package
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) SetPackage(v []*DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages {
	s.Package = v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages) Validate() error {
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

type DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage struct {
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GetSize() *int64 {
	return s.Size
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetExpiredTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetPackageId(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.PackageId = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetSize(v int64) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.Size = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) SetStartTime(v string) *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage {
	s.StartTime = &v
	return s
}

func (s *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackagesPackage) Validate() error {
	return dara.Validate(s)
}
