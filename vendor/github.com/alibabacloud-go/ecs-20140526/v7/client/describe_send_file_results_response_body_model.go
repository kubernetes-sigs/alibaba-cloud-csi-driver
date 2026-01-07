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
	SetNextToken(v string) *DescribeSendFileResultsResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeSendFileResultsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSendFileResultsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSendFileResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSendFileResultsResponseBody
	GetTotalCount() *int64
}

type DescribeSendFileResultsResponseBody struct {
	// The file sending records.
	Invocations *DescribeSendFileResultsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of file sending tasks queried.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeSendFileResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSendFileResultsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSendFileResultsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSendFileResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSendFileResultsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSendFileResultsResponseBody) SetInvocations(v *DescribeSendFileResultsResponseBodyInvocations) *DescribeSendFileResultsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetNextToken(v string) *DescribeSendFileResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetPageNumber(v int64) *DescribeSendFileResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetPageSize(v int64) *DescribeSendFileResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetRequestId(v string) *DescribeSendFileResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetTotalCount(v int64) *DescribeSendFileResultsResponseBody {
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
	// The content of the file.
	//
	// example:
	//
	// #!/bin/bash  echo "Current User is :"  echo $(ps | grep "$$" | awk \\"{print $2}\\")
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the file content. Valid values:
	//
	// 	- PlainText: The file content is not encoded.
	//
	// 	- Base64: The file content is encoded in Base64.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The time when the file sending task was created.
	//
	// example:
	//
	// 2019-12-20T06:15:54Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the file.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The group of the file.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions on the file.
	//
	// example:
	//
	// 777
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file.
	//
	// example:
	//
	// test
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The overall sending status of the file. The overall sending status of the file varies based on the sending status of the file on all destination instances. Valid values:
	//
	// 	- Pending: The file is being verified or sent. If the sending state of the file on at least one instance is Pending, the overall sending state of the file is Pending.
	//
	// 	- Running: The file is being sent to the instances. If the sending state of the file on at least one instance is Running, the overall sending state of the file is Running.
	//
	// 	- Success: If the sending state of the file on all instances is Success, the overall sending state of the file is Success.
	//
	// 	- If the sending state of the file on all instances is Failed, the overall sending state of the file is Failed. If the sending state of the file on one or more instances is one of the following values, the overall sending state of the file is Failed:
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
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The ID of the file sending task.
	//
	// example:
	//
	// f-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The destination instances.
	InvokeInstances *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Struct"`
	// The name of the file.
	//
	// example:
	//
	// test.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether a file in the destination directory is overwritten if the file has the same name as the sent file.
	//
	// example:
	//
	// false
	Overwrite *string `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The tags of the file sending task.
	Tags *DescribeSendFileResultsResponseBodyInvocationsInvocationTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The destination directory.
	//
	// example:
	//
	// /home/user
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The number of the destination instances.
	//
	// example:
	//
	// 1
	VmCount *int32 `json:"VmCount,omitempty" xml:"VmCount,omitempty"`
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

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetInvokeInstances() *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances {
	return s.InvokeInstances
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetName() *string {
	return s.Name
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetOverwrite() *string {
	return s.Overwrite
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetTags() *DescribeSendFileResultsResponseBodyInvocationsInvocationTags {
	return s.Tags
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetTargetDir() *string {
	return s.TargetDir
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetVmCount() *int32 {
	return s.VmCount
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

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvokeId(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvokeId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvokeInstances(v *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvokeInstances = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetName(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Name = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetOverwrite(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Overwrite = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetTags(v *DescribeSendFileResultsResponseBodyInvocationsInvocationTags) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Tags = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetTargetDir(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.TargetDir = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetVmCount(v int32) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.VmCount = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) Validate() error {
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

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances struct {
	InvokeInstance []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance `json:"InvokeInstance,omitempty" xml:"InvokeInstance,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances) GetInvokeInstance() []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	return s.InvokeInstance
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances) SetInvokeInstance(v []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances {
	s.InvokeInstance = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstances) Validate() error {
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

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance struct {
	// The creation time of the file sending task.
	//
	// example:
	//
	// 2019-12-20T06:15:54Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error code returned when the file failed to be sent to the instance. Valid values:
	//
	// 	- Null: The file is sent to the instance.
	//
	// 	- InstanceNotExists: The instance does not exist or has been released.
	//
	// 	- InstanceReleased: The instance is released while the file is being sent.
	//
	// 	- InstanceNotRunning: The instance is not running when the file sending task is being created.
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
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the file failed to be sent or the file sending task failed to be executed. Valid values:
	//
	// 	- Null: The file is sent to the instance.
	//
	// 	- the specified instance does not exists
	//
	// 	- the specified instance has been released
	//
	// 	- the instance is not running when create task
	//
	// 	- the specified account does not exists
	//
	// 	- the aliyun service is not running on the instance
	//
	// 	- the aliyun service in the instance does not response
	//
	// 	- the aliyun service in the instance is upgrading now
	//
	// 	- the aliyun service in the instance need upgrade
	//
	// 	- the command delivery has been timeout
	//
	// 	- the file creation is failed due to unknown error
	//
	// 	- the authority of file is invalid
	//
	// 	- File content is empty
	//
	// 	- the content of file is invalid
	//
	// 	- File already exists
	//
	// 	- File name is invalid
	//
	// 	- File path is invalid
	//
	// 	- Owner not exists
	//
	// 	- Group not exists
	//
	// 	- Mode is invalid
	//
	// example:
	//
	// the instance is not running when create task
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The time when the file sending task was completed.
	//
	// example:
	//
	// 2019-12-20T06:15:54Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-uf614fhehhz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the file sending task. Valid values:
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
	// The time when the file sending task started to be executed on the instance.
	//
	// example:
	//
	// 2019-12-20T06:15:54Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the task status was last updated.
	//
	// example:
	//
	// 2019-12-20T06:15:54Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetCreationTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetErrorCode(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetErrorInfo(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetFinishTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.FinishTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetInstanceId(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetInvocationStatus(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetStartTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) SetUpdateTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeInstancesInvokeInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationTags struct {
	Tag []*DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationTags) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTags) GetTag() []*DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag {
	return s.Tag
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTags) SetTag(v []*DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) *DescribeSendFileResultsResponseBodyInvocationsInvocationTags {
	s.Tag = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTags) Validate() error {
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

type DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag struct {
	// The tag key of the file sending task.
	//
	// example:
	//
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the file sending task.
	//
	// example:
	//
	// zhangsan
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) SetTagKey(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) SetTagValue(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationTagsTag) Validate() error {
	return dara.Validate(s)
}
