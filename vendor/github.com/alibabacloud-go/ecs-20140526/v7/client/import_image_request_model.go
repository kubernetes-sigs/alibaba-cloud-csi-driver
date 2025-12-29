// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *ImportImageRequest
	GetArchitecture() *string
	SetBootMode(v string) *ImportImageRequest
	GetBootMode() *string
	SetClientToken(v string) *ImportImageRequest
	GetClientToken() *string
	SetDescription(v string) *ImportImageRequest
	GetDescription() *string
	SetDetectionStrategy(v string) *ImportImageRequest
	GetDetectionStrategy() *string
	SetDiskDeviceMapping(v []*ImportImageRequestDiskDeviceMapping) *ImportImageRequest
	GetDiskDeviceMapping() []*ImportImageRequestDiskDeviceMapping
	SetDryRun(v bool) *ImportImageRequest
	GetDryRun() *bool
	SetFeatures(v *ImportImageRequestFeatures) *ImportImageRequest
	GetFeatures() *ImportImageRequestFeatures
	SetImageName(v string) *ImportImageRequest
	GetImageName() *string
	SetLicenseType(v string) *ImportImageRequest
	GetLicenseType() *string
	SetOSType(v string) *ImportImageRequest
	GetOSType() *string
	SetOwnerId(v int64) *ImportImageRequest
	GetOwnerId() *int64
	SetPlatform(v string) *ImportImageRequest
	GetPlatform() *string
	SetRegionId(v string) *ImportImageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ImportImageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ImportImageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportImageRequest
	GetResourceOwnerId() *int64
	SetRoleName(v string) *ImportImageRequest
	GetRoleName() *string
	SetStorageLocationArn(v string) *ImportImageRequest
	GetStorageLocationArn() *string
	SetTag(v []*ImportImageRequestTag) *ImportImageRequest
	GetTag() []*ImportImageRequestTag
}

type ImportImageRequest struct {
	// The system architecture. Valid values:
	//
	// 	- i386
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// Default value: x86_64.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The boot mode of the image. Valid values:
	//
	// 	- BIOS
	//
	// 	- UEFI
	//
	// Default value: BIOS. If you set `Architecture` to arm64, set this parameter to UEFI.
	//
	// > Make sure that you are aware of the boot modes supported by the specified image, as thehe modified boot mode needs to be supported by the image. This way, instances that use this image can start.
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.*	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The image description. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mode in which to check the image. If you do not specify this parameter, the image is not checked. Only the standard check mode is supported.
	//
	// >  This parameter is supported for most Linux and Windows operating system versions. For more information about image check items and operating system limits for image check, see [Overview](https://help.aliyun.com/document_detail/439819.html) and [Operating system limits for image check](https://help.aliyun.com/document_detail/475800.html).
	//
	// example:
	//
	// Standard
	DetectionStrategy *string `json:"DetectionStrategy,omitempty" xml:"DetectionStrategy,omitempty"`
	// Details about the custom images.
	DiskDeviceMapping []*ImportImageRequestDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized RAM users, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The attributes of the image.
	Features *ImportImageRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The image name. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `acs:` or `aliyun`. The name cannot contain `http://` or `https://`. The name can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// ImageTestName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the license used to activate the operating system after the image is imported. Valid values:
	//
	// 	- Auto: ECS checks the operating system of the image and allocates a license to the operating system. ECS first checks whether the operating system distribution specified by `Platform` has a license allocated through an official Alibaba Cloud channel. If yes, the allocated license is used. If no, the license that comes with the source operating system is used.
	//
	// 	- Aliyun: The license allocated through an official Alibaba Cloud channel is used for the operating system distribution specified by `Platform`.
	//
	// 	- BYOL: The license that comes with the source operating system is used. In this case, make sure that your license key is eligible for use in Alibaba Cloud.
	//
	// Default value: Auto.
	//
	// example:
	//
	// Auto
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The operating system platform. Valid values:
	//
	// 	- windows
	//
	// 	- linux
	//
	// Default value: linux.
	//
	// example:
	//
	// linux
	OSType  *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system distribution. Valid values:
	//
	// 	- Aliyun
	//
	// 	- Anolis
	//
	// 	- CentOS
	//
	// 	- Ubuntu
	//
	// 	- CoreOS
	//
	// 	- SUSE
	//
	// 	- Debian
	//
	// 	- OpenSUSE
	//
	// 	- FreeBSD
	//
	// 	- RedHat
	//
	// 	- Kylin
	//
	// 	- UOS
	//
	// 	- Fedora
	//
	// 	- Fedora CoreOS
	//
	// 	- CentOS Stream
	//
	// 	- AlmaLinux
	//
	// 	- Rocky Linux
	//
	// 	- Gentoo
	//
	// 	- Customized Linux
	//
	// 	- Others Linux
	//
	// 	- Windows Server 2022
	//
	// 	- Windows Server 2019
	//
	// 	- Windows Server 2016
	//
	// 	- Windows Server 2012
	//
	// 	- Windows Server 2008
	//
	// 	- Windows Server 2003
	//
	// Default value: Others Linux.
	//
	// example:
	//
	// Aliyun
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID of the source image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the image.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the RAM role used to import the image.
	//
	// example:
	//
	// AliyunECSImageImportDefaultRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the cloud box, which is used to uniquely identify a storage location in the cloud.
	//
	// >  Specify this parameter only if you import an image from OSS on CloudBox. Otherwise, you do not need to specify this parameter. For more information, see [What is OSS on CloudBox?](https://help.aliyun.com/document_detail/430190.html)
	//
	// The ARN must be in the following format: `arn:acs:cloudbox:{RegionId}:{AliUid}:cloudbox/{CloudBoxId}`. Replace `{RegionId}` with the region ID of the cloud box, `{AliUid}` with the ID of the Alibaba Cloud account to which the cloud box belongs, and `{CloudBoxId}` with the ID of the cloud box.
	//
	// example:
	//
	// arn:acs:cloudbox:cn-hangzhou:123456:cloudbox/cb-xx***123
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The image tags.
	Tag []*ImportImageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ImportImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequest) GoString() string {
	return s.String()
}

func (s *ImportImageRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *ImportImageRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *ImportImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ImportImageRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportImageRequest) GetDetectionStrategy() *string {
	return s.DetectionStrategy
}

func (s *ImportImageRequest) GetDiskDeviceMapping() []*ImportImageRequestDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *ImportImageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ImportImageRequest) GetFeatures() *ImportImageRequestFeatures {
	return s.Features
}

func (s *ImportImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ImportImageRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ImportImageRequest) GetOSType() *string {
	return s.OSType
}

func (s *ImportImageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportImageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ImportImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportImageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ImportImageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportImageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportImageRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ImportImageRequest) GetStorageLocationArn() *string {
	return s.StorageLocationArn
}

func (s *ImportImageRequest) GetTag() []*ImportImageRequestTag {
	return s.Tag
}

func (s *ImportImageRequest) SetArchitecture(v string) *ImportImageRequest {
	s.Architecture = &v
	return s
}

func (s *ImportImageRequest) SetBootMode(v string) *ImportImageRequest {
	s.BootMode = &v
	return s
}

func (s *ImportImageRequest) SetClientToken(v string) *ImportImageRequest {
	s.ClientToken = &v
	return s
}

func (s *ImportImageRequest) SetDescription(v string) *ImportImageRequest {
	s.Description = &v
	return s
}

func (s *ImportImageRequest) SetDetectionStrategy(v string) *ImportImageRequest {
	s.DetectionStrategy = &v
	return s
}

func (s *ImportImageRequest) SetDiskDeviceMapping(v []*ImportImageRequestDiskDeviceMapping) *ImportImageRequest {
	s.DiskDeviceMapping = v
	return s
}

func (s *ImportImageRequest) SetDryRun(v bool) *ImportImageRequest {
	s.DryRun = &v
	return s
}

func (s *ImportImageRequest) SetFeatures(v *ImportImageRequestFeatures) *ImportImageRequest {
	s.Features = v
	return s
}

func (s *ImportImageRequest) SetImageName(v string) *ImportImageRequest {
	s.ImageName = &v
	return s
}

func (s *ImportImageRequest) SetLicenseType(v string) *ImportImageRequest {
	s.LicenseType = &v
	return s
}

func (s *ImportImageRequest) SetOSType(v string) *ImportImageRequest {
	s.OSType = &v
	return s
}

func (s *ImportImageRequest) SetOwnerId(v int64) *ImportImageRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportImageRequest) SetPlatform(v string) *ImportImageRequest {
	s.Platform = &v
	return s
}

func (s *ImportImageRequest) SetRegionId(v string) *ImportImageRequest {
	s.RegionId = &v
	return s
}

func (s *ImportImageRequest) SetResourceGroupId(v string) *ImportImageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ImportImageRequest) SetResourceOwnerAccount(v string) *ImportImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportImageRequest) SetResourceOwnerId(v int64) *ImportImageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportImageRequest) SetRoleName(v string) *ImportImageRequest {
	s.RoleName = &v
	return s
}

func (s *ImportImageRequest) SetStorageLocationArn(v string) *ImportImageRequest {
	s.StorageLocationArn = &v
	return s
}

func (s *ImportImageRequest) SetTag(v []*ImportImageRequestTag) *ImportImageRequest {
	s.Tag = v
	return s
}

func (s *ImportImageRequest) Validate() error {
	if s.DiskDeviceMapping != nil {
		for _, item := range s.DiskDeviceMapping {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Features != nil {
		if err := s.Features.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportImageRequestDiskDeviceMapping struct {
	// The device name of disk N in the custom image.
	//
	// >  This parameter will be removed in the future. We recommend that you do not use this parameter to ensure future compatibility.
	//
	// example:
	//
	// null
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The size of disk N in the custom image. Unit: GiB.
	//
	// You can use this parameter to specify the sizes of the system disk and data disks in the custom image. When you specify the size of the system disk, make sure that the specified size is greater than or equal to the size of the imported image file. Unit: GiB. Valid values:
	//
	// 	- When the N value is 1, this parameter specifies the size of the system disk in the custom image. Valid values: 1 to 2048.
	//
	// 	- When the N value is an integer in the range of 2 to 17, this parameter specifies the size of a data disk in the custom image. Valid values: 1 to 2048.
	//
	// After the image file is uploaded to an OSS bucket, you can view the size of the image file in the OSS bucket.
	//
	// >  This parameter will be removed in the future. We recommend that you use `DiskDeviceMapping.N.DiskImageSize` to ensure future compatibility.
	//
	// example:
	//
	// 80
	DiskImSize *int32 `json:"DiskImSize,omitempty" xml:"DiskImSize,omitempty"`
	// The size of disk N in the custom image after the source image is imported.
	//
	// You can use this parameter to specify the sizes of the system disk and data disks in the custom image. When you specify the size of the system disk, make sure that the specified size is greater than or equal to the size of the imported image file. Unit: GiB. Valid values:
	//
	// 	- When the N value is 1, this parameter specifies the size of the system disk in the custom image. Valid values: 1 to 2048.
	//
	// 	- When the N value is an integer in the range of 2 to 17, this parameter specifies the size of a data disk in the custom image. Valid values: 1 to 2048.
	//
	// After the image file is uploaded to an OSS bucket, you can view the size of the image file in the OSS bucket.
	//
	// example:
	//
	// 80
	DiskImageSize *int32 `json:"DiskImageSize,omitempty" xml:"DiskImageSize,omitempty"`
	// The format of the source image. Valid values:
	//
	// 	- RAW
	//
	// 	- VHD
	//
	// 	- QCOW2
	//
	// 	- VMDK (invitational preview)
	//
	// This parameter is empty by default, which indicates that the system checks the image format and uses the check result as the value of this parameter.
	//
	// example:
	//
	// QCOW2
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The Object Storage Service (OSS) bucket where the image file is stored.
	//
	// >  Before you import images for the first time, you must use RAM to authorize ECS to access your OSS buckets. If ECS is not authorized to access your OSS buckets, the `NoSetRoletoECSServiceAcount` error code is returned when you call the ImportImage operation. For more information, see **Usage notes**.
	//
	// example:
	//
	// ecsimageos
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The name (key) of the object that the image file is stored as in the OSS bucket.
	//
	// example:
	//
	// CentOS_5.4_32.raw
	OSSObject *string `json:"OSSObject,omitempty" xml:"OSSObject,omitempty"`
}

func (s ImportImageRequestDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequestDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *ImportImageRequestDiskDeviceMapping) GetDevice() *string {
	return s.Device
}

func (s *ImportImageRequestDiskDeviceMapping) GetDiskImSize() *int32 {
	return s.DiskImSize
}

func (s *ImportImageRequestDiskDeviceMapping) GetDiskImageSize() *int32 {
	return s.DiskImageSize
}

func (s *ImportImageRequestDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *ImportImageRequestDiskDeviceMapping) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *ImportImageRequestDiskDeviceMapping) GetOSSObject() *string {
	return s.OSSObject
}

func (s *ImportImageRequestDiskDeviceMapping) SetDevice(v string) *ImportImageRequestDiskDeviceMapping {
	s.Device = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetDiskImSize(v int32) *ImportImageRequestDiskDeviceMapping {
	s.DiskImSize = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetDiskImageSize(v int32) *ImportImageRequestDiskDeviceMapping {
	s.DiskImageSize = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetFormat(v string) *ImportImageRequestDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetOSSBucket(v string) *ImportImageRequestDiskDeviceMapping {
	s.OSSBucket = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) SetOSSObject(v string) *ImportImageRequestDiskDeviceMapping {
	s.OSSObject = &v
	return s
}

func (s *ImportImageRequestDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}

type ImportImageRequestFeatures struct {
	// The metadata access mode version of the image. Valid values:
	//
	// 	- v1: You cannot set the metadata access mode to security hardening when you create instances from the image.
	//
	// 	- v2: You can set the metadata access mode to security hardening when you create instances from the image.
	//
	// Default value: v1.
	//
	// example:
	//
	// v2
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	// Specifies whether the image supports the Non-Volatile Memory Express (NVMe) protocol. Valid values:
	//
	// 	- supported: The image supports the NVMe protocol. Instances created from the image also support the NVMe protocol.
	//
	// 	- unsupported: The image does not support the NVMe protocol. Instances created from the image do not support the NVMe protocol.
	//
	// example:
	//
	// supported
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s ImportImageRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequestFeatures) GoString() string {
	return s.String()
}

func (s *ImportImageRequestFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *ImportImageRequestFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *ImportImageRequestFeatures) SetImdsSupport(v string) *ImportImageRequestFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *ImportImageRequestFeatures) SetNvmeSupport(v string) *ImportImageRequestFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *ImportImageRequestFeatures) Validate() error {
	return dara.Validate(s)
}

type ImportImageRequestTag struct {
	// The key of tag N of the image. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the image. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImportImageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequestTag) GoString() string {
	return s.String()
}

func (s *ImportImageRequestTag) GetKey() *string {
	return s.Key
}

func (s *ImportImageRequestTag) GetValue() *string {
	return s.Value
}

func (s *ImportImageRequestTag) SetKey(v string) *ImportImageRequestTag {
	s.Key = &v
	return s
}

func (s *ImportImageRequestTag) SetValue(v string) *ImportImageRequestTag {
	s.Value = &v
	return s
}

func (s *ImportImageRequestTag) Validate() error {
	return dara.Validate(s)
}
