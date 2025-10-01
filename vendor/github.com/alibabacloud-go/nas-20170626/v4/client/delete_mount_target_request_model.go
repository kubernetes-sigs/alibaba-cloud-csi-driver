// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteMountTargetRequest
	GetFileSystemId() *string
	SetMountTargetDomain(v string) *DeleteMountTargetRequest
	GetMountTargetDomain() *string
}

type DeleteMountTargetRequest struct {
	// The ID of the file system.
	//
	// 	- Sample ID of a General-purpose NAS file system: 31a8e4\\*\\*\\*\\*.
	//
	// 	- The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\\*\\*\\*\\*.
	//
	// 	- The IDs of Cloud Parallel File Storage (CPFS) file systems must start with `cpfs-`, for example, cpfs-125487\\*\\*\\*\\*.
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// 174494****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The domain name of the mount target.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174494b666-x****.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
}

func (s DeleteMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteMountTargetRequest) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DeleteMountTargetRequest) SetFileSystemId(v string) *DeleteMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteMountTargetRequest) SetMountTargetDomain(v string) *DeleteMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *DeleteMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
