// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeDataFlowTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDataFlowTasksResponseBody
	GetRequestId() *string
	SetTaskInfo(v *DescribeDataFlowTasksResponseBodyTaskInfo) *DescribeDataFlowTasksResponseBody
	GetTaskInfo() *DescribeDataFlowTasksResponseBodyTaskInfo
}

type DescribeDataFlowTasksResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskInfo  *DescribeDataFlowTasksResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DescribeDataFlowTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDataFlowTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataFlowTasksResponseBody) GetTaskInfo() *DescribeDataFlowTasksResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *DescribeDataFlowTasksResponseBody) SetNextToken(v string) *DescribeDataFlowTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBody) SetRequestId(v string) *DescribeDataFlowTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBody) SetTaskInfo(v *DescribeDataFlowTasksResponseBodyTaskInfo) *DescribeDataFlowTasksResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DescribeDataFlowTasksResponseBody) Validate() error {
	if s.TaskInfo != nil {
		if err := s.TaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataFlowTasksResponseBodyTaskInfo struct {
	Task []*DescribeDataFlowTasksResponseBodyTaskInfoTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfo) GetTask() []*DescribeDataFlowTasksResponseBodyTaskInfoTask {
	return s.Task
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfo) SetTask(v []*DescribeDataFlowTasksResponseBodyTaskInfoTask) *DescribeDataFlowTasksResponseBodyTaskInfo {
	s.Task = v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfo) Validate() error {
	if s.Task != nil {
		for _, item := range s.Task {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataFlowTasksResponseBodyTaskInfoTask struct {
	// example:
	//
	// KEEP_LATEST
	ConflictPolicy *string                                                     `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	CreateTime     *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataFlowId     *string                                                     `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	DataType       *string                                                     `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Directory      *string                                                     `json:"Directory,omitempty" xml:"Directory,omitempty"`
	DstDirectory   *string                                                     `json:"DstDirectory,omitempty" xml:"DstDirectory,omitempty"`
	EndTime        *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMsg       *string                                                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	FileSystemPath *string                                                     `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	FilesystemId   *string                                                     `json:"FilesystemId,omitempty" xml:"FilesystemId,omitempty"`
	FsPath         *string                                                     `json:"FsPath,omitempty" xml:"FsPath,omitempty"`
	Includes       *string                                                     `json:"Includes,omitempty" xml:"Includes,omitempty"`
	Originator     *string                                                     `json:"Originator,omitempty" xml:"Originator,omitempty"`
	Progress       *int64                                                      `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ProgressStats  *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats `json:"ProgressStats,omitempty" xml:"ProgressStats,omitempty" type:"Struct"`
	// Deprecated
	ReportPath           *string                                               `json:"ReportPath,omitempty" xml:"ReportPath,omitempty"`
	Reports              *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Struct"`
	SourceStorage        *string                                               `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	StartTime            *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskAction           *string                                               `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	TaskId               *string                                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TransferFileListPath *string                                               `json:"TransferFileListPath,omitempty" xml:"TransferFileListPath,omitempty"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTask) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetConflictPolicy() *string {
	return s.ConflictPolicy
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetDirectory() *string {
	return s.Directory
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetDstDirectory() *string {
	return s.DstDirectory
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetFilesystemId() *string {
	return s.FilesystemId
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetFsPath() *string {
	return s.FsPath
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetIncludes() *string {
	return s.Includes
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetOriginator() *string {
	return s.Originator
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetProgress() *int64 {
	return s.Progress
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetProgressStats() *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	return s.ProgressStats
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetReportPath() *string {
	return s.ReportPath
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetReports() *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports {
	return s.Reports
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetSourceStorage() *string {
	return s.SourceStorage
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) GetTransferFileListPath() *string {
	return s.TransferFileListPath
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetConflictPolicy(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.ConflictPolicy = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetCreateTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.CreateTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDataFlowId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.DataFlowId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDataType(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.DataType = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDirectory(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Directory = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetDstDirectory(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.DstDirectory = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetEndTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.EndTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetErrorMsg(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFileSystemPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFilesystemId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FilesystemId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetFsPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.FsPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetIncludes(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Includes = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetOriginator(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Originator = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetProgress(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Progress = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetProgressStats(v *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.ProgressStats = v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetReportPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.ReportPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetReports(v *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Reports = v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetSourceStorage(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.SourceStorage = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetStartTime(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.StartTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetStatus(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.Status = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetTaskAction(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.TaskAction = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetTaskId(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.TaskId = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) SetTransferFileListPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTask {
	s.TransferFileListPath = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTask) Validate() error {
	if s.ProgressStats != nil {
		if err := s.ProgressStats.Validate(); err != nil {
			return err
		}
	}
	if s.Reports != nil {
		if err := s.Reports.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats struct {
	ActualBytes  *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	ActualFiles  *int64 `json:"ActualFiles,omitempty" xml:"ActualFiles,omitempty"`
	AverageSpeed *int64 `json:"AverageSpeed,omitempty" xml:"AverageSpeed,omitempty"`
	BytesDone    *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	BytesTotal   *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	FilesDone    *int64 `json:"FilesDone,omitempty" xml:"FilesDone,omitempty"`
	FilesTotal   *int64 `json:"FilesTotal,omitempty" xml:"FilesTotal,omitempty"`
	RemainTime   *int64 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetActualFiles() *int64 {
	return s.ActualFiles
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetAverageSpeed() *int64 {
	return s.AverageSpeed
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetFilesDone() *int64 {
	return s.FilesDone
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetFilesTotal() *int64 {
	return s.FilesTotal
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) GetRemainTime() *int64 {
	return s.RemainTime
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetActualBytes(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.ActualBytes = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetActualFiles(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.ActualFiles = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetAverageSpeed(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.AverageSpeed = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetBytesDone(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.BytesDone = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetBytesTotal(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.BytesTotal = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetFilesDone(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.FilesDone = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetFilesTotal(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.FilesTotal = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) SetRemainTime(v int64) *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats {
	s.RemainTime = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats) Validate() error {
	return dara.Validate(s)
}

type DescribeDataFlowTasksResponseBodyTaskInfoTaskReports struct {
	Report []*DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Repeated"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTaskReports) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTaskReports) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports) GetReport() []*DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport {
	return s.Report
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports) SetReport(v []*DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports {
	s.Report = v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports) Validate() error {
	if s.Report != nil {
		for _, item := range s.Report {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) GetName() *string {
	return s.Name
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) GetPath() *string {
	return s.Path
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) SetName(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport {
	s.Name = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) SetPath(v string) *DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport {
	s.Path = &v
	return s
}

func (s *DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport) Validate() error {
	return dara.Validate(s)
}
