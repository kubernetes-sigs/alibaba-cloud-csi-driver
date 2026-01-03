// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *CreateImageRequest
	GetArchitecture() *string
	SetBootMode(v string) *CreateImageRequest
	GetBootMode() *string
	SetClientToken(v string) *CreateImageRequest
	GetClientToken() *string
	SetDescription(v string) *CreateImageRequest
	GetDescription() *string
	SetDetectionStrategy(v string) *CreateImageRequest
	GetDetectionStrategy() *string
	SetDiskDeviceMapping(v []*CreateImageRequestDiskDeviceMapping) *CreateImageRequest
	GetDiskDeviceMapping() []*CreateImageRequestDiskDeviceMapping
	SetDryRun(v bool) *CreateImageRequest
	GetDryRun() *bool
	SetFeatures(v *CreateImageRequestFeatures) *CreateImageRequest
	GetFeatures() *CreateImageRequestFeatures
	SetImageFamily(v string) *CreateImageRequest
	GetImageFamily() *string
	SetImageName(v string) *CreateImageRequest
	GetImageName() *string
	SetImageVersion(v string) *CreateImageRequest
	GetImageVersion() *string
	SetInstanceId(v string) *CreateImageRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateImageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateImageRequest
	GetOwnerId() *int64
	SetPlatform(v string) *CreateImageRequest
	GetPlatform() *string
	SetRegionId(v string) *CreateImageRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateImageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateImageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateImageRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *CreateImageRequest
	GetSnapshotId() *string
	SetTag(v []*CreateImageRequestTag) *CreateImageRequest
	GetTag() []*CreateImageRequestTag
}

type CreateImageRequest struct {
	// The system architecture of the system disk. If you specify a data disk snapshot to create the system disk of the custom image, use Architecture to specify the system architecture of the system disk. Valid values:
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
	// 	- BIOS: BIOS mode
	//
	// 	- UEFI: Unified Extensible Firmware Interface (UEFI) mode
	//
	// 	- UEFI-Preferred (default): BIOS mode and UEFI mode
	//
	// >  Before you specify this parameter, make sure that you are familiar with the boot modes supported by the image. If you specify a boot mode that is not supported by the image, ECS instances created from the image cannot start as expected. For information about the boot modes of images, see the [Boot modes of images](~~2244655#b9caa9b8bb1wf~~) section of the "Best practices for ECS instance boot modes" topic.
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The image description. The description must be 2 to 256 characters in length and cannot start with [http:// or https://.](http://https://。)
	//
	// example:
	//
	// ImageTestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mode in which to check the custom image. If you do not specify this parameter, the image is not checked. Only the standard check mode is supported.
	//
	// >  This parameter is supported for most Linux and Windows operating system versions. For information about image check items and operating system limits for image check, see [Overview of image check](https://help.aliyun.com/document_detail/439819.html) and [Operating system limits for image check](https://help.aliyun.com/document_detail/475800.html).
	//
	// example:
	//
	// Standard
	DetectionStrategy *string `json:"DetectionStrategy,omitempty" xml:"DetectionStrategy,omitempty"`
	// Details of the disks and snapshots from which the custom image is created. If you want to create a custom image based on a system disk snapshot and data disk snapshots, use this parameter to specify the snapshots.
	DiskDeviceMapping []*CreateImageRequestDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
	DryRun            *bool                                  `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The attributes of the custom image.
	Features *CreateImageRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The name of the image family. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with acs: or aliyun. The name cannot contain http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The name of the custom image. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// TestCentOS
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image version.
	//
	// >  If you specify an ECS instance that runs an Alibaba Cloud Marketplace image or a custom image derived from an Alibaba Cloud Marketplace image by using `InstanceId`, you must leave this parameter empty or set this parameter to the `ImageVersion` value of the image run by the specified ECS instance.
	//
	// example:
	//
	// 2017011017
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The ID of the ECS instance from which to create the custom image. To create a custom image from an ECS instance, you must specify this parameter.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system distribution for the system disk in the custom image. If you specify a data disk snapshot to create the system disk of the custom image, use Platform to specify the operating system distribution for the system disk. Valid values:
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
	// CentOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID of the custom image that you want to create. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the custom image. If you leave this parameter empty, the image is assigned to the default resource group.
	//
	// >  If you call the CreateImage operation as a Resource Access Management (RAM) user who does not have permissions on the default resource group and leave `ResourceGroupId` empty, the `Forbidden: User not authorized to operate on the specified resource` error message is returned. You must specify the ID of a resource group on which the RAM user has permissions or grant the RAM user permissions on the default resource group, and then call the CreateImage operation again.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the snapshot from which to create the custom image.
	//
	// >  To create a custom image from only a system disk snapshot of an ECS instance, you can specify this parameter or `DiskDeviceMapping.N.SnapshotId` to specify the snapshot ID. If you add data disk snapshots, you can use only `DiskDeviceMapping.N.SnapshotId` to specify snapshots.
	//
	// example:
	//
	// s-bp17441ohwkdca0****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The tags.
	Tag []*CreateImageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateImageRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *CreateImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImageRequest) GetDetectionStrategy() *string {
	return s.DetectionStrategy
}

func (s *CreateImageRequest) GetDiskDeviceMapping() []*CreateImageRequestDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *CreateImageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateImageRequest) GetFeatures() *CreateImageRequestFeatures {
	return s.Features
}

func (s *CreateImageRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImageRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *CreateImageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateImageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateImageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateImageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateImageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateImageRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateImageRequest) GetTag() []*CreateImageRequestTag {
	return s.Tag
}

func (s *CreateImageRequest) SetArchitecture(v string) *CreateImageRequest {
	s.Architecture = &v
	return s
}

func (s *CreateImageRequest) SetBootMode(v string) *CreateImageRequest {
	s.BootMode = &v
	return s
}

func (s *CreateImageRequest) SetClientToken(v string) *CreateImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetDetectionStrategy(v string) *CreateImageRequest {
	s.DetectionStrategy = &v
	return s
}

func (s *CreateImageRequest) SetDiskDeviceMapping(v []*CreateImageRequestDiskDeviceMapping) *CreateImageRequest {
	s.DiskDeviceMapping = v
	return s
}

func (s *CreateImageRequest) SetDryRun(v bool) *CreateImageRequest {
	s.DryRun = &v
	return s
}

func (s *CreateImageRequest) SetFeatures(v *CreateImageRequestFeatures) *CreateImageRequest {
	s.Features = v
	return s
}

func (s *CreateImageRequest) SetImageFamily(v string) *CreateImageRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetImageVersion(v string) *CreateImageRequest {
	s.ImageVersion = &v
	return s
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetOwnerAccount(v string) *CreateImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImageRequest) SetOwnerId(v int64) *CreateImageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImageRequest) SetPlatform(v string) *CreateImageRequest {
	s.Platform = &v
	return s
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetResourceGroupId(v string) *CreateImageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageRequest) SetResourceOwnerAccount(v string) *CreateImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImageRequest) SetResourceOwnerId(v int64) *CreateImageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotId(v string) *CreateImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequest) SetTag(v []*CreateImageRequestTag) *CreateImageRequest {
	s.Tag = v
	return s
}

func (s *CreateImageRequest) Validate() error {
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

type CreateImageRequestDiskDeviceMapping struct {
	// The device name of disk N in the custom image. Valid values:
	//
	// 	- The device name of the system disk must be /dev/xvda.
	//
	// 	- The device names of the data disks are unique and range from /dev/xvdb to /dev/xvdz in alphabetical order.
	//
	// example:
	//
	// /dev/vdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The type of disk N in the custom image. You can specify this parameter to create the system disk of the custom image from a data disk snapshot. If you do not specify this parameter, the disk type is determined by the corresponding snapshot. Valid values:
	//
	// 	- system: system disk. You can specify only one snapshot to use to create the system disk in the custom image.
	//
	// 	- data: data disk. You can specify up to 16 snapshots to use to create data disks in the custom image.
	//
	// example:
	//
	// system
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The size of disk N in the custom image. Unit: GiB. The valid values and default value of DiskDeviceMapping.N.Size vary based on the value of DiskDeviceMapping.N.SnapshotId.
	//
	// 	- If you leave DiskDeviceMapping.N.SnapshotId empty, DiskDeviceMapping.N.Size has the following valid values and default values:
	//
	//     	- For basic disks, the valid values range from 5 to 2000, and the default value is 5.
	//
	//     	- For other disks, the valid values range from 20 to 32768, and the default value is 20.
	//
	// 	- If you specify DiskDeviceMapping.N.SnapshotId, the value of DiskDeviceMapping.N.Size must be greater than or equal to the size of the specified snapshot. The default value of DiskDeviceMapping.N.Size is the size of the specified snapshot.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-bp17441ohwkdca0****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateImageRequestDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *CreateImageRequestDiskDeviceMapping) GetDevice() *string {
	return s.Device
}

func (s *CreateImageRequestDiskDeviceMapping) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateImageRequestDiskDeviceMapping) GetSize() *int32 {
	return s.Size
}

func (s *CreateImageRequestDiskDeviceMapping) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateImageRequestDiskDeviceMapping) SetDevice(v string) *CreateImageRequestDiskDeviceMapping {
	s.Device = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) SetDiskType(v string) *CreateImageRequestDiskDeviceMapping {
	s.DiskType = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) SetSize(v int32) *CreateImageRequestDiskDeviceMapping {
	s.Size = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) SetSnapshotId(v string) *CreateImageRequestDiskDeviceMapping {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequestDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}

type CreateImageRequestFeatures struct {
	// The image metadata access mode. Valid values:
	//
	// 	- v1: You cannot set the image metadata access mode to security hardening when you create instances from the image.
	//
	// 	- v2: You can set the image metadata access mode to security hardening when you create instances from the image.
	//
	// When you use a snapshot to create instances, the default value is set to 1. If you use an instance to create an image, the value of the ImdsSupport parameter is used by default.
	//
	// example:
	//
	// v2
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
}

func (s CreateImageRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestFeatures) GoString() string {
	return s.String()
}

func (s *CreateImageRequestFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *CreateImageRequestFeatures) SetImdsSupport(v string) *CreateImageRequestFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *CreateImageRequestFeatures) Validate() error {
	return dara.Validate(s)
}

type CreateImageRequestTag struct {
	// The key of tag N of the custom image. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the custom image. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// ValueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateImageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateImageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateImageRequestTag) SetKey(v string) *CreateImageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateImageRequestTag) SetValue(v string) *CreateImageRequestTag {
	s.Value = &v
	return s
}

func (s *CreateImageRequestTag) Validate() error {
	return dara.Validate(s)
}
