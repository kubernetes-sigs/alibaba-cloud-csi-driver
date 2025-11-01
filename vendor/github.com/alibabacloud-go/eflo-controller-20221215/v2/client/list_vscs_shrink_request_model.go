// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVscsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVscsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVscsShrinkRequest
	GetNextToken() *string
	SetNodeIdsShrink(v string) *ListVscsShrinkRequest
	GetNodeIdsShrink() *string
	SetResourceGroupId(v string) *ListVscsShrinkRequest
	GetResourceGroupId() *string
	SetTag(v []*ListVscsShrinkRequestTag) *ListVscsShrinkRequest
	GetTag() []*ListVscsShrinkRequestTag
	SetVscName(v string) *ListVscsShrinkRequest
	GetVscName() *string
}

type ListVscsShrinkRequest struct {
	// The maximum number of data entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of the nodes.
	NodeIdsShrink *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListVscsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VSC name.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
}

func (s ListVscsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVscsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVscsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVscsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVscsShrinkRequest) GetNodeIdsShrink() *string {
	return s.NodeIdsShrink
}

func (s *ListVscsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVscsShrinkRequest) GetTag() []*ListVscsShrinkRequestTag {
	return s.Tag
}

func (s *ListVscsShrinkRequest) GetVscName() *string {
	return s.VscName
}

func (s *ListVscsShrinkRequest) SetMaxResults(v int32) *ListVscsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVscsShrinkRequest) SetNextToken(v string) *ListVscsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListVscsShrinkRequest) SetNodeIdsShrink(v string) *ListVscsShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

func (s *ListVscsShrinkRequest) SetResourceGroupId(v string) *ListVscsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsShrinkRequest) SetTag(v []*ListVscsShrinkRequestTag) *ListVscsShrinkRequest {
	s.Tag = v
	return s
}

func (s *ListVscsShrinkRequest) SetVscName(v string) *ListVscsShrinkRequest {
	s.VscName = &v
	return s
}

func (s *ListVscsShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVscsShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVscsShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVscsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *ListVscsShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVscsShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVscsShrinkRequestTag) SetKey(v string) *ListVscsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *ListVscsShrinkRequestTag) SetValue(v string) *ListVscsShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *ListVscsShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
