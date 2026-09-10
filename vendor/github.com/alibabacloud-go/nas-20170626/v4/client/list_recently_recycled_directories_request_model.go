// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentlyRecycledDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ListRecentlyRecycledDirectoriesRequest
	GetFileSystemId() *string
	SetMaxResults(v int64) *ListRecentlyRecycledDirectoriesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListRecentlyRecycledDirectoriesRequest
	GetNextToken() *string
}

type ListRecentlyRecycledDirectoriesRequest struct {
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum number of directories to return per query.
	//
	// Valid values: 10 to 1000.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. You do not need to specify this parameter for the first query.
	//
	// If the query results are not completely returned, the NextToken parameter is returned with a value. You can specify the NextToken value in the next request to continue the query.
	//
	// example:
	//
	// 1256****25
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListRecentlyRecycledDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecentlyRecycledDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListRecentlyRecycledDirectoriesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListRecentlyRecycledDirectoriesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListRecentlyRecycledDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetFileSystemId(v string) *ListRecentlyRecycledDirectoriesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetMaxResults(v int64) *ListRecentlyRecycledDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) SetNextToken(v string) *ListRecentlyRecycledDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecentlyRecycledDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
