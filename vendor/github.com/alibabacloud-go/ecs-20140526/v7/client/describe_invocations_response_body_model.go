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
	SetNextToken(v string) *DescribeInvocationsResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeInvocationsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInvocationsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeInvocationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeInvocationsResponseBody
	GetTotalCount() *int64
}

type DescribeInvocationsResponseBody struct {
	// The ID of instance N. When you specify this parameter, the system queries all the execution records of all the commands that run on the instance.
	Invocations *DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
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
	// 	- Success: If the execution state on at least one instance is Success and the execution state on the other instances is Stopped or Success, the overall execution state is Success.
	//
	//     	- One-time task: The execution is complete, and the exit code is 0.
	//
	//     	- Scheduled task: The last execution is complete, the exit code is 0, and the specified period ends.
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
	// 	- Pending: The command is being verified or sent. If the execution state on at least one instance is Pending, the overall execution state is Pending.
	//
	// 	- Scheduled: The command that is set to run on a schedule is sent and waiting to be run. If the execution state on at least one instance is Scheduled, the overall execution state is Scheduled.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The command type. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances.
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances.
	//
	// 	- RunShellScript: shell command, applicable to Linux instances.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The command name. If you specify both this parameter and `InstanceId`, this parameter does not take effect.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the command is to be automatically run. Valid values:
	//
	// 	- true: The command is run by calling the `RunCommand` or `InvokeCommand` operation with `RepeatMode` set to `Period`, `NextRebootOnly`, or `EveryReboot`.
	//
	// 	- false: The command meets one of the following requirements:
	//
	//     	- The command is run by calling the `RunCommand` or `InvokeCommand` operation with `RepeatMode` set to `Once`.
	//
	//     	- The command task is canceled, stopped, or completed.
	//
	// Default value: false.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeInvocationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInvocationsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v *DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetNextToken(v string) *DescribeInvocationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageNumber(v int64) *DescribeInvocationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageSize(v int64) *DescribeInvocationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetTotalCount(v int64) *DescribeInvocationsResponseBody {
	s.TotalCount = &v
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
	// The size of the Output text that was truncated and discarded because the Output value exceeded 24 KB in size.
	//
	// example:
	//
	// cnBtIC1xYSB8IGdyZXAgdnNm****
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// testDescription
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The time when the command process ended.
	//
	// example:
	//
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command output.
	//
	// 	- If ContentEncoding is set to PlainText in the request, the original command output is returned.
	//
	// 	- If ContentEncoding is set to Base64 in the request, the Base64-encoded command output is returned.
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The execution status of the command on a single instance.
	//
	// >  We recommend that you ignore this parameter and check the value of `InvocationStatus` in the response to obtain the execution status.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The error message returned when the command failed to be sent or run. Valid values:
	//
	// 	- If this parameter is empty, the command was run as expected.
	//
	// 	- The security group rules denied access to the aliyun service.
	//
	// 	- The specified instance does not exist.
	//
	// 	- The specified instance was released during task execution.
	//
	// 	- The specified instance was not running during task execution.
	//
	// 	- The OS type of the instance does not support the specified command type.
	//
	// 	- The specified account does not exist.
	//
	// 	- The specified directory does not exist.
	//
	// 	- The cron expression is invalid.
	//
	// 	- The aliyun service is not running on the instance.
	//
	// 	- The aliyun service in the instance does not response.
	//
	// 	- The aliyun service in the instance is upgrading during task execution.
	//
	// 	- The aliyun service in the instance need to be upgraded to at least version to support the feature. indicates the earliest version that supports the feature. indicates the name of the feature.
	//
	// 	- The command delivery has been timeout.
	//
	// 	- The command execution has been timeout.
	//
	// 	- The command execution got an exception.
	//
	// 	- The command execution exit code is not zero.
	//
	// 	- The specified instance was released during task execution.
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The time when the command started to be run on the instance.
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The number of times that the command was run on the instance.
	//
	// 	- If the command is set to run only once, the value is 0 or 1.
	//
	// 	- If the command is set to run on a schedule, the value is the number of times that the command has been run on the instance.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The command execution Output delivers the object URI to OSS. This field is an empty string when the delivery fails or is in progress.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// Running
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The time when the command task was created.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The tags that are added to the command.
	InvokeInstances *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Struct"`
	// Indicates whether the command is to be automatically run.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The output delivery status of the command execution. Valid values:
	//
	// 	- InProgress: The delivery is in progress.
	//
	// 	- Finished: The delivery is complete.
	//
	// 	- Failed: The delivery failed.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// Specifies whether to return the command outputs in the response.
	//
	// 	- true: The command outputs are returned. When this parameter is set to true, you must specify `InvokeId`, `InstanceId`, or both.
	//
	// 	- false: The command outputs are not returned.
	//
	// Default value: false
	//
	// example:
	//
	// oss://testBucket/testPrefix
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The error code for the failure to send or run the command. Valid values:
	//
	// 	- If this parameter is empty, the command is run normally.
	//
	// 	- InstanceNotExists: The specified instance did not exist or was released.
	//
	// 	- InstanceReleased: The instance is released during command execution.
	//
	// 	- InstanceNotRunning: The instance was not running when the command started to be run.
	//
	// 	- CommandNotApplicable: The command was inapplicable to the specified instance.
	//
	// 	- AccountNotExists: The username specified to run the command did not exist.
	//
	// 	- DirectoryNotExists: The specified directory did not exist.
	//
	// 	- BadCronExpression: The specified cron expression for the execution schedule was invalid.
	//
	// 	- ClientNotRunning: Cloud Assistant Agent was not running.
	//
	// 	- ClientNotResponse: Cloud Assistant Agent does not respond.
	//
	// 	- ClientIsUpgrading: Cloud Assistant Agent is being upgraded.
	//
	// 	- ClientNeedUpgrade: Cloud Assistant Agent needed to be upgraded.
	//
	// 	- DeliveryTimeout: The request to send the command timed out.
	//
	// 	- ExecutionTimeout: The execution timed out.
	//
	// 	- ExecutionException: An exception occurred while the command was being executed.
	//
	// 	- ExecutionInterrupted: The command task was interrupted.
	//
	// 	- ExitCodeNonzero: The execution was complete, but the exit code was not 0.
	//
	// 	- SecurityGroupRuleDenied: Access to Cloud Assistant was denied by security group rules.
	//
	// 	- TaskConcurrencyLimit: The number of concurrent tasks exceeds the maximum limit.
	Tags *DescribeInvocationsResponseBodyInvocationsInvocationTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the execution status was updated.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// false
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The execution mode of the command. If you specify both this parameter and `InstanceId`, this parameter does not take effect. Valid values:
	//
	// 	- Once: The command is immediately run.
	//
	// 	- Period: The command is run on a schedule.
	//
	// 	- NextRebootOnly: The command is run the next time the instances start.
	//
	// 	- EveryReboot: The command is run every time the instances start.
	//
	// This parameter is empty by default, which indicates that commands run in all modes are queried.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The exit code of the execution. Valid values:
	//
	// 	- For Linux instances, the value is the exit code of the shell process.
	//
	// 	- For Windows instances, the value is the exit code of the batch or PowerShell process.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The execution status on a single instance. Valid values:
	//
	// 	- Pending: The command is being verified or sent.
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
	//     	- Scheduled task: The last execution was complete, the exit code was 0, and the specified period ended.
	//
	// 	- Failed:
	//
	//     	- One-time task: The execution was complete, but the exit code was not 0.
	//
	//     	- Scheduled task: The last execution was complete, but the exit code was not 0. The specified period is about to end.
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
	//     	- Scheduled task: The command is waiting to be run.
	//
	// example:
	//
	// /home/
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetContainerName() *string {
	return s.ContainerName
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeInstances() *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances {
	return s.InvokeInstances
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeStatus() *string {
	return s.InvokeStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetLauncher() *string {
	return s.Launcher
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetOssOutputDelivery() *string {
	return s.OssOutputDelivery
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetParameters() *string {
	return s.Parameters
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetTags() *DescribeInvocationsResponseBodyInvocationsInvocationTags {
	return s.Tags
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetTimed() *bool {
	return s.Timed
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetTimeout() *int64 {
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandId(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandName(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandType(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetContainerId(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.ContainerId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetContainerName(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.ContainerName = &v
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeInstances(v *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeInstances = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetLauncher(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Launcher = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetOssOutputDelivery(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.OssOutputDelivery = &v
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetTags(v *DescribeInvocationsResponseBodyInvocationsInvocationTags) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Tags = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetTerminationMode(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.TerminationMode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetTimed(v bool) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Timed = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetTimeout(v int64) *DescribeInvocationsResponseBodyInvocationsInvocation {
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
	if s.InvokeInstances != nil {
		if err := s.InvokeInstances.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances struct {
	InvokeInstance []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance `json:"InvokeInstance,omitempty" xml:"InvokeInstance,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances) GetInvokeInstance() []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	return s.InvokeInstance
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances) SetInvokeInstance(v []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances {
	s.InvokeInstance = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances) Validate() error {
	if s.InvokeInstance != nil {
		for _, item := range s.InvokeInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance struct {
	// The command description.
	//
	// example:
	//
	// 2019-12-20T06:15:54Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The value of tag N of the command. You can specify up to 20 tag values for the command. The tag value can be an empty string. It can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The instances on which the command was run.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The total number of the commands.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The custom parameters in the command.
	//
	// example:
	//
	// 2019-12-20T06:15:56Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// Finished
	InstanceInvokeStatus *string `json:"InstanceInvokeStatus,omitempty" xml:"InstanceInvokeStatus,omitempty"`
	// The key of tag N of the command. You can specify up to 20 tag keys for the command. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags added can be displayed in the response. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The overall execution status of the command task. The value of this parameter depends on the execution status of the command task on all the involved instances. Valid values:
	//
	// 	- Pending: The command is being verified or sent. When the execution state on at least one instance is Pending, the overall execution state is Pending.
	//
	// 	- Scheduled: The command that is set to run on a schedule was sent and waiting to be run. When the execution state on at least one instance is Scheduled, the overall execution state is Scheduled.
	//
	// 	- Running: The command is being run on the instances. When the execution state on at least one instance is Running, the overall execution state is Running.
	//
	// 	- Success: When the execution state on at least one instance is Success and the execution state on the other instances is Stopped or Success, the overall execution state is Success.
	//
	//     	- One-time task: The execution was complete, and the exit code was 0.
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
	// >  `InvokeStatus` in the response functions similarly to this parameter. We recommend that you check the value of this parameter.
	//
	// example:
	//
	// Finished
	OssOutputStatus *string `json:"OssOutputStatus,omitempty" xml:"OssOutputStatus,omitempty"`
	// Command to execute the Output OSS delivery configuration.
	//
	// example:
	//
	// oss://testBucket/testPrefix/output.txt
	OssOutputUri *string `json:"OssOutputUri,omitempty" xml:"OssOutputUri,omitempty"`
	// Indicates whether the command is to be automatically run.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the command task was created.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// Details about the command executions.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution states of the command.
	//
	// example:
	//
	// 2020-01-19T09:15:47Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// false
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The maximum timeout period for the command execution. Unit: seconds.
	//
	// When a command cannot be run, the command execution times out. When a command execution times out, Cloud Assistant Agent forcefully terminates the command process by canceling the process ID (PID) of the command.
	//
	// example:
	//
	// 2020-01-19T09:15:47Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetDropped() *int32 {
	return s.Dropped
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetInstanceInvokeStatus() *string {
	return s.InstanceInvokeStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetOssOutputStatus() *string {
	return s.OssOutputStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetOssOutputUri() *string {
	return s.OssOutputUri
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetOutput() *string {
	return s.Output
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetRepeats() *int32 {
	return s.Repeats
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetTimed() *bool {
	return s.Timed
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetDropped(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetExitCode(v int64) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetInstanceId(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetInstanceInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.InstanceInvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetOssOutputStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.OssOutputStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetOssOutputUri(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.OssOutputUri = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetStartTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetStopTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetTimed(v bool) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.Timed = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyInvocationsInvocationTags struct {
	Tag []*DescribeInvocationsResponseBodyInvocationsInvocationTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationTags) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTags) GetTag() []*DescribeInvocationsResponseBodyInvocationsInvocationTagsTag {
	return s.Tag
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTags) SetTag(v []*DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) *DescribeInvocationsResponseBodyInvocationsInvocationTags {
	s.Tag = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocationsInvocationTagsTag struct {
	// The command content.
	//
	// 	- If ContentEncoding is set to PlainText in the request, the original command content is returned.
	//
	// 	- If ContentEncoding is set to Base64 in the request, the Base64-encoded command content is returned.
	//
	// example:
	//
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The execution path of the command.
	//
	// example:
	//
	// zhangsan
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) SetTagKey(v string) *DescribeInvocationsResponseBodyInvocationsInvocationTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) SetTagValue(v string) *DescribeInvocationsResponseBodyInvocationsInvocationTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationTagsTag) Validate() error {
	return dara.Validate(s)
}
