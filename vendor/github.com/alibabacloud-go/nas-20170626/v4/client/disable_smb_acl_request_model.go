// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmbAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DisableSmbAclRequest
	GetFileSystemId() *string
}

type DisableSmbAclRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableSmbAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSmbAclRequest) GoString() string {
	return s.String()
}

func (s *DisableSmbAclRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DisableSmbAclRequest) SetFileSystemId(v string) *DisableSmbAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *DisableSmbAclRequest) Validate() error {
	return dara.Validate(s)
}
