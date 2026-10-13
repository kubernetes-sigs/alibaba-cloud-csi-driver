// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupRefreshTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeNodeGroupRefreshTaskShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNodeGroupRefreshTaskShrinkRequest
	GetNextToken() *string
	SetNodeGroupRefreshTaskId(v string) *DescribeNodeGroupRefreshTaskShrinkRequest
	GetNodeGroupRefreshTaskId() *string
	SetNodeStatusesShrink(v string) *DescribeNodeGroupRefreshTaskShrinkRequest
	GetNodeStatusesShrink() *string
}

type DescribeNodeGroupRefreshTaskShrinkRequest struct {
	// The maximum number of entries per page for a paged query. Valid values: 1 to 500. Default value: 100. For more information about paging, set this parameter together with NextToken.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request. This parameter is used to paginate through the node list in the current refresh task.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the refresh task.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-159136551662516768776
	NodeGroupRefreshTaskId *string `json:"NodeGroupRefreshTaskId,omitempty" xml:"NodeGroupRefreshTaskId,omitempty"`
	// The node refresh statuses to filter by. Valid values:
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
	NodeStatusesShrink *string `json:"NodeStatuses,omitempty" xml:"NodeStatuses,omitempty"`
}

func (s DescribeNodeGroupRefreshTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupRefreshTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) GetNodeGroupRefreshTaskId() *string {
	return s.NodeGroupRefreshTaskId
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) GetNodeStatusesShrink() *string {
	return s.NodeStatusesShrink
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) SetMaxResults(v int32) *DescribeNodeGroupRefreshTaskShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) SetNextToken(v string) *DescribeNodeGroupRefreshTaskShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) SetNodeGroupRefreshTaskId(v string) *DescribeNodeGroupRefreshTaskShrinkRequest {
	s.NodeGroupRefreshTaskId = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) SetNodeStatusesShrink(v string) *DescribeNodeGroupRefreshTaskShrinkRequest {
	s.NodeStatusesShrink = &v
	return s
}

func (s *DescribeNodeGroupRefreshTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
