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
	// The object that contains the script execution records.
	Invocations *DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// The ID of the request.
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
	if s.Invocations != nil {
		if err := s.Invocations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocations struct {
	// The command execution records.
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

type DescribeInvocationsResponseBodyInvocationsInvocation struct {
	// The content of the command.
	//
	// - If \\`ContentEncoding\\` is set to \\`PlainText\\`, the original script content is returned.
	//
	// - If \\`ContentEncoding\\` is set to \\`Base64\\`, the Base64-encoded script content is returned.
	//
	// example:
	//
	// cnBtIC1xYSB8IGdyZXAgdnNm****
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The description of the command.
	//
	// example:
	//
	// testDescription
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The name of the command.
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cron expression for the scheduled command.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The overall execution status of the command. This status is determined by the execution status on all involved instances. Valid values:
	//
	// - Pending: The system is verifying or sending the command. If the command is in the Pending state on at least one instance, the overall status is Pending.
	//
	// - Scheduled: The scheduled command is sent and is waiting to run. If the command is in the Scheduled state on at least one instance, the overall status is Scheduled.
	//
	// - Running: The command is running on the instances. If the command is in the Running state on at least one instance, the overall status is Running.
	//
	// - Success: The command was successfully executed. The command status on each instance is Stopped or Success, and the status on at least one instance is Success.
	//
	//   - For one-time tasks: The command execution is complete and the exit code is 0.
	//
	//   - For scheduled tasks: The last execution was successful with an exit code of 0, and all scheduled executions are complete.
	//
	// - Failed: The command execution failed. The command status on each instance is Stopped or Failed. The overall status is Failed if the command status on one or more instances is one of the following:
	//
	//   - The command failed to be verified (Invalid).
	//
	//   - The command failed to be sent (Aborted).
	//
	//   - The command execution is complete, but the exit code is not 0 (Failed).
	//
	//   - The command timed out (Timeout).
	//
	//   - An error occurred during the command execution (Error).
	//
	// - Stopping: The task is being stopped. If the command is in the Stopping state on at least one instance, the overall status is Stopping.
	//
	// - Stopped: The task was stopped. The overall status is Stopped if the command is in the Stopped state on all instances. The overall status is Stopped if the command status on the instances is one of the following:
	//
	//   - The task was canceled (Cancelled).
	//
	//   - The task was terminated (Terminated).
	//
	// - PartialFailed: The command was successfully executed on some instances but failed on others. The overall status is PartialFailed if the command status on the instances is Success, Failed, or Stopped.
	//
	// > The `InvokeStatus` parameter has a similar meaning. However, check the value of this parameter.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The ID of the command execution.
	//
	// example:
	//
	// t-ind3k9ytvvduoe8
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The command execution records.
	InvokeNodes *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The overall execution status of the command. This status is determined by the execution status on one or more instances. Valid values:
	//
	// - Running:
	//
	//   - Scheduled execution: The status is always Running before you manually stop the scheduled command.
	//
	//   - One-time execution: The overall status is Running if a command process is in progress.
	//
	// - Finished:
	//
	//   - Scheduled execution: A command process cannot be in the Finished state.
	//
	//   - One-time execution: The execution is complete on all instances. Alternatively, the command process is manually stopped on some instances and the execution is complete on the other instances.
	//
	// - Failed:
	//
	//   - Scheduled execution: A command process cannot be in the Failed state.
	//
	//   - One-time execution: The execution failed on all instances.
	//
	// - Stopped: The command is stopped.
	//
	// - Stopping: The command is being stopped.
	//
	// - PartialFailed: The execution failed on some instances. This value is not returned if you specify the `NodeId` parameter.
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
	// - Once: The command is immediately executed.
	//
	// - Period: The command is executed on a schedule.
	//
	// - NextRebootOnly: The command is automatically executed the next time the instance starts.
	//
	// - EveryReboot: The command is automatically executed every time the instance starts.
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
	// The name of the user who runs the command.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The directory where the command is run on the instance.
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
	if s.InvokeNodes != nil {
		if err := s.InvokeNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes struct {
	// The command execution records on the nodes.
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

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode struct {
	// The start time of the command execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The number of characters that are truncated and discarded because the \\`Output\\` value exceeds 24 KB in size.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The error code for a file sending failure. Valid values:
	//
	// - Empty: The file was sent as expected.
	//
	// - NodeNotExists: The specified instance does not exist or has been released.
	//
	// - NodeReleased: The instance was released while the file was being sent.
	//
	// - NodeNotRunning: The instance was not in the Running state when the file sending task was created.
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
	// - DeliveryTimeout: The file failed to be sent due to a timeout.
	//
	// - FileCreateFail: The file failed to be created.
	//
	// - FileAlreadyExists: A file with the same name exists in the same path.
	//
	// - FileContentInvalid: The file content is invalid.
	//
	// - FileNameInvalid: The file name is invalid.
	//
	// - FilePathInvalid: The file path is invalid.
	//
	// - FileAuthorityInvalid: The file permissions are invalid.
	//
	// - UserGroupNotExists: The user group specified for sending the file does not exist.
	//
	// example:
	//
	// NodeNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The details about the cause of a command sending or execution failure. Valid values:
	//
	// - Empty: The command was executed as expected.
	//
	// - the specified node does not exist: The specified instance does not exist or has been released.
	//
	// - the instance was released during the command execution: The instance was released during the command execution.
	//
	// - the instance is not running when create task: The instance was not in the Running state during the command execution.
	//
	// - the command is not applicable: The command is not applicable to the specified instance.
	//
	// - the specified account does not exist: The specified account does not exist.
	//
	// - the specified directory does not exist: The specified directory does not exist.
	//
	// - the cron job expression is invalid: The specified cron expression is invalid.
	//
	// - Cloud Assistant Agent is not running: Cloud Assistant Agent is not running.
	//
	// - Cloud Assistant Agent is not responding: Cloud Assistant Agent is not responding.
	//
	// - Cloud Assistant Agent is being upgraded: Cloud Assistant Agent is being upgraded.
	//
	// - Cloud Assistant Agent needs to be upgraded: Cloud Assistant Agent needs to be upgraded.
	//
	// - The command failed to be sent due to a timeout: The command failed to be sent due to a timeout.
	//
	// - The command execution timed out: The command execution timed out.
	//
	// - An exception occurred during the command execution: An exception occurred during the command execution.
	//
	// - The command execution was interrupted: The command execution was interrupted.
	//
	// - The command execution is complete, but the exit code is not 0: The command execution is complete, but the exit code is not 0.
	//
	// - The instance was released while the file was being sent: The instance was released while the file was being sent.
	//
	// example:
	//
	// the specified node does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command process. Valid values:
	//
	// - On a Linux instance, this is the exit code of the Shell process.
	//
	// - On a Windows instance, this is the exit code of the Batch or PowerShell process.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the execution was complete.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The execution status of the command on a single instance. Valid values:
	//
	// - Pending: The system is verifying or sending the command.
	//
	// - Invalid: The specified command type or parameter is incorrect.
	//
	// - Aborted: Failed to send the command to the instance. The instance must be in the Running state and the command must be sent within 1 minute.
	//
	// - Running: The command is running on the instance.
	//
	// - Success:
	//
	//   - For a one-time command: The execution is complete and the exit code is 0.
	//
	//   - For a scheduled command: The last execution was successful with an exit code of 0, and the specified period is over.
	//
	// - Failed:
	//
	//   - For a one-time command: The execution is complete, but the exit code is not 0.
	//
	//   - For a scheduled command: The last execution was successful, but the exit code was not 0. The scheduled execution will be aborted.
	//
	// - Error: An exception occurred during the command execution and the execution cannot continue.
	//
	// - Timeout: The command execution timed out.
	//
	// - Cancelled: The command execution was canceled. The command was not started.
	//
	// - Stopping: The task is being stopped.
	//
	// - Terminated: The command was terminated during execution.
	//
	// - Scheduled:
	//
	//   - For a one-time command: This status is not applicable and will not occur.
	//
	//   - For a scheduled command: The command is waiting to run.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The ID of the node.
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
	// The output of the command.
	//
	// - If \\`ContentEncoding\\` is set to \\`PlainText\\`, the original output is returned.
	//
	// - If \\`ContentEncoding\\` is set to \\`Base64\\`, the Base64-encoded output is returned.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of times the command has been executed on the instance.
	//
	// - If the command is a one-time execution, the value is 0 or 1.
	//
	// - If the command is a scheduled execution, the value is the number of times the command has been executed.
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
	// The time when \\`StopInvocation\\` was called to stop the command execution.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// Indicates whether the command will be automatically run in the future. Valid values:
	//
	// - true: The command is a scheduled command. The `RepeatMode` parameter was set to `Period`, `NextRebootOnly`, or `EveryReboot` when `RunCommand` or `InvokeCommand` was called.
	//
	// - false (default): The command is a one-time command or has finished.
	//
	//   - The `RepeatMode` parameter was set to `Once` when `RunCommand` or `InvokeCommand` was called.
	//
	//   - The command was canceled, stopped, or has finished running.
	//
	// example:
	//
	// false
	Timed *string `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The time when the record was updated.
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
