// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetTestResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListNetTestResultsRequest
	GetMaxResults() *int64
	SetNetTestType(v string) *ListNetTestResultsRequest
	GetNetTestType() *string
	SetNextToken(v string) *ListNetTestResultsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListNetTestResultsRequest
	GetResourceGroupId() *string
}

type ListNetTestResultsRequest struct {
	// The number of entries to return on each page. Maximum value: 100.
	//
	// Default value:
	//
	// - If you do not set this parameter or you set it to a value less than 20, the default value is 20.
	//
	// - If you set the value to greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The type of the network test.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the value of \\`NextToken\\` that is returned in the last response.
	//
	// example:
	//
	// 3a6b93229825ac667104463b5679****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxno4vh5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNetTestResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsRequest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListNetTestResultsRequest) GetNetTestType() *string {
	return s.NetTestType
}

func (s *ListNetTestResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetTestResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNetTestResultsRequest) SetMaxResults(v int64) *ListNetTestResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNetTestResultsRequest) SetNetTestType(v string) *ListNetTestResultsRequest {
	s.NetTestType = &v
	return s
}

func (s *ListNetTestResultsRequest) SetNextToken(v string) *ListNetTestResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNetTestResultsRequest) SetResourceGroupId(v string) *ListNetTestResultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNetTestResultsRequest) Validate() error {
	return dara.Validate(s)
}
