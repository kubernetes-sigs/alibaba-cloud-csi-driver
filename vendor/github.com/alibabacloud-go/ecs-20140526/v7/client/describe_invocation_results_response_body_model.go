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
	CommandId          *string                                                                               `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	ContainerId        *string                                                                               `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	ContainerName      *string                                                                               `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Dropped            *int32                                                                                `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	ErrorCode          *string                                                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo          *string                                                                               `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ExitCode           *int64                                                                                `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishedTime       *string                                                                               `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	InstanceId         *string                                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvocationStatus   *string                                                                               `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeId           *string                                                                               `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	InvokeRecordStatus *string                                                                               `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	Launcher           *string                                                                               `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	OssOutputDelivery  *string                                                                               `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OssOutputErrorCode *string                                                                               `json:"OssOutputErrorCode,omitempty" xml:"OssOutputErrorCode,omitempty"`
	OssOutputErrorInfo *string                                                                               `json:"OssOutputErrorInfo,omitempty" xml:"OssOutputErrorInfo,omitempty"`
	OssOutputStatus    *string                                                                               `json:"OssOutputStatus,omitempty" xml:"OssOutputStatus,omitempty"`
	OssOutputUri       *string                                                                               `json:"OssOutputUri,omitempty" xml:"OssOutputUri,omitempty"`
	Output             *string                                                                               `json:"Output,omitempty" xml:"Output,omitempty"`
	Repeats            *int32                                                                                `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	StartTime          *string                                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StopTime           *string                                                                               `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	Tags               *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResultTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TerminationMode    *string                                                                               `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	Username           *string                                                                               `json:"Username,omitempty" xml:"Username,omitempty"`
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

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetOssOutputErrorCode() *string {
	return s.OssOutputErrorCode
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) GetOssOutputErrorInfo() *string {
	return s.OssOutputErrorInfo
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

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetOssOutputErrorCode(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.OssOutputErrorCode = &v
	return s
}

func (s *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult) SetOssOutputErrorInfo(v string) *DescribeInvocationResultsResponseBodyInvocationInvocationResultsInvocationResult {
	s.OssOutputErrorInfo = &v
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
