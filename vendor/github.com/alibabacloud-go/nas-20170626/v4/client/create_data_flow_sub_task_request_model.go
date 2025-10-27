// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowSubTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDataFlowSubTaskRequest
	GetClientToken() *string
	SetCondition(v *CreateDataFlowSubTaskRequestCondition) *CreateDataFlowSubTaskRequest
	GetCondition() *CreateDataFlowSubTaskRequestCondition
	SetDataFlowId(v string) *CreateDataFlowSubTaskRequest
	GetDataFlowId() *string
	SetDataFlowTaskId(v string) *CreateDataFlowSubTaskRequest
	GetDataFlowTaskId() *string
	SetDryRun(v bool) *CreateDataFlowSubTaskRequest
	GetDryRun() *bool
	SetDstFilePath(v string) *CreateDataFlowSubTaskRequest
	GetDstFilePath() *string
	SetFileSystemId(v string) *CreateDataFlowSubTaskRequest
	GetFileSystemId() *string
	SetSrcFilePath(v string) *CreateDataFlowSubTaskRequest
	GetSrcFilePath() *string
}

type CreateDataFlowSubTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The check conditions. The check must be passed after the following conditions are specified.
	Condition *CreateDataFlowSubTaskRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// The ID of the data flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The ID of the data flow task.
	//
	// >  Only the IDs of data streaming tasks are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-38aa8e890f45****
	DataFlowTaskId *string `json:"DataFlowTaskId,omitempty" xml:"DataFlowTaskId,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no data streaming subtask is created and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the DataFlowSubTaskId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a data streaming subtask is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// /mnt/file.png
	DstFilePath *string `json:"DstFilePath,omitempty" xml:"DstFilePath,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// /test/file.png
	SrcFilePath *string `json:"SrcFilePath,omitempty" xml:"SrcFilePath,omitempty"`
}

func (s CreateDataFlowSubTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowSubTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowSubTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDataFlowSubTaskRequest) GetCondition() *CreateDataFlowSubTaskRequestCondition {
	return s.Condition
}

func (s *CreateDataFlowSubTaskRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *CreateDataFlowSubTaskRequest) GetDataFlowTaskId() *string {
	return s.DataFlowTaskId
}

func (s *CreateDataFlowSubTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDataFlowSubTaskRequest) GetDstFilePath() *string {
	return s.DstFilePath
}

func (s *CreateDataFlowSubTaskRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateDataFlowSubTaskRequest) GetSrcFilePath() *string {
	return s.SrcFilePath
}

func (s *CreateDataFlowSubTaskRequest) SetClientToken(v string) *CreateDataFlowSubTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowSubTaskRequest) SetCondition(v *CreateDataFlowSubTaskRequestCondition) *CreateDataFlowSubTaskRequest {
	s.Condition = v
	return s
}

func (s *CreateDataFlowSubTaskRequest) SetDataFlowId(v string) *CreateDataFlowSubTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CreateDataFlowSubTaskRequest) SetDataFlowTaskId(v string) *CreateDataFlowSubTaskRequest {
	s.DataFlowTaskId = &v
	return s
}

func (s *CreateDataFlowSubTaskRequest) SetDryRun(v bool) *CreateDataFlowSubTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowSubTaskRequest) SetDstFilePath(v string) *CreateDataFlowSubTaskRequest {
	s.DstFilePath = &v
	return s
}

func (s *CreateDataFlowSubTaskRequest) SetFileSystemId(v string) *CreateDataFlowSubTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowSubTaskRequest) SetSrcFilePath(v string) *CreateDataFlowSubTaskRequest {
	s.SrcFilePath = &v
	return s
}

func (s *CreateDataFlowSubTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataFlowSubTaskRequestCondition struct {
	// The modification time. The value must be a UNIX timestamp. Unit: ns.
	//
	// example:
	//
	// 1725897600000000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The file size. Unit: bytes.
	//
	// example:
	//
	// 68
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateDataFlowSubTaskRequestCondition) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowSubTaskRequestCondition) GoString() string {
	return s.String()
}

func (s *CreateDataFlowSubTaskRequestCondition) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *CreateDataFlowSubTaskRequestCondition) GetSize() *int64 {
	return s.Size
}

func (s *CreateDataFlowSubTaskRequestCondition) SetModifyTime(v int64) *CreateDataFlowSubTaskRequestCondition {
	s.ModifyTime = &v
	return s
}

func (s *CreateDataFlowSubTaskRequestCondition) SetSize(v int64) *CreateDataFlowSubTaskRequestCondition {
	s.Size = &v
	return s
}

func (s *CreateDataFlowSubTaskRequestCondition) Validate() error {
	return dara.Validate(s)
}
