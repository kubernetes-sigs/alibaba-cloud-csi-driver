// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLifecycleRetrieveJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleRetrieveJobs(v []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) *ListLifecycleRetrieveJobsResponseBody
	GetLifecycleRetrieveJobs() []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs
	SetPageNumber(v int32) *ListLifecycleRetrieveJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLifecycleRetrieveJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLifecycleRetrieveJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLifecycleRetrieveJobsResponseBody
	GetTotalCount() *int32
}

type ListLifecycleRetrieveJobsResponseBody struct {
	// The details about the data retrieval tasks.
	LifecycleRetrieveJobs []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs `json:"LifecycleRetrieveJobs,omitempty" xml:"LifecycleRetrieveJobs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of data retrieval tasks.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBody) GetLifecycleRetrieveJobs() []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	return s.LifecycleRetrieveJobs
}

func (s *ListLifecycleRetrieveJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLifecycleRetrieveJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLifecycleRetrieveJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLifecycleRetrieveJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetLifecycleRetrieveJobs(v []*ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) *ListLifecycleRetrieveJobsResponseBody {
	s.LifecycleRetrieveJobs = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageNumber(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetPageSize(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetRequestId(v string) *ListLifecycleRetrieveJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) SetTotalCount(v int32) *ListLifecycleRetrieveJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBody) Validate() error {
	if s.LifecycleRetrieveJobs != nil {
		for _, item := range s.LifecycleRetrieveJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs struct {
	// The time when the task was created.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2021-02-30T10:08:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of files that are read in the data retrieval task.
	//
	// example:
	//
	// 100
	DiscoveredFileCount *int64 `json:"DiscoveredFileCount,omitempty" xml:"DiscoveredFileCount,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the data retrieval task.
	//
	// example:
	//
	// lrj-nfstest-ia-160****853-hshvw
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The execution path of the data retrieval task.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The total number of files that are retrieved.
	//
	// example:
	//
	// 80
	RetrievedFileCount *int64 `json:"RetrievedFileCount,omitempty" xml:"RetrievedFileCount,omitempty"`
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
	// 	- InfrequentAccess: the IA storage class.
	//
	// 	- Archive: the Archive storage class.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The time when the task was updated.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2021-02-30T11:08:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) String() string {
	return dara.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetDiscoveredFileCount() *int64 {
	return s.DiscoveredFileCount
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetPaths() []*string {
	return s.Paths
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetRetrievedFileCount() *int64 {
	return s.RetrievedFileCount
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetStatus() *string {
	return s.Status
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetStorageType() *string {
	return s.StorageType
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetCreateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.CreateTime = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetDiscoveredFileCount(v int64) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.DiscoveredFileCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetFileSystemId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.FileSystemId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetJobId(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.JobId = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetPaths(v []*string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.Paths = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetRetrievedFileCount(v int64) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.RetrievedFileCount = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetStatus(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.Status = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetStorageType(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.StorageType = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) SetUpdateTime(v string) *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs {
	s.UpdateTime = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponseBodyLifecycleRetrieveJobs) Validate() error {
	return dara.Validate(s)
}
