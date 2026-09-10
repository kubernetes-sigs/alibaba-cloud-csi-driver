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
	// The FileId of the directory to query.
	//
	// If the recycle bin is empty, you can call this operation with FileId=2 (root directory inode) to verify the reachability of the operation or query the recycle bin content under the root directory. You can obtain other valid FileId values by calling the [ListRecentlyRecycledDirectories](https://help.aliyun.com/document_detail/2412173.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of files or directories returned per query.
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
	// If a single query does not return all files and directories, a non-empty NextToken is returned. You can specify the correct NextToken in subsequent queries to continue listing.
	//
	// example:
	//
	// 1256****25
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
