// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecycledDirectoriesAndFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntries(v []*ListRecycledDirectoriesAndFilesResponseBodyEntries) *ListRecycledDirectoriesAndFilesResponseBody
	GetEntries() []*ListRecycledDirectoriesAndFilesResponseBodyEntries
	SetNextToken(v string) *ListRecycledDirectoriesAndFilesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRecycledDirectoriesAndFilesResponseBody
	GetRequestId() *string
}

type ListRecycledDirectoriesAndFilesResponseBody struct {
	// The information about files or directories in the recycle bin.
	Entries []*ListRecycledDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// A pagination token.
	//
	// If all the files and directories are incompletely returned in a query, the return value of the NextToken parameter is not empty. In this case, you can specify a valid value for the NextToken parameter to continue the query.
	//
	// example:
	//
	// CKuO8QMSIjE2OTc3NzI0NjI5MTcyMTYyNDVfMzEzNTUyMF81MjEzODY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) GetEntries() []*ListRecycledDirectoriesAndFilesResponseBodyEntries {
	return s.Entries
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetEntries(v []*ListRecycledDirectoriesAndFilesResponseBodyEntries) *ListRecycledDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListRecycledDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecycledDirectoriesAndFilesResponseBodyEntries struct {
	// The time when the file or directory was last accessed.
	//
	// example:
	//
	// 2019-10-30T10:08:08Z
	ATime *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	// The time when the metadata was last modified.
	//
	// example:
	//
	// 2019-10-30T10:08:08Z
	CTime *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	// The time when the file or directory was deleted.
	//
	// example:
	//
	// 2021-05-30T10:08:08Z
	DeleteTime *string `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// The IDs of the files or directories.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The inode of the file or directory.
	//
	// example:
	//
	// 04***08
	Inode *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	// The time when the file or directory was last modified.
	//
	// example:
	//
	// 2019-10-30T10:08:08Z
	MTime *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	// The name of the file or directory before it was deleted.
	//
	// example:
	//
	// test001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// The value 0 is returned for this parameter if Directory is returned for the Type parameter.
	//
	// example:
	//
	// 1073741824
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the returned object. Valid values:
	//
	// 	- File
	//
	// 	- Directory
	//
	// example:
	//
	// File
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) String() string {
	return dara.Prettify(s)
}

func (s ListRecycledDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetATime() *string {
	return s.ATime
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetCTime() *string {
	return s.CTime
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetDeleteTime() *string {
	return s.DeleteTime
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetFileId() *string {
	return s.FileId
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetInode() *string {
	return s.Inode
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetMTime() *string {
	return s.MTime
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetName() *string {
	return s.Name
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetSize() *int64 {
	return s.Size
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) GetType() *string {
	return s.Type
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetATime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.ATime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetCTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.CTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetDeleteTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.DeleteTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetMTime(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.MTime = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetSize(v int64) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Size = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListRecycledDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

func (s *ListRecycledDirectoriesAndFilesResponseBodyEntries) Validate() error {
	return dara.Validate(s)
}
