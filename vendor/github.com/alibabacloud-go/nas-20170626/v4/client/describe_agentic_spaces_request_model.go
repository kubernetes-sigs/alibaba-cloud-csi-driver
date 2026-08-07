// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticSpacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeAgenticSpacesRequest
	GetFileSystemId() *string
	SetFilters(v []*DescribeAgenticSpacesRequestFilters) *DescribeAgenticSpacesRequest
	GetFilters() []*DescribeAgenticSpacesRequestFilters
	SetMaxResults(v int64) *DescribeAgenticSpacesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeAgenticSpacesRequest
	GetNextToken() *string
}

type DescribeAgenticSpacesRequest struct {
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The collection of filter key information.
	Filters []*DescribeAgenticSpacesRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of results per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// MTc3OTkzNTA0Mjg0NTc1MDI4OCM0MDQ0MzA****=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeAgenticSpacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticSpacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticSpacesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeAgenticSpacesRequest) GetFilters() []*DescribeAgenticSpacesRequestFilters {
	return s.Filters
}

func (s *DescribeAgenticSpacesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeAgenticSpacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAgenticSpacesRequest) SetFileSystemId(v string) *DescribeAgenticSpacesRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAgenticSpacesRequest) SetFilters(v []*DescribeAgenticSpacesRequestFilters) *DescribeAgenticSpacesRequest {
	s.Filters = v
	return s
}

func (s *DescribeAgenticSpacesRequest) SetMaxResults(v int64) *DescribeAgenticSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAgenticSpacesRequest) SetNextToken(v string) *DescribeAgenticSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAgenticSpacesRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAgenticSpacesRequestFilters struct {
	// The name of the filter key. Valid values:
	//
	// - AgenticSpaceIds
	//
	// example:
	//
	// AgenticSpaceIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key. Wildcards are not supported.
	//
	// - AgenticSpaceIds: only a single ID is supported.
	//
	// example:
	//
	// 06229oypxjgox0u****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAgenticSpacesRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticSpacesRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeAgenticSpacesRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeAgenticSpacesRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeAgenticSpacesRequestFilters) SetKey(v string) *DescribeAgenticSpacesRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeAgenticSpacesRequestFilters) SetValue(v string) *DescribeAgenticSpacesRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeAgenticSpacesRequestFilters) Validate() error {
	return dara.Validate(s)
}
