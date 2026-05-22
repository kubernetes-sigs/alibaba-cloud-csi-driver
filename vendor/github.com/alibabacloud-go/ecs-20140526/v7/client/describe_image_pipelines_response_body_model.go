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
	AddAccounts             *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAddAccounts     `json:"AddAccounts,omitempty" xml:"AddAccounts,omitempty" type:"Struct"`
	AdvancedOptions         *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetAdvancedOptions `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty" type:"Struct"`
	BaseImage               *string                                                                         `json:"BaseImage,omitempty" xml:"BaseImage,omitempty"`
	BaseImageType           *string                                                                         `json:"BaseImageType,omitempty" xml:"BaseImageType,omitempty"`
	BuildContent            *string                                                                         `json:"BuildContent,omitempty" xml:"BuildContent,omitempty"`
	CreationTime            *string                                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeleteInstanceOnFailure *bool                                                                           `json:"DeleteInstanceOnFailure,omitempty" xml:"DeleteInstanceOnFailure,omitempty"`
	Description             *string                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// Deprecated
	ImageName               *string                                                                            `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOptions            *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptions       `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	ImagePipelineId         *string                                                                            `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty"`
	ImportImageOptions      *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptions `json:"ImportImageOptions,omitempty" xml:"ImportImageOptions,omitempty" type:"Struct"`
	InstanceType            *string                                                                            `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetMaxBandwidthOut *int32                                                                             `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	Name                    *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	// Deprecated
	NvmeSupport     *string                                                                     `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
	RepairMode      *string                                                                     `json:"RepairMode,omitempty" xml:"RepairMode,omitempty"`
	ResourceGroupId *string                                                                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SystemDiskSize  *int32                                                                      `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	Tags            *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetTags        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TestContent     *string                                                                     `json:"TestContent,omitempty" xml:"TestContent,omitempty"`
	ToRegionIds     *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetToRegionIds `json:"ToRegionIds,omitempty" xml:"ToRegionIds,omitempty" type:"Struct"`
	VSwitchId       *string                                                                     `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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
	ImageNameSuffix      *string `json:"ImageNameSuffix,omitempty" xml:"ImageNameSuffix,omitempty"`
	RetainCloudAssistant *bool   `json:"RetainCloudAssistant,omitempty" xml:"RetainCloudAssistant,omitempty"`
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
	Description   *string                                                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageFamily   *string                                                                                   `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ImageFeatures *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageFeatures `json:"ImageFeatures,omitempty" xml:"ImageFeatures,omitempty" type:"Struct"`
	ImageName     *string                                                                                   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageTags     *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImageOptionsImageTags     `json:"ImageTags,omitempty" xml:"ImageTags,omitempty" type:"Struct"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Architecture        *string                                                                                              `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	BootMode            *string                                                                                              `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	Description         *string                                                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskDeviceMappings  *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	Features            *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsFeatures           `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	ImageName           *string                                                                                              `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImportImageTags     *DescribeImagePipelinesResponseBodyImagePipelineImagePipelineSetImportImageOptionsImportImageTags    `json:"ImportImageTags,omitempty" xml:"ImportImageTags,omitempty" type:"Struct"`
	LicenseType         *string                                                                                              `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	OSType              *string                                                                                              `json:"OSType,omitempty" xml:"OSType,omitempty"`
	Platform            *string                                                                                              `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RetainImportedImage *bool                                                                                                `json:"RetainImportedImage,omitempty" xml:"RetainImportedImage,omitempty"`
	RetentionStrategy   *string                                                                                              `json:"RetentionStrategy,omitempty" xml:"RetentionStrategy,omitempty"`
	RoleName            *string                                                                                              `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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
	DiskImageSize *int32  `json:"DiskImageSize,omitempty" xml:"DiskImageSize,omitempty"`
	Format        *string `json:"Format,omitempty" xml:"Format,omitempty"`
	OSSBucket     *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	OSSObject     *string `json:"OSSObject,omitempty" xml:"OSSObject,omitempty"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
