// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryOrFilePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *GetDirectoryOrFilePropertiesRequest
	GetFileSystemId() *string
	SetPath(v string) *GetDirectoryOrFilePropertiesRequest
	GetPath() *string
}

type GetDirectoryOrFilePropertiesRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of the directory.
	//
	// The path must start with a forward slash (/) and must be a path that exists in the mount target.
	//
	// This parameter is required.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetDirectoryOrFilePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryOrFilePropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetDirectoryOrFilePropertiesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetDirectoryOrFilePropertiesRequest) GetPath() *string {
	return s.Path
}

func (s *GetDirectoryOrFilePropertiesRequest) SetFileSystemId(v string) *GetDirectoryOrFilePropertiesRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesRequest) SetPath(v string) *GetDirectoryOrFilePropertiesRequest {
	s.Path = &v
	return s
}

func (s *GetDirectoryOrFilePropertiesRequest) Validate() error {
	return dara.Validate(s)
}
