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
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// rb-15****ed-r-1625****2441
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The job status. Valid values:
	//
	// 	- Running: The job is running.
	//
	// 	- Defragmenting: The job is defragmenting data.
	//
	// 	- PartialSuccess: The job is partially completed.
	//
	// 	- Success: The job is completed.
	//
	// 	- Fail: The job failed.
	//
	// 	- Cancelled: The job is canceled.
	//
	// 	- all (default)
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
