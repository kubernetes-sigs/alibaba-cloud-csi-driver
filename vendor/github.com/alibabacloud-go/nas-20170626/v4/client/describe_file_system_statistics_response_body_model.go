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
	// The statistics of file systems.
	FileSystemStatistics *DescribeFileSystemStatisticsResponseBodyFileSystemStatistics `json:"FileSystemStatistics,omitempty" xml:"FileSystemStatistics,omitempty" type:"Struct"`
	// The queried file systems.
	FileSystems *DescribeFileSystemStatisticsResponseBodyFileSystems `json:"FileSystems,omitempty" xml:"FileSystems,omitempty" type:"Struct"`
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
	// The number of expired file systems.
	//
	// example:
	//
	// 1
	ExpiredCount *int32 `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	// The number of expiring file systems.
	//
	// File systems whose expiration time is less than or equal to seven days away from the current time are counted.
	//
	// example:
	//
	// 1
	ExpiringCount *int32 `json:"ExpiringCount,omitempty" xml:"ExpiringCount,omitempty"`
	// The type of the file system.
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The storage usage of the file system.
	//
	// The value of this parameter is the maximum storage usage of the file system over the last hour.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 1611
	MeteredSize *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	// The number of file systems of the current type.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The capacity of the file system.
	//
	// Unit: GiB.
	//
	// example:
	//
	// 1
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- Subscription: The subscription billing method is used.
	//
	// 	- PayAsYouGo: The pay-as-you-go billing method is used.
	//
	// 	- Package: A storage plan is attached to the file system.
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the NAS file system was created.
	//
	// example:
	//
	// 2017-05-27T15:43:06CST
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the file system.
	//
	// example:
	//
	// 31a8e48eda
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the file system expires.
	//
	// example:
	//
	// 2017-08-27T15:43:06CST
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 109c04****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard: General-purpose NAS file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// 	- cpfs: CPFS file system
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The storage usage of the Infrequent Access (IA) storage medium.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 362832
	MeteredIASize *int64 `json:"MeteredIASize,omitempty" xml:"MeteredIASize,omitempty"`
	// The storage usage of the file system.
	//
	// The value of this parameter is the maximum storage usage of the file system over the last hour. Unit: bytes.
	//
	// example:
	//
	// 1611661312
	MeteredSize *int64 `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	// The information about storage plans.
	Packages *DescribeFileSystemStatisticsResponseBodyFileSystemsFileSystemPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Struct"`
	// The protocol type of the file system.
	//
	// Valid values:
	//
	// 	- NFS: Network File System (NFS)
	//
	// 	- SMB: Server Message Block (SMB)
	//
	// 	- cpfs: the protocol type supported by the CPFS file system
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the file system.
	//
	// This parameter is returned for Extreme NAS file systems and Cloud Parallel File Storage (CPFS) file systems. Valid values:
	//
	// 	- Pending: The file system is being created or modified.
	//
	// 	- Running: The file system is available. Before you create a mount target for the file system, make sure that the file system is in the Running state.
	//
	// 	- Stopped: The file system is unavailable.
	//
	// 	- Extending: The file system is being scaled out.
	//
	// 	- Stopping: The file system is being disabled.
	//
	// 	- Deleting: The file system is being deleted.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage type.
	//
	// Valid values:
	//
	// 	- Valid values for General-purpose NAS file systems: Capacity and Performance.
	//
	// 	- Valid values for Extreme NAS file systems: standard and advance.
	//
	// 	- Valid values for CPFS file systems: advance_100 (100 MB/s/TiB baseline) and advance_200 (200 MB/s/TiB baseline).
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// The end time of the validity period for the storage plan.
	//
	// example:
	//
	// 2020-01-05T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the storage plan.
	//
	// example:
	//
	// naspackage-0be9c4b624-37****
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The capacity of the storage plan.
	//
	// example:
	//
	// 107374182400
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The start time of the validity period for the storage plan.
	//
	// example:
	//
	// 2019-12-05T01:40:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
