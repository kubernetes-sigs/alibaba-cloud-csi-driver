// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntries(v *DescribeFilesetsResponseBodyEntries) *DescribeFilesetsResponseBody
	GetEntries() *DescribeFilesetsResponseBodyEntries
	SetFileSystemId(v string) *DescribeFilesetsResponseBody
	GetFileSystemId() *string
	SetNextToken(v string) *DescribeFilesetsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeFilesetsResponseBody
	GetRequestId() *string
}

type DescribeFilesetsResponseBody struct {
	// The fileset information.
	Entries *DescribeFilesetsResponseBodyEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Struct"`
	// Deprecated
	//
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// >  CPFS is not supported on the international site.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
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

func (s DescribeFilesetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBody) GetEntries() *DescribeFilesetsResponseBodyEntries {
	return s.Entries
}

func (s *DescribeFilesetsResponseBody) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFilesetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFilesetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFilesetsResponseBody) SetEntries(v *DescribeFilesetsResponseBodyEntries) *DescribeFilesetsResponseBody {
	s.Entries = v
	return s
}

func (s *DescribeFilesetsResponseBody) SetFileSystemId(v string) *DescribeFilesetsResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesetsResponseBody) SetNextToken(v string) *DescribeFilesetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeFilesetsResponseBody) SetRequestId(v string) *DescribeFilesetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFilesetsResponseBody) Validate() error {
	if s.Entries != nil {
		if err := s.Entries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFilesetsResponseBodyEntries struct {
	Entrie []*DescribeFilesetsResponseBodyEntriesEntrie `json:"Entrie,omitempty" xml:"Entrie,omitempty" type:"Repeated"`
}

func (s DescribeFilesetsResponseBodyEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesetsResponseBodyEntries) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBodyEntries) GetEntrie() []*DescribeFilesetsResponseBodyEntriesEntrie {
	return s.Entrie
}

func (s *DescribeFilesetsResponseBodyEntries) SetEntrie(v []*DescribeFilesetsResponseBodyEntriesEntrie) *DescribeFilesetsResponseBodyEntries {
	s.Entrie = v
	return s
}

func (s *DescribeFilesetsResponseBodyEntries) Validate() error {
	if s.Entrie != nil {
		for _, item := range s.Entrie {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFilesetsResponseBodyEntriesEntrie struct {
	// The time when the fileset was created.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2021-09-30T10:08:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Specifies whether to enable deletion protection to allow you to release the fileset by using the console or by calling the [DeleteFileset](https://help.aliyun.com/document_detail/2838077.html) operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter can protect filesets only against manual releases, but not against automatic releases.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The fileset description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The usage of the file quantity.
	//
	// >  Only CPFS for LINGJUN V2.7.0 and later support this parameter.
	//
	// example:
	//
	// 1024
	FileCountUsage *int64 `json:"FileCountUsage,omitempty" xml:"FileCountUsage,omitempty"`
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// >  CPFS is not supported on the international site.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The fileset path.
	//
	// example:
	//
	// pathtoroot/fset
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The fileset ID.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The quota information.
	//
	// >  Only CPFS for Lingjun V2.7.0 and later support this parameter.
	Quota *DescribeFilesetsResponseBodyEntriesEntrieQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// The capacity usage. Unit: bytes.
	//
	// >  Only CPFS for LINGJUN V2.7.0 and later support this parameter.
	//
	// example:
	//
	// 1024
	SpaceUsage *int64 `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	// The fileset status. Valid values:
	//
	// 	- CREATING: The fileset is being created.
	//
	// 	- CREATED: The fileset has been created and is running properly.
	//
	// 	- RELEASING: The fileset is being released.
	//
	// 	- RELEASED: The fileset has been deleted.
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the fileset was last updated.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2021-09-30T10:08:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeFilesetsResponseBodyEntriesEntrie) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesetsResponseBodyEntriesEntrie) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetDescription() *string {
	return s.Description
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetFileCountUsage() *int64 {
	return s.FileCountUsage
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetFsetId() *string {
	return s.FsetId
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetQuota() *DescribeFilesetsResponseBodyEntriesEntrieQuota {
	return s.Quota
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetSpaceUsage() *int64 {
	return s.SpaceUsage
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetStatus() *string {
	return s.Status
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetCreateTime(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.CreateTime = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetDeletionProtection(v bool) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetDescription(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.Description = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFileCountUsage(v int64) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FileCountUsage = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFileSystemId(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFileSystemPath(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetFsetId(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.FsetId = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetQuota(v *DescribeFilesetsResponseBodyEntriesEntrieQuota) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.Quota = v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetSpaceUsage(v int64) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.SpaceUsage = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetStatus(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.Status = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) SetUpdateTime(v string) *DescribeFilesetsResponseBodyEntriesEntrie {
	s.UpdateTime = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrie) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFilesetsResponseBodyEntriesEntrieQuota struct {
	// The file quantity quota. Valid values:
	//
	// 	- Minimum value: 10000.
	//
	// 	- Maximum value: 10000000000.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The capacity quota. Unit: bytes.
	//
	// 	- Minimum value: 10737418240 (10 GiB).
	//
	// 	- Step size: 1073741824 (1 GiB).
	//
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s DescribeFilesetsResponseBodyEntriesEntrieQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesetsResponseBodyEntriesEntrieQuota) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponseBodyEntriesEntrieQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *DescribeFilesetsResponseBodyEntriesEntrieQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *DescribeFilesetsResponseBodyEntriesEntrieQuota) SetFileCountLimit(v int64) *DescribeFilesetsResponseBodyEntriesEntrieQuota {
	s.FileCountLimit = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrieQuota) SetSizeLimit(v int64) *DescribeFilesetsResponseBodyEntriesEntrieQuota {
	s.SizeLimit = &v
	return s
}

func (s *DescribeFilesetsResponseBodyEntriesEntrieQuota) Validate() error {
	return dara.Validate(s)
}
