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
	// The ID of the data flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The ID of the data streaming task.
	//
	// This parameter is required.
	//
	// example:
	//
	// subTaskId-370kyfmyknxcyzw****
	DataFlowSubTaskId *string `json:"DataFlowSubTaskId,omitempty" xml:"DataFlowSubTaskId,omitempty"`
	// The ID of the data flow task.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-38aa8e890f45****
	DataFlowTaskId *string `json:"DataFlowTaskId,omitempty" xml:"DataFlowTaskId,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no data streaming task is created and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, service limits, and available File Storage NAS (NAS) resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a data streaming task is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
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
