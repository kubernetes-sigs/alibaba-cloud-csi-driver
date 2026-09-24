// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupRefreshTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodeGroupRefreshTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupRefreshTasksResponseBody
	GetNextToken() *string
	SetNodeGroupRefreshTasks(v []*ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) *ListNodeGroupRefreshTasksResponseBody
	GetNodeGroupRefreshTasks() []*ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks
	SetRequestId(v string) *ListNodeGroupRefreshTasksResponseBody
	GetRequestId() *string
}

type ListNodeGroupRefreshTasksResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. An empty value indicates that no more results exist.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of node group refresh tasks.
	NodeGroupRefreshTasks []*ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks `json:"NodeGroupRefreshTasks,omitempty" xml:"NodeGroupRefreshTasks,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 03668372-18FF-5959-98D9-6B36A4643C7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeGroupRefreshTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupRefreshTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupRefreshTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupRefreshTasksResponseBody) GetNodeGroupRefreshTasks() []*ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	return s.NodeGroupRefreshTasks
}

func (s *ListNodeGroupRefreshTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeGroupRefreshTasksResponseBody) SetMaxResults(v int32) *ListNodeGroupRefreshTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBody) SetNextToken(v string) *ListNodeGroupRefreshTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBody) SetNodeGroupRefreshTasks(v []*ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) *ListNodeGroupRefreshTasksResponseBody {
	s.NodeGroupRefreshTasks = v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBody) SetRequestId(v string) *ListNodeGroupRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBody) Validate() error {
	if s.NodeGroupRefreshTasks != nil {
		for _, item := range s.NodeGroupRefreshTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks struct {
	// The end time of the refresh task in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-20T10:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of failed nodes.
	//
	// example:
	//
	// 2
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The number of finished nodes, including succeeded, failed, and skipped nodes.
	//
	// example:
	//
	// 45
	FinishedCount *int64 `json:"FinishedCount,omitempty" xml:"FinishedCount,omitempty"`
	// The maximum disruptive action level allowed for the refresh operation.
	//
	// example:
	//
	// Refresh
	MaxDisruptiveAction *string `json:"MaxDisruptiveAction,omitempty" xml:"MaxDisruptiveAction,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-3525
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// task-159136551662516768776
	NodeGroupRefreshTaskId *string `json:"NodeGroupRefreshTaskId,omitempty" xml:"NodeGroupRefreshTaskId,omitempty"`
	// The start time of the refresh task in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-20T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - Pending: The refresh task is created and waiting to be executed.
	//
	// - InProgress: The refresh task is being processed.
	//
	// - Success: The refresh task is executed.
	//
	// - Failed: The refresh task failed.
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of nodes to refresh in this task.
	//
	// example:
	//
	// 100
	TotalNodeCount *int64 `json:"TotalNodeCount,omitempty" xml:"TotalNodeCount,omitempty"`
}

func (s ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GoString() string {
	return s.String()
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetEndTime() *string {
	return s.EndTime
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetFinishedCount() *int64 {
	return s.FinishedCount
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetMaxDisruptiveAction() *string {
	return s.MaxDisruptiveAction
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetNodeGroupRefreshTaskId() *string {
	return s.NodeGroupRefreshTaskId
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetStartTime() *string {
	return s.StartTime
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetStatus() *string {
	return s.Status
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) GetTotalNodeCount() *int64 {
	return s.TotalNodeCount
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetEndTime(v string) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.EndTime = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetFailedCount(v int64) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.FailedCount = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetFinishedCount(v int64) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.FinishedCount = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetMaxDisruptiveAction(v string) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.MaxDisruptiveAction = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetNodeGroupId(v string) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetNodeGroupRefreshTaskId(v string) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.NodeGroupRefreshTaskId = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetStartTime(v string) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.StartTime = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetStatus(v string) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.Status = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) SetTotalNodeCount(v int64) *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks {
	s.TotalNodeCount = &v
	return s
}

func (s *ListNodeGroupRefreshTasksResponseBodyNodeGroupRefreshTasks) Validate() error {
	return dara.Validate(s)
}
