// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBootable(v bool) *AttachDiskRequest
	GetBootable() *bool
	SetDeleteWithInstance(v bool) *AttachDiskRequest
	GetDeleteWithInstance() *bool
	SetDevice(v string) *AttachDiskRequest
	GetDevice() *string
	SetDiskId(v string) *AttachDiskRequest
	GetDiskId() *string
	SetForce(v bool) *AttachDiskRequest
	GetForce() *bool
	SetInstanceId(v string) *AttachDiskRequest
	GetInstanceId() *string
	SetKeyPairName(v string) *AttachDiskRequest
	GetKeyPairName() *string
	SetOwnerAccount(v string) *AttachDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachDiskRequest
	GetOwnerId() *int64
	SetPassword(v string) *AttachDiskRequest
	GetPassword() *string
	SetResourceOwnerAccount(v string) *AttachDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachDiskRequest
	GetResourceOwnerId() *int64
}

type AttachDiskRequest struct {
	// Specifies whether to attach the disk as the system disk. Valid values:
	//
	// 	- true: attaches the disk as the system disk.
	//
	// 	- false: does not attach the disk as the system disk.
	//
	// Default value: false.
	//
	// >  You can set `Bootable` to true only if the instance does not have a system disk.
	//
	// example:
	//
	// false
	Bootable *bool `json:"Bootable,omitempty" xml:"Bootable,omitempty"`
	// Specifies whether to release the disk when the instance is released. Valid values:
	//
	// 	- true: releases the disk when the instance is released.
	//
	// 	- false: does not release the disk when the instance is released. The disk is retained as a pay-as-you-go data disk.
	//
	// Default value: false.
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- If `OperationLocks` in the DescribeInstances response contains `"LockReason" : "security"` for the instance to which the disk is attached, the instance is locked for security reasons. Regardless of whether you set `DeleteWithInstance` to `false`, the DeleteWithInstance setting is ignored, and the disk is released when the instance is released.
	//
	// 	- If you want to attach an `elastic ephemeral disk`, you must set `DeleteWithInstance` to `true`.
	//
	// 	- You cannot specify DeleteWithInstance for disks for which the multi-attach feature is enabled.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The device name of the disk.
	//
	// >  This parameter will be removed in the future. We recommend that you use other parameters to ensure future compatibility.
	//
	// example:
	//
	// testDeviceName
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The ID of the disk. The disk specified by `DiskId` and the instance specified by `InstanceId` must reside in the same zone.
	//
	// >  For information about the limits on attaching a data disk and a system disk, see the "Usage notes" section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp1j4l5axzdy6ftk****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to force attach the disk to the instance. Valid values:
	//
	// 	- true: force attaches the disk to the instance.
	//
	// 	- false: does not force attach the disk to the instance.
	//
	// Default value: false.
	//
	// >  You can set this parameter to true only for Regional Enterprise SSDs (ESSDs) (cloud_regional_disk_auto).
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the instance to which you want to attach the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1dq5lozx5f4pmd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the SSH key pair that you bind to the Linux instance when you attach the system disk.
	//
	// 	- Windows instances do not support logons based on SSH key pairs. The `Password` parameter takes effect even if the KeyPairName parameter is specified.
	//
	// 	- For Linux instances, the username and password-based logon method is disabled by default.
	//
	// example:
	//
	// KeyPairTestName
	KeyPairName  *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password that is set when you attach the system disk. The password is applicable only to the administrator and root users. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	//     ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// For Windows instances, passwords cannot start with a forward slash (/).
	//
	// > If `Password` is configured, we recommend that you send requests over HTTPS to prevent password leaks.
	//
	// example:
	//
	// EcsV587!
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AttachDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDiskRequest) GoString() string {
	return s.String()
}

func (s *AttachDiskRequest) GetBootable() *bool {
	return s.Bootable
}

func (s *AttachDiskRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *AttachDiskRequest) GetDevice() *string {
	return s.Device
}

func (s *AttachDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *AttachDiskRequest) GetForce() *bool {
	return s.Force
}

func (s *AttachDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachDiskRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *AttachDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachDiskRequest) GetPassword() *string {
	return s.Password
}

func (s *AttachDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachDiskRequest) SetBootable(v bool) *AttachDiskRequest {
	s.Bootable = &v
	return s
}

func (s *AttachDiskRequest) SetDeleteWithInstance(v bool) *AttachDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *AttachDiskRequest) SetDevice(v string) *AttachDiskRequest {
	s.Device = &v
	return s
}

func (s *AttachDiskRequest) SetDiskId(v string) *AttachDiskRequest {
	s.DiskId = &v
	return s
}

func (s *AttachDiskRequest) SetForce(v bool) *AttachDiskRequest {
	s.Force = &v
	return s
}

func (s *AttachDiskRequest) SetInstanceId(v string) *AttachDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachDiskRequest) SetKeyPairName(v string) *AttachDiskRequest {
	s.KeyPairName = &v
	return s
}

func (s *AttachDiskRequest) SetOwnerAccount(v string) *AttachDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachDiskRequest) SetOwnerId(v int64) *AttachDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachDiskRequest) SetPassword(v string) *AttachDiskRequest {
	s.Password = &v
	return s
}

func (s *AttachDiskRequest) SetResourceOwnerAccount(v string) *AttachDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachDiskRequest) SetResourceOwnerId(v int64) *AttachDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachDiskRequest) Validate() error {
	return dara.Validate(s)
}
