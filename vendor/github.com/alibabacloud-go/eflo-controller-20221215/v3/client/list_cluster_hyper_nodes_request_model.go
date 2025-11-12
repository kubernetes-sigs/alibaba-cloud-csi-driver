// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterHyperNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterHyperNodesRequest
	GetClusterId() *string
	SetMaxResults(v int64) *ListClusterHyperNodesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListClusterHyperNodesRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListClusterHyperNodesRequest
	GetNodeGroupId() *string
	SetResourceGroupId(v string) *ListClusterHyperNodesRequest
	GetResourceGroupId() *string
	SetTags(v []*ListClusterHyperNodesRequestTags) *ListClusterHyperNodesRequest
	GetTags() []*ListClusterHyperNodesRequestTags
}

type ListClusterHyperNodesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// rg-acfmywpvugkh7kq
	ResourceGroupId *string                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*ListClusterHyperNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClusterHyperNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHyperNodesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterHyperNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterHyperNodesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListClusterHyperNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClusterHyperNodesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListClusterHyperNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClusterHyperNodesRequest) GetTags() []*ListClusterHyperNodesRequestTags {
	return s.Tags
}

func (s *ListClusterHyperNodesRequest) SetClusterId(v string) *ListClusterHyperNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterHyperNodesRequest) SetMaxResults(v int64) *ListClusterHyperNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClusterHyperNodesRequest) SetNextToken(v string) *ListClusterHyperNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListClusterHyperNodesRequest) SetNodeGroupId(v string) *ListClusterHyperNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListClusterHyperNodesRequest) SetResourceGroupId(v string) *ListClusterHyperNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClusterHyperNodesRequest) SetTags(v []*ListClusterHyperNodesRequestTags) *ListClusterHyperNodesRequest {
	s.Tags = v
	return s
}

func (s *ListClusterHyperNodesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterHyperNodesRequestTags struct {
	// example:
	//
	// my_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// my_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterHyperNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHyperNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListClusterHyperNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListClusterHyperNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListClusterHyperNodesRequestTags) SetKey(v string) *ListClusterHyperNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListClusterHyperNodesRequestTags) SetValue(v string) *ListClusterHyperNodesRequestTags {
	s.Value = &v
	return s
}

func (s *ListClusterHyperNodesRequestTags) Validate() error {
	return dara.Validate(s)
}
