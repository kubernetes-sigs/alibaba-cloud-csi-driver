// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmbAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcl(v *DescribeSmbAclResponseBodyAcl) *DescribeSmbAclResponseBody
	GetAcl() *DescribeSmbAclResponseBodyAcl
	SetRequestId(v string) *DescribeSmbAclResponseBody
	GetRequestId() *string
}

type DescribeSmbAclResponseBody struct {
	// The information about the ACL feature.
	Acl *DescribeSmbAclResponseBodyAcl `json:"Acl,omitempty" xml:"Acl,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSmbAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmbAclResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponseBody) GetAcl() *DescribeSmbAclResponseBodyAcl {
	return s.Acl
}

func (s *DescribeSmbAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSmbAclResponseBody) SetAcl(v *DescribeSmbAclResponseBodyAcl) *DescribeSmbAclResponseBody {
	s.Acl = v
	return s
}

func (s *DescribeSmbAclResponseBody) SetRequestId(v string) *DescribeSmbAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmbAclResponseBody) Validate() error {
	if s.Acl != nil {
		if err := s.Acl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSmbAclResponseBodyAcl struct {
	// Indicates whether the file system allows anonymous access. Valid values:
	//
	// 	- true: The file system allows anonymous access.
	//
	// 	- false: The file system does not allow anonymous access.
	//
	// example:
	//
	// true
	EnableAnonymousAccess *bool `json:"EnableAnonymousAccess,omitempty" xml:"EnableAnonymousAccess,omitempty"`
	// Indicates whether the ACL feature is enabled. Valid values:
	//
	// 	- true: The ACL feature is enabled.
	//
	// 	- false: The ACL feature is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether encryption in transit is enabled. Valid values:
	//
	// 	- true: Encryption in transit is enabled.
	//
	// 	- false: Encryption in transit is disabled.
	//
	// example:
	//
	// true
	EncryptData *bool `json:"EncryptData,omitempty" xml:"EncryptData,omitempty"`
	// The home directory of each user.
	//
	// example:
	//
	// /home
	HomeDirPath *string `json:"HomeDirPath,omitempty" xml:"HomeDirPath,omitempty"`
	// Indicates whether the file system denies access from non-encrypted clients. Valid values:
	//
	// 	- true: The file system denies access from non-encrypted clients.
	//
	// 	- false: The file system allows access from non-encrypted clients.
	//
	// example:
	//
	// true
	RejectUnencryptedAccess *bool `json:"RejectUnencryptedAccess,omitempty" xml:"RejectUnencryptedAccess,omitempty"`
	// The ID of a super admin.
	//
	// example:
	//
	// S-1-0-0
	SuperAdminSid *string `json:"SuperAdminSid,omitempty" xml:"SuperAdminSid,omitempty"`
}

func (s DescribeSmbAclResponseBodyAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmbAclResponseBodyAcl) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponseBodyAcl) GetEnableAnonymousAccess() *bool {
	return s.EnableAnonymousAccess
}

func (s *DescribeSmbAclResponseBodyAcl) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeSmbAclResponseBodyAcl) GetEncryptData() *bool {
	return s.EncryptData
}

func (s *DescribeSmbAclResponseBodyAcl) GetHomeDirPath() *string {
	return s.HomeDirPath
}

func (s *DescribeSmbAclResponseBodyAcl) GetRejectUnencryptedAccess() *bool {
	return s.RejectUnencryptedAccess
}

func (s *DescribeSmbAclResponseBodyAcl) GetSuperAdminSid() *string {
	return s.SuperAdminSid
}

func (s *DescribeSmbAclResponseBodyAcl) SetEnableAnonymousAccess(v bool) *DescribeSmbAclResponseBodyAcl {
	s.EnableAnonymousAccess = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetEnabled(v bool) *DescribeSmbAclResponseBodyAcl {
	s.Enabled = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetEncryptData(v bool) *DescribeSmbAclResponseBodyAcl {
	s.EncryptData = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetHomeDirPath(v string) *DescribeSmbAclResponseBodyAcl {
	s.HomeDirPath = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetRejectUnencryptedAccess(v bool) *DescribeSmbAclResponseBodyAcl {
	s.RejectUnencryptedAccess = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) SetSuperAdminSid(v string) *DescribeSmbAclResponseBodyAcl {
	s.SuperAdminSid = &v
	return s
}

func (s *DescribeSmbAclResponseBodyAcl) Validate() error {
	return dara.Validate(s)
}
