// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDirQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CancelDirQuotaRequest
	GetFileSystemId() *string
	SetPath(v string) *CancelDirQuotaRequest
	GetPath() *string
	SetUserId(v string) *CancelDirQuotaRequest
	GetUserId() *string
	SetUserType(v string) *CancelDirQuotaRequest
	GetUserType() *string
}

type CancelDirQuotaRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of a directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /data/sub1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The UID or GID of a user for whom you want to cancel the directory quota.
	//
	// This parameter is required and valid only if the UserType parameter is set to Uid or Gid.
	//
	// Examples:
	//
	// 	- If you want to cancel a quota for a user whose UID is 500, set the UserType parameter to Uid and set the UserId parameter to 500.
	//
	// 	- If you want to cancel a quota for a group whose GID is 100, set the UserType parameter to Gid and set the UserId parameter to 100.
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

func (s CancelDirQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDirQuotaRequest) GoString() string {
	return s.String()
}

func (s *CancelDirQuotaRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CancelDirQuotaRequest) GetPath() *string {
	return s.Path
}

func (s *CancelDirQuotaRequest) GetUserId() *string {
	return s.UserId
}

func (s *CancelDirQuotaRequest) GetUserType() *string {
	return s.UserType
}

func (s *CancelDirQuotaRequest) SetFileSystemId(v string) *CancelDirQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDirQuotaRequest) SetPath(v string) *CancelDirQuotaRequest {
	s.Path = &v
	return s
}

func (s *CancelDirQuotaRequest) SetUserId(v string) *CancelDirQuotaRequest {
	s.UserId = &v
	return s
}

func (s *CancelDirQuotaRequest) SetUserType(v string) *CancelDirQuotaRequest {
	s.UserType = &v
	return s
}

func (s *CancelDirQuotaRequest) Validate() error {
	return dara.Validate(s)
}
