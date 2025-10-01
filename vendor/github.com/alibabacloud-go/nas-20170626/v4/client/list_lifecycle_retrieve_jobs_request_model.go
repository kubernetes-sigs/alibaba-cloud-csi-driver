// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLifecycleRetrieveJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ListLifecycleRetrieveJobsRequest
	GetFileSystemId() *string
	SetPageNumber(v int32) *ListLifecycleRetrieveJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLifecycleRetrieveJobsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListLifecycleRetrieveJobsRequest
	GetStatus() *string
	SetStorageType(v string) *ListLifecycleRetrieveJobsRequest
	GetStorageType() *string
}

type ListLifecycleRetrieveJobsRequest struct {
	// The ID of the file system.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the data retrieval task. Valid values:
	//
	// 	- active: The task is running.
	//
	// 	- canceled: The task is canceled.
	//
	// 	- completed: The task is completed.
	//
	// 	- failed: The task has failed.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class.
	//
	// 	- InfrequentAccess: the Infrequent Access (IA) storage class.
	//
	// 	- Archive: the Archive storage class.
	//
	// >  If the StorageType parameter is not specified, data retrieval tasks of all types are returned.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ListLifecycleRetrieveJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLifecycleRetrieveJobsRequest) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListLifecycleRetrieveJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLifecycleRetrieveJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLifecycleRetrieveJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListLifecycleRetrieveJobsRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ListLifecycleRetrieveJobsRequest) SetFileSystemId(v string) *ListLifecycleRetrieveJobsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageNumber(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetPageSize(v int32) *ListLifecycleRetrieveJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetStatus(v string) *ListLifecycleRetrieveJobsRequest {
	s.Status = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) SetStorageType(v string) *ListLifecycleRetrieveJobsRequest {
	s.StorageType = &v
	return s
}

func (s *ListLifecycleRetrieveJobsRequest) Validate() error {
	return dara.Validate(s)
}
