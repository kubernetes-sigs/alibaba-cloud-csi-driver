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
	// This parameter is required if you create a mount target for a General-purpose NAS file system or an Extreme NAS file system.
	//
	// The default permission group for virtual private clouds (VPCs) is named DEFAULT_VPC_GROUP_NAME.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// Specifies whether to perform a dry run to check for existing mount targets. This parameter is valid only for CPFS file systems.
	//
	// If you set this parameter to true, the system checks whether the request parameters are valid and whether the requested resources are available. In this case, no mount target is created and no fee is incurred.
	//
	// 	- true: performs a dry run but does not create a mount target. In the dry run, the system checks the request format, service limits, available CPFS resources, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code `200` is returned. No value is returned for the `MountTargetDomain` parameter.
	//
	// 	- false (default): sends the request. If the request passes the dry run, a mount target is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to create an IPv6 domain name for the mount target.
	//
	// Valid values:
	//
	// 	- true: An IPv6 domain name is created for the mount target.
	//
	// 	- false (default): No IPv6 domain name is created for the mount target.
	//
	// > Only Extreme NAS file systems that reside in the Chinese mainland support IPv6. If you want to create an IPv6 domain name for the mount target, you must enable IPv6 for the file system.
	//
	// example:
	//
	// true
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// The ID of the file system.
	//
	// 	- Sample ID of a General-purpose NAS file system: 31a8e4\\*\\*\\*\\*.
	//
	// 	- The IDs of Extreme NAS file systems must start with `extreme-`, for example, extreme-0015\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-125487\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174494****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The network type of the mount target. Valid value: **Vpc**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is valid and required if the mount target resides in a VPC. Example: If you set the NetworkType parameter to VPC, you must specify the VSwitchId parameter.
	//
	// example:
	//
	// vsw-2zevmwkwyztjuoffg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is valid and required if the mount target resides in a VPC. Example: If you set the NetworkType parameter to VPC, you must specify the VpcId parameter.
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
