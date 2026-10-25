// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosticResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiagType(v string) *ListDiagnosticResultsRequest
	GetDiagType() *string
	SetMaxResults(v int64) *ListDiagnosticResultsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListDiagnosticResultsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListDiagnosticResultsRequest
	GetResourceGroupId() *string
}

type ListDiagnosticResultsRequest struct {
	// The diagnostic type.
	//
	// example:
	//
	// NetDiag
	DiagType *string `json:"DiagType,omitempty" xml:"DiagType,omitempty"`
	// The number of entries per page in a paged query. Maximum value: 100.
	//
	// Default value:
	//
	// - If this parameter is not specified or is set to a value less than 20, the default value is 20.
	//
	// - If this parameter is set to a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of NextToken that was returned from the previous API call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmywpvugk****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListDiagnosticResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsRequest) GetDiagType() *string {
	return s.DiagType
}

func (s *ListDiagnosticResultsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListDiagnosticResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiagnosticResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDiagnosticResultsRequest) SetDiagType(v string) *ListDiagnosticResultsRequest {
	s.DiagType = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetMaxResults(v int64) *ListDiagnosticResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetNextToken(v string) *ListDiagnosticResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetResourceGroupId(v string) *ListDiagnosticResultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDiagnosticResultsRequest) Validate() error {
	return dara.Validate(s)
}
