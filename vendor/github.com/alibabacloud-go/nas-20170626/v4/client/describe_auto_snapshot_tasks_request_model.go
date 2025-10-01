// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyIds(v string) *DescribeAutoSnapshotTasksRequest
	GetAutoSnapshotPolicyIds() *string
	SetFileSystemIds(v string) *DescribeAutoSnapshotTasksRequest
	GetFileSystemIds() *string
	SetFileSystemType(v string) *DescribeAutoSnapshotTasksRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeAutoSnapshotTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoSnapshotTasksRequest
	GetPageSize() *int32
}

type DescribeAutoSnapshotTasksRequest struct {
	// The IDs of automatic snapshot policies.
	//
	// You can specify a maximum of 100 policy IDs. If you want to query the tasks of multiple automatic snapshot policies, you must separate the policy IDs with commas (,).
	//
	// example:
	//
	// sp-extreme-233e6****,sp-extreme-233e6****, sp-extreme-233e6****
	AutoSnapshotPolicyIds *string `json:"AutoSnapshotPolicyIds,omitempty" xml:"AutoSnapshotPolicyIds,omitempty"`
	// The ID of the file system.
	//
	// You can specify a maximum of 100 file system IDs. If you want to query the snapshots of multiple file systems, you must separate the file system IDs with commas (,).
	//
	// example:
	//
	// extreme-233e6****,extreme -23vbp****,extreme -23vas****
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme NAS file systems.
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAutoSnapshotTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotTasksRequest) GetAutoSnapshotPolicyIds() *string {
	return s.AutoSnapshotPolicyIds
}

func (s *DescribeAutoSnapshotTasksRequest) GetFileSystemIds() *string {
	return s.FileSystemIds
}

func (s *DescribeAutoSnapshotTasksRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeAutoSnapshotTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoSnapshotTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoSnapshotTasksRequest) SetAutoSnapshotPolicyIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.AutoSnapshotPolicyIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemIds(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemIds = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetFileSystemType(v string) *DescribeAutoSnapshotTasksRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageNumber(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) SetPageSize(v int32) *DescribeAutoSnapshotTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotTasksRequest) Validate() error {
	return dara.Validate(s)
}
