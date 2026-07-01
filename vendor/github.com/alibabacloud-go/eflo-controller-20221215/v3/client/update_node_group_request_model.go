// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemMountEnabled(v bool) *UpdateNodeGroupRequest
	GetFileSystemMountEnabled() *bool
	SetImageId(v string) *UpdateNodeGroupRequest
	GetImageId() *string
	SetKeyPairName(v string) *UpdateNodeGroupRequest
	GetKeyPairName() *string
	SetLoginPassword(v string) *UpdateNodeGroupRequest
	GetLoginPassword() *string
	SetNewNodeGroupName(v string) *UpdateNodeGroupRequest
	GetNewNodeGroupName() *string
	SetNodeGroupId(v string) *UpdateNodeGroupRequest
	GetNodeGroupId() *string
	SetRamRoleName(v string) *UpdateNodeGroupRequest
	GetRamRoleName() *string
	SetUserData(v string) *UpdateNodeGroupRequest
	GetUserData() *string
}

type UpdateNodeGroupRequest struct {
	// Specifies whether to mount file storage on nodes.
	//
	// example:
	//
	// True
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The ID of the image for the node group. If you omit this parameter, the image remains unchanged.
	//
	// example:
	//
	// i1232142432432
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// test
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The login password for the nodes in the node group.
	//
	// example:
	//
	// Password
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The new name of the node group.
	//
	// example:
	//
	// test-update
	NewNodeGroupName *string `json:"NewNodeGroupName,omitempty" xml:"NewNodeGroupName,omitempty"`
	// The ID of the node group.
	//
	// example:
	//
	// i120021051733814190732
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// To query the RAM roles that you have created, call the ListRoles operation in the RAM API. The trust policy for the role must specify Intelligent Computing Lingjun as the trusted entity. Note: You cannot remove a role by clearing this parameter.
	//
	// example:
	//
	// xianwen-test-ram-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The user data. This script runs at node startup.
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupRequest) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *UpdateNodeGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateNodeGroupRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *UpdateNodeGroupRequest) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *UpdateNodeGroupRequest) GetNewNodeGroupName() *string {
	return s.NewNodeGroupName
}

func (s *UpdateNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateNodeGroupRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *UpdateNodeGroupRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateNodeGroupRequest) SetFileSystemMountEnabled(v bool) *UpdateNodeGroupRequest {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetImageId(v string) *UpdateNodeGroupRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetKeyPairName(v string) *UpdateNodeGroupRequest {
	s.KeyPairName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetLoginPassword(v string) *UpdateNodeGroupRequest {
	s.LoginPassword = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNewNodeGroupName(v string) *UpdateNodeGroupRequest {
	s.NewNodeGroupName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNodeGroupId(v string) *UpdateNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetRamRoleName(v string) *UpdateNodeGroupRequest {
	s.RamRoleName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetUserData(v string) *UpdateNodeGroupRequest {
	s.UserData = &v
	return s
}

func (s *UpdateNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
