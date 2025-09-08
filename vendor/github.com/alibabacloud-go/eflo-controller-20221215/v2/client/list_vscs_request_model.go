// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVscsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVscsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVscsRequest
	GetNextToken() *string
	SetNodeIds(v []*string) *ListVscsRequest
	GetNodeIds() []*string
	SetResourceGroupId(v string) *ListVscsRequest
	GetResourceGroupId() *string
	SetTag(v []*ListVscsRequestTag) *ListVscsRequest
	GetTag() []*ListVscsRequestTag
	SetVscName(v string) *ListVscsRequest
	GetVscName() *string
}

type ListVscsRequest struct {
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
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*ListVscsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VSC name.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
}

func (s ListVscsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVscsRequest) GoString() string {
	return s.String()
}

func (s *ListVscsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVscsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVscsRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *ListVscsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVscsRequest) GetTag() []*ListVscsRequestTag {
	return s.Tag
}

func (s *ListVscsRequest) GetVscName() *string {
	return s.VscName
}

func (s *ListVscsRequest) SetMaxResults(v int32) *ListVscsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVscsRequest) SetNextToken(v string) *ListVscsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVscsRequest) SetNodeIds(v []*string) *ListVscsRequest {
	s.NodeIds = v
	return s
}

func (s *ListVscsRequest) SetResourceGroupId(v string) *ListVscsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsRequest) SetTag(v []*ListVscsRequestTag) *ListVscsRequest {
	s.Tag = v
	return s
}

func (s *ListVscsRequest) SetVscName(v string) *ListVscsRequest {
	s.VscName = &v
	return s
}

func (s *ListVscsRequest) Validate() error {
	return dara.Validate(s)
}

type ListVscsRequestTag struct {
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

func (s ListVscsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVscsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVscsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVscsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVscsRequestTag) SetKey(v string) *ListVscsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVscsRequestTag) SetValue(v string) *ListVscsRequestTag {
	s.Value = &v
	return s
}

func (s *ListVscsRequestTag) Validate() error {
	return dara.Validate(s)
}
