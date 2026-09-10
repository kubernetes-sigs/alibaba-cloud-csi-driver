// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoriesAndFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryOnly(v bool) *ListDirectoriesAndFilesRequest
	GetDirectoryOnly() *bool
	SetFileSystemId(v string) *ListDirectoriesAndFilesRequest
	GetFileSystemId() *string
	SetMaxResults(v int64) *ListDirectoriesAndFilesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListDirectoriesAndFilesRequest
	GetNextToken() *string
	SetPath(v string) *ListDirectoriesAndFilesRequest
	GetPath() *string
	SetStorageType(v string) *ListDirectoriesAndFilesRequest
	GetStorageType() *string
}

type ListDirectoriesAndFilesRequest struct {
	// Specifies whether to query only directories.
	//
	// Valid values:
	//
	// - false (default): No. Both directories and files can be queried.
	//
	// - true: Yes. Only directories are queried.
	//
	// > When StorageType is set to All, DirectoryOnly must be set to true and cannot be set to false.
	//
	// example:
	//
	// false
	DirectoryOnly *bool `json:"DirectoryOnly,omitempty" xml:"DirectoryOnly,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of directories or files included in each query result.
	//
	// Valid values: 10 to 128.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the return results are truncated, you can use NextToken to initiate a new request to retrieve the content after the current truncation position.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The absolute path of the specified directory.
	//
	// The path must start with a forward slash (/) and must be an existing path in the mount target.
	//
	// This parameter is required.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage class type.
	//
	// - InfrequentAccess: IA storage class.
	//
	// - Archive: Archive storage class.
	//
	// - All: queries data of all storage classes.
	//
	// > When StorageType is set to All, you must set DirectoryOnly to true.
	//
	// This parameter is required.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ListDirectoriesAndFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoriesAndFilesRequest) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesRequest) GetDirectoryOnly() *bool {
	return s.DirectoryOnly
}

func (s *ListDirectoriesAndFilesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListDirectoriesAndFilesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListDirectoriesAndFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDirectoriesAndFilesRequest) GetPath() *string {
	return s.Path
}

func (s *ListDirectoriesAndFilesRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ListDirectoriesAndFilesRequest) SetDirectoryOnly(v bool) *ListDirectoriesAndFilesRequest {
	s.DirectoryOnly = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetFileSystemId(v string) *ListDirectoriesAndFilesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetMaxResults(v int64) *ListDirectoriesAndFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetNextToken(v string) *ListDirectoriesAndFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetPath(v string) *ListDirectoriesAndFilesRequest {
	s.Path = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) SetStorageType(v string) *ListDirectoriesAndFilesRequest {
	s.StorageType = &v
	return s
}

func (s *ListDirectoriesAndFilesRequest) Validate() error {
	return dara.Validate(s)
}
