// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupRefreshTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodeGroupRefreshTasksRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListNodeGroupRefreshTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupRefreshTasksRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListNodeGroupRefreshTasksRequest
	GetNodeGroupId() *string
	SetStatuses(v []*string) *ListNodeGroupRefreshTasksRequest
	GetStatuses() []*string
}

type ListNodeGroupRefreshTasksRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The maximum number of entries per page for paging. Valid values: 1 to 500. Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-3525
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The list of task statuses. Valid values:
	//
	// - Pending: The refresh task is created and waiting to be executed.
	//
	// - InProgress: The refresh task is being processed.
	//
	// - Success: The refresh task is executed.
	//
	// - Failed: The refresh task failed.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListNodeGroupRefreshTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupRefreshTasksRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeGroupRefreshTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupRefreshTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupRefreshTasksRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListNodeGroupRefreshTasksRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListNodeGroupRefreshTasksRequest) SetClusterId(v string) *ListNodeGroupRefreshTasksRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupRefreshTasksRequest) SetMaxResults(v int32) *ListNodeGroupRefreshTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupRefreshTasksRequest) SetNextToken(v string) *ListNodeGroupRefreshTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupRefreshTasksRequest) SetNodeGroupId(v string) *ListNodeGroupRefreshTasksRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupRefreshTasksRequest) SetStatuses(v []*string) *ListNodeGroupRefreshTasksRequest {
	s.Statuses = v
	return s
}

func (s *ListNodeGroupRefreshTasksRequest) Validate() error {
	return dara.Validate(s)
}
