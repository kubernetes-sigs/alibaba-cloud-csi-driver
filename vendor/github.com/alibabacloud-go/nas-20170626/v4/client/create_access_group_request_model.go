// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *CreateAccessGroupRequest
	GetAccessGroupName() *string
	SetAccessGroupType(v string) *CreateAccessGroupRequest
	GetAccessGroupType() *string
	SetDescription(v string) *CreateAccessGroupRequest
	GetDescription() *string
	SetFileSystemType(v string) *CreateAccessGroupRequest
	GetFileSystemType() *string
}

type CreateAccessGroupRequest struct {
	// The name of the permission group.
	//
	// Limits:
	//
	// 	- The name must be 3 to 64 characters in length.
	//
	// 	- The name must start with a letter and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be different from the name of the default permission group.
	//
	// The default permission group for virtual private clouds (VPCs) is named DEFAULT_VPC_GROUP_NAME.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The network type of the permission group. Valid value: **Vpc**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Vpc
	AccessGroupType *string `json:"AccessGroupType,omitempty" xml:"AccessGroupType,omitempty"`
	// The description of the permission group.
	//
	// Limits:
	//
	// 	- By default, the description of a permission group is the same as the name of the permission group. The description must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// vpctestaccessgroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard (default): General-purpose NAS file system.
	//
	// 	- extreme: Extreme NAS file system.
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s CreateAccessGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *CreateAccessGroupRequest) GetAccessGroupType() *string {
	return s.AccessGroupType
}

func (s *CreateAccessGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAccessGroupRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *CreateAccessGroupRequest) SetAccessGroupName(v string) *CreateAccessGroupRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateAccessGroupRequest) SetAccessGroupType(v string) *CreateAccessGroupRequest {
	s.AccessGroupType = &v
	return s
}

func (s *CreateAccessGroupRequest) SetDescription(v string) *CreateAccessGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAccessGroupRequest) SetFileSystemType(v string) *CreateAccessGroupRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAccessGroupRequest) Validate() error {
	return dara.Validate(s)
}
