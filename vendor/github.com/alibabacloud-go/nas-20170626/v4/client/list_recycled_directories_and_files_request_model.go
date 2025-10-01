// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecycledDirectoriesAndFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *ListRecycledDirectoriesAndFilesRequest
	GetFileId() *string
	SetFileSystemId(v string) *ListRecycledDirectoriesAndFilesRequest
	GetFileSystemId() *string
	SetMaxResults(v int64) *ListRecycledDirectoriesAndFilesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListRecycledDirectoriesAndFilesRequest
	GetNextToken() *string
}

type ListRecycledDirectoriesAndFilesRequest struct {
	// The ID of the directory that you want to query.
	//
	// You can call the [ListRecentlyRecycledDirectories ](https://help.aliyun.com/document_detail/2412173.html)operation to query the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of files or directories to return for each query.
	//
	// Valid values: 10 to 1000.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// If all the files and directories are incompletely returned in a query, the return value of the NextToken parameter is not empty. In this case, you can specify a valid value for the NextToken parameter to continue the query.
	//
	// example:
	//
	// CJyNARIsMTY5OTI2NjQ3NTEzMjY2OTMwOF8xODA5NF8ufnl0YkROTl9uZXcuaXB5bmI=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListRecycledDirectoriesAndFilesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListRecycledDirectoriesAndFilesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListRecycledDirectoriesAndFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListRecycledDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) SetNextToken(v string) *ListRecycledDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesRequest) Validate() error {
	return dara.Validate(s)
}
