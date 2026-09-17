// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDirQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileCountLimit(v int64) *SetDirQuotaRequest
	GetFileCountLimit() *int64
	SetFileSystemId(v string) *SetDirQuotaRequest
	GetFileSystemId() *string
	SetPath(v string) *SetDirQuotaRequest
	GetPath() *string
	SetQuotaType(v string) *SetDirQuotaRequest
	GetQuotaType() *string
	SetSizeLimit(v int64) *SetDirQuotaRequest
	GetSizeLimit() *int64
	SetUserId(v string) *SetDirQuotaRequest
	GetUserId() *string
	SetUserType(v string) *SetDirQuotaRequest
	GetUserType() *string
}

type SetDirQuotaRequest struct {
	// The maximum number of files in the directory.
	//
	// This includes files, directories, and special files.
	//
	//
	// When QuotaType is set to Enforcement, you must specify at least one of SizeLimit and FileCountLimit.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of the directory in the file system.
	//
	//  > - You can set a quota only for a directory that has been created in the NAS file system. The directory path for the quota is the absolute path in the NAS file system, not the local path on a compute node (for example, an ECS instance or container).
	//
	//  > - Directories whose path names contain Chinese characters are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data/sub1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The quota type.
	//
	// Valid values:
	//
	// - Accounting: statistical quota. Only tracks usage.
	//
	// - Enforcement: restrictive quota. When usage exceeds the limit, operations such as creating files or directories and appending data fail.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accounting
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The total capacity limit for files in the directory.
	//
	// Unit: GiB.
	//
	//
	// When QuotaType is set to Enforcement, you must specify at least one of SizeLimit and FileCountLimit.
	//
	// example:
	//
	// 1024
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	// The UID or GID to restrict.
	//
	// This parameter is required and valid only when UserType is set to Uid or Gid.
	//
	// Examples:
	//
	// - To restrict the user whose UID is 500, set UserType to Uid and UserId to 500.
	//
	// - To restrict the user group whose GID is 100, set UserType to Gid and UserId to 100.
	//
	// example:
	//
	// 500
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user type.
	//
	// Valid values:
	//
	// - Uid: user ID
	//
	// - Gid: user group ID
	//
	// - AllUsers: all users
	//
	// This parameter is required.
	//
	// example:
	//
	// Uid
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s SetDirQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDirQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetDirQuotaRequest) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *SetDirQuotaRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *SetDirQuotaRequest) GetPath() *string {
	return s.Path
}

func (s *SetDirQuotaRequest) GetQuotaType() *string {
	return s.QuotaType
}

func (s *SetDirQuotaRequest) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *SetDirQuotaRequest) GetUserId() *string {
	return s.UserId
}

func (s *SetDirQuotaRequest) GetUserType() *string {
	return s.UserType
}

func (s *SetDirQuotaRequest) SetFileCountLimit(v int64) *SetDirQuotaRequest {
	s.FileCountLimit = &v
	return s
}

func (s *SetDirQuotaRequest) SetFileSystemId(v string) *SetDirQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *SetDirQuotaRequest) SetPath(v string) *SetDirQuotaRequest {
	s.Path = &v
	return s
}

func (s *SetDirQuotaRequest) SetQuotaType(v string) *SetDirQuotaRequest {
	s.QuotaType = &v
	return s
}

func (s *SetDirQuotaRequest) SetSizeLimit(v int64) *SetDirQuotaRequest {
	s.SizeLimit = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserId(v string) *SetDirQuotaRequest {
	s.UserId = &v
	return s
}

func (s *SetDirQuotaRequest) SetUserType(v string) *SetDirQuotaRequest {
	s.UserType = &v
	return s
}

func (s *SetDirQuotaRequest) Validate() error {
	return dara.Validate(s)
}
