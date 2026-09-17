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
	// Specifies whether to allow anonymous access.
	//
	// - true: Anonymous access is allowed.
	//
	// - false (default): Anonymous access is not allowed.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// false
	EnableAnonymousAccess *bool `json:"EnableAnonymousAccess,omitempty" xml:"EnableAnonymousAccess,omitempty"`
	// Specifies whether to enable encryption in transit.
	//
	// - true: Encryption in transit is enabled.
	//
	// - false (default): Encryption in transit is not enabled.
	//
	// example:
	//
	// false
	EncryptData *bool `json:"EncryptData,omitempty" xml:"EncryptData,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The home folder path for each user. The file path format is as follows:
	//
	// - Use a forward slash (/) or backslash (\\) as the separator.
	//
	// - Each segment cannot contain `<>":|?*`.
	//
	// - The length of each segment ranges from 0 to 255.
	//
	// - The total length ranges from 0 to 32767.
	//
	// For example, if the user folder is `/home`, the file system performs automatic creation of the `/home/A` folder when user A performs logon. If `/home/A` already exists, this step is skipped.
	//
	// > User A must have the permission to create folders. Otherwise, the `/home/A` folder cannot be created.
	//
	// example:
	//
	// /home
	HomeDirPath *string `json:"HomeDirPath,omitempty" xml:"HomeDirPath,omitempty"`
	// The Base64-encoded string of the keytab file content.
	//
	// example:
	//
	// BQIAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAQAIqIx6v7p11oUAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAwAIqIx6v7p11oUAAABPAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAFwAQnQZWB3RAPHU7PMIJyBWePAAAAF8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQASACAGJ7F0s+bcBjf6jD5HlvlRLmPSOW+qDZe0Qk0lQcf8WwAAAE8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQARABDdFmanrSIatnDDh****
	Keytab *string `json:"Keytab,omitempty" xml:"Keytab,omitempty"`
	// The MD5-encrypted string of the keytab file content.
	//
	// example:
	//
	// E3CCF7E2416DF04FA958AA4513EA****
	KeytabMd5 *string `json:"KeytabMd5,omitempty" xml:"KeytabMd5,omitempty"`
	// Specifies whether to reject unencrypted clients.
	//
	// - true: Unencrypted clients are rejected.
	//
	// - false (default): Unencrypted clients are not rejected.
	//
	// example:
	//
	// false
	RejectUnencryptedAccess *bool `json:"RejectUnencryptedAccess,omitempty" xml:"RejectUnencryptedAccess,omitempty"`
	// The ID of the superuser. The ID must follow these rules:
	//
	// - Must start with `S`, and no other letters are allowed after the initial S.
	//
	// - Must contain at least three hyphens (-) as separators.
	//
	// For example, `S-1-5-22` or `S-1-5-22-23`.
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
