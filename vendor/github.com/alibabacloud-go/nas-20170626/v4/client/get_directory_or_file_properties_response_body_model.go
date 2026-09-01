// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryOrFilePropertiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntry(v *GetDirectoryOrFilePropertiesResponseBodyEntry) *GetDirectoryOrFilePropertiesResponseBody
	GetEntry() *GetDirectoryOrFilePropertiesResponseBodyEntry
	SetRequestId(v string) *GetDirectoryOrFilePropertiesResponseBody
	GetRequestId() *string
}

type GetDirectoryOrFilePropertiesResponseBody struct {
	// The properties of the directory or file.
	Entry *GetDirectoryOrFilePropertiesResponseBodyEntry `json:"Entry,omitempty" xml:"Entry,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponseBody) GetEntry() *GetDirectoryOrFilePropertiesResponseBodyEntry {
	return s.Entry
}

func (s *GetDirectoryOrFilePropertiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDirectoryOrFilePropertiesResponseBody) SetEntry(v *GetDirectoryOrFilePropertiesResponseBodyEntry) *GetDirectoryOrFilePropertiesResponseBody {
	s.Entry = v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBody) SetRequestId(v string) *GetDirectoryOrFilePropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBody) Validate() error {
	if s.Entry != nil {
		if err := s.Entry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectoryOrFilePropertiesResponseBodyEntry struct {
	// The time when the file was last accessed.
	//
	// The time is displayed in the ISO 8601 standard format and is returned in UTC. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is returned only if Type is set to File.
	//
	// example:
	//
	// 2021-02-01T10:08:08Z
	ATime *string `json:"ATime,omitempty" xml:"ATime,omitempty"`
	// The time when the metadata of the file was last modified.
	//
	// The time is displayed in the ISO 8601 standard format and is returned in UTC. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is returned only if Type is set to File.
	//
	// example:
	//
	// 2021-02-11T10:08:10Z
	CTime *string `json:"CTime,omitempty" xml:"CTime,omitempty"`
	// Indicates whether the directory contains an archive file.
	//
	// This parameter is returned only if Type is set to Directory.
	//
	// Valid values:
	//
	// - true: The directory contains at least one archive file.
	//
	// - false: The directory does not contain archive files.
	//
	// example:
	//
	// false
	HasArchiveFile *bool `json:"HasArchiveFile,omitempty" xml:"HasArchiveFile,omitempty"`
	// Indicates whether the directory contains an Infrequent Access (IA) file.
	//
	// This parameter is returned only if Type is set to Directory.
	//
	// Valid values:
	//
	// - true: The directory contains at least one IA file.
	//
	// - false: The directory does not contain IA files.
	//
	// example:
	//
	// true
	HasInfrequentAccessFile *bool `json:"HasInfrequentAccessFile,omitempty" xml:"HasInfrequentAccessFile,omitempty"`
	// The inode of the file or directory.
	//
	// example:
	//
	// 40
	Inode *string `json:"Inode,omitempty" xml:"Inode,omitempty"`
	// The time when the file was last modified.
	//
	// The time is displayed in the ISO 8601 standard format and is returned in UTC. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is returned only if Type is set to File.
	//
	// example:
	//
	// 2021-02-11T10:08:08Z
	MTime *string `json:"MTime,omitempty" xml:"MTime,omitempty"`
	// The name of the file or directory.
	//
	// example:
	//
	// file.txt
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OfflineDuration          *int64  `json:"OfflineDuration,omitempty" xml:"OfflineDuration,omitempty"`
	OfflineUnchangedDuration *int64  `json:"OfflineUnchangedDuration,omitempty" xml:"OfflineUnchangedDuration,omitempty"`
	// The time when the last data retrieval task was run.
	//
	// The time is displayed in the ISO 8601 standard format and is returned in UTC. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// This parameter is returned only if Type is set to File.
	//
	// example:
	//
	// 2021-02-11T10:08:08Z
	RetrieveTime *string `json:"RetrieveTime,omitempty" xml:"RetrieveTime,omitempty"`
	// The size of the file.
	//
	// Unit: bytes.
	//
	// This parameter is returned only if Type is set to File.
	//
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The storage type of the file.
	//
	// This parameter is returned only if Type is set to File.
	//
	// Valid values:
	//
	// - Standard: General-purpose NAS file systems
	//
	// - InfrequentAccess: IA storage medium
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The type of the returned entry.
	//
	// Valid values:
	//
	// - File: The entry is a file.
	//
	// - Directory: The entry is a directory.
	//
	// example:
	//
	// File
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesResponseBodyEntry) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetATime() *string {
	return s.ATime
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetCTime() *string {
	return s.CTime
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetHasArchiveFile() *bool {
	return s.HasArchiveFile
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetHasInfrequentAccessFile() *bool {
	return s.HasInfrequentAccessFile
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetInode() *string {
	return s.Inode
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetMTime() *string {
	return s.MTime
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetName() *string {
	return s.Name
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetOfflineDuration() *int64 {
	return s.OfflineDuration
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetOfflineUnchangedDuration() *int64 {
	return s.OfflineUnchangedDuration
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetRetrieveTime() *string {
	return s.RetrieveTime
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetSize() *int64 {
	return s.Size
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetStorageType() *string {
	return s.StorageType
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) GetType() *string {
	return s.Type
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetATime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.ATime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetCTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.CTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetHasArchiveFile(v bool) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.HasArchiveFile = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetHasInfrequentAccessFile(v bool) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.HasInfrequentAccessFile = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetInode(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Inode = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetMTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.MTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetName(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Name = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetOfflineDuration(v int64) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.OfflineDuration = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetOfflineUnchangedDuration(v int64) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.OfflineUnchangedDuration = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetRetrieveTime(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.RetrieveTime = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetSize(v int64) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Size = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetStorageType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.StorageType = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) SetType(v string) *GetDirectoryOrFilePropertiesResponseBodyEntry {
	s.Type = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesResponseBodyEntry) Validate() error {
	return dara.Validate(s)
}
