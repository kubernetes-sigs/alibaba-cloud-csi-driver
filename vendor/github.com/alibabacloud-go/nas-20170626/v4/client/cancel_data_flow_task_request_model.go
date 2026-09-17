// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelDataFlowTaskRequest
	GetClientToken() *string
	SetDataFlowId(v string) *CancelDataFlowTaskRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *CancelDataFlowTaskRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CancelDataFlowTaskRequest
	GetFileSystemId() *string
	SetTaskId(v string) *CancelDataFlowTaskRequest
	GetTaskId() *string
}

type CancelDataFlowTaskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
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
	// Specifies whether to perform a dry run for this cancel request to check parameter validity and task state without actually canceling the task.
	//
	// Valid values:
	//
	// - true: Sends a check request without actually canceling the task. The check items include whether required parameters are specified, the request format, and the task state. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request and directly cancels the task after the check passes.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS General-purpose: Must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: Must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The data flow task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-38aa8e890f45****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelDataFlowTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelDataFlowTaskRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *CancelDataFlowTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CancelDataFlowTaskRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CancelDataFlowTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelDataFlowTaskRequest) SetClientToken(v string) *CancelDataFlowTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetDataFlowId(v string) *CancelDataFlowTaskRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetDryRun(v bool) *CancelDataFlowTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetFileSystemId(v string) *CancelDataFlowTaskRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowTaskRequest) SetTaskId(v string) *CancelDataFlowTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelDataFlowTaskRequest) Validate() error {
	return dara.Validate(s)
}
