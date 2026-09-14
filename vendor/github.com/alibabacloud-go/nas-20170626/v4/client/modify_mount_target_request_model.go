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
	SetAccessPointAccessOnly(v bool) *ModifyMountTargetRequest
	GetAccessPointAccessOnly() *bool
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
	// The permission group attached to the mount target.
	//
	// example:
	//
	// classic-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// Specifies whether the VPC mount target supports access only through access points. This parameter applies only to CPFS for Lingjun file systems.
	//
	// example:
	//
	// false
	AccessPointAccessOnly *bool `json:"AccessPointAccessOnly,omitempty" xml:"AccessPointAccessOnly,omitempty"`
	// The IPv4/IPv6 dual-stack mount target.
	//
	// > Currently, only Extreme NAS in regions in the Chinese mainland supports IPv6.
	//
	// example:
	//
	// 174494b666-x****.dualstack.cn-hangzhou.nas.aliyuncs.com
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The file system ID.
	//
	// - General-purpose NAS: `31a8e4****`.
	//
	// - Extreme NAS: Must start with `extreme-`, such as `extreme-0015****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The IPv4 mount target.
	//
	// example:
	//
	// 1ca404a666-w****.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The mount target status.
	//
	// Valid values:
	//
	// - Active: active
	//
	// - Inactive: inactive
	//
	// > Only General-purpose NAS supports changing the mount target status.
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

func (s *ModifyMountTargetRequest) GetAccessPointAccessOnly() *bool {
	return s.AccessPointAccessOnly
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

func (s *ModifyMountTargetRequest) SetAccessPointAccessOnly(v bool) *ModifyMountTargetRequest {
	s.AccessPointAccessOnly = &v
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
