// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecycleBinJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListRecycleBinJobsResponseBodyJobs) *ListRecycleBinJobsResponseBody
	GetJobs() []*ListRecycleBinJobsResponseBodyJobs
	SetPageNumber(v int64) *ListRecycleBinJobsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRecycleBinJobsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListRecycleBinJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRecycleBinJobsResponseBody
	GetTotalCount() *int64
}

type ListRecycleBinJobsResponseBody struct {
	// The collection of task information in the recycle bin.
	Jobs []*ListRecycleBinJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of tasks per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E15E394-38A6-457A-A62A-D9797C9A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks in the recycle bin.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecycleBinJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecycleBinJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBody) GetJobs() []*ListRecycleBinJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListRecycleBinJobsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRecycleBinJobsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRecycleBinJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecycleBinJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRecycleBinJobsResponseBody) SetJobs(v []*ListRecycleBinJobsResponseBodyJobs) *ListRecycleBinJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetPageNumber(v int64) *ListRecycleBinJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetPageSize(v int64) *ListRecycleBinJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetRequestId(v string) *ListRecycleBinJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) SetTotalCount(v int64) *ListRecycleBinJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecycleBinJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecycleBinJobsResponseBodyJobs struct {
	// The time when the task was created. The time follows the ISO 8601 standard in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2021-05-30T10:08:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code.
	//
	// This value is valid only when Status is Fail or PartialSuccess.
	//
	// example:
	//
	// InvalidFileId.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// This value is valid only when JobStatus is Fail or PartialSuccess.
	//
	// example:
	//
	// The Target File or Directory does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The FileId of the file or directory associated with the task.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file or directory associated with the task.
	//
	// example:
	//
	// test001
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 8C****C54
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The execution progress of the task.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
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
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task type. Valid values:
	//
	// - Restore: A file restoration task.
	//
	// - Delete: A file deletion task.
	//
	// example:
	//
	// Restore
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecycleBinJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListRecycleBinJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetFileId() *string {
	return s.FileId
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetFileName() *string {
	return s.FileName
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetId() *string {
	return s.Id
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetProgress() *string {
	return s.Progress
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListRecycleBinJobsResponseBodyJobs) GetType() *string {
	return s.Type
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetCreateTime(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorCode(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorCode = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetErrorMessage(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.ErrorMessage = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileId = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetFileName(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.FileName = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetId(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetProgress(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Progress = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetStatus(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) SetType(v string) *ListRecycleBinJobsResponseBodyJobs {
	s.Type = &v
	return s
}

func (s *ListRecycleBinJobsResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}
