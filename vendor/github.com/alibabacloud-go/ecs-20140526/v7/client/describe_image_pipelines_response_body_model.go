// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImagePipeline(v *DescribeImagePipelinesResponseBodyImagePipeline) *DescribeImagePipelinesResponseBody
	GetImagePipeline() *DescribeImagePipelinesResponseBodyImagePipeline
	SetMaxResults(v int32) *DescribeImagePipelinesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImagePipelinesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeImagePipelinesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImagePipelinesResponseBody
	GetTotalCount() *int32
}

type DescribeImagePipelinesResponseBody struct {
	// Details of the image templates.
	ImagePipeline *DescribeImagePipelinesResponseBodyImagePipeline `json:"ImagePipeline,omitempty" xml:"ImagePipeline,omitempty" type:"Struct"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. For information about how to use the return value, see the "Usage notes" section of this topic.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of image templates returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImagePipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBody) GetImagePipeline() *DescribeImagePipelinesResponseBodyImagePipeline {
	return s.ImagePipeline
}

func (s *DescribeImagePipelinesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImagePipelinesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImagePipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImagePipelinesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImagePipelinesResponseBody) SetImagePipeline(v *DescribeImagePipelinesResponseBodyImagePipeline) *DescribeImagePipelinesResponseBody {
	s.ImagePipeline = v
	return s
}

func (s *DescribeImagePipelinesResponseBody) SetMaxResults(v int32) *DescribeImagePipelinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeImagePipelinesResponseBody) SetNextToken(v string) *DescribeImagePipelinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImagePipelinesResponseBody) SetRequestId(v string) *DescribeImagePipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagePipelinesResponseBody) SetTotalCount(v int32) *DescribeImagePipelinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImagePipelinesResponseBody) Validate() error {
	if s.ImagePipeline != nil {
		if err := s.ImagePipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipeline struct {
	ImagePipelineSet []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet `json:"ImagePipelineSet,omitempty" xml:"ImagePipelineSet,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesResponseBodyImagePipeline) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipeline) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipeline) GetImagePipelineSet() []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	return s.ImagePipelineSet
}

func (s *DescribeImagePipelinesResponseBodyImagePipeline) SetImagePipelineSet(v []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) *DescribeImagePipelinesResponseBodyImagePipeline {
	s.ImagePipelineSet = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipeline) Validate() error {
	if s.ImagePipelineSet != nil {
		for _, item := range s.ImagePipelineSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet struct {
	// The IDs of Alibaba Cloud accounts to which to share the image that will be created based on the image template.
	AddAccounts *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts `json:"AddAccounts,omitempty" xml:"AddAccounts,omitempty" type:"Struct"`
	// The advanced settings.
	AdvancedOptions *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty" type:"Struct"`
	// The source image.
	//
	// 	- When `BaseImageType` is set to IMAGE, the value of this parameter is the ID of a custom image.
	//
	// 	- When `BaseImageType` is set to IMAGE_FAMILY, the value of this parameter is the name of an image family.
	//
	// example:
	//
	// m-bp67acfmxazb4p****
	BaseImage *string `json:"BaseImage,omitempty" xml:"BaseImage,omitempty"`
	// The type of the source image. Valid values:
	//
	// 	- IMAGE: custom image
	//
	// 	- IMAGE_FAMILY: image family
	//
	// example:
	//
	// IMAGE
	BaseImageType *string `json:"BaseImageType,omitempty" xml:"BaseImageType,omitempty"`
	// The content of the image template.
	//
	// example:
	//
	// FROM IMAGE:m-bp67acfmxazb4p****
	BuildContent *string `json:"BuildContent,omitempty" xml:"BuildContent,omitempty"`
	// The time when the image template was created.
	//
	// example:
	//
	// 2020-11-24T06:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether to release the intermediate instance when the image fails to be created.
	//
	// example:
	//
	// true
	DeleteInstanceOnFailure *bool `json:"DeleteInstanceOnFailure,omitempty" xml:"DeleteInstanceOnFailure,omitempty"`
	// The description of the image template.
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
	ImageOptions *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The ID of the image template.
	//
	// example:
	//
	// ip-2ze5tsl5bp6nf2b3****
	ImagePipelineId *string `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty"`
	// The attributes and settings of the imported image.
	ImportImageOptions *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions `json:"ImportImageOptions,omitempty" xml:"ImportImageOptions,omitempty" type:"Struct"`
	// The instance type.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The size of the outbound public bandwidth for the intermediate instance. Unit: Mbit/s.
	//
	// example:
	//
	// 0
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The name of the image template.
	//
	// example:
	//
	// testImagePipeline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	//
	// Indicates whether the image created based on the image template supports the Non-Volatile Memory Express (NVMe) protocol.
	//
	// >  This parameter is no longer used. We recommend that you use ImageOptions.ImageFeatures.NvmeSupport.
	//
	// example:
	//
	// auto
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
	// The repair mode of the image template.
	//
	// Valid values:
	//
	// 	- Standard: the standard mode
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
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The system disk size of the intermediate instance. Unit: GiB
	//
	// example:
	//
	// 40
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags of the image template.
	Tags *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The content of the image test template.
	//
	// example:
	//
	// null
	TestContent *string `json:"TestContent,omitempty" xml:"TestContent,omitempty"`
	// The IDs of regions to which to distribute the image that will be created based on the image template.
	ToRegionIds *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds `json:"ToRegionIds,omitempty" xml:"ToRegionIds,omitempty" type:"Struct"`
	// The ID of the vSwitch in the virtual private cloud (VPC).
	//
	// example:
	//
	// vsw-bp67acfmxazb4p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetAddAccounts() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts {
	return s.AddAccounts
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetAdvancedOptions() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions {
	return s.AdvancedOptions
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetBaseImage() *string {
	return s.BaseImage
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetBaseImageType() *string {
	return s.BaseImageType
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetBuildContent() *string {
	return s.BuildContent
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetDeleteInstanceOnFailure() *bool {
	return s.DeleteInstanceOnFailure
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetDescription() *string {
	return s.Description
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetImageOptions() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions {
	return s.ImageOptions
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetImagePipelineId() *string {
	return s.ImagePipelineId
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetImportImageOptions() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	return s.ImportImageOptions
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetName() *string {
	return s.Name
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetRepairMode() *string {
	return s.RepairMode
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetTags() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags {
	return s.Tags
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetTestContent() *string {
	return s.TestContent
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetToRegionIds() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds {
	return s.ToRegionIds
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetAddAccounts(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.AddAccounts = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetAdvancedOptions(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.AdvancedOptions = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetBaseImage(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.BaseImage = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetBaseImageType(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.BaseImageType = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetBuildContent(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.BuildContent = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetCreationTime(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetDeleteInstanceOnFailure(v bool) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.DeleteInstanceOnFailure = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetDescription(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.Description = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetImageFamily(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.ImageFamily = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetImageName(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.ImageName = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetImageOptions(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.ImageOptions = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetImagePipelineId(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.ImagePipelineId = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetImportImageOptions(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.ImportImageOptions = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetInstanceType(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.InstanceType = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetInternetMaxBandwidthOut(v int32) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetName(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.Name = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetNvmeSupport(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.NvmeSupport = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetRepairMode(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.RepairMode = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetResourceGroupId(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetSystemDiskSize(v int32) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetTags(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.Tags = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetTestContent(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.TestContent = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetToRegionIds(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.ToRegionIds = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) SetVSwitchId(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet {
	s.VSwitchId = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSet) Validate() error {
	if s.AddAccounts != nil {
		if err := s.AddAccounts.Validate(); err != nil {
			return err
		}
	}
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
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.ToRegionIds != nil {
		if err := s.ToRegionIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts struct {
	AddAccount []*string `json:"AddAccount,omitempty" xml:"AddAccount,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts) GetAddAccount() []*string {
	return s.AddAccount
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts) SetAddAccount(v []*string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts {
	s.AddAccount = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions struct {
	// Indicates whether to disable the feature that automatically adds a suffix to the name of the image created based on the image template. Valid value:
	//
	// 	- disable
	//
	// example:
	//
	// disable
	ImageNameSuffix *string `json:"ImageNameSuffix,omitempty" xml:"ImageNameSuffix,omitempty"`
	// Indicates whether to retain Cloud Assistant. During the image building process, the system automatically installs Cloud Assistant in the intermediate instance to run commands. You can choose whether to retain Cloud Assistant in the new image created based on the image template. Valid values:
	//
	// 	- true: retains Cloud Assistant.
	//
	// 	- false: does not retain Cloud Assistant.
	//
	// >  This parameter does not affect Cloud Assistant that comes with your image.
	//
	// example:
	//
	// true
	RetainCloudAssistant *bool `json:"RetainCloudAssistant,omitempty" xml:"RetainCloudAssistant,omitempty"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) GetImageNameSuffix() *string {
	return s.ImageNameSuffix
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) GetRetainCloudAssistant() *bool {
	return s.RetainCloudAssistant
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) SetImageNameSuffix(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions {
	s.ImageNameSuffix = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) SetRetainCloudAssistant(v bool) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions {
	s.RetainCloudAssistant = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions struct {
	// The description of the image.
	//
	// example:
	//
	// description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image family.
	//
	// example:
	//
	// family
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The feature attributes of the image.
	ImageFeatures *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures `json:"ImageFeatures,omitempty" xml:"ImageFeatures,omitempty" type:"Struct"`
	// The prefix of the image name.
	//
	// example:
	//
	// imageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The tags of the image.
	ImageTags *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags `json:"ImageTags,omitempty" xml:"ImageTags,omitempty" type:"Struct"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) GetDescription() *string {
	return s.Description
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) GetImageFeatures() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures {
	return s.ImageFeatures
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) GetImageTags() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags {
	return s.ImageTags
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) SetDescription(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions {
	s.Description = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) SetImageFamily(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions {
	s.ImageFamily = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) SetImageFeatures(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions {
	s.ImageFeatures = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) SetImageName(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions {
	s.ImageName = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) SetImageTags(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions {
	s.ImageTags = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions) Validate() error {
	if s.ImageFeatures != nil {
		if err := s.ImageFeatures.Validate(); err != nil {
			return err
		}
	}
	if s.ImageTags != nil {
		if err := s.ImageTags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures struct {
	// Indicates whether the image supports the NVMe protocol. Valid values:
	//
	// 	- supported: The image supports the NVMe protocol. Instances created from the image also support the NVMe protocol.
	//
	// 	- unsupported: The image does not support the NVMe protocol. Instances created from the image do not support the NVMe protocol.
	//
	// 	- auto: The system automatically checks whether the image supports the NVMe protocol. The system automatically checks whether the NVMe driver is installed on your image before the image is built. If you install or uninstall the NVMe driver during the image building task, the check result may be incorrect. We recommend that you set the value to supported or unsupported based on the image building content.
	//
	// example:
	//
	// auto
	NvmeSupport *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures) SetNvmeSupport(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags struct {
	ImageTag []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag `json:"ImageTag,omitempty" xml:"ImageTag,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags) GetImageTag() []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag {
	return s.ImageTag
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags) SetImageTag(v []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags {
	s.ImageTag = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags) Validate() error {
	if s.ImageTag != nil {
		for _, item := range s.ImageTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag struct {
	// The tag key of the image.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the image.
	//
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) SetTagKey(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag {
	s.TagKey = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) SetTagValue(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag {
	s.TagValue = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTagsImageTag) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions struct {
	// The operating system architecture. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
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
	// example:
	//
	// BIOS
	BootMode    *string `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The information of disks from which the custom images are created.
	//
	// 	- When the value of N is 1, a custom image is created from the system disk.
	//
	// 	- When the value of N is an integer in the range of 2 to 17, a custom image is created from a data disk.
	DiskDeviceMappings *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	// The attributes of the custom image.
	Features        *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures        `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	ImageName       *string                                                                                           `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImportImageTags *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags `json:"ImportImageTags,omitempty" xml:"ImportImageTags,omitempty" type:"Struct"`
	// The type of the license to use to activate the operating system after the image is imported. Valid values:
	//
	// 	- Auto: ECS detects the operating system of the image and allocates a license to the operating system In this mode, the system first checks whether a license allocated by an official Alibaba Cloud channel is specified in the `Platform`. If a license allocated by an official Alibaba Cloud channel is specified, the system allocates the license to the imported image. If no such license is specified, the Bring Your Own License (BYOL) mode is used.
	//
	// 	- Aliyun: The license allocated through an official Alibaba Cloud channel is used for the operating system distribution specified by `Platform`.
	//
	// 	- BYOL: The license that comes with the source operating system is used. When you use the BYOL license, make sure that your license key is supported by Alibaba Cloud.
	//
	// example:
	//
	// Auto
	LicenseType *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	// The operating system type. Valid values:
	//
	// 	- windows: Windows operating systems
	//
	// 	- linux: Linux operating systems
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

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetBootMode() *string {
	return s.BootMode
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetDescription() *string {
	return s.Description
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetDiskDeviceMappings() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetFeatures() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures {
	return s.Features
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetImportImageTags() *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags {
	return s.ImportImageTags
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetLicenseType() *string {
	return s.LicenseType
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetOSType() *string {
	return s.OSType
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetRetainImportedImage() *bool {
	return s.RetainImportedImage
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetRetentionStrategy() *string {
	return s.RetentionStrategy
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetArchitecture(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.Architecture = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetBootMode(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.BootMode = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetDescription(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.Description = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetDiskDeviceMappings(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.DiskDeviceMappings = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetFeatures(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.Features = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetImageName(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.ImageName = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetImportImageTags(v *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.ImportImageTags = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetLicenseType(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.LicenseType = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetOSType(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.OSType = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetPlatform(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.Platform = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetRetainImportedImage(v bool) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.RetainImportedImage = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetRetentionStrategy(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.RetentionStrategy = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) SetRoleName(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions {
	s.RoleName = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions) Validate() error {
	if s.DiskDeviceMappings != nil {
		if err := s.DiskDeviceMappings.Validate(); err != nil {
			return err
		}
	}
	if s.Features != nil {
		if err := s.Features.Validate(); err != nil {
			return err
		}
	}
	if s.ImportImageTags != nil {
		if err := s.ImportImageTags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings struct {
	DiskDeviceMapping []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings) GetDiskDeviceMapping() []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings) SetDiskDeviceMapping(v []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings {
	s.DiskDeviceMapping = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings) Validate() error {
	if s.DiskDeviceMapping != nil {
		for _, item := range s.DiskDeviceMapping {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping struct {
	// The size of disk N in the custom image after the image is imported.
	//
	// You can use this parameter to specify the sizes of the system disk and data disks in the custom image. When you specify the size of the system disk, make sure that the specified size is greater than or equal to the size of the source image file. Unit: GiB. Valid values:
	//
	// 	- When N is set to 1, this parameter indicates the size of the system disk in the custom image. Valid values: 1 to 2048.
	//
	// 	- When N is set to an integer in the range of 2 to 17, this parameter indicates the size of a data disk in the custom image. Valid values: 1 to 2048.
	//
	// After the image file is uploaded to an OSS bucket, you can view the size of the image file in the OSS bucket.
	//
	// example:
	//
	// 40
	DiskImageSize *int32 `json:"DiskImageSize,omitempty" xml:"DiskImageSize,omitempty"`
	// The format of the image. Valid values:
	//
	// 	- RAW
	//
	// 	- VHD
	//
	// 	- QCOW2
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

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) GetDiskImageSize() *int32 {
	return s.DiskImageSize
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) GetOSSObject() *string {
	return s.OSSObject
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) SetDiskImageSize(v int32) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping {
	s.DiskImageSize = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) SetFormat(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) SetOSSBucket(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping {
	s.OSSBucket = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) SetOSSObject(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping {
	s.OSSObject = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappingsDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures struct {
	ImdsSupport *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	// Indicates whether the image supports the NVMe protocol. Valid values:
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

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) SetImdsSupport(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) SetNvmeSupport(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags struct {
	ImportImageTag []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag `json:"ImportImageTag,omitempty" xml:"ImportImageTag,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags) GetImportImageTag() []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag {
	return s.ImportImageTag
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags) SetImportImageTag(v []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags {
	s.ImportImageTag = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags) Validate() error {
	if s.ImportImageTag != nil {
		for _, item := range s.ImportImageTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) GetKey() *string {
	return s.Key
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) GetValue() *string {
	return s.Value
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) SetKey(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag {
	s.Key = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) SetValue(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag {
	s.Value = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTagsImportImageTag) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags struct {
	Tag []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags) GetTag() []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag {
	return s.Tag
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags) SetTag(v []*DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags {
	s.Tag = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags) Validate() error {
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

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) SetTagKey(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) SetTagValue(v string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds struct {
	ToRegionId []*string `json:"ToRegionId,omitempty" xml:"ToRegionId,omitempty" type:"Repeated"`
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds) GetToRegionId() []*string {
	return s.ToRegionId
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds) SetToRegionId(v []*string) *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds {
	s.ToRegionId = v
	return s
}

func (s *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds) Validate() error {
	return dara.Validate(s)
}
