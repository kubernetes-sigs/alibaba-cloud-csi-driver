// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupDriftedNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodeGroupDriftedNodesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupDriftedNodesRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListNodeGroupDriftedNodesRequest
	GetNodeGroupId() *string
	SetNodeIds(v []*string) *ListNodeGroupDriftedNodesRequest
	GetNodeIds() []*string
}

type ListNodeGroupDriftedNodesRequest struct {
	// The maximum number of entries per page for a paged query. Valid values: 1 to 500. Default value: 100.
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
	// The ID of the node group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3525
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// Limits the check scope. If not specified, all nodes in the node group are checked. <warning>If the model is a super node, pass the TrayNode ID.</warning>
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s ListNodeGroupDriftedNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupDriftedNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupDriftedNodesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupDriftedNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupDriftedNodesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListNodeGroupDriftedNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *ListNodeGroupDriftedNodesRequest) SetMaxResults(v int32) *ListNodeGroupDriftedNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupDriftedNodesRequest) SetNextToken(v string) *ListNodeGroupDriftedNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupDriftedNodesRequest) SetNodeGroupId(v string) *ListNodeGroupDriftedNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupDriftedNodesRequest) SetNodeIds(v []*string) *ListNodeGroupDriftedNodesRequest {
	s.NodeIds = v
	return s
}

func (s *ListNodeGroupDriftedNodesRequest) Validate() error {
	return dara.Validate(s)
}
