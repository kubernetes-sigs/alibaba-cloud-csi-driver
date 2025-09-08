// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterNodesRequest
	GetClusterId() *string
	SetMaxResults(v int64) *ListClusterNodesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListClusterNodesRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListClusterNodesRequest
	GetNodeGroupId() *string
	SetResourceGroupId(v string) *ListClusterNodesRequest
	GetResourceGroupId() *string
	SetTags(v []*ListClusterNodesRequestTags) *ListClusterNodesRequest
	GetTags() []*ListClusterNodesRequestTags
}

type ListClusterNodesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call.
	//
	// example:
	//
	// AAAAAdQ3Z+oPlg49gsr2y8jb6wY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxkxkllss
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListClusterNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterNodesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListClusterNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClusterNodesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListClusterNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClusterNodesRequest) GetTags() []*ListClusterNodesRequestTags {
	return s.Tags
}

func (s *ListClusterNodesRequest) SetClusterId(v string) *ListClusterNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterNodesRequest) SetMaxResults(v int64) *ListClusterNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClusterNodesRequest) SetNextToken(v string) *ListClusterNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListClusterNodesRequest) SetNodeGroupId(v string) *ListClusterNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListClusterNodesRequest) SetResourceGroupId(v string) *ListClusterNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClusterNodesRequest) SetTags(v []*ListClusterNodesRequestTags) *ListClusterNodesRequest {
	s.Tags = v
	return s
}

func (s *ListClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}

type ListClusterNodesRequestTags struct {
	// The tag key for the node.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for the node.
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListClusterNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListClusterNodesRequestTags) SetKey(v string) *ListClusterNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListClusterNodesRequestTags) SetValue(v string) *ListClusterNodesRequestTags {
	s.Value = &v
	return s
}

func (s *ListClusterNodesRequestTags) Validate() error {
	return dara.Validate(s)
}
