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
	// A client-generated token that you can use to ensure the idempotence of the request. The token must be unique across different requests.
	//
	// The `ClientToken` value must be an ASCII string of 64 characters or less. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the `ClientToken`. The request ID is unique for each request.
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
	// Specifies whether to perform a dry run for the request.
	//
	// A dry run checks for parameter validity and resource availability without actually canceling the task or incurring charges.
	//
	// Valid values:
	//
	// - `true`: Performs a dry run. The system checks the request for potential issues, including missing parameters, invalid formats, and service limits. If the check fails, the system returns an error message; otherwise, it returns a success code.
	//
	// - `false` (default): Sends a normal request. After the request passes the check, the task is canceled.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - For a general-purpose CPFS instance, the ID must start with `cpfs-`, for example, `cpfs-125487****`.
	//
	// - For a CPFS for AI Computing instance, the ID must start with `bmcpfs-`, for example, `bmcpfs-0015****`.
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
