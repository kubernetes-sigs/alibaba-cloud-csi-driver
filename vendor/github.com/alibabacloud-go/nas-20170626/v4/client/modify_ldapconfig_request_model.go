// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLDAPConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindDN(v string) *ModifyLDAPConfigRequest
	GetBindDN() *string
	SetFileSystemId(v string) *ModifyLDAPConfigRequest
	GetFileSystemId() *string
	SetSearchBase(v string) *ModifyLDAPConfigRequest
	GetSearchBase() *string
	SetURI(v string) *ModifyLDAPConfigRequest
	GetURI() *string
}

type ModifyLDAPConfigRequest struct {
	// The LDAP entry.
	//
	// example:
	//
	// cn=alibaba,dc=com
	BindDN *string `json:"BindDN,omitempty" xml:"BindDN,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 109c042666
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The LDAP search base.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc=example
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	// The LDAP service address.
	//
	// This parameter is required.
	//
	// example:
	//
	// ldap://ldap.example.example
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s ModifyLDAPConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigRequest) GetBindDN() *string {
	return s.BindDN
}

func (s *ModifyLDAPConfigRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyLDAPConfigRequest) GetSearchBase() *string {
	return s.SearchBase
}

func (s *ModifyLDAPConfigRequest) GetURI() *string {
	return s.URI
}

func (s *ModifyLDAPConfigRequest) SetBindDN(v string) *ModifyLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetFileSystemId(v string) *ModifyLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetSearchBase(v string) *ModifyLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

func (s *ModifyLDAPConfigRequest) SetURI(v string) *ModifyLDAPConfigRequest {
	s.URI = &v
	return s
}

func (s *ModifyLDAPConfigRequest) Validate() error {
	return dara.Validate(s)
}
