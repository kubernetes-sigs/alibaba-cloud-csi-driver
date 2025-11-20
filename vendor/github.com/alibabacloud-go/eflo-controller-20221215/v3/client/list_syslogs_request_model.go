// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyslogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromTime(v string) *ListSyslogsRequest
	GetFromTime() *string
	SetNextToken(v string) *ListSyslogsRequest
	GetNextToken() *string
	SetNodeId(v string) *ListSyslogsRequest
	GetNodeId() *string
	SetQuery(v string) *ListSyslogsRequest
	GetQuery() *string
	SetReverse(v bool) *ListSyslogsRequest
	GetReverse() *bool
	SetToTime(v string) *ListSyslogsRequest
	GetToTime() *string
}

type ListSyslogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1659745800
	FromTime *string `json:"FromTime,omitempty" xml:"FromTime,omitempty"`
	// example:
	//
	// 392e8b4a03ed171433cc39f5b464ec9d
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-nwy37atbj44
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// *
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1665369329
	ToTime *string `json:"ToTime,omitempty" xml:"ToTime,omitempty"`
}

func (s ListSyslogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsRequest) GoString() string {
	return s.String()
}

func (s *ListSyslogsRequest) GetFromTime() *string {
	return s.FromTime
}

func (s *ListSyslogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSyslogsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListSyslogsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListSyslogsRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListSyslogsRequest) GetToTime() *string {
	return s.ToTime
}

func (s *ListSyslogsRequest) SetFromTime(v string) *ListSyslogsRequest {
	s.FromTime = &v
	return s
}

func (s *ListSyslogsRequest) SetNextToken(v string) *ListSyslogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSyslogsRequest) SetNodeId(v string) *ListSyslogsRequest {
	s.NodeId = &v
	return s
}

func (s *ListSyslogsRequest) SetQuery(v string) *ListSyslogsRequest {
	s.Query = &v
	return s
}

func (s *ListSyslogsRequest) SetReverse(v bool) *ListSyslogsRequest {
	s.Reverse = &v
	return s
}

func (s *ListSyslogsRequest) SetToTime(v string) *ListSyslogsRequest {
	s.ToTime = &v
	return s
}

func (s *ListSyslogsRequest) Validate() error {
	return dara.Validate(s)
}
