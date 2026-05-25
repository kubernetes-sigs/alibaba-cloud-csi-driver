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
	CommandContent     *string                                                              `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandDescription *string                                                              `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	CommandId          *string                                                              `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	CommandName        *string                                                              `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	CommandType        *string                                                              `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	ContainerId        *string                                                              `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	ContainerName      *string                                                              `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	CreationTime       *string                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Frequency          *string                                                              `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InvocationStatus   *string                                                              `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeId           *string                                                              `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	InvokeInstances    *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Struct"`
	InvokeStatus       *string                                                              `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	Launcher           *string                                                              `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	OssOutputDelivery  *string                                                              `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	Parameters         *string                                                              `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RepeatMode         *string                                                              `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	Tags               *DescribeInvocationsResponseBodyInvocationsInvocationTags            `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TerminationMode    *string                                                              `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	Timed              *bool                                                                `json:"Timed,omitempty" xml:"Timed,omitempty"`
	Timeout            *int64                                                               `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Username           *string                                                              `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkingDir         *string                                                              `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
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
	CreationTime         *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Dropped              *int32  `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	ErrorCode            *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo            *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ExitCode             *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishTime           *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceInvokeStatus *string `json:"InstanceInvokeStatus,omitempty" xml:"InstanceInvokeStatus,omitempty"`
	InvocationStatus     *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	OssOutputErrorCode   *string `json:"OssOutputErrorCode,omitempty" xml:"OssOutputErrorCode,omitempty"`
	OssOutputErrorInfo   *string `json:"OssOutputErrorInfo,omitempty" xml:"OssOutputErrorInfo,omitempty"`
	OssOutputStatus      *string `json:"OssOutputStatus,omitempty" xml:"OssOutputStatus,omitempty"`
	OssOutputUri         *string `json:"OssOutputUri,omitempty" xml:"OssOutputUri,omitempty"`
	Output               *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Repeats              *int32  `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StopTime             *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	Timed                *bool   `json:"Timed,omitempty" xml:"Timed,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetOssOutputErrorCode() *string {
	return s.OssOutputErrorCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetOssOutputErrorInfo() *string {
	return s.OssOutputErrorInfo
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetOssOutputErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.OssOutputErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetOssOutputErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.OssOutputErrorInfo = &v
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
