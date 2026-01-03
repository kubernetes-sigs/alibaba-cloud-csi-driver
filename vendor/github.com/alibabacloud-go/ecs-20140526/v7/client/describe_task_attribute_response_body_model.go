// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *DescribeTaskAttributeResponseBody
	GetCreationTime() *string
	SetFailedCount(v int32) *DescribeTaskAttributeResponseBody
	GetFailedCount() *int32
	SetFinishedTime(v string) *DescribeTaskAttributeResponseBody
	GetFinishedTime() *string
	SetOperationProgressSet(v *DescribeTaskAttributeResponseBodyOperationProgressSet) *DescribeTaskAttributeResponseBody
	GetOperationProgressSet() *DescribeTaskAttributeResponseBodyOperationProgressSet
	SetRegionId(v string) *DescribeTaskAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeTaskAttributeResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *DescribeTaskAttributeResponseBody
	GetSuccessCount() *int32
	SetSupportCancel(v string) *DescribeTaskAttributeResponseBody
	GetSupportCancel() *string
	SetTaskAction(v string) *DescribeTaskAttributeResponseBody
	GetTaskAction() *string
	SetTaskId(v string) *DescribeTaskAttributeResponseBody
	GetTaskId() *string
	SetTaskProcess(v string) *DescribeTaskAttributeResponseBody
	GetTaskProcess() *string
	SetTaskStatus(v string) *DescribeTaskAttributeResponseBody
	GetTaskStatus() *string
	SetTotalCount(v int32) *DescribeTaskAttributeResponseBody
	GetTotalCount() *int32
}

type DescribeTaskAttributeResponseBody struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2015-11-23T02:13Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The number of failed tasks.
	//
	// example:
	//
	// 0
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 2015-11-23T02:19Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The return data of the task.
	OperationProgressSet *DescribeTaskAttributeResponseBodyOperationProgressSet `json:"OperationProgressSet,omitempty" xml:"OperationProgressSet,omitempty" type:"Struct"`
	// The region ID of the task.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of completed tasks.
	//
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// Indicates whether the task can be canceled by calling the [CancelTask](https://help.aliyun.com/document_detail/25624.html) operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SupportCancel *string `json:"SupportCancel,omitempty" xml:"SupportCancel,omitempty"`
	// The name of the operation that generated the task.
	//
	// example:
	//
	// ExportImage
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-ce946ntx4wr****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The progress of the task.
	//
	// example:
	//
	// 100%
	TaskProcess *string `json:"TaskProcess,omitempty" xml:"TaskProcess,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTaskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeTaskAttributeResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DescribeTaskAttributeResponseBody) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeTaskAttributeResponseBody) GetOperationProgressSet() *DescribeTaskAttributeResponseBodyOperationProgressSet {
	return s.OperationProgressSet
}

func (s *DescribeTaskAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTaskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskAttributeResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeTaskAttributeResponseBody) GetSupportCancel() *string {
	return s.SupportCancel
}

func (s *DescribeTaskAttributeResponseBody) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTaskAttributeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskAttributeResponseBody) GetTaskProcess() *string {
	return s.TaskProcess
}

func (s *DescribeTaskAttributeResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTaskAttributeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTaskAttributeResponseBody) SetCreationTime(v string) *DescribeTaskAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetFailedCount(v int32) *DescribeTaskAttributeResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetFinishedTime(v string) *DescribeTaskAttributeResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetOperationProgressSet(v *DescribeTaskAttributeResponseBodyOperationProgressSet) *DescribeTaskAttributeResponseBody {
	s.OperationProgressSet = v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetRegionId(v string) *DescribeTaskAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetRequestId(v string) *DescribeTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetSuccessCount(v int32) *DescribeTaskAttributeResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetSupportCancel(v string) *DescribeTaskAttributeResponseBody {
	s.SupportCancel = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetTaskAction(v string) *DescribeTaskAttributeResponseBody {
	s.TaskAction = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetTaskId(v string) *DescribeTaskAttributeResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetTaskProcess(v string) *DescribeTaskAttributeResponseBody {
	s.TaskProcess = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetTaskStatus(v string) *DescribeTaskAttributeResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) SetTotalCount(v int32) *DescribeTaskAttributeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTaskAttributeResponseBody) Validate() error {
	if s.OperationProgressSet != nil {
		if err := s.OperationProgressSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTaskAttributeResponseBodyOperationProgressSet struct {
	OperationProgress []*DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress `json:"OperationProgress,omitempty" xml:"OperationProgress,omitempty" type:"Repeated"`
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSet) GoString() string {
	return s.String()
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSet) GetOperationProgress() []*DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress {
	return s.OperationProgress
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSet) SetOperationProgress(v []*DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) *DescribeTaskAttributeResponseBodyOperationProgressSet {
	s.OperationProgress = v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSet) Validate() error {
	if s.OperationProgress != nil {
		for _, item := range s.OperationProgress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress struct {
	// The error code.
	//
	// example:
	//
	// ParameterInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified RegionId parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The status of the operation.
	//
	// example:
	//
	// Success
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// The type of resource information.
	RelatedItemSet *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet `json:"RelatedItemSet,omitempty" xml:"RelatedItemSet,omitempty" type:"Struct"`
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) GoString() string {
	return s.String()
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) GetRelatedItemSet() *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet {
	return s.RelatedItemSet
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) SetErrorCode(v string) *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress {
	s.ErrorCode = &v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) SetErrorMsg(v string) *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) SetOperationStatus(v string) *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress {
	s.OperationStatus = &v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) SetRelatedItemSet(v *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet) *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress {
	s.RelatedItemSet = v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgress) Validate() error {
	if s.RelatedItemSet != nil {
		if err := s.RelatedItemSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet struct {
	RelatedItem []*DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem `json:"RelatedItem,omitempty" xml:"RelatedItem,omitempty" type:"Repeated"`
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet) GoString() string {
	return s.String()
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet) GetRelatedItem() []*DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	return s.RelatedItem
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet) SetRelatedItem(v []*DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet {
	s.RelatedItem = v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSet) Validate() error {
	if s.RelatedItem != nil {
		for _, item := range s.RelatedItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem struct {
	// The name of the related item.
	//
	// example:
	//
	// OSSObject
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the related item.
	//
	// example:
	//
	// MYOSSPRE_m-23f8tcp***_t-23ym6mv***.vhd
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GoString() string {
	return s.String()
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GetName() *string {
	return s.Name
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) GetValue() *string {
	return s.Value
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) SetName(v string) *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	s.Name = &v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) SetValue(v string) *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem {
	s.Value = &v
	return s
}

func (s *DescribeTaskAttributeResponseBodyOperationProgressSetOperationProgressRelatedItemSetRelatedItem) Validate() error {
	return dara.Validate(s)
}
