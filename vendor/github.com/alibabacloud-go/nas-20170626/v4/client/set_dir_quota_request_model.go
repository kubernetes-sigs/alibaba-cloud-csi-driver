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
	// The number of files that a user can create in the directory.
	//
	// This number includes the number of files, subdirectories, and special files.
	//
	// If you set the QuotaType parameter to Enforcement, you must specify at least one of the SizeLimit and FileCountLimit parameters.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of the directory in the file system.
	//
	// > 	- You can set quotas only for the directories that have been created in a NAS file system. The path of the directory that you specify for a quota is the absolute path of the directory in the NAS file system, but not the local path of the compute node, such as an Elastic Compute Service (ECS) instance or a container.
	//
	// > 	- Directories whose names contain Chinese characters are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data/sub1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of the quota.
	//
	// Valid values:
	//
	// 	- Accounting: a statistical quota. If you set this parameter to Accounting, NAS calculates only the storage usage of the directory.
	//
	// 	- Enforcement: a restricted quota. If you set this parameter to Enforcement and the storage usage exceeds the quota, you can no longer create files or subdirectories for the directory, or write data to the directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Accounting
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The size of files that a user can create in the directory.
	//
	// Unit: GiB.
	//
	// If you set the QuotaType parameter to Enforcement, you must specify at least one of the SizeLimit and FileCountLimit parameters.
	//
	// example:
	//
	// 1024
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	// The UID or GID of the user for whom you want to set a directory quota.
	//
	// This parameter is required and valid only if the UserType parameter is set to Uid or Gid.
	//
	// Examples:
	//
	// 	- If you want to set a directory quota for a user whose UID is 500, set the UserType parameter to Uid and set the UserId parameter to 500.
	//
	// 	- If you want to set a directory quota for a user group whose GID is 100, set the UserType parameter to Gid and set the UserId parameter to 100.
	//
	// example:
	//
	// 500
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of the user.
	//
	// Valid values:
	//
	// 	- Uid: user ID
	//
	// 	- Gid: user group ID
	//
	// 	- AllUsers: all users
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
