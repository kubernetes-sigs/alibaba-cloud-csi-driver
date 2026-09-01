// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecycleBinJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ListRecycleBinJobsRequest
	GetFileSystemId() *string
	SetJobId(v string) *ListRecycleBinJobsRequest
	GetJobId() *string
	SetPageNumber(v int64) *ListRecycleBinJobsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRecycleBinJobsRequest
	GetPageSize() *int64
	SetStatus(v string) *ListRecycleBinJobsRequest
	GetStatus() *string
}

type ListRecycleBinJobsRequest struct {
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// rb-15****ed-r-1625****2441
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number of the current page in a paged query.
	//
	// Start value (default value): 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task status. Valid values:
	//
	// - Running: The task is running.
	//
	// - Defragmenting: Data is being defragmented.
	//
	// - PartialSuccess: The task partially succeeded.
	//
	// - Success: The task succeeded.
	//
	// - Fail: The task failed.
	//
	// - Cancelled: The task is canceled.
	//
	// - All (default): All statuses.
	//
	// example:
	//
	// All
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRecycleBinJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecycleBinJobsRequest) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListRecycleBinJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListRecycleBinJobsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRecycleBinJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRecycleBinJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListRecycleBinJobsRequest) SetFileSystemId(v string) *ListRecycleBinJobsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetJobId(v string) *ListRecycleBinJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetPageNumber(v int64) *ListRecycleBinJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetPageSize(v int64) *ListRecycleBinJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecycleBinJobsRequest) SetStatus(v string) *ListRecycleBinJobsRequest {
	s.Status = &v
	return s
}

func (s *ListRecycleBinJobsRequest) Validate() error {
	return dara.Validate(s)
}
