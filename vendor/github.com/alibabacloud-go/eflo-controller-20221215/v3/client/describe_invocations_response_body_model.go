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
	// The script execution records.
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
	if s.Invocations != nil {
		if err := s.Invocations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocations struct {
	// The file delivery record.
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
	// The command content.
	//
	// - If ContentEncoding is set to PlainText, the original script content is returned.
	//
	// - If ContentEncoding is set to Base64, the Base64-encoded script content is returned.
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
	// The time when the task was created. The time is displayed in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The execution time of the scheduled command.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The overall execution status of the command. The overall execution status depends on the common execution status across all instances in this invocation. Valid values:
	//
	// - Pending: The system is validating or sending the command. When the command execution status on at least one instance is Pending, the overall execution status is Pending.
	//
	// - Scheduled: The scheduled command has been sent and is waiting to run. When the command execution status on at least one instance is Scheduled, the overall execution status is Scheduled.
	//
	// - Running: The command is running on the instance. When the command execution status on at least one instance is Running, the overall execution status is Running.
	//
	// - Success: The command execution status on each instance is Stopped or Success, and the command execution status on at least one instance is Success. The overall execution status is Success.
	//
	//     - One-time task: The command execution is complete and the exit code is 0.
	//
	//     - Scheduled task: The last execution was successful with an exit code of 0, and all specified schedules have been completed.
	//
	// - Failed: The command execution status on each instance is Stopped or Failed. The overall execution status is Failed. The return value is Failed when the command execution status on an instance is one or more of the following:
	//
	//     - Command validation failed (Invalid).
	//
	//     - Command delivery failed (Aborted).
	//
	//     - Command execution completed with a non-zero exit code (Failed).
	//
	//     - Command execution timed out (Timeout).
	//
	//     - Command execution encountered an exception (Error).
	//
	// - Stopping: The task is being stopped. When the command execution status on at least one instance is Stopping, the overall execution status is Stopping.
	//
	// - Stopped: The task has been stopped. When the command execution status on all instances is Stopped, the overall execution status is Stopped. The return value is Stopped when the command execution status on an instance is one of the following:
	//
	//     - The task was canceled (Cancelled).
	//
	//     - The task was terminated (Terminated).
	//
	// - PartialFailed: Some instances executed successfully while others failed. When the command execution status on each instance is Success, Failed, or Stopped, the overall execution status is PartialFailed.
	//
	// > The InvokeStatus parameter in the response has a similar meaning, but we recommend that you check this parameter instead.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The command execution ID.
	//
	// example:
	//
	// t-ind3k9ytvvduoe8
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The command execution records.
	InvokeNodes *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The overall execution status of the command. The overall execution status depends on the common execution status of one or more instances in the invocation. Valid values:
	//
	//
	//
	// - Running:
	//
	//     - Scheduled execution: The execution status remains Running until you manually stop the scheduled command.
	//
	//     - One-time execution: The overall execution status is Running if the command process is running on any instance.
	//
	// - Finished:
	//
	//     - Scheduled execution: The command process cannot have a status of Finished.
	//
	//     - One-time execution: All instances have completed execution, or the command process was manually stopped on some instances while the remaining instances completed execution.
	//
	// - Failed:
	//
	//     - Scheduled execution: The command process cannot have a status of Failed.
	//
	//     - One-time execution: All instances failed to execute the command.
	//
	// - Stopped: The command was stopped.
	//
	// - Stopping: The command is being stopped.
	//
	// - PartialFailed: Partial failure. This value does not take effect if the InstanceId parameter is also specified.
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
	// - Once: Immediately executes the command.
	//
	// - Period: Executes the command on a schedule.
	//
	// - NextRebootOnly: Automatically executes the command the next time the instance starts.
	//
	// - EveryReboot: Automatically executes the command every time the instance starts.
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
	// The username used to execute the command.
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
	if s.InvokeNodes != nil {
		if err := s.InvokeNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes struct {
	// The command execution record on a node.
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
	// The start time of the command execution. The time is displayed in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The length of the truncated and discarded text after the text length in the Output field exceeds 24 KB.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The error code for the file delivery failure. Valid values:
	//
	// - Empty: The file was delivered normally.
	//
	// - NodeNotExists: The specified instance does not exist or has been released.
	//
	// - NodeReleased: The instance was released during file delivery.
	//
	// - NodeNotRunning: The instance was not running when the file delivery task was created.
	//
	// - AccountNotExists: The specified account does not exist.
	//
	// - ClientNotRunning: The Cloud Assistant Agent is not running.
	//
	// - ClientNotResponse: The Cloud Assistant Agent is not responding.
	//
	// - ClientIsUpgrading: The Cloud Assistant Agent is being upgraded.
	//
	// - ClientNeedUpgrade: The Cloud Assistant Agent needs to be upgraded.
	//
	// - DeliveryTimeout: The file delivery timed out.
	//
	// - FileCreateFail: The file creation failed.
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
	// - UserGroupNotExists: The user group specified during file delivery does not exist.
	//
	// example:
	//
	// NodeNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The details of the command delivery failure or execution failure. Valid values:
	//
	// - Empty: The command executed normally.
	//
	// - the specified node does not exists: The specified instance does not exist or has been released.
	//
	// - the node has node when create task: The instance was released during command execution.
	//
	// - the node is not running when create task: The instance was not running when the command was executed.
	//
	// - the command is not applicable: The command is not applicable to the specified instance.
	//
	// - the specified account does not exists: The specified account does not exist.
	//
	// - the specified directory does not exists: The specified directory does not exist.
	//
	// - the cron job expression is invalid: The specified execution schedule expression is invalid.
	//
	// - the aliyun service is not running on the instance: The Cloud Assistant Agent is not running.
	//
	// - the aliyun service in the instance does not response: The Cloud Assistant Agent is not responding.
	//
	// - the aliyun service in the node is upgrading now: The Cloud Assistant Agent is being upgraded.
	//
	// - the aliyun service in the node need upgrade: The Cloud Assistant Agent needs to be upgraded.
	//
	// - the command delivery has been timeout: The command delivery timed out.
	//
	// - the command execution has been timeout: The command execution timed out.
	//
	// - the command execution got an exception: An exception occurred during command execution.
	//
	// - the command execution has been interrupted: The command execution was interrupted.
	//
	// - the command execution exit code is not zero: The command execution completed with a non-zero exit code.
	//
	// - the specified node has been released: The instance was released during file delivery.
	//
	// example:
	//
	// the specified node does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command process. Valid values:
	//
	// - For Linux instances, the value is the exit code of the Shell process.
	//
	// - For Windows instances, the value is the exit code of the Bat or PowerShell process.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The completion time. The time is displayed in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The command execution status on a single instance. Valid values:
	//
	// -  Pending: The system is validating or sending the command.
	//
	// -  Invalid: The specified command type or parameters are invalid.
	//
	// -  Aborted: Failed to send the command to the instance. The instance must be running, and the command must be deliverable within one minute.
	//
	// -  Running: The command is running on the instance.
	//
	// -  Success:
	//
	//     -  One-time command: The command execution is complete and the exit code is 0.
	//
	//     -  Scheduled command: The last execution was successful with an exit code of 0, and the specified schedule has ended.
	//
	// -  Failed:
	//
	//     -  One-time command: The command execution is complete and the exit code is non-zero.
	//
	//     -  Scheduled command: The last execution was successful with a non-zero exit code, and the specified schedule will be aborted.
	//
	// -  Error: An exception occurred during command execution and the command cannot continue.
	//
	// -  Timeout: The command execution timed out.
	//
	// -  Cancelled: The command execution action was canceled and the command was never started.
	//
	// -  Stopping: The task is being stopped.
	//
	// -  Terminated: The command was terminated during execution.
	//
	// -  Scheduled:
	//
	//     -  One-time command: Not applicable. This status does not occur.
	//
	//     -  Scheduled command: Waiting to run.
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
	// The command execution status on a single instance.
	//
	// example:
	//
	// Finished
	NodeInvokeStatus *string `json:"NodeInvokeStatus,omitempty" xml:"NodeInvokeStatus,omitempty"`
	// The command output.
	//
	// - If ContentEncoding is set to PlainText, the original output is returned.
	//
	// - If ContentEncoding is set to Base64, the Base64-encoded output is returned.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of times the command was executed on the instance.
	//
	// -  If the execution mode is one-time execution, the value is 0 or 1.
	//
	// -  If the execution mode is scheduled execution, the value is the number of times the command has been executed.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// The start time. The time is displayed in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when StopInvocation was called to stop the command execution. The time is displayed in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// Indicates whether the queried command will be automatically executed in the future. Valid values:
	//
	// - true: The command was executed by calling RunCommand or InvokeCommand with the RepeatMode parameter set to Period, NextRebootOnly, or EveryReboot.
	//
	// - false (default): The queried command is in one of the following states:
	//
	//     - The command was executed by calling RunCommand or InvokeCommand with the RepeatMode parameter set to Once.
	//
	//     - The command has been canceled, stopped, or has completed execution.
	//
	// example:
	//
	// false
	Timed *string `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The update time. The time is displayed in ISO 8601 format. Z in the example indicates UTC.
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
