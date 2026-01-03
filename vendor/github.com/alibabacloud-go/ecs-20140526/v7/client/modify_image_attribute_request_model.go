// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBootMode(v string) *ModifyImageAttributeRequest
	GetBootMode() *string
	SetDescription(v string) *ModifyImageAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyImageAttributeRequest
	GetDryRun() *bool
	SetFeatures(v *ModifyImageAttributeRequestFeatures) *ModifyImageAttributeRequest
	GetFeatures() *ModifyImageAttributeRequestFeatures
	SetImageFamily(v string) *ModifyImageAttributeRequest
	GetImageFamily() *string
	SetImageId(v string) *ModifyImageAttributeRequest
	GetImageId() *string
	SetImageName(v string) *ModifyImageAttributeRequest
	GetImageName() *string
	SetLicenseType(v string) *ModifyImageAttributeRequest
	GetLicenseType() *string
	SetOwnerAccount(v string) *ModifyImageAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyImageAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyImageAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyImageAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyImageAttributeRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ModifyImageAttributeRequest
	GetStatus() *string
}

type ModifyImageAttributeRequest struct {
	// The new boot mode of the image. Valid values:
	//
	// 	- BIOS: BIOS mode
	//
	// 	- UEFI: Unified Extensible Firmware Interface (UEFI) mode
	//
	// 	- UEFI-Preferred: BIOS mode and UEFI mode
	//
	// >  Before you change this parameter, make sure that you are familiar with the boot modes supported by the image. If you specify a boot mode that is not supported by the image, ECS instances created from the image cannot start as expected. For information about the boot modes of images, see the [Boot modes of custom images](~~2244655#b9caa9b8bb1wf~~) section of the "Best practices for ECS instance boot modes" topic.
	//
	// example:
	//
	// BIOS
	BootMode *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	// The new description of the custom image. The description must be 2 to 256 characters in length It cannot start with [http:// or https://.](http://https://。)
	//
	// This parameter is empty by default, which specifies that the original description is retained.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The attributes of the custom image.
	//
	// if can be null:
	// true
	Features *ModifyImageAttributeRequestFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	// The name of the image family. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. [It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).](http://https://。、（.）、（:）、（_）（-）。)
	//
	// By default, this parameter is empty.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp18ygjuqnwhechc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the custom image. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with acs: or aliyun. [It cannot contain http:// or https://. It can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).](http://https://。、（.）、（:）、（_）（-）。)
	//
	// By default, this parameter is empty. In this case, the original name is retained.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the license that is used to activate the operating system after the image is imported. Set the value to BYOL.
	//
	// BYOL: The license that comes with the source operating system is used. When you use the BYOL license, make sure that your license key is supported by Alibaba Cloud.
	//
	// example:
	//
	// Auto
	LicenseType  *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the custom image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new state of the custom image. Valid values:
	//
	// 	- Deprecated: puts the image into the Deprecated state. If the custom image is shared, you must unshare it before you can put it into the Deprecated state. Images in the Deprecated state cannot be shared or copied, but can be used to create instances or replace system disks.
	//
	// 	- Available: puts the image into the Available state. You can restore an image from the Deprecated state to the Available state.
	//
	// > If you want to roll back a custom image in the image family to a previous version, you can put the latest available custom image into the Deprecated state. If no custom images are in the Available state within the image family, an image family cannot be used to create instances. Proceed with caution if only a single custom image is in the Available state within the image family.
	//
	// example:
	//
	// Deprecated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) GetBootMode() *string {
	return s.BootMode
}

func (s *ModifyImageAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyImageAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyImageAttributeRequest) GetFeatures() *ModifyImageAttributeRequestFeatures {
	return s.Features
}

func (s *ModifyImageAttributeRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *ModifyImageAttributeRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageAttributeRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ModifyImageAttributeRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *ModifyImageAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyImageAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyImageAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImageAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyImageAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyImageAttributeRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyImageAttributeRequest) SetBootMode(v string) *ModifyImageAttributeRequest {
	s.BootMode = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetDescription(v string) *ModifyImageAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetDryRun(v bool) *ModifyImageAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetFeatures(v *ModifyImageAttributeRequestFeatures) *ModifyImageAttributeRequest {
	s.Features = v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageFamily(v string) *ModifyImageAttributeRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageName(v string) *ModifyImageAttributeRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetLicenseType(v string) *ModifyImageAttributeRequest {
	s.LicenseType = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetOwnerAccount(v string) *ModifyImageAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetOwnerId(v int64) *ModifyImageAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetRegionId(v string) *ModifyImageAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetResourceOwnerAccount(v string) *ModifyImageAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetResourceOwnerId(v int64) *ModifyImageAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetStatus(v string) *ModifyImageAttributeRequest {
	s.Status = &v
	return s
}

func (s *ModifyImageAttributeRequest) Validate() error {
	if s.Features != nil {
		if err := s.Features.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyImageAttributeRequestFeatures struct {
	// The image metadata access mode. Valid values:
	//
	// 	- v1: You cannot set the image metadata access mode to security hardening when you create instances from the image.
	//
	// 	- v2: You can set the image metadata access mode to security hardening when you create instances from the image.
	//
	//     **
	//
	//     **Note*	- You cannot change the value of ImdsSupport from v2 to v1 for an image. To change the value of ImdsSupport from v2 to v1 for an image, use the snapshots associated with the image to create an image and set ImdsSupport to v1 for the new image.
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
	// if can be null:
	// true
	//
	// example:
	//
	// supported
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s ModifyImageAttributeRequestFeatures) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageAttributeRequestFeatures) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequestFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *ModifyImageAttributeRequestFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *ModifyImageAttributeRequestFeatures) SetImdsSupport(v string) *ModifyImageAttributeRequestFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *ModifyImageAttributeRequestFeatures) SetNvmeSupport(v string) *ModifyImageAttributeRequestFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *ModifyImageAttributeRequestFeatures) Validate() error {
	return dara.Validate(s)
}
