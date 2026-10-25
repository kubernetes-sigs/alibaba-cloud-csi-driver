// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaultReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListFaultReportsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListFaultReportsRequest
	GetNextToken() *string
	SetNodes(v []*string) *ListFaultReportsRequest
	GetNodes() []*string
	SetStatus(v string) *ListFaultReportsRequest
	GetStatus() *string
}

type ListFaultReportsRequest struct {
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
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The status of the fault report.
	//
	// example:
	//
	// Processing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFaultReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFaultReportsRequest) GoString() string {
	return s.String()
}

func (s *ListFaultReportsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListFaultReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFaultReportsRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *ListFaultReportsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFaultReportsRequest) SetMaxResults(v int64) *ListFaultReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFaultReportsRequest) SetNextToken(v string) *ListFaultReportsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFaultReportsRequest) SetNodes(v []*string) *ListFaultReportsRequest {
	s.Nodes = v
	return s
}

func (s *ListFaultReportsRequest) SetStatus(v string) *ListFaultReportsRequest {
	s.Status = &v
	return s
}

func (s *ListFaultReportsRequest) Validate() error {
	return dara.Validate(s)
}
