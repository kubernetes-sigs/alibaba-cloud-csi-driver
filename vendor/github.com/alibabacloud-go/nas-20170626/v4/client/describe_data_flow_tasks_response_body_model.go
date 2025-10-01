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
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about data flow tasks.
	TaskInfo *DescribeDataFlowTasksResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDataFlowTasksResponseBodyTaskInfoTask struct {
	// The conflict policy for files with the same name. Valid values:
	//
	// 	- SKIP_THE_FILE: skips files with the same name.
	//
	// 	- KEEP_LATEST: compares the update time and keeps the latest version.
	//
	// 	- OVERWRITE_EXISTING: forcibly overwrites the existing file.
	//
	// example:
	//
	// KEEP_LATEST
	ConflictPolicy *string `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2021-08-04 18:27:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the data flow.
	//
	// example:
	//
	// dfid-194433a5be3****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The type of data on which operations are performed by the data flow task. Valid values:
	//
	// 	- Metadata: the metadata of a file, including the timestamp, ownership, and permission information of the file. If you select Metadata, only the metadata of the file is imported. You can only query the file. When you access the file data, the file is loaded from the source storage as required.
	//
	// 	- Data: the data blocks of the file.
	//
	// 	- MetaAndData: the metadata and data blocks of the file.
	//
	// >  CPFS for LINGJUN supports only the MetaAndData type.
	//
	// example:
	//
	// Metadata
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The directory in which the data flow task is executed.
	//
	// example:
	//
	// /path_in_cpfs/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The directory mapped to the data flow task.
	//
	// example:
	//
	// /path_in_cpfs/
	DstDirectory *string `json:"DstDirectory,omitempty" xml:"DstDirectory,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2021-08-04 18:27:35
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cause of the task exception.
	//
	// >  If this parameter is not returned or the return value is empty, no error occurs.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The directory of the fileset in the CPFS file system.
	//
	// Limits:
	//
	// 	- The directory must be 2 to 1024 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// 	- The directory must start and end with a forward slash (/).
	//
	// 	- The directory must be a fileset directory in the CPFS file system.
	//
	// >  Only CPFS supports this parameter.
	//
	// example:
	//
	// /a/b/c/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FilesystemId *string `json:"FilesystemId,omitempty" xml:"FilesystemId,omitempty"`
	// The path of the smart directory.
	//
	// example:
	//
	// /aa/
	FsPath *string `json:"FsPath,omitempty" xml:"FsPath,omitempty"`
	// Filter the directories under directory and transfer the folder contents contained in the filtered directory.
	//
	// example:
	//
	// ["/test/","/test1/"]
	Includes *string `json:"Includes,omitempty" xml:"Includes,omitempty"`
	// The initiator of the data flow task. Valid values:
	//
	// 	- User: The task is initiated by a user.
	//
	// 	- System: The task is automatically initiated by CPFS based on the automatic update interval.
	//
	// >  Only CPFS supports this parameter.
	//
	// example:
	//
	// User
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// The progress of the data flow task. The number of operations that have been performed by the data flow task.
	//
	// example:
	//
	// 240
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The progress of the data flow task.
	ProgressStats *DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats `json:"ProgressStats,omitempty" xml:"ProgressStats,omitempty" type:"Struct"`
	// Deprecated
	//
	// The save path of data flow task reports in the CPFS file system.
	//
	// 	- The task reports for a CPFS file system are generated in the `.dataflow_report` directory of the CPFS file system.
	//
	// 	- CPFS for LINGJUN returns an OSS download link for you to download the task reports.
	//
	// example:
	//
	// /path_in_cpfs/reportfile.cvs
	ReportPath *string `json:"ReportPath,omitempty" xml:"ReportPath,omitempty"`
	// The reports.
	//
	// >  Streaming tasks do not support reports.
	Reports *DescribeDataFlowTasksResponseBodyTaskInfoTaskReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Struct"`
	// The access path of the source storage. Format: `<storage type>://[<account id>:]<path>`.
	//
	// Parameters:
	//
	// 	- storage type: Only Object Storage Service (OSS) is supported.
	//
	// 	- account id: the UID of the account of the source storage.
	//
	// 	- path: the name of the OSS bucket. Limits:
	//
	//     	- The name can contain only lowercase letters, digits, and hyphens (-). The name must start and end with a lowercase letter or digit.
	//
	//     	- The name can be up to 128 characters in length.
	//
	//     	- The name must be encoded in UTF-8.
	//
	// >
	//
	// 	- The OSS bucket must be an existing bucket in the region.
	//
	// 	- Only CPFS for LINGJUN V2.6.0 and later support the account id parameter.
	//
	// example:
	//
	// oss://bucket1
	SourceStorage *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	// The time when the task started.
	//
	// example:
	//
	// 2021-08-04 18:27:35
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the data flow task. Valid values:
	//
	// 	- Pending: The data flow task has been created and has not started.
	//
	// 	- Executing: The data flow task is being executed.
	//
	// 	- Failed: The data flow task failed to be executed. You can view the cause of the failure in the data flow task report.
	//
	// 	- Completed: The data flow task is completed. You can check that all the files have been correctly transferred in the data flow task report.
	//
	// 	- Canceled: The data flow task is canceled and is not completed.
	//
	// 	- Canceling: The data flow task is being canceled.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the data flow task. Valid values:
	//
	// 	- Import: imports data stored in the source storage to a CPFS file system.
	//
	// 	- Export: exports specified data from a CPFS file system to the source storage.
	//
	// 	- StreamImport: imports the specified data from the source storage to a CPFS file system in streaming mode.
	//
	// 	- StreamExport: exports specified data from a CPFS file system to the source storage in streaming mode.
	//
	// 	- Evict: releases the data blocks of a file in a CPFS file system. After the eviction, only the metadata of the file is retained in the CPFS file system. You can still query the file. However, the data blocks of the file are cleared and do not occupy the storage space in the CPFS file system. When you access the file data, the file is loaded from the source storage as required.
	//
	// 	- Inventory: obtains the inventory list managed by a data flow from the CPFS file system, providing the cache status of inventories in the data flow.
	//
	// >  Only CPFS for LINGJUN V2.6.0 and later support StreamImport and StreamExport.
	//
	// example:
	//
	// Import
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The ID of the data flow task.
	//
	// example:
	//
	// taskId-12345678
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TransferFileListPath *string `json:"TransferFileListPath,omitempty" xml:"TransferFileListPath,omitempty"`
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
	return dara.Validate(s)
}

type DescribeDataFlowTasksResponseBodyTaskInfoTaskProgressStats struct {
	// The actual amount of data for which the data flow task is complete. Unit: bytes.
	//
	// example:
	//
	// 131092971520
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The actual number of files for which the data flow task is complete.
	//
	// example:
	//
	// 3
	ActualFiles *int64 `json:"ActualFiles,omitempty" xml:"ActualFiles,omitempty"`
	// The average flow velocity. Unit: bytes/s.
	//
	// example:
	//
	// 342279299
	AverageSpeed *int64 `json:"AverageSpeed,omitempty" xml:"AverageSpeed,omitempty"`
	// The amount of data (including skipped data) for which the data flow task is complete. Unit: bytes.
	//
	// example:
	//
	// 131092971520
	BytesDone *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	// The amount of data scanned on the source. Unit: bytes.
	//
	// example:
	//
	// 131092971520
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// The number of files (including skipped files) for which the data flow task is complete.
	//
	// example:
	//
	// 3
	FilesDone *int64 `json:"FilesDone,omitempty" xml:"FilesDone,omitempty"`
	// The number of files scanned on the source.
	//
	// example:
	//
	// 3
	FilesTotal *int64 `json:"FilesTotal,omitempty" xml:"FilesTotal,omitempty"`
	// The estimated remaining execution time. Unit: seconds.
	//
	// example:
	//
	// 437
	RemainTime *int64 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
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
	return dara.Validate(s)
}

type DescribeDataFlowTasksResponseBodyTaskInfoTaskReportsReport struct {
	// The name of the report.
	//
	// 	- CPFS:
	//
	//     TotalFilesReport: task reports.
	//
	// 	- CPFS for LINGJUN:
	//
	//     	- FailedFilesReport: failed file reports.
	//
	//     	- SkippedFilesReport: skipped file reports.
	//
	//     	- SuccessFilesReport: successful file reports.
	//
	// example:
	//
	// TotalFilesReport
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The report URL.
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
