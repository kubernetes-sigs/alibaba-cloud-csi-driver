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
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the commands.
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
	// The command output.
	//
	// If ContentEncoding is set to PlainText in the request, the original command output is returned. If ContentEncoding is set to Base64 in the request, the Base64-encoded command output is returned.
	//
	// example:
	//
	// Base64
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file.
	//
	// PlainText: The file content is not encoded. Base64: The file content is encoded in Base64. Default value: PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The time when the file sending task was created.
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The command description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group of the file.
	//
	// example:
	//
	// root
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The overall sending status of the file. The overall sending status of the file varies based on the sending status of the file on all destination instances. Valid values:
	//
	// 	- Pending: The file is being verified or sent. If the sending state of the file on at least one instance is Pending, the overall sending state of the file is Pending.
	//
	// 	- Running: The file is being sent to the instance. If the sending state of the file on at least one instance is Running, the overall sending state of the file is Running.
	//
	// 	- Success: If the sending state of the file on all instances is Success, the overall sending state of the file is Success.
	//
	// 	- Failed: If the sending state of the file on all instances is Failed, the overall sending state of the file is Failed. If the sending state of the file on one or more instances is one of the following values, the overall sending state of the file is Failed:
	//
	//     	- Invalid: The file is invalid.
	//
	//     	- Aborted: The file failed to be sent to the instances.
	//
	//     	- Failed: The file failed to be created on the instances.
	//
	//     	- Timeout: The file sending task timed out.
	//
	//     	- Error: An error occurred and interrupted the file sending task.
	//
	// 	- PartialFailed: The file sending task was completed on some instances but failed on other instances. If the sending state of the file is Success on some instances and is Failed on other instances, the overall sending state of the file is PartialFailed.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The file sending records.
	InvokeNodes *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The name of the file sending task.
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
	// Indicates whether a file was overwritten in the destination directory if the file has the same name as the sent file.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The destination directory.
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
	// The file sending records on a node.
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
	// The time when the file sending task was created.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error code returned when the file failed to be sent to the instance. Valid values:
	//
	// Null: The file is sent to the instance. NodeNotExists: The specified instance does not exist or has been released. NodeReleased: The instance was released while the file was being sent. NodeNotRunning: The instance was not running while the file sending task was being created. AccountNotExists: The specified account does not exist. ClientNotRunning: Cloud Assistant Agent is not running. ClientNotResponse: Cloud Assistant Agent did not respond. ClientIsUpgrading: Cloud Assistant Agent was being upgraded. ClientNeedUpgrade: Cloud Assistant Agent needs to be upgraded. DeliveryTimeout: The file sending task timed out. FileCreateFail: The file failed to be created. FileAlreadyExists: A file with the same name exists in the specified directory. FileContentInvalid: The file content is invalid. FileNameInvalid: The file name is invalid. FilePathInvalid: The specified directory is invalid. FileAuthorityInvalid: The specified permissions on the file are invalid. UserGroupNotExists: The specified user group does not exist.
	//
	// example:
	//
	// AccountNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the command failed to be sent or run. Valid values:
	//
	// 	- null: The command is run as expected.
	//
	// 	- the specified instance does not exists: The specified instance does not exist or is released.
	//
	// 	- the node has released when create task: The instance is released when the command is being run.
	//
	// 	- the node is not running when create task: The instance is not in the Running state while the command is being run.
	//
	// 	- the command is not applicable: The command is not applicable to the specified instance.
	//
	// 	- the specified account does not exists: The specified account does not exist.
	//
	// 	- the specified directory does not exists: The specified directory does not exist.
	//
	// 	- the cron job expression is invalid: The cron expression that specifies the execution time is invalid.
	//
	// 	- the aliyun service is not running on the instance: Cloud Assistant Agent is not running.
	//
	// 	- the aliyun service in the instance does not response: Cloud Assistant Agent does not respond.
	//
	// 	- the aliyun service in the node is upgrading now: Cloud Assistant Agent is being upgraded.
	//
	// 	- the aliyun service in the node need upgrade: Cloud Assistant Agent needs to be upgraded.
	//
	// 	- the command delivery has been timeout: The request to send the command timed out.
	//
	// 	- the command execution has been timeout: The command execution timed out.
	//
	// 	- the command execution got an exception: An exception occurred when the command is being run.
	//
	// 	- the command execution has been interrupted: The command execution is interrupted.
	//
	// 	- the command execution exit code is not zero: The command execution completes, but the exit code is not 0.
	//
	// 	- the specified instance has been released: The instance is released while the file is being sent.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The time when the file sending task ends. The time is in the 2020-05-22T17:04:18 format.
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The status of the file sending task on an instance. Valid values:
	//
	// 	- Pending: The file is being verified or sent.
	//
	// 	- Invalid: The file is invalid.
	//
	// 	- Running: The file is being sent to the instance.
	//
	// 	- Aborted: The file failed to be sent to the instance.
	//
	// 	- Success: The file is sent.
	//
	// 	- Failed: The file failed to be created on the instance.
	//
	// 	- Error: An error occurred and interrupted the file sending task.
	//
	// 	- Timeout: The file sending task timed out.
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
	// The start time.
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The update time.
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
