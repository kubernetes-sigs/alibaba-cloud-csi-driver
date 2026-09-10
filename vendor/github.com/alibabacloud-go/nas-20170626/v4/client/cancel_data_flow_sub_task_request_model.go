// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowSubTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelDataFlowSubTaskRequest
	GetClientToken() *string
	SetDataFlowId(v string) *CancelDataFlowSubTaskRequest
	GetDataFlowId() *string
	SetDataFlowSubTaskId(v string) *CancelDataFlowSubTaskRequest
	GetDataFlowSubTaskId() *string
	SetDataFlowTaskId(v string) *CancelDataFlowSubTaskRequest
	GetDataFlowTaskId() *string
	SetDryRun(v bool) *CancelDataFlowSubTaskRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CancelDataFlowSubTaskRequest
	GetFileSystemId() *string
}

type CancelDataFlowSubTaskRequest struct {
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
	// The data flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The data flow streaming task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// subTaskId-370kyfmyknxcyzw****
	DataFlowSubTaskId *string `json:"DataFlowSubTaskId,omitempty" xml:"DataFlowSubTaskId,omitempty"`
	// The data flow task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-38aa8e890f45****
	DataFlowTaskId *string `json:"DataFlowTaskId,omitempty" xml:"DataFlowTaskId,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating an instance or incurring fees.
	//
	// Valid values:
	//
	// - true: Sends a check request without creating an instance. The check items include whether required parameters are specified, the request format, business limitations, and NAS inventory. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s CancelDataFlowSubTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowSubTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowSubTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelDataFlowSubTaskRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *CancelDataFlowSubTaskRequest) GetDataFlowSubTaskId() *string {
	return s.DataFlowSubTaskId
}

func (s *CancelDataFlowSubTaskRequest) GetDataFlowTaskId() *string {
	return s.DataFlowTaskId
}

func (s *CancelDataFlowSubTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CancelDataFlowSubTaskRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CancelDataFlowSubTaskRequest) SetClientToken(v string) *CancelDataFlowSubTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowSubTaskRequest) SetDataFlowId(v string) *CancelDataFlowSubTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowSubTaskRequest) SetDataFlowSubTaskId(v string) *CancelDataFlowSubTaskRequest {
	s.DataFlowSubTaskId = &v
	return s
}

func (s *CancelDataFlowSubTaskRequest) SetDataFlowTaskId(v string) *CancelDataFlowSubTaskRequest {
	s.DataFlowTaskId = &v
	return s
}

func (s *CancelDataFlowSubTaskRequest) SetDryRun(v bool) *CancelDataFlowSubTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowSubTaskRequest) SetFileSystemId(v string) *CancelDataFlowSubTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowSubTaskRequest) Validate() error {
	return dara.Validate(s)
}
