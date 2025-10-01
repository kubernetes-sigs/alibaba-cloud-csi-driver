// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLDAPConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindDN(v string) *CreateLDAPConfigRequest
	GetBindDN() *string
	SetFileSystemId(v string) *CreateLDAPConfigRequest
	GetFileSystemId() *string
	SetSearchBase(v string) *CreateLDAPConfigRequest
	GetSearchBase() *string
	SetURI(v string) *CreateLDAPConfigRequest
	GetURI() *string
}

type CreateLDAPConfigRequest struct {
	// An LDAP entry.
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
	// 109c04****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// An LDAP search base.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc=example
	SearchBase *string `json:"SearchBase,omitempty" xml:"SearchBase,omitempty"`
	// An LDAP URI.
	//
	// This parameter is required.
	//
	// example:
	//
	// ldap://ldap.example.example
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateLDAPConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigRequest) GetBindDN() *string {
	return s.BindDN
}

func (s *CreateLDAPConfigRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateLDAPConfigRequest) GetSearchBase() *string {
	return s.SearchBase
}

func (s *CreateLDAPConfigRequest) GetURI() *string {
	return s.URI
}

func (s *CreateLDAPConfigRequest) SetBindDN(v string) *CreateLDAPConfigRequest {
	s.BindDN = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetFileSystemId(v string) *CreateLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetSearchBase(v string) *CreateLDAPConfigRequest {
	s.SearchBase = &v
	return s
}

func (s *CreateLDAPConfigRequest) SetURI(v string) *CreateLDAPConfigRequest {
	s.URI = &v
	return s
}

func (s *CreateLDAPConfigRequest) Validate() error {
	return dara.Validate(s)
}
