// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *ModifyMountTargetRequest
	GetAccessGroupName() *string
	SetDualStackMountTargetDomain(v string) *ModifyMountTargetRequest
	GetDualStackMountTargetDomain() *string
	SetFileSystemId(v string) *ModifyMountTargetRequest
	GetFileSystemId() *string
	SetMountTargetDomain(v string) *ModifyMountTargetRequest
	GetMountTargetDomain() *string
	SetStatus(v string) *ModifyMountTargetRequest
	GetStatus() *string
}

type ModifyMountTargetRequest struct {
	// The name of the permission group that is attached to the mount target.
	//
	// example:
	//
	// classic-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	//
	// >  Only Extreme NAS file systems that reside in the Chinese mainland support IPv6.
	//
	// example:
	//
	// 174494b666-x****.dualstack.cn-hangzhou.nas.aliyuncs.com
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The ID of the file system.
	//
	// 	- Sample ID of a General-purpose NAS file system: `31a8e4****`.
	//
	// 	- The IDs of Extreme NAS file systems must start with `extreme-`, for example, `extreme-0015****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The IPv4 domain name of the mount target.
	//
	// example:
	//
	// 1ca404a666-w****.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The status of the mount target.
	//
	// Valid values:
	//
	// 	- Active: The mount target is available.
	//
	// 	- Inactive: The mount target is unavailable.
	//
	// >  Only General-purpose File Storage NAS (NAS) file systems support changing the mount target status.
	//
	// example:
	//
	// Inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyMountTargetRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *ModifyMountTargetRequest) GetDualStackMountTargetDomain() *string {
	return s.DualStackMountTargetDomain
}

func (s *ModifyMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyMountTargetRequest) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *ModifyMountTargetRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyMountTargetRequest) SetAccessGroupName(v string) *ModifyMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *ModifyMountTargetRequest) SetDualStackMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *ModifyMountTargetRequest) SetFileSystemId(v string) *ModifyMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyMountTargetRequest) SetMountTargetDomain(v string) *ModifyMountTargetRequest {
	s.MountTargetDomain = &v
	return s
}

func (s *ModifyMountTargetRequest) SetStatus(v string) *ModifyMountTargetRequest {
	s.Status = &v
	return s
}

func (s *ModifyMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
