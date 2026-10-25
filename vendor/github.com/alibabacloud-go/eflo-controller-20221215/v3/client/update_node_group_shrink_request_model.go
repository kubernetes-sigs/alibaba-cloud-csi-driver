// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemMountEnabled(v bool) *UpdateNodeGroupShrinkRequest
	GetFileSystemMountEnabled() *bool
	SetImageId(v string) *UpdateNodeGroupShrinkRequest
	GetImageId() *string
	SetKeyPairName(v string) *UpdateNodeGroupShrinkRequest
	GetKeyPairName() *string
	SetLoginPassword(v string) *UpdateNodeGroupShrinkRequest
	GetLoginPassword() *string
	SetNewNodeGroupName(v string) *UpdateNodeGroupShrinkRequest
	GetNewNodeGroupName() *string
	SetNodeGroupId(v string) *UpdateNodeGroupShrinkRequest
	GetNodeGroupId() *string
	SetRamRoleName(v string) *UpdateNodeGroupShrinkRequest
	GetRamRoleName() *string
	SetSystemDiskShrink(v string) *UpdateNodeGroupShrinkRequest
	GetSystemDiskShrink() *string
	SetUserData(v string) *UpdateNodeGroupShrinkRequest
	GetUserData() *string
}

type UpdateNodeGroupShrinkRequest struct {
	// Specifies whether file storage mounting is supported.
	//
	// example:
	//
	// True
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The default image ID of the node group. If you do not specify this parameter, the image ID remains unchanged.
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
	// The logon password for the machines in the node group.
	//
	// example:
	//
	// Password
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The name of the node group.
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
	// The name of the RAM role for the node. You can call the ListRoles API operation of RAM to query the RAM roles that you have created. The trusted entity of the role must be Lingjun AI Computing Service. <warning>Submitting an empty string "" clears the RAM role of the node group.</warning>
	//
	// example:
	//
	// xianwen-test-ram-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The system cloud disk configuration.
	SystemDiskShrink *string `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The custom executable shell script. It must be Base64-encoded, and the maximum size of the raw data is 16 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateNodeGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupShrinkRequest) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *UpdateNodeGroupShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateNodeGroupShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *UpdateNodeGroupShrinkRequest) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *UpdateNodeGroupShrinkRequest) GetNewNodeGroupName() *string {
	return s.NewNodeGroupName
}

func (s *UpdateNodeGroupShrinkRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateNodeGroupShrinkRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *UpdateNodeGroupShrinkRequest) GetSystemDiskShrink() *string {
	return s.SystemDiskShrink
}

func (s *UpdateNodeGroupShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateNodeGroupShrinkRequest) SetFileSystemMountEnabled(v bool) *UpdateNodeGroupShrinkRequest {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetImageId(v string) *UpdateNodeGroupShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetKeyPairName(v string) *UpdateNodeGroupShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetLoginPassword(v string) *UpdateNodeGroupShrinkRequest {
	s.LoginPassword = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetNewNodeGroupName(v string) *UpdateNodeGroupShrinkRequest {
	s.NewNodeGroupName = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetNodeGroupId(v string) *UpdateNodeGroupShrinkRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetRamRoleName(v string) *UpdateNodeGroupShrinkRequest {
	s.RamRoleName = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetSystemDiskShrink(v string) *UpdateNodeGroupShrinkRequest {
	s.SystemDiskShrink = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) SetUserData(v string) *UpdateNodeGroupShrinkRequest {
	s.UserData = &v
	return s
}

func (s *UpdateNodeGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
