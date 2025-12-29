// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTasksResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTasksResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeTasksResponseBody
	GetRequestId() *string
	SetTaskSet(v *DescribeTasksResponseBodyTaskSet) *DescribeTasksResponseBody
	GetTaskSet() *DescribeTasksResponseBodyTaskSet
	SetTotalCount(v int32) *DescribeTasksResponseBody
	GetTotalCount() *int32
}

type DescribeTasksResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
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
	// Details about the tasks.
	TaskSet *DescribeTasksResponseBodyTaskSet `json:"TaskSet,omitempty" xml:"TaskSet,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTasksResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTasksResponseBody) GetTaskSet() *DescribeTasksResponseBodyTaskSet {
	return s.TaskSet
}

func (s *DescribeTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageSize(v int32) *DescribeTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRegionId(v string) *DescribeTasksResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTaskSet(v *DescribeTasksResponseBodyTaskSet) *DescribeTasksResponseBody {
	s.TaskSet = v
	return s
}

func (s *DescribeTasksResponseBody) SetTotalCount(v int32) *DescribeTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTasksResponseBody) Validate() error {
	if s.TaskSet != nil {
		if err := s.TaskSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTasksResponseBodyTaskSet struct {
	Task []*DescribeTasksResponseBodyTaskSetTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBodyTaskSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyTaskSet) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTaskSet) GetTask() []*DescribeTasksResponseBodyTaskSetTask {
	return s.Task
}

func (s *DescribeTasksResponseBodyTaskSet) SetTask(v []*DescribeTasksResponseBodyTaskSetTask) *DescribeTasksResponseBodyTaskSet {
	s.Task = v
	return s
}

func (s *DescribeTasksResponseBodyTaskSet) Validate() error {
	if s.Task != nil {
		for _, item := range s.Task {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTasksResponseBodyTaskSetTask struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2020-11-24T12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The time when the task ended.
	//
	// example:
	//
	// 2020-11-24T12:50Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// m-bp1i8huqm5u7****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Indicates whether the task can be canceled.
	//
	// example:
	//
	// true
	SupportCancel *string `json:"SupportCancel,omitempty" xml:"SupportCancel,omitempty"`
	// The name of the operation that generates the task.
	//
	// example:
	//
	// ImportImage
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-bp1hvgwromzv32iq****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeTasksResponseBodyTaskSetTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyTaskSetTask) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTaskSetTask) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeTasksResponseBodyTaskSetTask) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeTasksResponseBodyTaskSetTask) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTasksResponseBodyTaskSetTask) GetSupportCancel() *string {
	return s.SupportCancel
}

func (s *DescribeTasksResponseBodyTaskSetTask) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTasksResponseBodyTaskSetTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTasksResponseBodyTaskSetTask) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTasksResponseBodyTaskSetTask) SetCreationTime(v string) *DescribeTasksResponseBodyTaskSetTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTaskSetTask) SetFinishedTime(v string) *DescribeTasksResponseBodyTaskSetTask {
	s.FinishedTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTaskSetTask) SetResourceId(v string) *DescribeTasksResponseBodyTaskSetTask {
	s.ResourceId = &v
	return s
}

func (s *DescribeTasksResponseBodyTaskSetTask) SetSupportCancel(v string) *DescribeTasksResponseBodyTaskSetTask {
	s.SupportCancel = &v
	return s
}

func (s *DescribeTasksResponseBodyTaskSetTask) SetTaskAction(v string) *DescribeTasksResponseBodyTaskSetTask {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksResponseBodyTaskSetTask) SetTaskId(v string) *DescribeTasksResponseBodyTaskSetTask {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyTaskSetTask) SetTaskStatus(v string) *DescribeTasksResponseBodyTaskSetTask {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksResponseBodyTaskSetTask) Validate() error {
	return dara.Validate(s)
}
