// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeTaskResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeTaskResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *DescribeTaskResponseBody
	GetCreateTime() *string
	SetMessage(v string) *DescribeTaskResponseBody
	GetMessage() *string
	SetNodeIds(v []*string) *DescribeTaskResponseBody
	GetNodeIds() []*string
	SetRequestId(v string) *DescribeTaskResponseBody
	GetRequestId() *string
	SetSteps(v []*DescribeTaskResponseBodySteps) *DescribeTaskResponseBody
	GetSteps() []*DescribeTaskResponseBodySteps
	SetTaskState(v string) *DescribeTaskResponseBody
	GetTaskState() *string
	SetTaskType(v string) *DescribeTaskResponseBody
	GetTaskType() *string
	SetUpdateTime(v string) *DescribeTaskResponseBody
	GetUpdateTime() *string
}

type DescribeTaskResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The start time. The value is in ISO 8601 format. Z in the example indicates UTC.
	//
	// example:
	//
	// 2022-11-30T02:00:00.852Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The task failure message.
	//
	// example:
	//
	// Releasing [prod_main_mid_26e234cf] in region [cn-beijing] with weight [0]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of node IDs.
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A7FD7411-9395-52E8-AF42-EB3A4A55446D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution steps.
	Steps []*DescribeTaskResponseBodySteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// The task state.
	//
	// example:
	//
	// running
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The task type.
	//
	// example:
	//
	// cut_cluster
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The update time. Uses ISO 8601 format. The Z in the example indicates UTC.
	//
	// example:
	//
	// 2022-11-30T03:40:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeTaskResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTaskResponseBody) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DescribeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskResponseBody) GetSteps() []*DescribeTaskResponseBodySteps {
	return s.Steps
}

func (s *DescribeTaskResponseBody) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeTaskResponseBody) SetClusterId(v string) *DescribeTaskResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetClusterName(v string) *DescribeTaskResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeTaskResponseBody) SetCreateTime(v string) *DescribeTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetMessage(v string) *DescribeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBody) SetNodeIds(v []*string) *DescribeTaskResponseBody {
	s.NodeIds = v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetSteps(v []*DescribeTaskResponseBodySteps) *DescribeTaskResponseBody {
	s.Steps = v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskState(v string) *DescribeTaskResponseBody {
	s.TaskState = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskType(v string) *DescribeTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBody) SetUpdateTime(v string) *DescribeTaskResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) Validate() error {
	if s.Steps != nil {
		for _, item := range s.Steps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTaskResponseBodySteps struct {
	// The step failure message.
	//
	// example:
	//
	// get taskinfo failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The stage tag.
	//
	// example:
	//
	// node scale-in
	StageTag *string `json:"StageTag,omitempty" xml:"StageTag,omitempty"`
	// The start time. Z in the example indicates UTC.
	//
	// example:
	//
	// 2022-11-30T2:00:00.852Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The step name.
	//
	// example:
	//
	// create_vpd
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The step execution state.
	//
	// example:
	//
	// execution_success
	StepState *string `json:"StepState,omitempty" xml:"StepState,omitempty"`
	// The step type.
	//
	// example:
	//
	// normal
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
	// The subtasks.
	SubTasks []*DescribeTaskResponseBodyStepsSubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	// The update time. Uses ISO 8601 format. The Z in the example indicates UTC.
	//
	// example:
	//
	// 2022-11-30T02:20:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBodySteps) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBodySteps) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodySteps) GetMessage() *string {
	return s.Message
}

func (s *DescribeTaskResponseBodySteps) GetStageTag() *string {
	return s.StageTag
}

func (s *DescribeTaskResponseBodySteps) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTaskResponseBodySteps) GetStepName() *string {
	return s.StepName
}

func (s *DescribeTaskResponseBodySteps) GetStepState() *string {
	return s.StepState
}

func (s *DescribeTaskResponseBodySteps) GetStepType() *string {
	return s.StepType
}

func (s *DescribeTaskResponseBodySteps) GetSubTasks() []*DescribeTaskResponseBodyStepsSubTasks {
	return s.SubTasks
}

func (s *DescribeTaskResponseBodySteps) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeTaskResponseBodySteps) SetMessage(v string) *DescribeTaskResponseBodySteps {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStageTag(v string) *DescribeTaskResponseBodySteps {
	s.StageTag = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStartTime(v string) *DescribeTaskResponseBodySteps {
	s.StartTime = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepName(v string) *DescribeTaskResponseBodySteps {
	s.StepName = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepState(v string) *DescribeTaskResponseBodySteps {
	s.StepState = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepType(v string) *DescribeTaskResponseBodySteps {
	s.StepType = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetSubTasks(v []*DescribeTaskResponseBodyStepsSubTasks) *DescribeTaskResponseBodySteps {
	s.SubTasks = v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetUpdateTime(v string) *DescribeTaskResponseBodySteps {
	s.UpdateTime = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) Validate() error {
	if s.SubTasks != nil {
		for _, item := range s.SubTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTaskResponseBodyStepsSubTasks struct {
	// The creation time. The Z in the example indicates UTC.
	//
	// example:
	//
	// 2022-11-30T2:00:00.852Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The subtask failure message.
	//
	// example:
	//
	// Releasing [prod_main_mid_26e234cf] in region [cn-beijing] with weight [0]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The node ID.
	//
	// example:
	//
	// i158805051661047928377
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task execution state.
	//
	// example:
	//
	// running
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The task type.
	//
	// example:
	//
	// cut_node_sub_task
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The update time. Uses ISO 8601 format. The Z in the example indicates UTC.
	//
	// example:
	//
	// 2022-11-30T02:20:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBodyStepsSubTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBodyStepsSubTasks) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetMessage() *string {
	return s.Message
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetCreateTime(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetMessage(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskId(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskState(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskState = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskType(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetUpdateTime(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.UpdateTime = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) Validate() error {
	return dara.Validate(s)
}
