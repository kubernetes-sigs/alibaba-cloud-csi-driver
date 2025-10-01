// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmbAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAnonymousAccess(v bool) *ModifySmbAclRequest
	GetEnableAnonymousAccess() *bool
	SetEncryptData(v bool) *ModifySmbAclRequest
	GetEncryptData() *bool
	SetFileSystemId(v string) *ModifySmbAclRequest
	GetFileSystemId() *string
	SetHomeDirPath(v string) *ModifySmbAclRequest
	GetHomeDirPath() *string
	SetKeytab(v string) *ModifySmbAclRequest
	GetKeytab() *string
	SetKeytabMd5(v string) *ModifySmbAclRequest
	GetKeytabMd5() *string
	SetRejectUnencryptedAccess(v bool) *ModifySmbAclRequest
	GetRejectUnencryptedAccess() *bool
	SetSuperAdminSid(v string) *ModifySmbAclRequest
	GetSuperAdminSid() *string
}

type ModifySmbAclRequest struct {
	// Specifies whether to allow anonymous access. Valid values:
	//
	// 	- true: The file system allows anonymous access.
	//
	// 	- false (default): The file system denies anonymous access.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// false
	EnableAnonymousAccess *bool `json:"EnableAnonymousAccess,omitempty" xml:"EnableAnonymousAccess,omitempty"`
	// Specifies whether to enable encryption in transit. Valid values:
	//
	// 	- true: enables encryption in transit.
	//
	// 	- false (default): disables encryption in transit.
	//
	// example:
	//
	// false
	EncryptData *bool `json:"EncryptData,omitempty" xml:"EncryptData,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The home directory of each user. Each user-specific home directory must meet the following requirements:
	//
	// 	- Each segment starts with a forward slash (/) or a backward slash (\\\\).
	//
	// 	- Each segment does not contain the following special characters: `<>":|?*`.
	//
	// 	- Each segment is 0 to 255 characters in length.
	//
	// 	- The total length is 0 to 32,767 characters.
	//
	// For example, if you create a user named A and the home directory is `/home`, the file system automatically creates a directory named `/home/A` when User A logs on to the file system. If the `/home/A` directory already exists, the file system does not create the directory.
	//
	// > User A must have the permissions to create folders in the \\home directory. Otherwise, the file system cannot create the `/home/A` directory when User A logs on to the file system.
	//
	// example:
	//
	// /home
	HomeDirPath *string `json:"HomeDirPath,omitempty" xml:"HomeDirPath,omitempty"`
	// The string that is generated after the system encodes the keytab file by using Base64.
	//
	// example:
	//
	// BQIAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAQAIqIx6v7p11oUAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAwAIqIx6v7p11oUAAABPAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAFwAQnQZWB3RAPHU7PMIJyBWePAAAAF8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQASACAGJ7F0s+bcBjf6jD5HlvlRLmPSOW+qDZe0Qk0lQcf8WwAAAE8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQARABDdFmanrSIatnDDh****
	Keytab *string `json:"Keytab,omitempty" xml:"Keytab,omitempty"`
	// The string that is generated after the system encodes the keytab file by using MD5.
	//
	// example:
	//
	// E3CCF7E2416DF04FA958AA4513EA****
	KeytabMd5 *string `json:"KeytabMd5,omitempty" xml:"KeytabMd5,omitempty"`
	// Specifies whether to deny access from non-encrypted clients. Valid values:
	//
	// 	- true: The file system denies access from non-encrypted clients.
	//
	// 	- false (default): The file system allows access from non-encrypted clients.
	//
	// example:
	//
	// false
	RejectUnencryptedAccess *bool `json:"RejectUnencryptedAccess,omitempty" xml:"RejectUnencryptedAccess,omitempty"`
	// The ID of a super admin. The ID must meet the following requirements:
	//
	// 	- The ID starts with `S` and does not contain letters except S.
	//
	// 	- The ID contains at least three hyphens (-) as delimiters.
	//
	// Examples: `S-1-5-22` and `S-1-5-22-23`.
	//
	// example:
	//
	// S-1-5-22
	SuperAdminSid *string `json:"SuperAdminSid,omitempty" xml:"SuperAdminSid,omitempty"`
}

func (s ModifySmbAclRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySmbAclRequest) GoString() string {
	return s.String()
}

func (s *ModifySmbAclRequest) GetEnableAnonymousAccess() *bool {
	return s.EnableAnonymousAccess
}

func (s *ModifySmbAclRequest) GetEncryptData() *bool {
	return s.EncryptData
}

func (s *ModifySmbAclRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifySmbAclRequest) GetHomeDirPath() *string {
	return s.HomeDirPath
}

func (s *ModifySmbAclRequest) GetKeytab() *string {
	return s.Keytab
}

func (s *ModifySmbAclRequest) GetKeytabMd5() *string {
	return s.KeytabMd5
}

func (s *ModifySmbAclRequest) GetRejectUnencryptedAccess() *bool {
	return s.RejectUnencryptedAccess
}

func (s *ModifySmbAclRequest) GetSuperAdminSid() *string {
	return s.SuperAdminSid
}

func (s *ModifySmbAclRequest) SetEnableAnonymousAccess(v bool) *ModifySmbAclRequest {
	s.EnableAnonymousAccess = &v
	return s
}

func (s *ModifySmbAclRequest) SetEncryptData(v bool) *ModifySmbAclRequest {
	s.EncryptData = &v
	return s
}

func (s *ModifySmbAclRequest) SetFileSystemId(v string) *ModifySmbAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifySmbAclRequest) SetHomeDirPath(v string) *ModifySmbAclRequest {
	s.HomeDirPath = &v
	return s
}

func (s *ModifySmbAclRequest) SetKeytab(v string) *ModifySmbAclRequest {
	s.Keytab = &v
	return s
}

func (s *ModifySmbAclRequest) SetKeytabMd5(v string) *ModifySmbAclRequest {
	s.KeytabMd5 = &v
	return s
}

func (s *ModifySmbAclRequest) SetRejectUnencryptedAccess(v bool) *ModifySmbAclRequest {
	s.RejectUnencryptedAccess = &v
	return s
}

func (s *ModifySmbAclRequest) SetSuperAdminSid(v string) *ModifySmbAclRequest {
	s.SuperAdminSid = &v
	return s
}

func (s *ModifySmbAclRequest) Validate() error {
	return dara.Validate(s)
}
