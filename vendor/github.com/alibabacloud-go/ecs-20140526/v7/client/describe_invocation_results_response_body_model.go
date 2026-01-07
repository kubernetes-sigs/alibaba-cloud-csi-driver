// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocation(v *DescribeInvocationResultsResponseBodyInvocation) *DescribeInvocationResultsResponseBody
	GetInvocation() *DescribeInvocationResultsResponseBodyInvocation
	SetRequestId(v string) *DescribeInvocationResultsResponseBody
	GetRequestId() *string
}

type DescribeInvocationResultsResponseBody struct {
	// The execution status of the command task. Valid values:
	//
	// 	- Running:
	//
	//     	- Scheduled task: Before you stop the scheduled execution of the command, the execution state is always Running.
	//
	//     	- One-time task: If the command is being run on instances, the execution state is Running.
	//
	// 	- Finished:
	//
	//     	- Scheduled task: The execution state can never be Finished.
	//
	//     	- One-time task: The execution is complete on all instances, or the execution is stopped on some instances and is complete on the other instances.
	//
	// 	- Success:
	//
	//     	- One-time task: The execution is complete, and the exit code is 0.
	//
	//     	- Scheduled task: The last execution is complete, the exit code is 0, and the specified period ends.
	//
	// 	- Failed:
	//
	//     	- Scheduled task: The execution state can never be Failed.
	//
	//     	- One-time task: The execution fails on all instances.
	//
	// 	- PartialFailed:
	//
	//     	- Scheduled task: The execution state can never be PartialFailed.
	//
	//     	- One-time task: The execution fails on some instances.
	//
	// 	- Stopped: The task is stopped.
	//
	// 	- Stopping: The task is being stopped.
	Invocation *DescribeInvocationResultsResponseBodyInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Struct"`
	// The ID of the command.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsResponseBody) GetInvocation() *DescribeInvocationResultsResponseBodyInvocation {
	return s.Invocation
}

func (s *DescribeInvocationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationResultsResponseBody) SetInvocation(v *DescribeInvocationResultsResponseBodyInvocation) *DescribeInvocationResultsResponseBody {
	s.Invocation = v
	return s
}

func (s *DescribeInvocationResultsResponseBody) SetRequestId(v string) *DescribeInvocationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationResultsResponseBody) Validate() error {
	if s.Invocation != nil {
		if err := s.Invocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationResultsResponseBodyInvocation struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	InvocationResults *DescribeInvocationResultsResponseBodyInvocationInvocationResults `json:"InvocationResults,omitempty" xml:"InvocationResults,omitempty" type:"Struct"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The encoding mode of the `CommandContent` and `Output` values in the response. Valid values:
	//
	// 	- PlainText: returns the original command content and command output.
	//
	// 	- Base64: returns the Base64-encoded command content and command output.
	//
	// Default value: Base64.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specifies whether to return the results of historical scheduled executions. Valid values:
	//
	// 	- true: returns the results of historical scheduled executions. If you set this parameter to true, you must set InvokeId to the ID of a task that is run on a schedule (RepeatMode set to Period) or on each system startup (RepeatMode set to EveryReboot).
	//
	// 	- false: does not return the results of historical scheduled executions.
	//
	// Default value: false.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvocationResultsResponseBodyInvocation) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsResponseBodyInvocation) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsResponseBodyInvocation) GetInvocationResults() *DescribeInvocationResultsResponseBodyInvocationInvocationResults {
	return s.InvocationResults
}

func (s *DescribeInvocationResultsResponseBodyInvocation) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationResultsResponseBodyInvocation) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInvocationResultsResponseBodyInvocation) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInvocationResultsResponseBodyInvocation) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInvocationResultsResponseBodyInvocation) SetInvocationResults(v *DescribeInvocationResultsResponseBodyInvocationInvocationResults) *DescribeInvocationResultsResponseBodyInvocation {
	s.InvocationResults = v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocation) SetNextToken(v string) *DescribeInvocationResultsResponseBodyInvocation {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocation) SetPageNumber(v int64) *DescribeInvocationResultsResponseBodyInvocation {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocation) SetPageSize(v int64) *DescribeInvocationResultsResponseBodyInvocation {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocation) SetTotalCount(v int64) *DescribeInvocationResultsResponseBodyInvocation {
	s.TotalCount = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocation) Validate() error {
	if s.InvocationResults != nil {
		if err := s.InvocationResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationResultsResponseBodyInvocationInvocationResults struct {
	InvocationResult []*DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult `json:"InvocationResult,omitempty" xml:"InvocationResult,omitempty" type:"Repeated"`
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResults) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResults) GetInvocationResult() []*DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	return s.InvocationResult
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResults) SetInvocationResult(v []*DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) *DescribeInvocationResultsResponseBodyInvocationInvocationResults {
	s.InvocationResult = v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResults) Validate() error {
	if s.InvocationResult != nil {
		for _, item := range s.InvocationResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult struct {
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
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// Command to execute the Output OSS delivery configuration.
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
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
	//     	- Scheduled task: The last execution was complete, but the exit code was not 0. The specified period was about to end.
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
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The time when the command started to be run on the instance.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The key of tag N of the command task. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags added can be displayed in the response. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The total number of the commands.
	//
	// example:
	//
	// 2019-12-20T06:15:56Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The value of tag N of the command task. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The tag of the command task.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The number of times that the command was run on the instance.
	//
	// 	- If the command is set to run only once, the value is 0 or 1.
	//
	// 	- If the command is set to run on a schedule, the value is the number of times that the command has been run on the instance.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The page number.
	//
	// example:
	//
	// Running
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The exit code of the command task.
	//
	// 	- For Linux instances, the value is the exit code of the shell command.
	//
	// 	- For Windows instances, the value is the exit code of the batch or PowerShell command.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The tags of the command task.
	//
	// example:
	//
	// oss://testBucket/testPrefix
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	// The execution results.
	//
	// example:
	//
	// Finished
	OssOutputStatus *string `json:"OssOutputStatus,omitempty" xml:"OssOutputStatus,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// oss://testBucket/testPrefix/output.txt
	OssOutputUri *string `json:"OssOutputUri,omitempty" xml:"OssOutputUri,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// MTU6MzA6MDEK
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// Details about the execution results.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// root
	//
	// example:
	//
	// 2020-01-19T09:15:47Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The time when the command task was completed. If the command task times out, the end time is equal to the start time of the command task specified by `StartTime` plus the timeout period specified by `Timeout`.
	Tags *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The execution status of the command. Valid values:
	//
	// 	- Running:
	//
	//     	- Scheduled task: Before you stop the scheduled execution of the command, the execution state is always Running.
	//
	//     	- One-time task: If the command is being run on instances, the execution state is Running.
	//
	// 	- Finished:
	//
	//     	- Scheduled task: The execution state can never be Finished.
	//
	//     	- One-time task: The execution was complete on all instances, or the execution was stopped on some instances and was complete on the other instances.
	//
	// 	- Failed:
	//
	//     	- Scheduled task: The execution state can never be Failed.
	//
	//     	- One-time task: The execution failed on all instances.
	//
	// 	- PartialFailed:
	//
	//     	- Scheduled task: The execution state can never be PartialFailed.
	//
	//     	- One-time task: The execution failed on some instances.
	//
	// 	- Stopped: The task was stopped.
	//
	// 	- Stopping: The task is being stopped.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// The size of the Output text that was truncated and discarded because the `Output` value exceeded 24 KB in size.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetDropped() *int32 {
	return s.Dropped
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetLauncher() *string {
	return s.Launcher
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetOssOutputDelivery() *string {
	return s.OssOutputDelivery
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetOssOutputStatus() *string {
	return s.OssOutputStatus
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetOssOutputUri() *string {
	return s.OssOutputUri
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetOutput() *string {
	return s.Output
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetRepeats() *int32 {
	return s.Repeats
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetTags() *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags {
	return s.Tags
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetUsername() *string {
	return s.Username
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetCommandId(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.CommandId = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetContainerId(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.ContainerId = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetContainerName(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.ContainerName = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetDropped(v int32) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetErrorCode(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetErrorInfo(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetExitCode(v int64) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetFinishedTime(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.FinishedTime = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetInstanceId(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetInvocationStatus(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetInvokeId(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetInvokeRecordStatus(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetLauncher(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.Launcher = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetOssOutputDelivery(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.OssOutputDelivery = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetOssOutputStatus(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.OssOutputStatus = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetOssOutputUri(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.OssOutputUri = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetOutput(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.Output = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetRepeats(v int32) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetStartTime(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetStopTime(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetTags(v *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.Tags = v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetTerminationMode(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.TerminationMode = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetUsername(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.Username = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags struct {
	Tag []*DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags) GetTag() []*DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag {
	return s.Tag
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags) SetTag(v []*DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags {
	s.Tag = v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags) Validate() error {
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

type DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag struct {
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
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The username used to run the command on the instance.
	//
	// example:
	//
	// zhangsan
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) SetTagKey(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) SetTagValue(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTagsTag) Validate() error {
	return dara.Validate(s)
}
