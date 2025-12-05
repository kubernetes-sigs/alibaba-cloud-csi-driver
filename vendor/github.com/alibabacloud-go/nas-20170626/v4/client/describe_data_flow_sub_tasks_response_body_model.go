// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowSubTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataFlowSubTask(v *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask) *DescribeDataFlowSubTasksResponseBody
	GetDataFlowSubTask() *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask
	SetNextToken(v string) *DescribeDataFlowSubTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDataFlowSubTasksResponseBody
	GetRequestId() *string
}

type DescribeDataFlowSubTasksResponseBody struct {
	// The details about data streaming tasks.
	DataFlowSubTask *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask `json:"DataFlowSubTask,omitempty" xml:"DataFlowSubTask,omitempty" type:"Struct"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// pUJaUwAAAABhdGUyNTk1MQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataFlowSubTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksResponseBody) GetDataFlowSubTask() *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask {
	return s.DataFlowSubTask
}

func (s *DescribeDataFlowSubTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDataFlowSubTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataFlowSubTasksResponseBody) SetDataFlowSubTask(v *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask) *DescribeDataFlowSubTasksResponseBody {
	s.DataFlowSubTask = v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBody) SetNextToken(v string) *DescribeDataFlowSubTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBody) SetRequestId(v string) *DescribeDataFlowSubTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBody) Validate() error {
	if s.DataFlowSubTask != nil {
		if err := s.DataFlowSubTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataFlowSubTasksResponseBodyDataFlowSubTask struct {
	DataFlowSubTask []*DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask `json:"DataFlowSubTask,omitempty" xml:"DataFlowSubTask,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTask) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask) GetDataFlowSubTask() []*DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	return s.DataFlowSubTask
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask) SetDataFlowSubTask(v []*DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask {
	s.DataFlowSubTask = v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTask) Validate() error {
	if s.DataFlowSubTask != nil {
		for _, item := range s.DataFlowSubTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask struct {
	// The time when the data streaming task was created.
	//
	// example:
	//
	// 2024-07-01 19:59:29
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the data flow.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The ID of the data streaming task.
	//
	// example:
	//
	// subTaskId-370kyfmyknxcyzw****
	DataFlowSubTaskId *string `json:"DataFlowSubTaskId,omitempty" xml:"DataFlowSubTaskId,omitempty"`
	// The ID of the data flow task.
	//
	// example:
	//
	// task-38aa8e890f45****
	DataFlowTaskId *string `json:"DataFlowTaskId,omitempty" xml:"DataFlowTaskId,omitempty"`
	// The path of the destination file. Limits:
	//
	// 	- The path must be 1 to 1,023 characters in length.
	//
	// 	- The path must be encoded in UTF-8.
	//
	// 	- The path must start with a forward slash (/).
	//
	// 	- The path must end with the file name.
	//
	// example:
	//
	// /mnt/file.png
	DstFilePath *string `json:"DstFilePath,omitempty" xml:"DstFilePath,omitempty"`
	// The time when the data streaming task ended.
	//
	// example:
	//
	// 2024-07-04 11:14:22
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error message returned when the task failed.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The file information.
	FileDetail *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail `json:"FileDetail,omitempty" xml:"FileDetail,omitempty" type:"Struct"`
	// The ID of the file system.
	//
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The progress of the data streaming task. Valid values: 0 to 10000.
	//
	// example:
	//
	// 10000
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The progress information about data streaming tasks.
	ProgressStats *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats `json:"ProgressStats,omitempty" xml:"ProgressStats,omitempty" type:"Struct"`
	// The path of the source file. Limits:
	//
	// 	- The path must be 1 to 1,023 characters in length.
	//
	// 	- The path must be encoded in UTF-8.
	//
	// 	- The path must start with a forward slash (/).
	//
	// 	- The path must end with the file name.
	//
	// example:
	//
	// /test/file.png
	SrcFilePath *string `json:"SrcFilePath,omitempty" xml:"SrcFilePath,omitempty"`
	// The time when the data streaming task started.
	//
	// example:
	//
	// 2024-07-03 10:43:16
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the data streaming task. Valid values:
	//
	// 	- EXPIRED: The task is terminated.
	//
	// 	- CREATED: The task is created.
	//
	// 	- RUNNING: The task is running.
	//
	// 	- COMPLETE: The task is complete.
	//
	// 	- CANCELING: The task is being canceled.
	//
	// 	- FAILED: The task failed to be executed.
	//
	// 	- CANCELED: The task is canceled.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetDataFlowSubTaskId() *string {
	return s.DataFlowSubTaskId
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetDataFlowTaskId() *string {
	return s.DataFlowTaskId
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetDstFilePath() *string {
	return s.DstFilePath
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetFileDetail() *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail {
	return s.FileDetail
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetProgressStats() *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats {
	return s.ProgressStats
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetSrcFilePath() *string {
	return s.SrcFilePath
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetCreateTime(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.CreateTime = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetDataFlowId(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.DataFlowId = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetDataFlowSubTaskId(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.DataFlowSubTaskId = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetDataFlowTaskId(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.DataFlowTaskId = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetDstFilePath(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.DstFilePath = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetEndTime(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.EndTime = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetErrorMsg(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetFileDetail(v *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.FileDetail = v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetFileSystemId(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetProgress(v int32) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.Progress = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetProgressStats(v *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.ProgressStats = v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetSrcFilePath(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.SrcFilePath = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetStartTime(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.StartTime = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) SetStatus(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask {
	s.Status = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTask) Validate() error {
	if s.FileDetail != nil {
		if err := s.FileDetail.Validate(); err != nil {
			return err
		}
	}
	if s.ProgressStats != nil {
		if err := s.ProgressStats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail struct {
	// The checksum. Format example: crc64:123456.
	//
	// example:
	//
	// crc64:850309505450944****
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// The time when the file was modified. The value is a UNIX timestamp. Unit: ns.
	//
	// example:
	//
	// 1721167603
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 68
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) GetChecksum() *string {
	return s.Checksum
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) GetSize() *int64 {
	return s.Size
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) SetChecksum(v string) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail {
	s.Checksum = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) SetModifyTime(v int64) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) SetSize(v int64) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail {
	s.Size = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskFileDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats struct {
	// The actual amount of data for which the data flow task is complete. Unit: bytes.
	//
	// example:
	//
	// 68
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The average flow velocity. Unit: bytes/s.
	//
	// example:
	//
	// 34
	AverageSpeed *int64 `json:"AverageSpeed,omitempty" xml:"AverageSpeed,omitempty"`
	// The amount of data (including skipped data) for which the data flow task is complete. Unit: bytes.
	//
	// example:
	//
	// 68
	BytesDone *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	// The amount of data scanned on the source. Unit: bytes.
	//
	// example:
	//
	// 68
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) GetAverageSpeed() *int64 {
	return s.AverageSpeed
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) SetActualBytes(v int64) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats {
	s.ActualBytes = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) SetAverageSpeed(v int64) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats {
	s.AverageSpeed = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) SetBytesDone(v int64) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats {
	s.BytesDone = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) SetBytesTotal(v int64) *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats {
	s.BytesTotal = &v
	return s
}

func (s *DescribeDataFlowSubTasksResponseBodyDataFlowSubTaskDataFlowSubTaskProgressStats) Validate() error {
	return dara.Validate(s)
}
