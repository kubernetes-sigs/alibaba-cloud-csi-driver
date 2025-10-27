// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateFileRequest
	GetFileSystemId() *string
	SetOwner(v string) *CreateFileRequest
	GetOwner() *string
	SetOwnerAccessInheritable(v bool) *CreateFileRequest
	GetOwnerAccessInheritable() *bool
	SetPath(v string) *CreateFileRequest
	GetPath() *string
	SetType(v string) *CreateFileRequest
	GetType() *string
}

type CreateFileRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the portable account. The ID must be a 16-digit string. The string can contain digits and lowercase letters.
	//
	// example:
	//
	// 378cc7630f26****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Specifies whether to share the directory. Valid values:
	//
	// 	- false (default): does not share the directory.
	//
	// 	- true: shares the directory.
	//
	// > 	- This parameter takes effect only if the Type parameter is set to Directory and the Owner parameter is not empty.
	//
	// > 	- The permissions on a directory can be inherited by the owner. The owner has read and write permissions on the subdirectories and subfiles created in the directory, even if they are created by others.
	//
	// example:
	//
	// false
	OwnerAccessInheritable *bool `json:"OwnerAccessInheritable,omitempty" xml:"OwnerAccessInheritable,omitempty"`
	// The absolute path of the directory or file. The path must start and end with a forward slash (/) and must be 2 to 1024 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- File
	//
	// 	- Directory
	//
	// This parameter is required.
	//
	// example:
	//
	// File
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateFileRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateFileRequest) GetOwnerAccessInheritable() *bool {
	return s.OwnerAccessInheritable
}

func (s *CreateFileRequest) GetPath() *string {
	return s.Path
}

func (s *CreateFileRequest) GetType() *string {
	return s.Type
}

func (s *CreateFileRequest) SetFileSystemId(v string) *CreateFileRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateFileRequest) SetOwner(v string) *CreateFileRequest {
	s.Owner = &v
	return s
}

func (s *CreateFileRequest) SetOwnerAccessInheritable(v bool) *CreateFileRequest {
	s.OwnerAccessInheritable = &v
	return s
}

func (s *CreateFileRequest) SetPath(v string) *CreateFileRequest {
	s.Path = &v
	return s
}

func (s *CreateFileRequest) SetType(v string) *CreateFileRequest {
	s.Type = &v
	return s
}

func (s *CreateFileRequest) Validate() error {
	return dara.Validate(s)
}
