// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocations(v *DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody
	GetInvocations() *DescribeInvocationsResponseBodyInvocations
	SetRequestId(v string) *DescribeInvocationsResponseBody
	GetRequestId() *string
}

type DescribeInvocationsResponseBody struct {
	// The command execution record.
	Invocations *DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) GetInvocations() *DescribeInvocationsResponseBodyInvocations {
	return s.Invocations
}

func (s *DescribeInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v *DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyInvocations struct {
	// The file sending records.
	Invocation []*DescribeInvocationsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvocation() []*DescribeInvocationsResponseBodyInvocationsInvocation {
	return s.Invocation
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocation(v []*DescribeInvocationsResponseBodyInvocationsInvocation) *DescribeInvocationsResponseBodyInvocations {
	s.Invocation = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyInvocationsInvocation struct {
	// The executed command.
	//
	// 	- If ContentEncoding is set to PlainText in the request, the original command content is returned.
	//
	// 	- If ContentEncoding is set to Base64 in the request, the Base64-encoded command content is returned.
	//
	// example:
	//
	// cnBtIC1xYSB8IGdyZXAgdnNm****
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command description.
	//
	// example:
	//
	// testDescription
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The command name.
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The time when the command task was created.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The schedule on which the command was run.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The overall execution state of the command task. The value of this parameter depends on the execution states of the command task on all the involved instances. Valid values:
	//
	// 	- Pending: The command was being verified or sent. If the execution state on at least one instance is Pending, the overall execution state is Pending.
	//
	// 	- Scheduled: The command that is set to run on a schedule is sent and waiting to be run. If the execution state on at least one instance is Scheduled, the overall execution state is Scheduled.
	//
	// 	- Running: The command is being run on the instance. When the execution state on at least one instance is Running, the overall execution state is Running.
	//
	// 	- Success: When the execution state on at least one instance is Success and the execution state on the other instances is Stopped or Success, the overall execution state is Success.
	//
	//     	- One-time task: The execution is complete, and the exit code is 0.
	//
	//     	- Scheduled task: The last execution was complete, the exit code was 0, and the specified period ended.
	//
	// 	- Failed: When the execution state on all instances is Stopped or Failed, the overall execution state is Failed. When the execution state on an instance is one of the following values, Failed is returned as the overall execution state:
	//
	//     	- Invalid: The command is invalid.
	//
	//     	- Aborted: The command failed to be sent.
	//
	//     	- Failed: The execution was complete, but the exit code was not 0.
	//
	//     	- Timeout: The execution timed out.
	//
	//     	- Error: An error occurred while the command was being run.
	//
	// 	- Stopping: The command task is being stopped. When the execution state on at least one instance is Stopping, the overall execution state is Stopping.
	//
	// 	- Stopped: The task was stopped. When the execution state on all instances is Stopped, the overall execution state is Stopped. When the execution state on an instance is one of the following values, Stopped is returned as the overall execution state:
	//
	//     	- Cancelled: The task was canceled.
	//
	//     	- Terminated: The task was terminated.
	//
	// 	- PartialFailed: The execution was complete on some instances and failed on other instances. When the execution state is Success on some instances and is Failed or Stopped on the other instances, the overall execution state is PartialFailed.
	//
	// >  The value of the `InvokeStatus` response parameter is similar to the value of InvocationStatus. We recommend that you ignore InvokeStatus and check the value of InvocationStatus.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The execution ID.
	//
	// example:
	//
	// t-ind3k9ytvvduoe8
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The command execution records.
	InvokeNodes *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The overall execution status of the command task. The value of this parameter depends on the execution states of the command task on all involved instances. Valid values:
	//
	// 	- Running:
	//
	//     	- Scheduled task: Before you stop the scheduled execution of the command, the overall execution state is always Running.
	//
	//     	- One-time task: If the command is being run on instances, the overall execution state is Running.
	//
	// 	- Finished:
	//
	//     	- Scheduled task: The overall execution state can never be Finished.
	//
	//     	- One-time task: The execution is complete on all instances, or the execution is stopped on some instances and is complete on the other instances.
	//
	// 	- Failed:
	//
	//     	- Scheduled task: The overall execution state can never be Failed.
	//
	//     	- One-time task: The execution failed on all instances.
	//
	// 	- Stopped: The task is stopped.
	//
	// 	- Stopping: The task is being stopped.
	//
	// 	- PartialFailed: The task fails on some instances. If you specify both this parameter and `InstanceId`, this parameter does not take effect.
	//
	// example:
	//
	// Running
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The custom parameters in the command.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The execution mode of the command. Valid values:
	//
	// 	- Once: The command is run immediately.
	//
	// 	- Period: The command is run on a schedule.
	//
	// 	- NextRebootOnly: The command is run the next time the instances start.
	//
	// 	- EveryReboot: runs the command every time the instances start.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The timeout period for the command execution. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username that is used to run the command.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The working directory of the command on the instance.
	//
	// example:
	//
	// /home
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocation) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandContent() *string {
	return s.CommandContent
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandDescription() *string {
	return s.CommandDescription
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetFrequency() *string {
	return s.Frequency
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeNodes() *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes {
	return s.InvokeNodes
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeStatus() *string {
	return s.InvokeStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetParameters() *string {
	return s.Parameters
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetUsername() *string {
	return s.Username
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandDescription(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandDescription = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandName(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetFrequency(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Frequency = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeNodes(v *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeNodes = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetParameters(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Parameters = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetRepeatMode(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.RepeatMode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetTimeout(v int32) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Timeout = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetUsername(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Username = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetWorkingDir(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.WorkingDir = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes struct {
	// The command execution records of the node.
	InvokeNode []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) GetInvokeNode() []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	return s.InvokeNode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode struct {
	// The start time of the execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The size of the Output text that was truncated and discarded because the Output value exceeded 24 KB in size.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The error code returned when the file failed to be sent to the instance. Valid values:
	//
	// 	- Null: The file is sent to the instance.
	//
	// 	- NodeNotExists: The specified instance does not exist or has been released.
	//
	// 	- NodeReleased: The instance was released while the file was being sent.
	//
	// 	- NodeNotRunning: The instance was not running while the file sending task was being created.
	//
	// 	- AccountNotExists: The specified account does not exist.
	//
	// 	- ClientNotRunning: Cloud Assistant Agent is not running.
	//
	// 	- ClientNotResponse: Cloud Assistant Agent does not respond.
	//
	// 	- ClientIsUpgrading: Cloud Assistant Agent is being upgraded.
	//
	// 	- ClientNeedUpgrade: Cloud Assistant Agent needs to be upgraded.
	//
	// 	- DeliveryTimeout: The file sending task timed out.
	//
	// 	- FileCreateFail: The file failed to be created.
	//
	// 	- FileAlreadyExists: A file with the same name exists in the specified directory.
	//
	// 	- FileContentInvalid: The file content is invalid.
	//
	// 	- FileNameInvalid: The file name is invalid.
	//
	// 	- FilePathInvalid: The specified directory is invalid.
	//
	// 	- FileAuthorityInvalid: The specified permissions on the file are invalid.
	//
	// 	- UserGroupNotExists: The specified user group does not exist.
	//
	// example:
	//
	// NodeNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the command cannot be sent or run.
	//
	// 	- If this parameter is empty, the command was run as expected.
	//
	// 	- the specified node does not exists: The specified instance does not exist or is released.
	//
	// 	- the node has node when create task: The instance is released when the command is being run.
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
	// 	- the specified node has been released: The instance is released while the file is being sent.
	//
	// example:
	//
	// the specified node does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the execution. Valid values:
	//
	// 	- For Linux instances, the value is the exit code of the shell process.
	//
	// 	- For Windows instances, the value is the exit code of the batch or PowerShell process.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The end time of the execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The execution status of the command on a single instance. Valid values:
	//
	// 	- Pending: The command was being verified or sent.
	//
	// 	- Invalid: The specified command type or parameter is invalid.
	//
	// 	- Aborted: The command failed to be sent to the instance. To send a command to an instance, make sure that the instance is in the Running state and the command can be sent to the instance within 1 minute.
	//
	// 	- Running: The command is being run on the instance.
	//
	// 	- Success:
	//
	//     	- One-time task: The execution was complete, and the exit code was 0.
	//
	//     	- Recurring execution: The previous occurrence of the execution is complete, and the exit code is 0. The specified recurring period during which the command is run ends.
	//
	// 	- Failed:
	//
	//     	- One-time task: The execution was complete, but the exit code was not 0.
	//
	//     	- Recurring execution: The previous occurrence of the execution is complete, but the exit code is not 0. The specified recurring period during which the command is run is about to end.
	//
	// 	- Error: The execution cannot proceed due to an exception.
	//
	// 	- Timeout: The execution timed out.
	//
	// 	- Cancelled: The execution was canceled before it started.
	//
	// 	- Stopping: The command task is being stopped.
	//
	// 	- Terminated: The execution was terminated before completion.
	//
	// 	- Scheduled:
	//
	//     	- One-time task: The execution state can never be Scheduled.
	//
	//     	- Recurring execution: The command is waiting to be run.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-lbj36wkp70b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The execution status of the command on a single instance.
	//
	// example:
	//
	// Finished
	NodeInvokeStatus *string `json:"NodeInvokeStatus,omitempty" xml:"NodeInvokeStatus,omitempty"`
	// The command output.
	//
	// 	- If ContentEncoding is set to PlainText in the request, the original command output is returned.
	//
	// 	- If ContentEncoding is set to Base64 in the request, the Base64-encoded command output is returned.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of times that the command was run on the instance.
	//
	// 	- If the command is set to run only once, the value is 0 or 1.
	//
	// 	- If the command is set to run on a schedule, the value is the number of times that the command has been run on the instance.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the command task was stopped. If you call the StopInvocation operation to stop the command task, the value of this parameter is the time when the operation is called.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// Indicates whether the command is to be automatically run. Valid values:
	//
	// 	- true: The command is run by calling the `RunCommand` or `InvokeCommand` operation with `RepeatMode` set to `Period`, `NextRebootOnly`, or `EveryReboot`.
	//
	// 	- false (default): The command meets the following requirements.
	//
	//     	- The command is run by calling the `RunCommand` or `InvokeCommand` operation with `RepeatMode` set to `Once`.
	//
	//     	- The command task is canceled, stopped, or completed.
	//
	// example:
	//
	// false
	Timed *string `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The update time of the execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetDropped() *int32 {
	return s.Dropped
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetExitCode() *int32 {
	return s.ExitCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetNodeInvokeStatus() *string {
	return s.NodeInvokeStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetOutput() *string {
	return s.Output
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetRepeats() *int32 {
	return s.Repeats
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetTimed() *string {
	return s.Timed
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetDropped(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetExitCode(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeId(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeInvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStartTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStopTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetTimed(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Timed = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) Validate() error {
	return dara.Validate(s)
}
