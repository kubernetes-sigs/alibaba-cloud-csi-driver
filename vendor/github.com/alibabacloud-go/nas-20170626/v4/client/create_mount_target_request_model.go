// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *CreateMountTargetRequest
	GetAccessGroupName() *string
	SetDryRun(v bool) *CreateMountTargetRequest
	GetDryRun() *bool
	SetEnableIpv6(v bool) *CreateMountTargetRequest
	GetEnableIpv6() *bool
	SetFileSystemId(v string) *CreateMountTargetRequest
	GetFileSystemId() *string
	SetNetworkType(v string) *CreateMountTargetRequest
	GetNetworkType() *string
	SetSecurityGroupId(v string) *CreateMountTargetRequest
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *CreateMountTargetRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateMountTargetRequest
	GetVpcId() *string
}

type CreateMountTargetRequest struct {
	// The name of the permission group.
	//
	// This parameter is required if the file system is a General-purpose NAS or Extreme NAS file system.
	//
	// Default permission group: DEFAULT_VPC_GROUP_NAME (the default permission group for VPCs).
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// Specifies whether to check for existing mount targets. Only CPFS file systems are supported.
	//
	// A dry run checks parameter validity and inventory without actually creating a mount target or incurring fees.
	//
	// - true: sends a dry run request without creating a mount target. The check items include required parameters, request format, business limits, and CPFS inventory. If the check fails, the corresponding error is returned. If the check passes, `200 HttpCode` is returned, but `MountTargetDomain` is empty.
	//
	// - false (default): sends a normal request. After the check passes, a mount target is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to create an IPv6 mount target.
	//
	// Valid values:
	//
	// - true: creates an IPv6 mount target.
	//
	// - false (default): does not create an IPv6 mount target.
	//
	// > IPv6 is supported only by Extreme NAS file systems in all regions in the Chinese mainland. The file system must have IPv6 enabled.
	//
	// example:
	//
	// true
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The file system ID.
	//
	// - General-purpose NAS: 31a8e4\\*\\*\\*\\*.
	//
	// - Extreme NAS: The ID must start with `extreme-`, such as extreme-0015\\*\\*\\*\\*.
	//
	// - Cloud Parallel File Storage (CPFS): The ID must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174494****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The network type of the mount target. Set the value to **Vpc**, which indicates a virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// Vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required and valid only when the network type is VPC.
	//
	// For example:
	//
	// If NetworkType is set to VPC, VSwitchId is required.
	//
	// example:
	//
	// vsw-2zevmwkwyztjuoffg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// This parameter is required and valid only when the network type is VPC.
	//
	// For example:
	//
	// If NetworkType is set to VPC, VpcId is required.
	//
	// example:
	//
	// vpc-2zesj9afh3y518k9o****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateMountTargetRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *CreateMountTargetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateMountTargetRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *CreateMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateMountTargetRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateMountTargetRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateMountTargetRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateMountTargetRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateMountTargetRequest) SetAccessGroupName(v string) *CreateMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateMountTargetRequest) SetDryRun(v bool) *CreateMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateMountTargetRequest) SetEnableIpv6(v bool) *CreateMountTargetRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *CreateMountTargetRequest) SetFileSystemId(v string) *CreateMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateMountTargetRequest) SetNetworkType(v string) *CreateMountTargetRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateMountTargetRequest) SetSecurityGroupId(v string) *CreateMountTargetRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMountTargetRequest) SetVSwitchId(v string) *CreateMountTargetRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateMountTargetRequest) SetVpcId(v string) *CreateMountTargetRequest {
	s.VpcId = &v
	return s
}

func (s *CreateMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
