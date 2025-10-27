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
	// The information about the jobs of the recycle bin.
	Jobs []*ListRecycleBinJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of jobs returned per page.
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
	// The total number of jobs.
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
	return dara.Validate(s)
}

type ListRecycleBinJobsResponseBodyJobs struct {
	// The time when the job was created.
	//
	// example:
	//
	// 2021-05-30T10:08:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code returned.
	//
	// A valid value is returned only if you set the Status parameter to Fail or PartialSuccess.
	//
	// example:
	//
	// InvalidFileId.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// A valid value is returned only if you set the Status parameter to Fail or PartialSuccess.
	//
	// example:
	//
	// The Target File or Directory does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the file or directory in the job.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file or directory that is associated with the job.
	//
	// example:
	//
	// test001
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 8C****C54
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The progress of the job.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the job. Valid values:
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
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the job. Valid values:
	//
	// 	- Restore: a file restoration job
	//
	// 	- Delete: a file deletion job
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
