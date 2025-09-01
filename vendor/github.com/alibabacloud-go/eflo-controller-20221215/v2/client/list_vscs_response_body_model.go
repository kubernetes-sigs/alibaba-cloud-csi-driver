// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVscsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVscsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVscsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVscsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVscsResponseBody
	GetTotalCount() *int32
	SetVscs(v []*ListVscsResponseBodyVscs) *ListVscsResponseBody
	GetVscs() []*ListVscsResponseBodyVscs
}

type ListVscsResponseBody struct {
	// No response is returned. The TotalCount parameter is used.
	//
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token. It can be used in the next request to retrieve a new page of results. If this parameter is empty, no next page exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 03668372-18FF-5959-98D9-6B36A4643C7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VSCs.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The VSCs.
	Vscs []*ListVscsResponseBodyVscs `json:"Vscs,omitempty" xml:"Vscs,omitempty" type:"Repeated"`
}

func (s ListVscsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVscsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVscsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVscsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVscsResponseBody) GetVscs() []*ListVscsResponseBodyVscs {
	return s.Vscs
}

func (s *ListVscsResponseBody) SetMaxResults(v int32) *ListVscsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVscsResponseBody) SetNextToken(v string) *ListVscsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVscsResponseBody) SetRequestId(v string) *ListVscsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVscsResponseBody) SetTotalCount(v int32) *ListVscsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVscsResponseBody) SetVscs(v []*ListVscsResponseBodyVscs) *ListVscsResponseBody {
	s.Vscs = v
	return s
}

func (s *ListVscsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVscsResponseBodyVscs struct {
	// The ID of the Lingjun node.
	//
	// example:
	//
	// e01-cn-fzh47xd7u08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm2zkwhkns57i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The VSC status.
	//
	// Valid values:
	//
	// 	- Creating
	//
	// 	- Normal
	//
	// 	- Deleting
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListVscsResponseBodyVscsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VSC ID.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// The custom name of the VSC.
	//
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// The VSC type. Valid values: primary and standard.
	//
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s ListVscsResponseBodyVscs) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponseBodyVscs) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscs) GetNodeId() *string {
	return s.NodeId
}

func (s *ListVscsResponseBodyVscs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVscsResponseBodyVscs) GetStatus() *string {
	return s.Status
}

func (s *ListVscsResponseBodyVscs) GetTags() []*ListVscsResponseBodyVscsTags {
	return s.Tags
}

func (s *ListVscsResponseBodyVscs) GetVscId() *string {
	return s.VscId
}

func (s *ListVscsResponseBodyVscs) GetVscName() *string {
	return s.VscName
}

func (s *ListVscsResponseBodyVscs) GetVscType() *string {
	return s.VscType
}

func (s *ListVscsResponseBodyVscs) SetNodeId(v string) *ListVscsResponseBodyVscs {
	s.NodeId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetResourceGroupId(v string) *ListVscsResponseBodyVscs {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetStatus(v string) *ListVscsResponseBodyVscs {
	s.Status = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetTags(v []*ListVscsResponseBodyVscsTags) *ListVscsResponseBodyVscs {
	s.Tags = v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscId(v string) *ListVscsResponseBodyVscs {
	s.VscId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscName(v string) *ListVscsResponseBodyVscs {
	s.VscName = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscType(v string) *ListVscsResponseBodyVscs {
	s.VscType = &v
	return s
}

func (s *ListVscsResponseBodyVscs) Validate() error {
	return dara.Validate(s)
}

type ListVscsResponseBodyVscsTags struct {
	// The tag key.
	//
	// example:
	//
	// key001
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value001
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVscsResponseBodyVscsTags) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponseBodyVscsTags) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListVscsResponseBodyVscsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListVscsResponseBodyVscsTags) SetTagKey(v string) *ListVscsResponseBodyVscsTags {
	s.TagKey = &v
	return s
}

func (s *ListVscsResponseBodyVscsTags) SetTagValue(v string) *ListVscsResponseBodyVscsTags {
	s.TagValue = &v
	return s
}

func (s *ListVscsResponseBodyVscsTags) Validate() error {
	return dara.Validate(s)
}
