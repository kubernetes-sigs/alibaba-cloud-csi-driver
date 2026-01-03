// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImagePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddAccount(v []*int64) *CreateImagePipelineRequest
	GetAddAccount() []*int64
	SetAdvancedOptions(v *CreateImagePipelineRequestAdvancedOptions) *CreateImagePipelineRequest
	GetAdvancedOptions() *CreateImagePipelineRequestAdvancedOptions
	SetBaseImage(v string) *CreateImagePipelineRequest
	GetBaseImage() *string
	SetBaseImageType(v string) *CreateImagePipelineRequest
	GetBaseImageType() *string
	SetBuildContent(v string) *CreateImagePipelineRequest
	GetBuildContent() *string
	SetClientToken(v string) *CreateImagePipelineRequest
	GetClientToken() *string
	SetDeleteInstanceOnFailure(v bool) *CreateImagePipelineRequest
	GetDeleteInstanceOnFailure() *bool
	SetDescription(v string) *CreateImagePipelineRequest
	GetDescription() *string
	SetImageFamily(v string) *CreateImagePipelineRequest
	GetImageFamily() *string
	SetImageName(v string) *CreateImagePipelineRequest
	GetImageName() *string
	SetImageOptions(v *CreateImagePipelineRequestImageOptions) *CreateImagePipelineRequest
	GetImageOptions() *CreateImagePipelineRequestImageOptions
	SetImportImageOptions(v *CreateImagePipelineRequestImportImageOptions) *CreateImagePipelineRequest
	GetImportImageOptions() *CreateImagePipelineRequestImportImageOptions
	SetInstanceType(v string) *CreateImagePipelineRequest
	GetInstanceType() *string
	SetInternetMaxBandwidthOut(v int32) *CreateImagePipelineRequest
	GetInternetMaxBandwidthOut() *int32
	SetName(v string) *CreateImagePipelineRequest
	GetName() *string
	SetNvmeSupport(v string) *CreateImagePipelineRequest
	GetNvmeSupport() *string
	SetOwnerAccount(v string) *CreateImagePipelineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateImagePipelineRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateImagePipelineRequest
	GetRegionId() *string
	SetRepairMode(v string) *CreateImagePipelineRequest
	GetRepairMode() *string
	SetResourceGroupId(v string) *CreateImagePipelineRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateImagePipelineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateImagePipelineRequest
	GetResourceOwnerId() *int64
	SetSystemDiskSize(v int32) *CreateImagePipelineRequest
	GetSystemDiskSize() *int32
	SetTag(v []*CreateImagePipelineRequestTag) *CreateImagePipelineRequest
	GetTag() []*CreateImagePipelineRequestTag
	SetTestContent(v string) *CreateImagePipelineRequest
	GetTestContent() *string
	SetToRegionId(v []*string) *CreateImagePipelineRequest
	GetToRegionId() []*string
	SetVSwitchId(v string) *CreateImagePipelineRequest
	GetVSwitchId() *string
}

type CreateImagePipelineRequest struct {
	// The IDs of Alibaba Cloud accounts to which to share the image that will be created based on the image template. You can specify up to 20 account IDs.
	//
	// example:
	//
	// 1234567890
	AddAccount []*int64 `json:"AddAccount,omitempty" xml:"AddAccount,omitempty" type:"Repeated"`
	// The advanced settings.
	AdvancedOptions *CreateImagePipelineRequestAdvancedOptions `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty" type:"Struct"`
	// The source image.
	//
	// 	- If you set `BaseImageType` to IMAGE, set BaseImage to the ID of a custom image.
	//
	// 	- If you set `BaseImageType` to IMAGE_FAMILY, set BaseImage to the name of an image family.
	//
	// 	- If you set `BaseImageType` to OSS, you do not need to specify BaseImage.
	//
	// example:
	//
	// m-bp67acfmxazb4p****
	BaseImage *string `json:"BaseImage,omitempty" xml:"BaseImage,omitempty"`
	// The type of the source image. Valid values:
	//
	// 	- IMAGE: image
	//
	// 	- IMAGE_FAMILY: image family
	//
	// 	- OSS: Object Storage Service (OSS) object
	//
	// This parameter is required.
	//
	// example:
	//
	// IMAGE
	BaseImageType *string `json:"BaseImageType,omitempty" xml:"BaseImageType,omitempty"`
	// The build content in the image template. The content cannot exceed 16 KB in size. For information about the commands supported by Image Builder, see [Commands supported by Image Builder](https://help.aliyun.com/document_detail/200206.html).
	//
	// example:
	//
	// FROM IMAGE:m-bp67acfmxazb4p****
	BuildContent *string `json:"BuildContent,omitempty" xml:"BuildContent,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.***	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to release the intermediate instance when the image cannot be created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// > If the intermediate instance cannot be started, the instance is released by default.
	//
	// example:
	//
	// true
	DeleteInstanceOnFailure *bool `json:"DeleteInstanceOnFailure,omitempty" xml:"DeleteInstanceOnFailure,omitempty"`
	// The description of the image template. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated
	//
	// The family of the image created based on the image template.
	//
	// >  This parameter is no longer used. We recommend that you use ImageOptions.ImageFamily.
	//
	// example:
	//
	// null
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// Deprecated
	//
	// The name prefix of the image created based on the image template.
	//
	// >  This parameter is no longer used. We recommend that you use ImageOptions.ImageName.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The attributes of the image created based on the image template.
	ImageOptions *CreateImagePipelineRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The attributes and settings of the image that you want to import. If you set `BaseImageType` to OSS, you must specify this parameter.
	ImportImageOptions *CreateImagePipelineRequestImportImageOptions `json:"ImportImageOptions,omitempty" xml:"ImportImageOptions,omitempty" type:"Struct"`
	// The instance type. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query instance types.
	//
	// If you do not configure this parameter, an instance type that provides the fewest vCPUs and memory resources is automatically selected. This configuration is subject to resource availability of instance types. For example, the ecs.g6.large instance type is automatically selected. If available ecs.g6.large resources are insufficient, the ecs.g6.xlarge instance type is selected.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The size of the outbound public bandwidth for the intermediate instance. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The name of the launch template. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// >  If you do not specify `Name`, the return value of `ImagePipelineId` is used.
	//
	// example:
	//
	// testImagePipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// Specifies whether the image created based on the image template supports the NVMe protocol.
	//
	// >  This parameter is no longer used. We recommend that you use ImageOptions.ImageFeatures.NvmeSupport.
	//
	// example:
	//
	// auto
	NvmeSupport  *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repair mode of the image template.
	//
	// Valid values:
	//
	// 	- Standard: the standard mode.
	//
	//     Supported check items in Linux operating systems:
	//
	//     	- GUESTOS.CloudInit
	//
	//     	- GUESTOS.Dhcp
	//
	//     	- GUESTOS.Virtio
	//
	//     	- GUESTOS.OnlineResizeFS
	//
	//     	- GUESTOS.Grub
	//
	//     	- GUESTOS.Fstab
	//
	//     Supported check items in Windows operating systems:
	//
	//     	- GUESTOS.Virtio
	//
	//     	- GUESTOS.Update
	//
	//     	- GUESTOS.Hotfix
	//
	//     	- GUESTOS.Server
	//
	// >  As the check and repair capabilities continue to improve, the number of check items may increase. For more information about check items, see [Overview of image check](https://help.aliyun.com/document_detail/439819.html).
	//
	// example:
	//
	// null
	RepairMode *string `json:"RepairMode,omitempty" xml:"RepairMode,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The system disk size of the intermediate instance. Unit: GiB. Valid values: 20 to 500.
	//
	// Default value: 40.
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags to add to the template.
	Tag []*CreateImagePipelineRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The test content in the image template. The content cannot exceed 16 KB in size. For information about the commands supported by Image Builder, see [Commands supported by Image Builder](https://help.aliyun.com/document_detail/200206.html).
	//
	// example:
	//
	// null
	TestContent *string `json:"TestContent,omitempty" xml:"TestContent,omitempty"`
	// The IDs of regions to which you want to distribute the image that is created based on the image template. You can specify up to 20 region IDs.
	//
	// If you do not specify this parameter, the image is created only in the current region.
	//
	// example:
	//
	// cn-hangzhou
	ToRegionId []*string `json:"ToRegionId,omitempty" xml:"ToRegionId,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	//
	// If you do not specify this parameter, a new VPC and vSwitch are created. Make sure that the VPC quota in your account is sufficient. For more information, see [Limits and quotas](https://help.aliyun.com/document_detail/27750.html).
	//
	// example:
	//
	// vsw-bp67acfmxazb4p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateImagePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequest) GetAddAccount() []*int64 {
	return s.AddAccount
}

func (s *CreateImagePipelineRequest) GetAdvancedOptions() *CreateImagePipelineRequestAdvancedOptions {
	return s.AdvancedOptions
}

func (s *CreateImagePipelineRequest) GetBaseImage() *string {
	return s.BaseImage
}

func (s *CreateImagePipelineRequest) GetBaseImageType() *string {
	return s.BaseImageType
}

func (s *CreateImagePipelineRequest) GetBuildContent() *string {
	return s.BuildContent
}

func (s *CreateImagePipelineRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImagePipelineRequest) GetDeleteInstanceOnFailure() *bool {
	return s.DeleteInstanceOnFailure
}

func (s *CreateImagePipelineRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImagePipelineRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateImagePipelineRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImagePipelineRequest) GetImageOptions() *CreateImagePipelineRequestImageOptions {
	return s.ImageOptions
}

func (s *CreateImagePipelineRequest) GetImportImageOptions() *CreateImagePipelineRequestImportImageOptions {
	return s.ImportImageOptions
}

func (s *CreateImagePipelineRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateImagePipelineRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateImagePipelineRequest) GetName() *string {
	return s.Name
}

func (s *CreateImagePipelineRequest) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *CreateImagePipelineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateImagePipelineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateImagePipelineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImagePipelineRequest) GetRepairMode() *string {
	return s.RepairMode
}

func (s *CreateImagePipelineRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateImagePipelineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateImagePipelineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateImagePipelineRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateImagePipelineRequest) GetTag() []*CreateImagePipelineRequestTag {
	return s.Tag
}

func (s *CreateImagePipelineRequest) GetTestContent() *string {
	return s.TestContent
}

func (s *CreateImagePipelineRequest) GetToRegionId() []*string {
	return s.ToRegionId
}

func (s *CreateImagePipelineRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateImagePipelineRequest) SetAddAccount(v []*int64) *CreateImagePipelineRequest {
	s.AddAccount = v
	return s
}

func (s *CreateImagePipelineRequest) SetAdvancedOptions(v *CreateImagePipelineRequestAdvancedOptions) *CreateImagePipelineRequest {
	s.AdvancedOptions = v
	return s
}

func (s *CreateImagePipelineRequest) SetBaseImage(v string) *CreateImagePipelineRequest {
	s.BaseImage = &v
	return s
}

func (s *CreateImagePipelineRequest) SetBaseImageType(v string) *CreateImagePipelineRequest {
	s.BaseImageType = &v
	return s
}

func (s *CreateImagePipelineRequest) SetBuildContent(v string) *CreateImagePipelineRequest {
	s.BuildContent = &v
	return s
}

func (s *CreateImagePipelineRequest) SetClientToken(v string) *CreateImagePipelineRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImagePipelineRequest) SetDeleteInstanceOnFailure(v bool) *CreateImagePipelineRequest {
	s.DeleteInstanceOnFailure = &v
	return s
}

func (s *CreateImagePipelineRequest) SetDescription(v string) *CreateImagePipelineRequest {
	s.Description = &v
	return s
}

func (s *CreateImagePipelineRequest) SetImageFamily(v string) *CreateImagePipelineRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateImagePipelineRequest) SetImageName(v string) *CreateImagePipelineRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImagePipelineRequest) SetImageOptions(v *CreateImagePipelineRequestImageOptions) *CreateImagePipelineRequest {
	s.ImageOptions = v
	return s
}

func (s *CreateImagePipelineRequest) SetImportImageOptions(v *CreateImagePipelineRequestImportImageOptions) *CreateImagePipelineRequest {
	s.ImportImageOptions = v
	return s
}

func (s *CreateImagePipelineRequest) SetInstanceType(v string) *CreateImagePipelineRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateImagePipelineRequest) SetInternetMaxBandwidthOut(v int32) *CreateImagePipelineRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateImagePipelineRequest) SetName(v string) *CreateImagePipelineRequest {
	s.Name = &v
	return s
}

func (s *CreateImagePipelineRequest) SetNvmeSupport(v string) *CreateImagePipelineRequest {
	s.NvmeSupport = &v
	return s
}

func (s *CreateImagePipelineRequest) SetOwnerAccount(v string) *CreateImagePipelineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImagePipelineRequest) SetOwnerId(v int64) *CreateImagePipelineRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImagePipelineRequest) SetRegionId(v string) *CreateImagePipelineRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImagePipelineRequest) SetRepairMode(v string) *CreateImagePipelineRequest {
	s.RepairMode = &v
	return s
}

func (s *CreateImagePipelineRequest) SetResourceGroupId(v string) *CreateImagePipelineRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImagePipelineRequest) SetResourceOwnerAccount(v string) *CreateImagePipelineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImagePipelineRequest) SetResourceOwnerId(v int64) *CreateImagePipelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateImagePipelineRequest) SetSystemDiskSize(v int32) *CreateImagePipelineRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateImagePipelineRequest) SetTag(v []*CreateImagePipelineRequestTag) *CreateImagePipelineRequest {
	s.Tag = v
	return s
}

func (s *CreateImagePipelineRequest) SetTestContent(v string) *CreateImagePipelineRequest {
	s.TestContent = &v
	return s
}

func (s *CreateImagePipelineRequest) SetToRegionId(v []*string) *CreateImagePipelineRequest {
	s.ToRegionId = v
	return s
}

func (s *CreateImagePipelineRequest) SetVSwitchId(v string) *CreateImagePipelineRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateImagePipelineRequest) Validate() error {
	if s.AdvancedOptions != nil {
		if err := s.AdvancedOptions.Validate(); err != nil {
			return err
		}
	}
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.ImportImageOptions != nil {
		if err := s.ImportImageOptions.Validate(); err != nil {
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

type CreateImagePipelineRequestAdvancedOptions struct {
	// Specifies whether to disable the feature that automatically adds a suffix to the name of the image created based on the image template. Valid value:
	//
	// 	- disable
	//
	// example:
	//
	// disable
	ImageNameSuffix *string `json:"ImageNameSuffix,omitempty" xml:"ImageNameSuffix,omitempty"`
	// Specifies whether to retain Cloud Assistant Agent that is installed during the image building process. During the image building process, the system automatically installs Cloud Assistant Agent on the intermediate instance to run commands. You can choose whether to retain Cloud Assistant Agent that is installed during the image building process in the new image. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  The setting of this parameter does not affect Cloud Assistant Agent that comes with your image.
	//
	// example:
	//
	// true
	RetainCloudAssistant *bool `json:"RetainCloudAssistant,omitempty" xml:"RetainCloudAssistant,omitempty"`
}

func (s CreateImagePipelineRequestAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestAdvancedOptions) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestAdvancedOptions) GetImageNameSuffix() *string {
	return s.ImageNameSuffix
}

func (s *CreateImagePipelineRequestAdvancedOptions) GetRetainCloudAssistant() *bool {
	return s.RetainCloudAssistant
}

func (s *CreateImagePipelineRequestAdvancedOptions) SetImageNameSuffix(v string) *CreateImagePipelineRequestAdvancedOptions {
	s.ImageNameSuffix = &v
	return s
}

func (s *CreateImagePipelineRequestAdvancedOptions) SetRetainCloudAssistant(v bool) *CreateImagePipelineRequestAdvancedOptions {
	s.RetainCloudAssistant = &v
	return s
}

func (s *CreateImagePipelineRequestAdvancedOptions) Validate() error {
	return dara.Validate(s)
}

type CreateImagePipelineRequestImageOptions struct {
	// The description of the image. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image family. The image family name must be 2 to 128 characters in length. The name must start with a letter and cannot start with acs: or aliyun. The name cannot contain http:// or https:// and can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// family
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The feature attributes of the image.
	ImageFeatures *CreateImagePipelineRequestImageOptionsImageFeatures `json:"ImageFeatures,omitempty" xml:"ImageFeatures,omitempty" type:"Struct"`
	// The prefix of the image name. The prefix must be 2 to 64 characters in length. The prefix must start with a letter and cannot start with `http://` or `https://`. The prefix can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// The system generates the final image name that consists of the specified prefix and the ID of the build task (`ExecutionId`) in the format of `{ImageName}_{ExecutionId}`.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The tags to add to the image.
	ImageTags []*CreateImagePipelineRequestImageOptionsImageTags `json:"ImageTags,omitempty" xml:"ImageTags,omitempty" type:"Repeated"`
}

func (s CreateImagePipelineRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestImageOptions) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestImageOptions) GetDescription() *string {
	return s.Description
}

func (s *CreateImagePipelineRequestImageOptions) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateImagePipelineRequestImageOptions) GetImageFeatures() *CreateImagePipelineRequestImageOptionsImageFeatures {
	return s.ImageFeatures
}

func (s *CreateImagePipelineRequestImageOptions) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImagePipelineRequestImageOptions) GetImageTags() []*CreateImagePipelineRequestImageOptionsImageTags {
	return s.ImageTags
}

func (s *CreateImagePipelineRequestImageOptions) SetDescription(v string) *CreateImagePipelineRequestImageOptions {
	s.Description = &v
	return s
}

func (s *CreateImagePipelineRequestImageOptions) SetImageFamily(v string) *CreateImagePipelineRequestImageOptions {
	s.ImageFamily = &v
	return s
}

func (s *CreateImagePipelineRequestImageOptions) SetImageFeatures(v *CreateImagePipelineRequestImageOptionsImageFeatures) *CreateImagePipelineRequestImageOptions {
	s.ImageFeatures = v
	return s
}

func (s *CreateImagePipelineRequestImageOptions) SetImageName(v string) *CreateImagePipelineRequestImageOptions {
	s.ImageName = &v
	return s
}

func (s *CreateImagePipelineRequestImageOptions) SetImageTags(v []*CreateImagePipelineRequestImageOptionsImageTags) *CreateImagePipelineRequestImageOptions {
	s.ImageTags = v
	return s
}

func (s *CreateImagePipelineRequestImageOptions) Validate() error {
	if s.ImageFeatures != nil {
		if err := s.ImageFeatures.Validate(); err != nil {
			return err
		}
	}
	if s.ImageTags != nil {
		for _, item := range s.ImageTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImagePipelineRequestImageOptionsImageFeatures struct {
	// Specifies whether the image created based on the image template supports the NVMe protocol. Valid values:
	//
	// 	- supported: The image supports the NVMe protocol. Instances created from the image also support the NVMe protocol.
	//
	// 	- unsupported: The image does not support the NVMe protocol. Instances created from the image do not support the NVMe protocol.
	//
	// 	- auto: The system automatically detects whether the image supports the NVMe protocol. The system automatically detects whether the NVMe driver is installed on your image before the new image is built. If you install or uninstall the NVMe driver during the image building process, the detection result may be incorrect. We recommend that you set the value to supported or unsupported based on the image building content.
	//
	// example:
	//
	// auto
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s CreateImagePipelineRequestImageOptionsImageFeatures) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestImageOptionsImageFeatures) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestImageOptionsImageFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *CreateImagePipelineRequestImageOptionsImageFeatures) SetNvmeSupport(v string) *CreateImagePipelineRequestImageOptionsImageFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *CreateImagePipelineRequestImageOptionsImageFeatures) Validate() error {
	return dara.Validate(s)
}

type CreateImagePipelineRequestImageOptionsImageTags struct {
	// The key of tag N to add to the image. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the image. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImagePipelineRequestImageOptionsImageTags) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestImageOptionsImageTags) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestImageOptionsImageTags) GetKey() *string {
	return s.Key
}

func (s *CreateImagePipelineRequestImageOptionsImageTags) GetValue() *string {
	return s.Value
}

func (s *CreateImagePipelineRequestImageOptionsImageTags) SetKey(v string) *CreateImagePipelineRequestImageOptionsImageTags {
	s.Key = &v
	return s
}

func (s *CreateImagePipelineRequestImageOptionsImageTags) SetValue(v string) *CreateImagePipelineRequestImageOptionsImageTags {
	s.Value = &v
	return s
}

func (s *CreateImagePipelineRequestImageOptionsImageTags) Validate() error {
	return dara.Validate(s)
}

type CreateImagePipelineRequestImportImageOptions struct {
	// The system architecture of the system disk. If you specify a data disk snapshot to create the system disk of the image, use Architecture to specify the system architecture of the system disk. Valid values:
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
	// The new boot mode of the image. Valid values:
	//
	// 	- BIOS: BIOS mode
	//
	// 	- UEFI: Unified Extensible Firmware Interface (UEFI) mode
	//
	// Default value: BIOS. If you set Architecture to `arm64`, set this parameter to UEFI.
	//
	// >  Before you specify this parameter, make sure that you are familiar with the boot modes supported by the image. If you specify a boot mode that is not supported by the image, ECS instances created from the image cannot start as expected. For information about the boot modes of images, see the [Boot modes of images](~~2244655#b9caa9b8bb1wf~~) section of the "Best practices for ECS instance boot modes" topic.
	//
	// example:
	//
	// BIOS
	BootMode    *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information of disks from which the custom images are created.
	//
	// 	- When the N value is 1, this parameter creates a custom image from the system disk.
	//
	// 	- When the N value is an integer in the range of 2 to 17, this parameter creates a custom image from a data disk.
	DiskDeviceMappings []*CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Repeated"`
	// The attributes of the image.
	Features        *CreateImagePipelineRequestImportImageOptionsFeatures          `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	ImageName       *string                                                        `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImportImageTags []*CreateImagePipelineRequestImportImageOptionsImportImageTags `json:"ImportImageTags,omitempty" xml:"ImportImageTags,omitempty" type:"Repeated"`
	// The type of the license to use to activate the operating system after the image is imported. Valid values:
	//
	// 	- Auto: ECS detects the operating system of the image and allocates a license to the operating system. In this mode, the system first checks whether a license allocated by an official Alibaba Cloud channel is available for the operating system version specified by `Platform`. If a license allocated by an official Alibaba Cloud channel is available for the operating system version, the system allocates the license to the imported image. If no such license is available, the Bring Your Own License (BYOL) mode is used.
	//
	// 	- Aliyun: The license allocated by an official Alibaba Cloud channel for the operating system version specified by `Platform` is used.
	//
	// 	- BYOL: The license that comes with the source operating system is used. When you use the BYOL license, make sure that your license key is supported by Alibaba Cloud.
	//
	// Default value: Auto.
	//
	// example:
	//
	// Auto
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The operating system type. Valid value:
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
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The version of the operating system. Valid values:
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
	// 	- Other Windows
	//
	// Default value: Others Linux when the operating system type is linux, and Other Windows when the operating system type is windows.
	//
	// example:
	//
	// Aliyun
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// >  This parameter is in invitational preview.
	//
	// example:
	//
	// false
	RetainImportedImage *bool   `json:"RetainImportedImage,omitempty" xml:"RetainImportedImage,omitempty"`
	RetentionStrategy   *string `json:"RetentionStrategy,omitempty" xml:"RetentionStrategy,omitempty"`
	RoleName            *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateImagePipelineRequestImportImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestImportImageOptions) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestImportImageOptions) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateImagePipelineRequestImportImageOptions) GetBootMode() *string {
	return s.BootMode
}

func (s *CreateImagePipelineRequestImportImageOptions) GetDescription() *string {
	return s.Description
}

func (s *CreateImagePipelineRequestImportImageOptions) GetDiskDeviceMappings() []*CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *CreateImagePipelineRequestImportImageOptions) GetFeatures() *CreateImagePipelineRequestImportImageOptionsFeatures {
	return s.Features
}

func (s *CreateImagePipelineRequestImportImageOptions) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImagePipelineRequestImportImageOptions) GetImportImageTags() []*CreateImagePipelineRequestImportImageOptionsImportImageTags {
	return s.ImportImageTags
}

func (s *CreateImagePipelineRequestImportImageOptions) GetLicenseType() *string {
	return s.LicenseType
}

func (s *CreateImagePipelineRequestImportImageOptions) GetOSType() *string {
	return s.OSType
}

func (s *CreateImagePipelineRequestImportImageOptions) GetPlatform() *string {
	return s.Platform
}

func (s *CreateImagePipelineRequestImportImageOptions) GetRetainImportedImage() *bool {
	return s.RetainImportedImage
}

func (s *CreateImagePipelineRequestImportImageOptions) GetRetentionStrategy() *string {
	return s.RetentionStrategy
}

func (s *CreateImagePipelineRequestImportImageOptions) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateImagePipelineRequestImportImageOptions) SetArchitecture(v string) *CreateImagePipelineRequestImportImageOptions {
	s.Architecture = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetBootMode(v string) *CreateImagePipelineRequestImportImageOptions {
	s.BootMode = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetDescription(v string) *CreateImagePipelineRequestImportImageOptions {
	s.Description = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetDiskDeviceMappings(v []*CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) *CreateImagePipelineRequestImportImageOptions {
	s.DiskDeviceMappings = v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetFeatures(v *CreateImagePipelineRequestImportImageOptionsFeatures) *CreateImagePipelineRequestImportImageOptions {
	s.Features = v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetImageName(v string) *CreateImagePipelineRequestImportImageOptions {
	s.ImageName = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetImportImageTags(v []*CreateImagePipelineRequestImportImageOptionsImportImageTags) *CreateImagePipelineRequestImportImageOptions {
	s.ImportImageTags = v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetLicenseType(v string) *CreateImagePipelineRequestImportImageOptions {
	s.LicenseType = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetOSType(v string) *CreateImagePipelineRequestImportImageOptions {
	s.OSType = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetPlatform(v string) *CreateImagePipelineRequestImportImageOptions {
	s.Platform = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetRetainImportedImage(v bool) *CreateImagePipelineRequestImportImageOptions {
	s.RetainImportedImage = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetRetentionStrategy(v string) *CreateImagePipelineRequestImportImageOptions {
	s.RetentionStrategy = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) SetRoleName(v string) *CreateImagePipelineRequestImportImageOptions {
	s.RoleName = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptions) Validate() error {
	if s.DiskDeviceMappings != nil {
		for _, item := range s.DiskDeviceMappings {
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
	if s.ImportImageTags != nil {
		for _, item := range s.ImportImageTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings struct {
	// The size of disk N in the custom image after the source image is imported.
	//
	// You can use this parameter to specify the sizes of the system disk and data disks in the custom image. When you specify the size of the system disk, make sure that the specified size is greater than or equal to the size of the source image file. Unit: GiB. Valid values:
	//
	// 	- When the N value is 1, this parameter specifies the size of the system disk in the custom image. Valid values: 1 to 2048.
	//
	// 	- When the N value is an integer in the range of 2 to 17, this parameter creates a custom image from a data disk. Valid values: 1 to 2048.
	//
	// After the image file is uploaded to an OSS bucket, you can view the size of the image file in the OSS bucket.
	//
	// example:
	//
	// 40
	DiskImageSize *int32 `json:"DiskImageSize,omitempty" xml:"DiskImageSize,omitempty"`
	// The format of the source image. Valid values:
	//
	// 	- RAW
	//
	// 	- VHD
	//
	// 	- QCOW2
	//
	// This parameter is empty by default, which indicates that the system checks the format of the image and uses the check result as the value of this parameter.
	//
	// example:
	//
	// RAW
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The Object Storage Service (OSS) bucket where the image file is stored.
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

func (s CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) GetDiskImageSize() *int32 {
	return s.DiskImageSize
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) GetFormat() *string {
	return s.Format
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) GetOSSObject() *string {
	return s.OSSObject
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) SetDiskImageSize(v int32) *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings {
	s.DiskImageSize = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) SetFormat(v string) *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings {
	s.Format = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) SetOSSBucket(v string) *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings {
	s.OSSBucket = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) SetOSSObject(v string) *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings {
	s.OSSObject = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsDiskDeviceMappings) Validate() error {
	return dara.Validate(s)
}

type CreateImagePipelineRequestImportImageOptionsFeatures struct {
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	// Specifies whether the imported source image supports the Non-Volatile Memory Express (NVMe) protocol. Valid value:
	//
	// 	- supported Instances created from the image also support the NVMe protocol.
	//
	// 	- unsupported Instances created from the image do not support the NVMe protocol.
	//
	// Default value: unsupported.
	//
	// example:
	//
	// supported
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s CreateImagePipelineRequestImportImageOptionsFeatures) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestImportImageOptionsFeatures) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestImportImageOptionsFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *CreateImagePipelineRequestImportImageOptionsFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *CreateImagePipelineRequestImportImageOptionsFeatures) SetImdsSupport(v string) *CreateImagePipelineRequestImportImageOptionsFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsFeatures) SetNvmeSupport(v string) *CreateImagePipelineRequestImportImageOptionsFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsFeatures) Validate() error {
	return dara.Validate(s)
}

type CreateImagePipelineRequestImportImageOptionsImportImageTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImagePipelineRequestImportImageOptionsImportImageTags) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestImportImageOptionsImportImageTags) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestImportImageOptionsImportImageTags) GetKey() *string {
	return s.Key
}

func (s *CreateImagePipelineRequestImportImageOptionsImportImageTags) GetValue() *string {
	return s.Value
}

func (s *CreateImagePipelineRequestImportImageOptionsImportImageTags) SetKey(v string) *CreateImagePipelineRequestImportImageOptionsImportImageTags {
	s.Key = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsImportImageTags) SetValue(v string) *CreateImagePipelineRequestImportImageOptionsImportImageTags {
	s.Value = &v
	return s
}

func (s *CreateImagePipelineRequestImportImageOptionsImportImageTags) Validate() error {
	return dara.Validate(s)
}

type CreateImagePipelineRequestTag struct {
	// The key of tag N. Valid values of N: 1 to 20. You cannot specify empty strings as tag keys. The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value must be 0 to 128 characters in length. It cannot start with `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImagePipelineRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineRequestTag) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateImagePipelineRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateImagePipelineRequestTag) SetKey(v string) *CreateImagePipelineRequestTag {
	s.Key = &v
	return s
}

func (s *CreateImagePipelineRequestTag) SetValue(v string) *CreateImagePipelineRequestTag {
	s.Value = &v
	return s
}

func (s *CreateImagePipelineRequestTag) Validate() error {
	return dara.Validate(s)
}
