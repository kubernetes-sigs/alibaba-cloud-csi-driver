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
	// Whether to list only directories.
	//
	// Valid values:
	//
	// - `false` (default): Lists both directories and files.
	//
	// - `true`: Lists only directories.
	//
	// > If you set `StorageType` to `All`, you must set `DirectoryOnly` to `true`.
	//
	// example:
	//
	// false
	DirectoryOnly *bool `json:"DirectoryOnly,omitempty" xml:"DirectoryOnly,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum number of directories or files to return per page.
	//
	// Value range: 10–128
	//
	// Default value: 100
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A continuation token used to retrieve the next page of results when the response is truncated.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The absolute path of the directory.
	//
	// The path must start with a forward slash (/) and exist on the mount target.
	//
	// This parameter is required.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage type.
	//
	// - `InfrequentAccess`: infrequent access.
	//
	// - `Archive`: archive storage.
	//
	// - `All`: all storage types.
	//
	// > If you set `StorageType` to `All`, you must set `DirectoryOnly` to `true`.
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
