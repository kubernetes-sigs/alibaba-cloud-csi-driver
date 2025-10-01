// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *ModifyAccessGroupRequest
	GetAccessGroupName() *string
	SetDescription(v string) *ModifyAccessGroupRequest
	GetDescription() *string
	SetFileSystemType(v string) *ModifyAccessGroupRequest
	GetFileSystemType() *string
}

type ModifyAccessGroupRequest struct {
	// The name of the permission group.
	//
	// Limits:
	//
	// 	- The name must be 3 to 64 characters in length.
	//
	// 	- The name must start with a letter and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The description of the permission group.
	//
	// Limits:
	//
	// 	- By default, the description of the permission group is the same as the name of the permission group. The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// vpc-test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s ModifyAccessGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *ModifyAccessGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAccessGroupRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *ModifyAccessGroupRequest) SetAccessGroupName(v string) *ModifyAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetDescription(v string) *ModifyAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccessGroupRequest) SetFileSystemType(v string) *ModifyAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

func (s *ModifyAccessGroupRequest) Validate() error {
	return dara.Validate(s)
}
