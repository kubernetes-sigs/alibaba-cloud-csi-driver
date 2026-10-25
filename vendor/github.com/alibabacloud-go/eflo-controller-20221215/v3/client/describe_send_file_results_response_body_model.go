// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSendFileResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocations(v *DescribeSendFileResultsResponseBodyInvocations) *DescribeSendFileResultsResponseBody
	GetInvocations() *DescribeSendFileResultsResponseBodyInvocations
	SetRequestId(v string) *DescribeSendFileResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeSendFileResultsResponseBody
	GetTotalCount() *string
}

type DescribeSendFileResultsResponseBody struct {
	// The file sending records.
	Invocations *DescribeSendFileResultsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of commands.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSendFileResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBody) GetInvocations() *DescribeSendFileResultsResponseBodyInvocations {
	return s.Invocations
}

func (s *DescribeSendFileResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSendFileResultsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeSendFileResultsResponseBody) SetInvocations(v *DescribeSendFileResultsResponseBodyInvocations) *DescribeSendFileResultsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetRequestId(v string) *DescribeSendFileResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetTotalCount(v string) *DescribeSendFileResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) Validate() error {
	if s.Invocations != nil {
		if err := s.Invocations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSendFileResultsResponseBodyInvocations struct {
	// The command execution ID.
	Invocation []*DescribeSendFileResultsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocations) GetInvocation() []*DescribeSendFileResultsResponseBodyInvocationsInvocation {
	return s.Invocation
}

func (s *DescribeSendFileResultsResponseBodyInvocations) SetInvocation(v []*DescribeSendFileResultsResponseBodyInvocationsInvocation) *DescribeSendFileResultsResponseBodyInvocations {
	s.Invocation = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocations) Validate() error {
	if s.Invocation != nil {
		for _, item := range s.Invocation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSendFileResultsResponseBodyInvocationsInvocation struct {
	// The output information after command execution.
	//
	// If ContentEncoding is set to PlainText, the original output information is returned.
	//
	// If ContentEncoding is set to Base64, the Base64-encoded output information is returned.
	//
	// example:
	//
	// Base64
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The file content type. Valid values:
	//
	// - PlainText: Plain text.
	//
	// - Base64: Base64 encoding.
	//
	// Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The time when the sending task was created. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The command description.
	//
	// example:
	//
	// Description information
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group of the file.
	//
	// example:
	//
	// root
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions of the file.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The user of the file.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The overall sending status of the file. The overall status depends on the combined execution status of all instances in this sending operation. Valid values:
	//
	// - Pending: The system is validating or sending the file. If the file sending status of at least one instance is Pending, the overall status is Pending.
	//
	// - Running: The file is being sent on the instance. If the file sending status of at least one instance is Running, the overall status is Running.
	//
	// - Success: If the file sending status on all instances is Success, the overall status is Success.
	//
	// - Failed: If the file sending status on all instances is Failed, the overall status is Failed. The return value is Failed when one or more of the following statuses occur on an instance:
	//
	//     - The specified file parameters are invalid and validation failed (Invalid).
	//
	//     - Failed to send the file to the instance (Aborted).
	//
	//     - Failed to create the file on the instance (Failed).
	//
	//     - File sending timed out (Timeout).
	//
	//     - An exception occurred during file sending and the operation cannot continue (Error).
	//
	// - PartialFailed: File sending succeeded on some instances but failed on others. If the file sending status on all instances is either Success or Failed, the overall status is PartialFailed.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The file sending records.
	InvokeNodes *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The file sending name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Specifies whether to overwrite the file if a file with the same name already exists in the target directory. Valid values:
	//
	// - true: Overwrite the file.
	//
	// - false: Do not overwrite the file.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The target path.
	//
	// example:
	//
	// /home/user
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocation) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetContent() *string {
	return s.Content
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetContentType() *string {
	return s.ContentType
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetDescription() *string {
	return s.Description
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetFileGroup() *string {
	return s.FileGroup
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetFileMode() *string {
	return s.FileMode
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetFileOwner() *string {
	return s.FileOwner
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetInvokeNodes() *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes {
	return s.InvokeNodes
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetName() *string {
	return s.Name
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetTargetDir() *string {
	return s.TargetDir
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetContent(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Content = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetContentType(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.ContentType = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetCreationTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.CreationTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetDescription(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Description = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileGroup(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileGroup = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileMode(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileMode = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileOwner(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileOwner = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvocationStatus(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvokeNodes(v *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvokeNodes = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetName(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Name = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetNodeCount(v int32) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.NodeCount = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetOverwrite(v bool) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Overwrite = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetTargetDir(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.TargetDir = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) Validate() error {
	if s.InvokeNodes != nil {
		if err := s.InvokeNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes struct {
	// The file sending record of a node.
	InvokeNode []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) GetInvokeNode() []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	return s.InvokeNode
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) Validate() error {
	if s.InvokeNode != nil {
		for _, item := range s.InvokeNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode struct {
	// The time when the file sending task was created. The time is in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error code for the file sending failure. Valid values:
	//
	// - Empty: File sending was normal.
	//
	// - NodeNotExists: The specified instance does not exist or has been released.
	//
	// - NodeReleased: The instance was released during file sending.
	//
	// - NodeNotRunning: The instance was not running when the file sending task was created.
	//
	// - AccountNotExists: The specified account does not exist.
	//
	// - ClientNotRunning: Cloud Assistant Agent is not running.
	//
	// - ClientNotResponse: Cloud Assistant Agent is not responding.
	//
	// - ClientIsUpgrading: Cloud Assistant Agent is being upgraded.
	//
	// - ClientNeedUpgrade: Cloud Assistant Agent needs to be upgraded.
	//
	// - DeliveryTimeout: File sending timed out.
	//
	// - FileCreateFail: File creation failed.
	//
	// - FileAlreadyExists: A file with the same name already exists in the same path.
	//
	// - FileContentInvalid: The file content is invalid.
	//
	// - FileNameInvalid: The file name is invalid.
	//
	// - FilePathInvalid: The file path is invalid.
	//
	// - FileAuthorityInvalid: The file permissions are invalid.
	//
	// - UserGroupNotExists: The user group specified during file sending does not exist.
	//
	// example:
	//
	// AccountNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The details of the reason why the command failed to be sent or executed. Valid values:
	//
	// - Empty: The command was executed normally.
	//
	// - the specified instance does not exists: The specified instance does not exist or has been released.
	//
	// - the node has released when create task: The instance was released during command execution.
	//
	// - the node is not running when create task: The instance was not running when the command was executed.
	//
	// - the command is not applicable: The command is not applicable to the specified instance.
	//
	// - the specified account does not exists: The specified account does not exist.
	//
	// - the specified directory does not exists: The specified directory does not exist.
	//
	// - the cron job expression is invalid: The specified execution time expression is invalid.
	//
	// - the aliyun service is not running on the instance: Cloud Assistant Agent is not running.
	//
	// - the aliyun service in the instance does not response: Cloud Assistant Agent is not responding.
	//
	// - the aliyun service in the node is upgrading now: Cloud Assistant Agent is being upgraded.
	//
	// - the aliyun service in the node need upgrade: Cloud Assistant Agent needs to be upgraded.
	//
	// - the command delivery has been timeout: Sending the command timed out.
	//
	// - the command execution has been timeout: Command execution timed out.
	//
	// - the command execution got an exception: An exception occurred during command execution.
	//
	// - the command execution has been interrupted: Command execution was interrupted.
	//
	// - the command execution exit code is not zero: Command execution ended with a non-zero exit code.
	//
	// - the specified instance has been released: The instance was released during file sending.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The end time, in the format: "2020-05-22T17:04:18".
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The file sending status on a single instance. Valid values:
	//
	// - Pending: The system is validating or sending the file.
	//
	// - Invalid: The specified file parameters are invalid and validation failed.
	//
	// - Running: The file is being sent on the instance.
	//
	// - Aborted: Failed to send the file to the instance.
	//
	// - Success: File sending completed.
	//
	// - Failed: Failed to create the file on the instance.
	//
	// - Error: An exception occurred during file sending and the operation cannot continue.
	//
	// - Timeout: File sending timed out.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-9lb3c15m81j
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The start time. The time is in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The update time. The time is in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetCreationTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.CreationTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorCode(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorInfo(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetFinishTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.FinishTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetInvocationStatus(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeId(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStartTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StartTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetUpdateTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) Validate() error {
	return dara.Validate(s)
}
