// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodeGroupsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListNodeGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupsRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListNodeGroupsRequest
	GetNodeGroupId() *string
}

type ListNodeGroupsRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100.
	//
	// Default value:
	//
	// 	- If the value is not set or is set to less than 20, the default value is 20.
	//
	// 	- If the value is set to greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The NextToken for the next page. Include this return value when requesting the next page.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The node group ID.
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s ListNodeGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupsRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListNodeGroupsRequest) SetClusterId(v string) *ListNodeGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupsRequest) SetMaxResults(v int32) *ListNodeGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNextToken(v string) *ListNodeGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupId(v string) *ListNodeGroupsRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListNodeGroupsRequest) Validate() error {
	return dara.Validate(s)
}
