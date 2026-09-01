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
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The verification condition. The specified conditions must pass verification.
	Condition *CreateDataFlowSubTaskRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// The data flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The data flow task ID.
	//
	// >Only data flow streaming task IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-38aa8e890f45****
	DataFlowTaskId *string `json:"DataFlowTaskId,omitempty" xml:"DataFlowTaskId,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a check request without creating the data flow. The check items include whether required parameters are specified, the request format, and business limit dependencies. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but DataFlowSubTaskId is empty.
	//
	// - false (default): Sends a normal request and directly creates the instance after the check passes.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The destination file path.
	//
	// Limits:
	//
	// - The value must be 1 to 1,023 characters in length.
	//
	// - The value must be encoded in UTF-8.
	//
	// - The value must start with a forward slash (/).
	//
	// - The value must end with a file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// /mnt/file.png
	DstFilePath *string `json:"DstFilePath,omitempty" xml:"DstFilePath,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The source file path.
	//
	// Limits:
	//
	// - The value must be 1 to 1,023 characters in length.
	//
	// - The value must be encoded in UTF-8.
	//
	// - The value must start with a forward slash (/).
	//
	// - The value must end with a file name.
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
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataFlowSubTaskRequestCondition struct {
	// The modification time as a UNIX timestamp. Unit: ns.
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
