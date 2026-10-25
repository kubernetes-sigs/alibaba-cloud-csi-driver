// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaultReportsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListFaultReportsShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListFaultReportsShrinkRequest
	GetNextToken() *string
	SetNodesShrink(v string) *ListFaultReportsShrinkRequest
	GetNodesShrink() *string
	SetStatus(v string) *ListFaultReportsShrinkRequest
	GetStatus() *string
}

type ListFaultReportsShrinkRequest struct {
	// The maximum number of results to return.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// DA4F08D4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of nodes.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The status of the fault report.
	//
	// example:
	//
	// Processing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFaultReportsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFaultReportsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFaultReportsShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListFaultReportsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFaultReportsShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *ListFaultReportsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFaultReportsShrinkRequest) SetMaxResults(v int64) *ListFaultReportsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFaultReportsShrinkRequest) SetNextToken(v string) *ListFaultReportsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListFaultReportsShrinkRequest) SetNodesShrink(v string) *ListFaultReportsShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *ListFaultReportsShrinkRequest) SetStatus(v string) *ListFaultReportsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListFaultReportsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
