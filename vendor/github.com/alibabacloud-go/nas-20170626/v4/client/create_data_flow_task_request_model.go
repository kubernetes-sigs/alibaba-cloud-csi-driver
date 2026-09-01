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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The conflict policy for files with the same name.
	//
	// Valid values:
	//
	// - SKIP_THE_FILE: skips files with the same name.
	//
	// - KEEP_LATEST: compares the update time and keeps the latest version.
	//
	// - OVERWRITE_EXISTING: forcibly overwrites files with the same name.
	//
	// > This parameter is required when the file system type is CPFS for Lingjun.
	//
	// example:
	//
	// SKIP_THE_FILE
	ConflictPolicy *string `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	// Specifies whether to enable automatic creation of the folder if it does not exist.
	//
	// Valid values:
	//
	// - true: enables automatic creation of the folder.
	//
	// - false (default): does not enable automatic creation of the folder.
	//
	// > - This parameter takes effect when TaskAction is set to Import.
	//
	// > - Only CPFS for Lingjun 2.6.0 and later support this feature.
	//
	// example:
	//
	// false
	CreateDirIfNotExist *bool `json:"CreateDirIfNotExist,omitempty" xml:"CreateDirIfNotExist,omitempty"`
	// The data flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The type of data on which the data flow task operates.
	//
	// Valid values:
	//
	// - Metadata: the metadata of files, including the timestamp, ownership, permission, and other attributes. If you select Metadata, only the metadata of files is imported. You can view the file, but when you access the file data, the data is loaded from the source storage on demand.
	//
	// - Data: the data blocks of files.
	//
	// - MetaAndData: the metadata and data blocks of files.
	//
	// > When TaskAction is set to Evict, the DataType parameter is required.
	//
	// example:
	//
	// Metadata
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The source directory of data.
	//
	// Limits:
	//
	// - The value must be 1 to 1,023 characters in length.
	//
	// - The value must be encoded in UTF-8.
	//
	// - The value must start and end with a forward slash (/).
	//
	// - Only one directory can be specified at a time.
	//
	// - When TaskAction is set to Export, this directory must be a relative path within FileSystemPath.
	//
	// - When TaskAction is set to Import, this directory must be a relative path within SourceStoragePath.
	//
	// - When TaskAction is set to StreamExport, this directory must be a relative path within FileSystemPath.
	//
	// - When TaskAction is set to StreamImport, this directory must be a relative path within SourceStoragePath.
	//
	// > StreamImport and StreamExport are supported only in CPFS for Lingjun 2.6.0 and later.
	//
	// Directory, EntryList, and TransferFileListPath are mutually exclusive parameters. You can specify only one of them.
	//
	// example:
	//
	// /path_in_cpfs/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// The dry run checks parameter validity and whether required resources are available. The dry run does not create an instance or incur fees.
	//
	// Valid values:
	//
	// - true: performs a dry run without creating the instance. The system checks whether the required parameters are specified, whether the request format is valid, whether service limits are reached, and whether the required NAS resources are available. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned, but TaskId is empty.
	//
	// - false (default): performs a dry run and sends the request. If the request passes the dry run, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The target directory to which the data flow task maps.
	//
	// Limits:
	//
	//  - The value must start and end with a forward slash (/). /../ is not supported.
	//
	//  - The value must be 1 to 1,023 characters in length.
	//
	//  - The value must be encoded in UTF-8.
	//
	//  - Only one directory can be specified at a time.
	//
	//  - When TaskAction is set to Export, this directory must be a relative path within SourceStoragePath.
	//
	//  - When TaskAction is set to Import, this directory must be a relative path within FileSystemPath.
	//
	//  - When TaskAction is set to StreamExport, this directory must be a relative path within SourceStoragePath.
	//
	//  - When TaskAction is set to StreamImport, this directory must be a relative path within FileSystemPath.
	//
	// > StreamImport and StreamExport are supported only in CPFS for Lingjun 2.6.0 and later.
	//
	// example:
	//
	// /path_in_cpfs/
	DstDirectory *string `json:"DstDirectory,omitempty" xml:"DstDirectory,omitempty"`
	// The list of files on which the data flow task is executed.
	//
	// Limits:
	//
	// - The value must be encoded in UTF-8.
	//
	// - The total length of the file list must be less than 64 KB.
	//
	// - The file list is in JSON format.
	//
	// - The path of each file must be 1 to 1,023 characters in length and must start with a forward slash (/).
	//
	// - When TaskAction is set to Import, each element in the list represents an OSS object name.
	//
	// - When TaskAction is set to Export, each element in the list represents a CPFS file path.
	//
	// > Directory, EntryList, and TransferFileListPath are mutually exclusive parameters. You can specify only one of them.
	//
	// example:
	//
	// ["/path_in_cpfs/file1", "/path_in_cpfs/file2"]
	EntryList *string `json:"EntryList,omitempty" xml:"EntryList,omitempty"`
	// The file system ID.
	//
	// - General-purpose CPFS: must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Filters directories under the specified directory and transfers the content of the included folders.
	//
	// > - This parameter takes effect only when the Directory parameter is specified.
	//
	// > - The path of each folder must be 1 to 1,023 characters in length and must start and end with a forward slash (/). The total length must not exceed 3,000 characters.
	//
	// > - Only CPFS for Lingjun supports this feature.
	//
	// example:
	//
	// ["/test/","/test1/"]
	Includes *string `json:"Includes,omitempty" xml:"Includes,omitempty"`
	// If you specify SrcTaskId, enter the data flow task ID. The system copies the TaskAction, DataType, and EntryList parameter information from the specified data flow task, and you do not need to specify these parameters separately.
	//
	// > Data flow streaming tasks are not supported.
	//
	// example:
	//
	// task-29ee8e890f45****
	SrcTaskId *string `json:"SrcTaskId,omitempty" xml:"SrcTaskId,omitempty"`
	// The data flow node type.
	//
	// Valid values:
	//
	// - Import: performs data import from the source storage to CPFS.
	//
	// - Export: exports specified data from CPFS to the source storage.
	//
	// - StreamImport: batch imports specified data from the source storage to CPFS.
	//
	// - StreamExport: batch exports specified data from CPFS to the source storage.
	//
	// - Evict: releases data blocks of files on CPFS. After the release, only metadata is retained on CPFS. You can still query the file, but the data blocks are purged and do not occupy storage capacity on CPFS. When you access the file data, the data is loaded from the source storage on demand.
	//
	// - Inventory: obtains the file checklist managed by the data stream on CPFS. The checklist provides the cache status of files in the data flow.
	//
	// > CPFS for Lingjun supports only Import, Export, StreamImport, and StreamExport. StreamImport and StreamExport are supported only in CPFS for Lingjun 2.6.0 and later.
	//
	// example:
	//
	// Import
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The OSS directory. Data is synchronized based on the content of CSV files in the OSS directory. Limits:
	//
	// - The value must start and end with a forward slash (/).
	//
	// - The value is case-sensitive.
	//
	// - The value must be 1 to 1,023 characters in length.
	//
	// - The value must be encoded in UTF-8.
	//
	//
	// >- TransferFileListPath, Directory, and EntryList are mutually exclusive parameters. You can specify only one of them.
	//
	// >- This parameter specifies an existing path in OSS. The \\*.csv files in the path are stored in OSS.
	//
	// > - TransferFileListPath supports only Import and Export.
	//
	// > - In the Import scenario, the files or directories specified in the CSV files are imported from OSS to CPFS.
	//
	// > - In the Export scenario, the files or directories specified in the CSV files are exported from CPFS to OSS.
	//
	// > - The CSV file format must include the Name and Type columns. Name is a relative path, and Type supports two values: dir and file. If Type is dir, the Name value must end with a forward slash (/).
	//
	// >- Only CPFS for Lingjun supports this feature.
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
