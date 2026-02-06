// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoriesAndFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntries(v []*ListDirectoriesAndFilesResponseBodyEntries) *ListDirectoriesAndFilesResponseBody
	GetEntries() []*ListDirectoriesAndFilesResponseBodyEntries
	SetNextToken(v string) *ListDirectoriesAndFilesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDirectoriesAndFilesResponseBody
	GetRequestId() *string
}

type ListDirectoriesAndFilesResponseBody struct {
	// The details about the files or directories.
	Entries []*ListDirectoriesAndFilesResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDirectoriesAndFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBody) GetEntries() []*ListDirectoriesAndFilesResponseBodyEntries {
	return s.Entries
}

func (s *ListDirectoriesAndFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDirectoriesAndFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDirectoriesAndFilesResponseBody) SetEntries(v []*ListDirectoriesAndFilesResponseBodyEntries) *ListDirectoriesAndFilesResponseBody {
	s.Entries = v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetNextToken(v string) *ListDirectoriesAndFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) SetRequestId(v string) *ListDirectoriesAndFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBody) Validate() error {
	if s.Entries != nil {
		for _, item := range s.Entries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDirectoriesAndFilesResponseBodyEntries struct {
	// The time when the file was queried.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	//
	// example:
	//
	// 2021-02-01T10:08:08Z
	Atime *string `json:"Atime,omitempty" xml:"Atime,omitempty"`
	// The time when the raw data was modified.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	//
	// example:
	//
	// 2021-02-11T10:08:10Z
	Ctime *string `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	// The ID of the directory or file.
	//
	// example:
	//
	// 66
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// Indicates whether the directory contains files stored in the Archive storage class.
	//
	// This parameter is returned and valid only if the value of the Type parameter is Directory.
	//
	// Valid values:
	//
	// 	- true: The directory contains files stored in the Archive storage class.
	//
	// 	- false: The directory does not contain files stored in the Archive storage class.
	//
	// example:
	//
	// true
	HasArchiveFile *string `json:"HasArchiveFile,omitempty" xml:"HasArchiveFile,omitempty"`
	// Indicates whether the directory contains files stored in the IA storage class.
	//
	// This parameter is returned and valid only if the value of the Type parameter is Directory.
	//
	// Valid values:
	//
	// 	- true: The directory contains files stored in the IA storage class.
	//
	// 	- false: The directory does not contain files stored in the IA storage class.
	//
	// example:
	//
	// true
	HasInfrequentAccessFile *bool `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	// The file or directory inode.
	//
	// example:
	//
	// 66
	Inode *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	// The time when the file was modified.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	//
	// example:
	//
	// 2021-02-11T10:08:08Z
	Mtime *string `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	// The name of the file or directory.
	//
	// example:
	//
	// file.txt
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OfflineDuration          *int64  `json:"OfflineDuration,omitempty" xml:"OfflineDuration,omitempty"`
	OfflineUnchangedDuration *int64  `json:"OfflineUnchangedDuration,omitempty" xml:"OfflineUnchangedDuration,omitempty"`
	// The ID of the portable account. This parameter is returned and valid only if the value of the ProtocolType parameter is SMB and RAM-based access control is enabled.
	//
	// example:
	//
	// 37862c****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The time when the last data retrieval task was run.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	//
	// example:
	//
	// 2021-02-11T10:08:08Z
	RetrieveTime *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	// The size of the file.
	//
	// Unit: bytes.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	//
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The storage class.
	//
	// This parameter is returned and valid only if the value of the Type parameter is File.
	//
	// Valid values:
	//
	// 	- InfrequentAccess: the IA storage class.
	//
	// 	- Archive: the Archive storage class.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The type of the query result.
	//
	// Valid values:
	//
	// 	- File
	//
	// 	- Directory
	//
	// example:
	//
	// Directory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDirectoriesAndFilesResponseBodyEntries) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoriesAndFilesResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetAtime() *string {
	return s.Atime
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetCtime() *string {
	return s.Ctime
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetFileId() *string {
	return s.FileId
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetHasArchiveFile() *string {
	return s.HasArchiveFile
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetHasInfrequentAccessFile() *bool {
	return s.HasInfrequentAccessFile
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetInode() *string {
	return s.Inode
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetMtime() *string {
	return s.Mtime
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetName() *string {
	return s.Name
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetOfflineDuration() *int64 {
	return s.OfflineDuration
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetOfflineUnchangedDuration() *int64 {
	return s.OfflineUnchangedDuration
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetOwner() *string {
	return s.Owner
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetRetrieveTime() *string {
	return s.RetrieveTime
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetSize() *int64 {
	return s.Size
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetStorageType() *string {
	return s.StorageType
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) GetType() *string {
	return s.Type
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetAtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Atime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetCtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Ctime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetFileId(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.FileId = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetHasArchiveFile(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.HasArchiveFile = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetHasInfrequentAccessFile(v bool) *ListDirectoriesAndFilesResponseBodyEntries {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetInode(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Inode = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetMtime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Mtime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetName(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Name = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetOfflineDuration(v int64) *ListDirectoriesAndFilesResponseBodyEntries {
	s.OfflineDuration = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetOfflineUnchangedDuration(v int64) *ListDirectoriesAndFilesResponseBodyEntries {
	s.OfflineUnchangedDuration = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetOwner(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Owner = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetRetrieveTime(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.RetrieveTime = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetSize(v int64) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Size = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetStorageType(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.StorageType = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) SetType(v string) *ListDirectoriesAndFilesResponseBodyEntries {
	s.Type = &v
	return s
}

func (s *ListDirectoriesAndFilesResponseBodyEntries) Validate() error {
	return dara.Validate(s)
}
