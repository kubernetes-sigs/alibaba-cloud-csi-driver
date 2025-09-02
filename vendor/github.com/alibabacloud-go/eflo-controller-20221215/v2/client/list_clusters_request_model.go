// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListClustersRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListClustersRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListClustersRequest
	GetResourceGroupId() *string
	SetTags(v []*ListClustersRequestTags) *ListClustersRequest
	GetTags() []*ListClustersRequestTags
}

type ListClustersRequest struct {
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
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bg6wyoox6jq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListClustersRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersRequest) GetTags() []*ListClustersRequestTags {
	return s.Tags
}

func (s *ListClustersRequest) SetMaxResults(v int64) *ListClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClustersRequest) SetNextToken(v string) *ListClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListClustersRequest) SetResourceGroupId(v string) *ListClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersRequest) SetTags(v []*ListClustersRequestTags) *ListClustersRequest {
	s.Tags = v
	return s
}

func (s *ListClustersRequest) Validate() error {
	return dara.Validate(s)
}

type ListClustersRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// key_aa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value_aa
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequestTags) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListClustersRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListClustersRequestTags) SetKey(v string) *ListClustersRequestTags {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTags) SetValue(v string) *ListClustersRequestTags {
	s.Value = &v
	return s
}

func (s *ListClustersRequestTags) Validate() error {
	return dara.Validate(s)
}
