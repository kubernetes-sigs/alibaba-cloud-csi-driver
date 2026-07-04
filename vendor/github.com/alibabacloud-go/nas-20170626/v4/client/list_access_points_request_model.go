// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ListAccessPointsRequest
	GetFileSystemId() *string
	SetFilters(v []*ListAccessPointsRequestFilters) *ListAccessPointsRequest
	GetFilters() []*ListAccessPointsRequestFilters
	SetMaxResults(v int32) *ListAccessPointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccessPointsRequest
	GetNextToken() *string
}

type ListAccessPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0913nx15amuix9a****
	FileSystemId *string                           `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Filters      []*ListAccessPointsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTY4NzcxOTcwMjAzMDk2Nzc0MyM4MDM4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessPointsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListAccessPointsRequest) GetFilters() []*ListAccessPointsRequestFilters {
	return s.Filters
}

func (s *ListAccessPointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccessPointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessPointsRequest) SetFileSystemId(v string) *ListAccessPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListAccessPointsRequest) SetFilters(v []*ListAccessPointsRequestFilters) *ListAccessPointsRequest {
	s.Filters = v
	return s
}

func (s *ListAccessPointsRequest) SetMaxResults(v int32) *ListAccessPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccessPointsRequest) SetNextToken(v string) *ListAccessPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessPointsRequest) Validate() error {
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

type ListAccessPointsRequestFilters struct {
	// example:
	//
	// AccessPointId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccessPointsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListAccessPointsRequestFilters) GetName() *string {
	return s.Name
}

func (s *ListAccessPointsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ListAccessPointsRequestFilters) SetName(v string) *ListAccessPointsRequestFilters {
	s.Name = &v
	return s
}

func (s *ListAccessPointsRequestFilters) SetValue(v string) *ListAccessPointsRequestFilters {
	s.Value = &v
	return s
}

func (s *ListAccessPointsRequestFilters) Validate() error {
	return dara.Validate(s)
}
