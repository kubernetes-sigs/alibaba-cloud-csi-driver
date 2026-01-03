// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSystemDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *ReplaceSystemDiskRequestSystemDisk) *ReplaceSystemDiskRequest
	GetSystemDisk() *ReplaceSystemDiskRequestSystemDisk
	SetArchitecture(v string) *ReplaceSystemDiskRequest
	GetArchitecture() *string
	SetArn(v []*ReplaceSystemDiskRequestArn) *ReplaceSystemDiskRequest
	GetArn() []*ReplaceSystemDiskRequestArn
	SetClientToken(v string) *ReplaceSystemDiskRequest
	GetClientToken() *string
	SetDiskId(v string) *ReplaceSystemDiskRequest
	GetDiskId() *string
	SetEncryptAlgorithm(v string) *ReplaceSystemDiskRequest
	GetEncryptAlgorithm() *string
	SetEncrypted(v bool) *ReplaceSystemDiskRequest
	GetEncrypted() *bool
	SetImageId(v string) *ReplaceSystemDiskRequest
	GetImageId() *string
	SetInstanceId(v string) *ReplaceSystemDiskRequest
	GetInstanceId() *string
	SetKMSKeyId(v string) *ReplaceSystemDiskRequest
	GetKMSKeyId() *string
	SetKeyPairName(v string) *ReplaceSystemDiskRequest
	GetKeyPairName() *string
	SetOwnerAccount(v string) *ReplaceSystemDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReplaceSystemDiskRequest
	GetOwnerId() *int64
	SetPassword(v string) *ReplaceSystemDiskRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *ReplaceSystemDiskRequest
	GetPasswordInherit() *bool
	SetPlatform(v string) *ReplaceSystemDiskRequest
	GetPlatform() *string
	SetResourceOwnerAccount(v string) *ReplaceSystemDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReplaceSystemDiskRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *ReplaceSystemDiskRequest
	GetSecurityEnhancementStrategy() *string
	SetUseAdditionalService(v bool) *ReplaceSystemDiskRequest
	GetUseAdditionalService() *bool
}

type ReplaceSystemDiskRequest struct {
	SystemDisk *ReplaceSystemDiskRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// i386
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// This parameter is not available for public use.
	Arn []*ReplaceSystemDiskRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but make sure that the token is unique across requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// >  This parameter is deprecated. To improve compatibility, we recommend that you use `ImageId`.
	//
	// example:
	//
	// d-bp67acfmxazb4ph****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// > This parameter is not available for public use.
	//
	// example:
	//
	// hide
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  When you use a shared encrypted image to create the disk based on an encrypted snapshot, you must set Encrypted to true to ensure that the disk uses an encryption key of your own.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the image used to replace the system disk. This parameter is required.
	//
	// example:
	//
	// m-bp67acfmxazb4ph****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the instance whose operating system you want to replace.
	//
	// >  Make sure that the instance is in the `Stopped` (`Stopped`) state.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the KMS key to use for the system disk.
	//
	// example:
	//
	// e522b26d-abf6-4e0d-b5da-04b7******3c
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The name of the key pair.
	//
	// > This parameter is applicable only to Linux instances. You can bind an SSH key pair to the instance as a logon credential. After you bind the SSH key pair, the username and password logon method is disabled for the instance.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName  *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to reset the password for the instance. The password must be 8 to 30 characters in length and contain at least three of the following items: uppercase letters, lowercase letters, digits, and special characters. Special characters include:
	//
	//     ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// The passwords of Windows instances cannot start with a forward slash (/).
	//
	// This parameter is empty by default, which indicates that the current password remains unchanged.
	//
	// > If you specify `Password`, we recommend that you send requests over HTTPS to prevent password leaks.
	//
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the preset password of the image.
	//
	// Default value: false
	//
	// > If the PasswordInherit parameter is specified, you must leave the Password parameter empty. Before you use this parameter, make sure that a password is preset for the image.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// CentOS
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to use Security Center Basic after the system disk is replaced. Valid values:
	//
	// 	- Active: uses Security Center Basic after the system disk is re-initialized. This value is applicable only to public images.
	//
	// 	- Deactive: does not use Security Center Basic after the system disk is re-initialized. This value is applicable to all images.
	//
	// Default value: Deactive.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// Specifies whether to use the system configurations for virtual machines provided by Alibaba Cloud. System configurations for Windows: NTP and KMS. System configurations for Linux: NTP and YUM.
	//
	// > This parameter takes effect only when you attach a system disk whose device name is /dev/xvda.
	//
	// example:
	//
	// true
	UseAdditionalService *bool `json:"UseAdditionalService,omitempty" xml:"UseAdditionalService,omitempty"`
}

func (s ReplaceSystemDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSystemDiskRequest) GoString() string {
	return s.String()
}

func (s *ReplaceSystemDiskRequest) GetSystemDisk() *ReplaceSystemDiskRequestSystemDisk {
	return s.SystemDisk
}

func (s *ReplaceSystemDiskRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *ReplaceSystemDiskRequest) GetArn() []*ReplaceSystemDiskRequestArn {
	return s.Arn
}

func (s *ReplaceSystemDiskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceSystemDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ReplaceSystemDiskRequest) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *ReplaceSystemDiskRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *ReplaceSystemDiskRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ReplaceSystemDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReplaceSystemDiskRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *ReplaceSystemDiskRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ReplaceSystemDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReplaceSystemDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReplaceSystemDiskRequest) GetPassword() *string {
	return s.Password
}

func (s *ReplaceSystemDiskRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *ReplaceSystemDiskRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ReplaceSystemDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReplaceSystemDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReplaceSystemDiskRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *ReplaceSystemDiskRequest) GetUseAdditionalService() *bool {
	return s.UseAdditionalService
}

func (s *ReplaceSystemDiskRequest) SetSystemDisk(v *ReplaceSystemDiskRequestSystemDisk) *ReplaceSystemDiskRequest {
	s.SystemDisk = v
	return s
}

func (s *ReplaceSystemDiskRequest) SetArchitecture(v string) *ReplaceSystemDiskRequest {
	s.Architecture = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetArn(v []*ReplaceSystemDiskRequestArn) *ReplaceSystemDiskRequest {
	s.Arn = v
	return s
}

func (s *ReplaceSystemDiskRequest) SetClientToken(v string) *ReplaceSystemDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetDiskId(v string) *ReplaceSystemDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetEncryptAlgorithm(v string) *ReplaceSystemDiskRequest {
	s.EncryptAlgorithm = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetEncrypted(v bool) *ReplaceSystemDiskRequest {
	s.Encrypted = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetImageId(v string) *ReplaceSystemDiskRequest {
	s.ImageId = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetInstanceId(v string) *ReplaceSystemDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetKMSKeyId(v string) *ReplaceSystemDiskRequest {
	s.KMSKeyId = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetKeyPairName(v string) *ReplaceSystemDiskRequest {
	s.KeyPairName = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetOwnerAccount(v string) *ReplaceSystemDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetOwnerId(v int64) *ReplaceSystemDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetPassword(v string) *ReplaceSystemDiskRequest {
	s.Password = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetPasswordInherit(v bool) *ReplaceSystemDiskRequest {
	s.PasswordInherit = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetPlatform(v string) *ReplaceSystemDiskRequest {
	s.Platform = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetResourceOwnerAccount(v string) *ReplaceSystemDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetResourceOwnerId(v int64) *ReplaceSystemDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetSecurityEnhancementStrategy(v string) *ReplaceSystemDiskRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *ReplaceSystemDiskRequest) SetUseAdditionalService(v bool) *ReplaceSystemDiskRequest {
	s.UseAdditionalService = &v
	return s
}

func (s *ReplaceSystemDiskRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceSystemDiskRequestSystemDisk struct {
	// The capacity of the new system disk. Unit: GiB. Valid values:
	//
	// 	- Basic disk: Max{20, Size of the image specified by ImageId} to 500.
	//
	// 	- Enterprise SSD (ESSD):
	//
	//     	- PL0 ESSD: Max{1, Size of the image specified by ImageId} to 2048.
	//
	//     	- PL1 ESSD: Max{20, Size of the image specified by ImageId} to 2048.
	//
	//     	- PL2 ESSD: Max{461, Size of the image specified by ImageId} to 2048.
	//
	//     	- PL3 ESSD: Max{1261, Size of the image specified by ImageId} to 2048.
	//
	// 	- ESSD AutoPL disk: Max{1, Size of the image specified by ImageId} to 2048.
	//
	// 	- Other disk categories: Max{20, Size of the image specified by ImageId} to 2048.
	//
	// Default value: Max{40, Size of the image specified by ImageId}.
	//
	// >  If the capacity of the new system disk exceeds `Max{20, Capacity of the original system disk}`, you are charged for the excess capacity.
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ReplaceSystemDiskRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSystemDiskRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ReplaceSystemDiskRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *ReplaceSystemDiskRequestSystemDisk) SetSize(v int32) *ReplaceSystemDiskRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *ReplaceSystemDiskRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type ReplaceSystemDiskRequestArn struct {
	// > This parameter is unavailable.
	//
	// example:
	//
	// 0
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is not available for public use.
	//
	// example:
	//
	// null
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is not available for public use.
	//
	// example:
	//
	// null
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s ReplaceSystemDiskRequestArn) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSystemDiskRequestArn) GoString() string {
	return s.String()
}

func (s *ReplaceSystemDiskRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *ReplaceSystemDiskRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *ReplaceSystemDiskRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *ReplaceSystemDiskRequestArn) SetAssumeRoleFor(v int64) *ReplaceSystemDiskRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *ReplaceSystemDiskRequestArn) SetRoleType(v string) *ReplaceSystemDiskRequestArn {
	s.RoleType = &v
	return s
}

func (s *ReplaceSystemDiskRequestArn) SetRolearn(v string) *ReplaceSystemDiskRequestArn {
	s.Rolearn = &v
	return s
}

func (s *ReplaceSystemDiskRequestArn) Validate() error {
	return dara.Validate(s)
}
