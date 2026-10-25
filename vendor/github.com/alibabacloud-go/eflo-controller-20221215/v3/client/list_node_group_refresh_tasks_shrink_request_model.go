// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupRefreshTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodeGroupRefreshTasksShrinkRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListNodeGroupRefreshTasksShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupRefreshTasksShrinkRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListNodeGroupRefreshTasksShrinkRequest
	GetNodeGroupId() *string
	SetStatusesShrink(v string) *ListNodeGroupRefreshTasksShrinkRequest
	GetStatusesShrink() *string
}

type ListNodeGroupRefreshTasksShrinkRequest struct {
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
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s ListNodeGroupRefreshTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupRefreshTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) SetClusterId(v string) *ListNodeGroupRefreshTasksShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) SetMaxResults(v int32) *ListNodeGroupRefreshTasksShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) SetNextToken(v string) *ListNodeGroupRefreshTasksShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) SetNodeGroupId(v string) *ListNodeGroupRefreshTasksShrinkRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) SetStatusesShrink(v string) *ListNodeGroupRefreshTasksShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListNodeGroupRefreshTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
