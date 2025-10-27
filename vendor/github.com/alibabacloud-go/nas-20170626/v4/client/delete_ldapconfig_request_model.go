// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLDAPConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteLDAPConfigRequest
	GetFileSystemId() *string
}

type DeleteLDAPConfigRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404a348
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteLDAPConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLDAPConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteLDAPConfigRequest) SetFileSystemId(v string) *DeleteLDAPConfigRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteLDAPConfigRequest) Validate() error {
	return dara.Validate(s)
}
