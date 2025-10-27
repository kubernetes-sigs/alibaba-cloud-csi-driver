// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNfsAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DisableNfsAclRequest
	GetFileSystemId() *string
}

type DisableNfsAclRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 91fcdxxxx
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableNfsAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableNfsAclRequest) GoString() string {
	return s.String()
}

func (s *DisableNfsAclRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DisableNfsAclRequest) SetFileSystemId(v string) *DisableNfsAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *DisableNfsAclRequest) Validate() error {
	return dara.Validate(s)
}
