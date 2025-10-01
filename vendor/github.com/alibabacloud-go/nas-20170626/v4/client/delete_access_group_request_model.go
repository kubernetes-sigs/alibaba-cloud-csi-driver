// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *DeleteAccessGroupRequest
	GetAccessGroupName() *string
	SetFileSystemType(v string) *DeleteAccessGroupRequest
	GetFileSystemType() *string
}

type DeleteAccessGroupRequest struct {
	// The name of the permission group to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard (default): General-purpose NAS file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DeleteAccessGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DeleteAccessGroupRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DeleteAccessGroupRequest) SetAccessGroupName(v string) *DeleteAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DeleteAccessGroupRequest) SetFileSystemType(v string) *DeleteAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

func (s *DeleteAccessGroupRequest) Validate() error {
	return dara.Validate(s)
}
