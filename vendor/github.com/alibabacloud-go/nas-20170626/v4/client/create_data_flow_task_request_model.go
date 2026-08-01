// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDataFlowTaskRequest
	GetClientToken() *string
	SetConflictPolicy(v string) *CreateDataFlowTaskRequest
	GetConflictPolicy() *string
	SetCreateDirIfNotExist(v bool) *CreateDataFlowTaskRequest
	GetCreateDirIfNotExist() *bool
	SetDataFlowId(v string) *CreateDataFlowTaskRequest
	GetDataFlowId() *string
	SetDataType(v string) *CreateDataFlowTaskRequest
	GetDataType() *string
	SetDirectory(v string) *CreateDataFlowTaskRequest
	GetDirectory() *string
	SetDryRun(v bool) *CreateDataFlowTaskRequest
	GetDryRun() *bool
	SetDstDirectory(v string) *CreateDataFlowTaskRequest
	GetDstDirectory() *string
	SetEntryList(v string) *CreateDataFlowTaskRequest
	GetEntryList() *string
	SetFileSystemId(v string) *CreateDataFlowTaskRequest
	GetFileSystemId() *string
	SetIncludes(v string) *CreateDataFlowTaskRequest
	GetIncludes() *string
	SetSrcTaskId(v string) *CreateDataFlowTaskRequest
	GetSrcTaskId() *string
	SetTaskAction(v string) *CreateDataFlowTaskRequest
	GetTaskAction() *string
	SetTransferFileListPath(v string) *CreateDataFlowTaskRequest
	GetTransferFileListPath() *string
}

type CreateDataFlowTaskRequest struct {
	// A client-generated token that ensures the idempotence of the request. The token must be unique across different requests.
	//
	// `ClientToken` can contain only ASCII characters and must not exceed 64 characters. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the `RequestId` of the API request as the `ClientToken`. The `RequestId` may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The conflict policy for files with the same name.
	//
	// Valid values:
	//
	// - SKIP_THE_FILE: Skips files with the same name.
	//
	// - KEEP_LATEST: Compares update times and keeps the latest version.
	//
	// - OVERWRITE_EXISTING: Forcibly overwrites files with the same name.
	//
	// > This parameter is required if the file system is a CPFS AI-Computing Edition instance.
	//
	// example:
	//
	// SKIP_THE_FILE
	ConflictPolicy *string `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	// Specifies whether to automatically create the directory if it does not exist.
	//
	// Valid values:
	//
	// - true: Automatically creates the directory.
	//
	// - false (default): Does not automatically create the directory.
	//
	// > 	- This parameter takes effect only when `TaskAction` is set to `Import`.
	//
	// >
	//
	// > 	- This parameter is supported only by CPFS AI-Computing Edition V2.6.0 and later.
	//
	// example:
	//
	// false
	CreateDirIfNotExist *bool `json:"CreateDirIfNotExist,omitempty" xml:"CreateDirIfNotExist,omitempty"`
	// The ID of the data flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The data type that the data flow task operates on.
	//
	// Valid values:
	//
	// - Metadata: The metadata of the file, including attributes such as timestamp, ownership, and permissions. If you select `Metadata`, only the file metadata is imported. You can see the file, but when you access the file data, it is loaded from the source storage on demand.
	//
	// - Data: The data blocks of the file.
	//
	// - MetaAndData: The metadata and data blocks of the file.
	//
	// example:
	//
	// Metadata
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The source directory of the data.
	//
	// Limits:
	//
	// - The length must be 1 to 1,023 characters.
	//
	// - The directory must be UTF-8 encoded.
	//
	// - The directory must start and end with a forward slash (`/`).
	//
	// - Only one directory can be specified at a time.
	//
	// - If `TaskAction` is `Export`, this directory must be a relative path within `FileSystemPath`.
	//
	// - If `TaskAction` is `Import`, this directory must be a relative path within `SourceStoragePath`.
	//
	// - If `TaskAction` is `StreamExport`, this directory must be a relative path within `FileSystemPath`.
	//
	// - If `TaskAction` is `StreamImport`, this directory must be a relative path within `SourceStoragePath`.
	//
	// > `StreamImport` and `StreamExport` are supported only by CPFS AI-Computing Edition V2.6.0 and later.
	//
	// example:
	//
	// /path_in_cpfs/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// Specifies whether to perform a dry run for this creation request.
	//
	// A dry run checks parameter validity and inventory without creating an instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a check request without creating the instance. The system checks for required parameters, request format, business limits, and NAS inventory. If the check fails, an error is returned. If the check passes, an HTTP 200 status code is returned, but `TaskId` is empty.
	//
	// - false (default): Sends a normal request and creates the instance after the check passes.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The destination directory for the data flow task mapping.
	//
	// Limits:
	//
	// - The directory must start and end with a forward slash (`/`). The `/../` sequence is not supported.
	//
	// - The length must be 1 to 1,023 characters.
	//
	// - The directory must be UTF-8 encoded.
	//
	// - Only one directory can be specified at a time.
	//
	// - If `TaskAction` is `Export`, this directory must be a relative path within `SourceStoragePath`.
	//
	// - If `TaskAction` is `Import`, this directory must be a relative path within `FileSystemPath`.
	//
	// - If `TaskAction` is `StreamExport`, this directory must be a relative path within `SourceStoragePath`.
	//
	// - If `TaskAction` is `StreamImport`, this directory must be a relative path within `FileSystemPath`.
	//
	// > `StreamImport` and `StreamExport` are supported only by CPFS AI-Computing Edition V2.6.0 and later.
	//
	// example:
	//
	// /path_in_cpfs/
	DstDirectory *string `json:"DstDirectory,omitempty" xml:"DstDirectory,omitempty"`
	// The list of files for the data flow task to execute.
	//
	// Limits:
	//
	// - The list must be UTF-8 encoded.
	//
	// - The total length of the file list must be less than 64 KB.
	//
	// - The file list must be in JSON format.
	//
	// - The path of a single file must be 1 to 1,023 characters in length and must start with a forward slash (`/`).
	//
	// - If `TaskAction` is `Import`, each element in the list represents an OSS Object name.
	//
	// - If `TaskAction` is `Export`, each element in the list represents a CPFS file path.
	//
	// example:
	//
	// ["/path_in_cpfs/file1", "/path_in_cpfs/file2"]
	EntryList *string `json:"EntryList,omitempty" xml:"EntryList,omitempty"`
	// The ID of the file system.
	//
	// - CPFS General Purpose Edition: The ID must start with `cpfs-`, such as `cpfs-125487****`.
	//
	// - CPFS AI-Computing Edition: The ID must start with `bmcpfs-`, such as `bmcpfs-0015****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Filters the subdirectories under the `Directory` parameter and transfers the content of the filtered subdirectories.
	//
	// > - This parameter takes effect only when the `Directory` parameter is specified.
	//
	// >
	//
	// > - The path of a single folder must be 1 to 1,023 characters in length and must start and end with a forward slash (`/`). The total length cannot exceed 3,000 characters.
	//
	// >
	//
	// > - This feature is supported only by CPFS AI-Computing Edition.
	//
	// example:
	//
	// ["/test/","/test1/"]
	Includes *string `json:"Includes,omitempty" xml:"Includes,omitempty"`
	// If you specify `SrcTaskId`, enter the ID of a data flow task. The system copies the `TaskAction`, `DataType`, and `EntryList` parameter information from the specified task. You do not need to specify these parameters.
	//
	// > Data flow streaming tasks are not supported.
	//
	// example:
	//
	// task-29ee8e890f45****
	SrcTaskId *string `json:"SrcTaskId,omitempty" xml:"SrcTaskId,omitempty"`
	// The type of the data flow task.
	//
	// Valid values:
	//
	// - Import: Imports specified data from the source storage to the CPFS file system.
	//
	// - Export: Exports specified data from the CPFS file system to the source storage.
	//
	// - StreamImport: Imports a large amount of specified data from the source storage to the CPFS file system.
	//
	// - StreamExport: Exports a large amount of specified data from the CPFS file system to the source storage.
	//
	// - Evict: Releases the data blocks of a file from the CPFS file system. After the release, only the metadata of the file is retained. You can still query the file, but its data blocks are cleared and no longer occupy storage capacity. When you access the file data, it is loaded from the source storage on demand.
	//
	// - Inventory: Obtains the inventory of files managed by a data flow for the CPFS file system. This provides the cache status of files in the data flow.
	//
	// > CPFS AI-Computing Edition supports only `Import`, `Export`, `StreamImport`, and `StreamExport`. `StreamImport` and `StreamExport` are supported only by CPFS AI-Computing Edition V2.6.0 and later.
	//
	// example:
	//
	// Import
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// Specifies an OSS directory. Data is synchronized based on the content of the CSV files in this directory. The following limits apply.
	//
	// - The path must start and end with a forward slash (`/`).
	//
	// - The path is case-sensitive.
	//
	// - The length must be between 1 and 1,023 characters.
	//
	// - The path must be UTF-8 encoded.
	//
	// > 	- `TransferFileListPath`, `Directory`, and `EntryList` are mutually exclusive. You can specify only one of them.
	//
	// >
	//
	// > 	- This parameter specifies an existing path in OSS. The `*.csv` files are stored in this path.
	//
	// >
	//
	// > 	- `TransferFileListPath` supports only the `Import` and `Export` features.
	//
	// >
	//
	// > 	- For an `Import` task, the files or directories specified in the CSV file are imported from OSS to the CPFS file system.
	//
	// >
	//
	// > 	- For an `Export` task, the files or directories specified in the CSV file are exported from the CPFS file system to OSS.
	//
	// >
	//
	// > 	- The CSV file must contain `Name` and `Type` columns. `Name` is the relative path. `Type` can be `dir` or `file`. If `Type` is `dir`, the `Name` value must end with a forward slash (`/`).
	//
	// >
	//
	// > 	- This feature is supported only by CPFS AI-Computing Edition.
	//
	// example:
	//
	// /test_oss_path/
	TransferFileListPath *string `json:"TransferFileListPath,omitempty" xml:"TransferFileListPath,omitempty"`
}

func (s CreateDataFlowTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDataFlowTaskRequest) GetConflictPolicy() *string {
	return s.ConflictPolicy
}

func (s *CreateDataFlowTaskRequest) GetCreateDirIfNotExist() *bool {
	return s.CreateDirIfNotExist
}

func (s *CreateDataFlowTaskRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *CreateDataFlowTaskRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateDataFlowTaskRequest) GetDirectory() *string {
	return s.Directory
}

func (s *CreateDataFlowTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDataFlowTaskRequest) GetDstDirectory() *string {
	return s.DstDirectory
}

func (s *CreateDataFlowTaskRequest) GetEntryList() *string {
	return s.EntryList
}

func (s *CreateDataFlowTaskRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateDataFlowTaskRequest) GetIncludes() *string {
	return s.Includes
}

func (s *CreateDataFlowTaskRequest) GetSrcTaskId() *string {
	return s.SrcTaskId
}

func (s *CreateDataFlowTaskRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *CreateDataFlowTaskRequest) GetTransferFileListPath() *string {
	return s.TransferFileListPath
}

func (s *CreateDataFlowTaskRequest) SetClientToken(v string) *CreateDataFlowTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetConflictPolicy(v string) *CreateDataFlowTaskRequest {
	s.ConflictPolicy = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetCreateDirIfNotExist(v bool) *CreateDataFlowTaskRequest {
	s.CreateDirIfNotExist = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDataFlowId(v string) *CreateDataFlowTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDataType(v string) *CreateDataFlowTaskRequest {
	s.DataType = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDirectory(v string) *CreateDataFlowTaskRequest {
	s.Directory = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDryRun(v bool) *CreateDataFlowTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetDstDirectory(v string) *CreateDataFlowTaskRequest {
	s.DstDirectory = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetEntryList(v string) *CreateDataFlowTaskRequest {
	s.EntryList = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetFileSystemId(v string) *CreateDataFlowTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetIncludes(v string) *CreateDataFlowTaskRequest {
	s.Includes = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetSrcTaskId(v string) *CreateDataFlowTaskRequest {
	s.SrcTaskId = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetTaskAction(v string) *CreateDataFlowTaskRequest {
	s.TaskAction = &v
	return s
}

func (s *CreateDataFlowTaskRequest) SetTransferFileListPath(v string) *CreateDataFlowTaskRequest {
	s.TransferFileListPath = &v
	return s
}

func (s *CreateDataFlowTaskRequest) Validate() error {
	return dara.Validate(s)
}
