// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSnapshotsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v *DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody
	GetSnapshots() *DescribeSnapshotsResponseBodySnapshots
	SetTotalCount(v int32) *DescribeSnapshotsResponseBody
	GetTotalCount() *int32
}

type DescribeSnapshotsResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about snapshots.
	Snapshots *DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The total number of snapshots returned.
	//
	// example:
	//
	// 36
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotsResponseBody) GetSnapshots() *DescribeSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeSnapshotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnapshotsResponseBody) SetPageNumber(v int32) *DescribeSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetPageSize(v int32) *DescribeSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v *DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetTotalCount(v int32) *DescribeSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSnapshotsResponseBodySnapshots struct {
	Snapshot []*DescribeSnapshotsResponseBodySnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshot() []*DescribeSnapshotsResponseBodySnapshotsSnapshot {
	return s.Snapshot
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshot(v []*DescribeSnapshotsResponseBodySnapshotsSnapshot) *DescribeSnapshotsResponseBodySnapshots {
	s.Snapshot = v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}

type DescribeSnapshotsResponseBodySnapshotsSnapshot struct {
	// The time when snapshot creation was complete.
	//
	// The time follows the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard in UTC. The time is displayed in the `yyyy-MM-ddThh:mmZ` format.
	//
	// >  This parameter is valid only when the snapshot is created. During snapshot creation, the value of this parameter is the same as that of CreateTime.
	//
	// example:
	//
	// 2014-07-24T13:10:52Z
	CompletedTime *string `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	// The time when the snapshot was created.
	//
	// The time follows the [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard in UTC. The time is displayed in the `yyyy-MM-ddThh:mmZ` format.
	//
	// example:
	//
	// 2014-07-24T13:00:52Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the snapshot.
	//
	// example:
	//
	// FinanceDept
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the snapshot is encrypted.
	//
	// Valid values:
	//
	// 	- 0: The snapshot is not encrypted.
	//
	// 	- 1: The snapshot is encrypted.
	//
	// example:
	//
	// 1
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The type of the file system.
	//
	// example:
	//
	// extreme
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The progress of the snapshot creation. The value of this parameter is expressed as a percentage.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The remaining time that is required to create the snapshot.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 38
	RemainTime *int32 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	// The retention period of the auto snapshot.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// 	- \\-1: Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	//
	// 	- 1 to 65536: Auto snapshots are retained for the specified days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-extreme-snapsho****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The snapshot name.
	//
	// If you specify a name to create a snapshot, the name of the snapshot is returned. Otherwise, no value is returned for this parameter.
	//
	// example:
	//
	// FinanceJoshua
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The snapshot type. Valid values:
	//
	// 	- auto: automatically created snapshots
	//
	// 	- user: manually created snapshots
	//
	// example:
	//
	// user
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The ID of the source file system.
	//
	// This parameter is retained even if the source file system of the snapshot is deleted.
	//
	// example:
	//
	// extreme-012****
	SourceFileSystemId *string `json:"SourceFileSystemId,omitempty" xml:"SourceFileSystemId,omitempty"`
	// The capacity of the source file system.
	//
	// Unit: GiB.
	//
	// example:
	//
	// 2000
	SourceFileSystemSize *int64 `json:"SourceFileSystemSize,omitempty" xml:"SourceFileSystemSize,omitempty"`
	// The version of the source file system.
	//
	// example:
	//
	// 1
	SourceFileSystemVersion *string `json:"SourceFileSystemVersion,omitempty" xml:"SourceFileSystemVersion,omitempty"`
	// The status of the snapshot.
	//
	// Valid values:
	//
	// 	- progressing: The snapshot is being created.
	//
	// 	- accomplished: The snapshot is created.
	//
	// 	- failed: The snapshot fails to be created.
	//
	// example:
	//
	// accomplished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetCompletedTime() *string {
	return s.CompletedTime
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetDescription() *string {
	return s.Description
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetRemainTime() *int32 {
	return s.RemainTime
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceFileSystemId() *string {
	return s.SourceFileSystemId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceFileSystemSize() *int64 {
	return s.SourceFileSystemSize
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceFileSystemVersion() *string {
	return s.SourceFileSystemVersion
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetCompletedTime(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.CompletedTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetCreateTime(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.CreateTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetEncryptType(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.EncryptType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetFileSystemType(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.FileSystemType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRetentionDays(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RetentionDays = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotType(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemSize(v int64) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceFileSystemVersion(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceFileSystemVersion = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) Validate() error {
	return dara.Validate(s)
}
