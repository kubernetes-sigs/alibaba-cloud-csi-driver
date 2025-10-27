// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeSnapshotsRequest
	GetFileSystemId() *string
	SetFileSystemType(v string) *DescribeSnapshotsRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeSnapshotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotsRequest
	GetPageSize() *int32
	SetSnapshotIds(v string) *DescribeSnapshotsRequest
	GetSnapshotIds() *string
	SetSnapshotName(v string) *DescribeSnapshotsRequest
	GetSnapshotName() *string
	SetSnapshotType(v string) *DescribeSnapshotsRequest
	GetSnapshotType() *string
	SetStatus(v string) *DescribeSnapshotsRequest
	GetStatus() *string
}

type DescribeSnapshotsRequest struct {
	// The ID of the file system.
	//
	// example:
	//
	// extreme-22f****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme File Storage NAS (NAS) file systems.
	//
	// example:
	//
	// extreme
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The snapshot IDs.
	//
	// You can specify a maximum of 100 snapshot IDs. You must separate snapshot IDs with commas (,).
	//
	// example:
	//
	// s-extreme-67pxwk9aevrkr****,s-extreme-snapsho****,s-extreme-6tmsbas6ljhwh****
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	// The snapshot name.
	//
	// example:
	//
	// FinanceJoshua
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The type of the snapshot.
	//
	// Valid values:
	//
	// 	- auto: auto snapshot
	//
	// 	- user: manual snapshot
	//
	// 	- all (default): all snapshot types
	//
	// example:
	//
	// all
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
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
	// 	- all (default): all snapshot states.
	//
	// example:
	//
	// all
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeSnapshotsRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeSnapshotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsRequest) GetSnapshotIds() *string {
	return s.SnapshotIds
}

func (s *DescribeSnapshotsRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsRequest) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *DescribeSnapshotsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotsRequest) SetFileSystemId(v string) *DescribeSnapshotsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetFileSystemType(v string) *DescribeSnapshotsRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageNumber(v int32) *DescribeSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageSize(v int32) *DescribeSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotIds(v string) *DescribeSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotName(v string) *DescribeSnapshotsRequest {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotType(v string) *DescribeSnapshotsRequest {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetStatus(v string) *DescribeSnapshotsRequest {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
