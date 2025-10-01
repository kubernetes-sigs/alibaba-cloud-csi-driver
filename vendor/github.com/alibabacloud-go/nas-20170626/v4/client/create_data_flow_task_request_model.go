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
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The value of RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The conflict policy for files with the same name. Valid values:
	//
	// 	- SKIP_THE_FILE: skips files with the same name.
	//
	// 	- KEEP_LATEST: compares the update time and keeps the latest version.
	//
	// 	- OVERWRITE_EXISTING: forcibly overwrites the existing file.
	//
	// >  This parameter is required for CPFS for LINGJUN file systems.
	//
	// example:
	//
	// SKIP_THE_FILE
	ConflictPolicy *string `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	// Specifies whether to automatically create a directory if no directory exists. Valid values:
	//
	// 	- true: automatically creates a directory.
	//
	// 	- false (default): does not automatically create a directory.
	//
	// > - This parameter is required if the TaskAction parameter is set to Import.
	//
	// > - Only CPFS for LINGJUN V2.6.0 and later support this parameter.
	//
	// example:
	//
	// false
	CreateDirIfNotExist *bool `json:"CreateDirIfNotExist,omitempty" xml:"CreateDirIfNotExist,omitempty"`
	// The ID of the dataflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The type of data on which operations are performed by the dataflow task.
	//
	// Valid values:
	//
	// 	- Metadata: the metadata of a file, including the timestamp, ownership, and permission information of the file. If you select Metadata, only the metadata of the file is imported. You can only query the file. When you access the file data, the file is loaded from the source storage as required.
	//
	// 	- Data: the data blocks of a file.
	//
	// 	- MetaAndData: the metadata and data blocks of the file.
	//
	// example:
	//
	// Metadata
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The source directory of the data.
	//
	// Limits:
	//
	// 	- The directory must be 1 to 1,023 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// 	- The directory must start and end with a forward slash (/).
	//
	// 	- Only one directory can be listed at a time.
	//
	// 	- If the TaskAction parameter is set to Export, the directory must be a relative path within the FileSystemPath.
	//
	// 	- If the TaskAction parameter is set to Import, the directory must be a relative path within the SourceStoragePath.
	//
	// 	- If the TaskAction parameter is set to StreamExport, the directory must be a relative path within the FileSystemPath.
	//
	// 	- If the TaskAction parameter is set to StreamImport, the directory must be a relative path within the SourceStoragePath.
	//
	// >  Only CPFS for LINGJUN V2.6.0 and later support StreamImport and StreamExport.
	//
	// example:
	//
	// /path_in_cpfs/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no data flow task is created and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, service limits, and available File Storage NAS (NAS) resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the TaskId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a data flow task is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The directory mapped to the data flow task. Limits:
	//
	// 	- The directory must start and end with a forward slash (/). The directory cannot be /../.
	//
	// 	- The directory must be 1 to 1,023 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// 	- Only one directory can be listed at a time.
	//
	// 	- If the TaskAction parameter is set to Export, the directory must be a relative path within the SourceStoragePath.
	//
	// 	- If the TaskAction parameter is set to Import, the directory must be a relative path within the FileSystemPath.
	//
	// 	- If the TaskAction parameter is set to StreamExport, the directory must be a relative path within the SourceStoragePath.
	//
	// 	- If the TaskAction parameter is set to StreamImport, the directory must be a relative path within the FileSystemPath.
	//
	// >  Only CPFS for LINGJUN V2.6.0 and later support StreamImport and StreamExport.
	//
	// example:
	//
	// /path_in_cpfs/
	DstDirectory *string `json:"DstDirectory,omitempty" xml:"DstDirectory,omitempty"`
	// The list of files that are executed by the data flow task.
	//
	// Limits:
	//
	// 	- The list must be encoded in UTF-8.
	//
	// 	- The total length of the file list cannot exceed 64 KB.
	//
	// 	- The file list is in JSON format.
	//
	// 	- The path of a single file must be 1 to 1,023 characters in length and must start with a forward slash (/).
	//
	// 	- If the TaskAction parameter is set to Import, each element in the list represents an OSS object name.
	//
	// 	- If the TaskAction parameter is set to Export, each element in the list represents a CPFS file path.
	//
	// example:
	//
	// ["/path_in_cpfs/file1", "/path_in_cpfs/file2"]
	EntryList *string `json:"EntryList,omitempty" xml:"EntryList,omitempty"`
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-125487\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-0015\\*\\*\\*\\*.
	//
	// >  CPFS is not supported on the international site.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Filters subdirectories and transfers their contents.
	//
	// > 	- This parameter takes effect only when the Directory parameter is specified.
	//
	// > 	- The path length of a single folder must be 1 to 1023 characters, start and end with a forward slash (/), and the total length must not exceed 3000 characters.
	//
	// >	- Only CPFS for Lingjun supports this parameter.
	//
	// example:
	//
	// ["/test/","/test1/"]
	Includes *string `json:"Includes,omitempty" xml:"Includes,omitempty"`
	// If you specify SrcTaskId, you must enter the ID of the dataflow task. The system copies the TaskAction, DataType, and EntryList parameters from the destination dataflow task. You do not need to specify them.
	//
	// >  Streaming dataflow tasks are not supported.
	//
	// example:
	//
	// task-27aa8e890f45****
	SrcTaskId *string `json:"SrcTaskId,omitempty" xml:"SrcTaskId,omitempty"`
	// The type of the data flow task.
	//
	// Valid values:
	//
	// 	- Import: imports data stored in the source storage to a CPFS file system.
	//
	// 	- Export: exports specified data from a CPFS file system to the source storage.
	//
	// 	- StreamImport: batch imports the specified data from the source storage to a CPFS file system.
	//
	// 	- StreamExport: batch exports specified data from a CPFS file system to the source storage.
	//
	// >  Only CPFS for LINGJUN V2.6.0 and later support StreamImport and StreamExport.
	//
	// example:
	//
	// Import
	TaskAction           *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
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
