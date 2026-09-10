// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataInsightDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectory(v *ListDataInsightDirectoriesResponseBodyDirectory) *ListDataInsightDirectoriesResponseBody
	GetDirectory() *ListDataInsightDirectoriesResponseBodyDirectory
	SetFileSystemId(v string) *ListDataInsightDirectoriesResponseBody
	GetFileSystemId() *string
	SetMaxResults(v int32) *ListDataInsightDirectoriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataInsightDirectoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataInsightDirectoriesResponseBody
	GetRequestId() *string
}

type ListDataInsightDirectoriesResponseBody struct {
	Directory *ListDataInsightDirectoriesResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// /subDir
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataInsightDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataInsightDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataInsightDirectoriesResponseBody) GetDirectory() *ListDataInsightDirectoriesResponseBodyDirectory {
	return s.Directory
}

func (s *ListDataInsightDirectoriesResponseBody) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListDataInsightDirectoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataInsightDirectoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataInsightDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataInsightDirectoriesResponseBody) SetDirectory(v *ListDataInsightDirectoriesResponseBodyDirectory) *ListDataInsightDirectoriesResponseBody {
	s.Directory = v
	return s
}

func (s *ListDataInsightDirectoriesResponseBody) SetFileSystemId(v string) *ListDataInsightDirectoriesResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBody) SetMaxResults(v int32) *ListDataInsightDirectoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBody) SetNextToken(v string) *ListDataInsightDirectoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBody) SetRequestId(v string) *ListDataInsightDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBody) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataInsightDirectoriesResponseBodyDirectory struct {
	// example:
	//
	// 10240
	DirCapacity *int64 `json:"DirCapacity,omitempty" xml:"DirCapacity,omitempty"`
	// example:
	//
	// 10240
	DirCapacityOffline *int64 `json:"DirCapacityOffline,omitempty" xml:"DirCapacityOffline,omitempty"`
	// example:
	//
	// 10240
	DirCapacityOnline *int64 `json:"DirCapacityOnline,omitempty" xml:"DirCapacityOnline,omitempty"`
	// example:
	//
	// 2343232
	FileCount *int64 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// example:
	//
	// 204800
	FileCountOffline *int64 `json:"FileCountOffline,omitempty" xml:"FileCountOffline,omitempty"`
	// example:
	//
	// 204800
	FileCountOnline *int64                                                           `json:"FileCountOnline,omitempty" xml:"FileCountOnline,omitempty"`
	SubDirectories  []*ListDataInsightDirectoriesResponseBodyDirectorySubDirectories `json:"SubDirectories,omitempty" xml:"SubDirectories,omitempty" type:"Repeated"`
}

func (s ListDataInsightDirectoriesResponseBodyDirectory) String() string {
	return dara.Prettify(s)
}

func (s ListDataInsightDirectoriesResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) GetDirCapacity() *int64 {
	return s.DirCapacity
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) GetDirCapacityOffline() *int64 {
	return s.DirCapacityOffline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) GetDirCapacityOnline() *int64 {
	return s.DirCapacityOnline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) GetFileCount() *int64 {
	return s.FileCount
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) GetFileCountOffline() *int64 {
	return s.FileCountOffline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) GetFileCountOnline() *int64 {
	return s.FileCountOnline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) GetSubDirectories() []*ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	return s.SubDirectories
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) SetDirCapacity(v int64) *ListDataInsightDirectoriesResponseBodyDirectory {
	s.DirCapacity = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) SetDirCapacityOffline(v int64) *ListDataInsightDirectoriesResponseBodyDirectory {
	s.DirCapacityOffline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) SetDirCapacityOnline(v int64) *ListDataInsightDirectoriesResponseBodyDirectory {
	s.DirCapacityOnline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) SetFileCount(v int64) *ListDataInsightDirectoriesResponseBodyDirectory {
	s.FileCount = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) SetFileCountOffline(v int64) *ListDataInsightDirectoriesResponseBodyDirectory {
	s.FileCountOffline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) SetFileCountOnline(v int64) *ListDataInsightDirectoriesResponseBodyDirectory {
	s.FileCountOnline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) SetSubDirectories(v []*ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) *ListDataInsightDirectoriesResponseBodyDirectory {
	s.SubDirectories = v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectory) Validate() error {
	if s.SubDirectories != nil {
		for _, item := range s.SubDirectories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataInsightDirectoriesResponseBodyDirectorySubDirectories struct {
	// example:
	//
	// 2026-07-23T12:47:14Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 345518080
	DirCapacity *int64 `json:"DirCapacity,omitempty" xml:"DirCapacity,omitempty"`
	// example:
	//
	// 0
	DirCapacityOffline *int64 `json:"DirCapacityOffline,omitempty" xml:"DirCapacityOffline,omitempty"`
	// example:
	//
	// 345518080
	DirCapacityOnline *int64 `json:"DirCapacityOnline,omitempty" xml:"DirCapacityOnline,omitempty"`
	// example:
	//
	// 1
	DirLevel *int32 `json:"DirLevel,omitempty" xml:"DirLevel,omitempty"`
	// example:
	//
	// /dir_l1_n000
	DirName *string `json:"DirName,omitempty" xml:"DirName,omitempty"`
	// example:
	//
	// 84846
	FileCount *int64 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// example:
	//
	// 2343
	FileCountOffline *int64 `json:"FileCountOffline,omitempty" xml:"FileCountOffline,omitempty"`
	// example:
	//
	// 84355
	FileCountOnline *int64 `json:"FileCountOnline,omitempty" xml:"FileCountOnline,omitempty"`
	// example:
	//
	// 2026-07-29T03:41:12Z
	LastAccessTime *string `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	// example:
	//
	// 2026-07-29T03:41:12Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) String() string {
	return dara.Prettify(s)
}

func (s ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GoString() string {
	return s.String()
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetDirCapacity() *int64 {
	return s.DirCapacity
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetDirCapacityOffline() *int64 {
	return s.DirCapacityOffline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetDirCapacityOnline() *int64 {
	return s.DirCapacityOnline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetDirLevel() *int32 {
	return s.DirLevel
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetDirName() *string {
	return s.DirName
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetFileCount() *int64 {
	return s.FileCount
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetFileCountOffline() *int64 {
	return s.FileCountOffline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetFileCountOnline() *int64 {
	return s.FileCountOnline
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetLastAccessTime() *string {
	return s.LastAccessTime
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetCreatedAt(v string) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.CreatedAt = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetDirCapacity(v int64) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.DirCapacity = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetDirCapacityOffline(v int64) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.DirCapacityOffline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetDirCapacityOnline(v int64) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.DirCapacityOnline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetDirLevel(v int32) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.DirLevel = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetDirName(v string) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.DirName = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetFileCount(v int64) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.FileCount = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetFileCountOffline(v int64) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.FileCountOffline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetFileCountOnline(v int64) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.FileCountOnline = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetLastAccessTime(v string) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.LastAccessTime = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) SetUpdatedAt(v string) *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories {
	s.UpdatedAt = &v
	return s
}

func (s *ListDataInsightDirectoriesResponseBodyDirectorySubDirectories) Validate() error {
	return dara.Validate(s)
}
