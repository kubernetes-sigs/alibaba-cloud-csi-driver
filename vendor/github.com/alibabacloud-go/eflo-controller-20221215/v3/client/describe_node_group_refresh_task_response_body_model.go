// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupRefreshTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetEndTime() *string
	SetFailedCount(v int64) *DescribeNodeGroupRefreshTaskResponseBody
	GetFailedCount() *int64
	SetFinishedCount(v int64) *DescribeNodeGroupRefreshTaskResponseBody
	GetFinishedCount() *int64
	SetMaxDisruptiveAction(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetMaxDisruptiveAction() *string
	SetMaxResults(v int32) *DescribeNodeGroupRefreshTaskResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetNextToken() *string
	SetNodeGroupId(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetNodeGroupId() *string
	SetNodeGroupRefreshTaskId(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetNodeGroupRefreshTaskId() *string
	SetNodes(v []*DescribeNodeGroupRefreshTaskResponseBodyNodes) *DescribeNodeGroupRefreshTaskResponseBody
	GetNodes() []*DescribeNodeGroupRefreshTaskResponseBodyNodes
	SetRequestId(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeNodeGroupRefreshTaskResponseBody
	GetStatus() *string
	SetTotalNodeCount(v int64) *DescribeNodeGroupRefreshTaskResponseBody
	GetTotalNodeCount() *int64
}

type DescribeNodeGroupRefreshTaskResponseBody struct {
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
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. An empty value of NextToken indicates that no more results exist.
	//
	// example:
	//
	// 7ed93fda-5e7f-436a-ae5a-bd8e6b04e36b
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the node group.
	//
	// example:
	//
	// ng-3525
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The ID of the refresh task.
	//
	// example:
	//
	// task-159136551662516768776
	NodeGroupRefreshTaskId *string `json:"NodeGroupRefreshTaskId,omitempty" xml:"NodeGroupRefreshTaskId,omitempty"`
	// The list of nodes.
	Nodes []*DescribeNodeGroupRefreshTaskResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the refresh task in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-20T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - Pending: the refresh task is created and waiting to be executed.
	//
	// - InProgress: the refresh task is being processed.
	//
	// - Success: the refresh task is executed.
	//
	// - Failed: the refresh task failed to be executed.
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of nodes to be refreshed in the task.
	//
	// example:
	//
	// 100
	TotalNodeCount *int64 `json:"TotalNodeCount,omitempty" xml:"TotalNodeCount,omitempty"`
}

func (s DescribeNodeGroupRefreshTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupRefreshTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetFinishedCount() *int64 {
	return s.FinishedCount
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetMaxDisruptiveAction() *string {
	return s.MaxDisruptiveAction
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetNodeGroupRefreshTaskId() *string {
	return s.NodeGroupRefreshTaskId
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetNodes() []*DescribeNodeGroupRefreshTaskResponseBodyNodes {
	return s.Nodes
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) GetTotalNodeCount() *int64 {
	return s.TotalNodeCount
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetEndTime(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetFailedCount(v int64) *DescribeNodeGroupRefreshTaskResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetFinishedCount(v int64) *DescribeNodeGroupRefreshTaskResponseBody {
	s.FinishedCount = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetMaxDisruptiveAction(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.MaxDisruptiveAction = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetMaxResults(v int32) *DescribeNodeGroupRefreshTaskResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetNextToken(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetNodeGroupId(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetNodeGroupRefreshTaskId(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.NodeGroupRefreshTaskId = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetNodes(v []*DescribeNodeGroupRefreshTaskResponseBodyNodes) *DescribeNodeGroupRefreshTaskResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetRequestId(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetStartTime(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetStatus(v string) *DescribeNodeGroupRefreshTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) SetTotalNodeCount(v int64) *DescribeNodeGroupRefreshTaskResponseBody {
	s.TotalNodeCount = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBody) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodeGroupRefreshTaskResponseBodyNodes struct {
	// The action level actually executed on the node. If multiple properties are refreshed on the node, the highest required action level is used. Valid values:
	//
	// - Refresh: in-place refresh.
	//
	// - Reboot: restart.
	//
	// - Reimage: reimage.
	//
	// If the entire node is skipped, this value is empty.
	//
	// example:
	//
	// Refresh
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The error code.
	//
	// example:
	//
	// NodeInMaintenance
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The node has been confirmed for maintenance and there is no need to confirm the maintenance again. sn: 21B917666,status:ClusterNodeRepairing
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// node-001
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The list of property drifts for the node, including both executed and skipped properties.
	PropertyDrifts []*DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts `json:"PropertyDrifts,omitempty" xml:"PropertyDrifts,omitempty" type:"Repeated"`
	// The node refresh status. Valid values:
	//
	// - Pending: the node is waiting to be refreshed.
	//
	// - InProgress: the node is being refreshed.
	//
	// - Success: the node is refreshed.
	//
	// - Failed: the node failed to be refreshed.
	//
	// - Skipped: all properties to be refreshed on the node exceeded the MaxDisruptiveAction constraint and were skipped.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNodeGroupRefreshTaskResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupRefreshTaskResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) GetAction() *string {
	return s.Action
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) GetPropertyDrifts() []*DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts {
	return s.PropertyDrifts
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) GetStatus() *string {
	return s.Status
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) SetAction(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodes {
	s.Action = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) SetErrorCode(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodes {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) SetErrorMessage(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodes {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) SetNodeId(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) SetPropertyDrifts(v []*DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) *DescribeNodeGroupRefreshTaskResponseBodyNodes {
	s.PropertyDrifts = v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) SetStatus(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodes) Validate() error {
	if s.PropertyDrifts != nil {
		for _, item := range s.PropertyDrifts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts struct {
	// The current value of the node property. Complex types are serialized as JSON strings.
	//
	// example:
	//
	// old-role
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// The minimum action required to apply the target value: Refresh / Reboot / Reimage. For more information, see the MaxDisruptiveAction parameter description in the RefreshNodeGroupNodes operation.
	//
	// example:
	//
	// Refresh
	MinRequiredAction *string `json:"MinRequiredAction,omitempty" xml:"MinRequiredAction,omitempty"`
	// The property path in dot notation.
	//
	// example:
	//
	// RamRoleName
	PropertyPath *string `json:"PropertyPath,omitempty" xml:"PropertyPath,omitempty"`
	// Indicates whether the property was skipped because it exceeded the MaxDisruptiveAction constraint.
	//
	// example:
	//
	// false
	Skipped *bool `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The target value of the node property. Complex types are serialized as JSON strings.
	//
	// example:
	//
	// new-role
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) GetActualValue() *string {
	return s.ActualValue
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) GetMinRequiredAction() *string {
	return s.MinRequiredAction
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) GetPropertyPath() *string {
	return s.PropertyPath
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) GetSkipped() *bool {
	return s.Skipped
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) GetTargetValue() *string {
	return s.TargetValue
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) SetActualValue(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts {
	s.ActualValue = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) SetMinRequiredAction(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts {
	s.MinRequiredAction = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) SetPropertyPath(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts {
	s.PropertyPath = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) SetSkipped(v bool) *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts {
	s.Skipped = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) SetTargetValue(v string) *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts {
	s.TargetValue = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskResponseBodyNodesPropertyDrifts) Validate() error {
	return dara.Validate(s)
}
