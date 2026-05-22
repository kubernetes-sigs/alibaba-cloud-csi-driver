// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v *DescribeImagesResponseBodyImages) *DescribeImagesResponseBody
	GetImages() *DescribeImagesResponseBodyImages
	SetPageNumber(v int32) *DescribeImagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeImagesResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeImagesResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImagesResponseBody
	GetTotalCount() *int32
}

type DescribeImagesResponseBody struct {
	Images *DescribeImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 66189103-EDB2-43E2-BB60-BFF2B62F4EB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of images.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBody) GetImages() *DescribeImagesResponseBodyImages {
	return s.Images
}

func (s *DescribeImagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeImagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImagesResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImagesResponseBody) SetImages(v *DescribeImagesResponseBodyImages) *DescribeImagesResponseBody {
	s.Images = v
	return s
}

func (s *DescribeImagesResponseBody) SetPageNumber(v int32) *DescribeImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesResponseBody) SetPageSize(v int32) *DescribeImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRegionId(v string) *DescribeImagesResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeImagesResponseBody) SetRequestId(v string) *DescribeImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagesResponseBody) SetTotalCount(v int32) *DescribeImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImagesResponseBody) Validate() error {
	if s.Images != nil {
		if err := s.Images.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagesResponseBodyImages struct {
	Image []*DescribeImagesResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImages) GetImage() []*DescribeImagesResponseBodyImagesImage {
	return s.Image
}

func (s *DescribeImagesResponseBodyImages) SetImage(v []*DescribeImagesResponseBodyImagesImage) *DescribeImagesResponseBodyImages {
	s.Image = v
	return s
}

func (s *DescribeImagesResponseBodyImages) Validate() error {
	if s.Image != nil {
		for _, item := range s.Image {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagesResponseBodyImagesImage struct {
	Architecture            *string                                                  `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	BootMode                *string                                                  `json:"BootMode,omitempty" xml:"BootMode,omitempty"`
	CreationTime            *string                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description             *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DetectionOptions        *DescribeImagesResponseBodyImagesImageDetectionOptions   `json:"DetectionOptions,omitempty" xml:"DetectionOptions,omitempty" type:"Struct"`
	DiskDeviceMappings      *DescribeImagesResponseBodyImagesImageDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	Features                *DescribeImagesResponseBodyImagesImageFeatures           `json:"Features,omitempty" xml:"Features,omitempty" type:"Struct"`
	ImageFamily             *string                                                  `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	ImageId                 *string                                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName               *string                                                  `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageOwnerAlias         *string                                                  `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageOwnerId            *int64                                                   `json:"ImageOwnerId,omitempty" xml:"ImageOwnerId,omitempty"`
	ImageVersion            *string                                                  `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	IsCopied                *bool                                                    `json:"IsCopied,omitempty" xml:"IsCopied,omitempty"`
	IsPublic                *bool                                                    `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	IsSelfShared            *string                                                  `json:"IsSelfShared,omitempty" xml:"IsSelfShared,omitempty"`
	IsSubscribed            *bool                                                    `json:"IsSubscribed,omitempty" xml:"IsSubscribed,omitempty"`
	IsSupportCloudinit      *bool                                                    `json:"IsSupportCloudinit,omitempty" xml:"IsSupportCloudinit,omitempty"`
	IsSupportIoOptimized    *bool                                                    `json:"IsSupportIoOptimized,omitempty" xml:"IsSupportIoOptimized,omitempty"`
	LicenseType             *string                                                  `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	LoginAsNonRootSupported *bool                                                    `json:"LoginAsNonRootSupported,omitempty" xml:"LoginAsNonRootSupported,omitempty"`
	OSName                  *string                                                  `json:"OSName,omitempty" xml:"OSName,omitempty"`
	OSNameEn                *string                                                  `json:"OSNameEn,omitempty" xml:"OSNameEn,omitempty"`
	OSType                  *string                                                  `json:"OSType,omitempty" xml:"OSType,omitempty"`
	Platform                *string                                                  `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductCode             *string                                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Progress                *string                                                  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ResourceGroupId         *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Size                    *int32                                                   `json:"Size,omitempty" xml:"Size,omitempty"`
	Status                  *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName            *string                                                  `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	Tags                    *DescribeImagesResponseBodyImagesImageTags               `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Usage                   *string                                                  `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImage) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeImagesResponseBodyImagesImage) GetBootMode() *string {
	return s.BootMode
}

func (s *DescribeImagesResponseBodyImagesImage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeImagesResponseBodyImagesImage) GetDescription() *string {
	return s.Description
}

func (s *DescribeImagesResponseBodyImagesImage) GetDetectionOptions() *DescribeImagesResponseBodyImagesImageDetectionOptions {
	return s.DetectionOptions
}

func (s *DescribeImagesResponseBodyImagesImage) GetDiskDeviceMappings() *DescribeImagesResponseBodyImagesImageDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *DescribeImagesResponseBodyImagesImage) GetFeatures() *DescribeImagesResponseBodyImagesImageFeatures {
	return s.Features
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageOwnerId() *int64 {
	return s.ImageOwnerId
}

func (s *DescribeImagesResponseBodyImagesImage) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *DescribeImagesResponseBodyImagesImage) GetIsCopied() *bool {
	return s.IsCopied
}

func (s *DescribeImagesResponseBodyImagesImage) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *DescribeImagesResponseBodyImagesImage) GetIsSelfShared() *string {
	return s.IsSelfShared
}

func (s *DescribeImagesResponseBodyImagesImage) GetIsSubscribed() *bool {
	return s.IsSubscribed
}

func (s *DescribeImagesResponseBodyImagesImage) GetIsSupportCloudinit() *bool {
	return s.IsSupportCloudinit
}

func (s *DescribeImagesResponseBodyImagesImage) GetIsSupportIoOptimized() *bool {
	return s.IsSupportIoOptimized
}

func (s *DescribeImagesResponseBodyImagesImage) GetLicenseType() *string {
	return s.LicenseType
}

func (s *DescribeImagesResponseBodyImagesImage) GetLoginAsNonRootSupported() *bool {
	return s.LoginAsNonRootSupported
}

func (s *DescribeImagesResponseBodyImagesImage) GetOSName() *string {
	return s.OSName
}

func (s *DescribeImagesResponseBodyImagesImage) GetOSNameEn() *string {
	return s.OSNameEn
}

func (s *DescribeImagesResponseBodyImagesImage) GetOSType() *string {
	return s.OSType
}

func (s *DescribeImagesResponseBodyImagesImage) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeImagesResponseBodyImagesImage) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeImagesResponseBodyImagesImage) GetProgress() *string {
	return s.Progress
}

func (s *DescribeImagesResponseBodyImagesImage) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeImagesResponseBodyImagesImage) GetSize() *int32 {
	return s.Size
}

func (s *DescribeImagesResponseBodyImagesImage) GetStatus() *string {
	return s.Status
}

func (s *DescribeImagesResponseBodyImagesImage) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeImagesResponseBodyImagesImage) GetTags() *DescribeImagesResponseBodyImagesImageTags {
	return s.Tags
}

func (s *DescribeImagesResponseBodyImagesImage) GetUsage() *string {
	return s.Usage
}

func (s *DescribeImagesResponseBodyImagesImage) SetArchitecture(v string) *DescribeImagesResponseBodyImagesImage {
	s.Architecture = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetBootMode(v string) *DescribeImagesResponseBodyImagesImage {
	s.BootMode = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetCreationTime(v string) *DescribeImagesResponseBodyImagesImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetDescription(v string) *DescribeImagesResponseBodyImagesImage {
	s.Description = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetDetectionOptions(v *DescribeImagesResponseBodyImagesImageDetectionOptions) *DescribeImagesResponseBodyImagesImage {
	s.DetectionOptions = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetDiskDeviceMappings(v *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) *DescribeImagesResponseBodyImagesImage {
	s.DiskDeviceMappings = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetFeatures(v *DescribeImagesResponseBodyImagesImageFeatures) *DescribeImagesResponseBodyImagesImage {
	s.Features = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageFamily(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageFamily = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageId(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageName(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageOwnerAlias(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageOwnerId(v int64) *DescribeImagesResponseBodyImagesImage {
	s.ImageOwnerId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetImageVersion(v string) *DescribeImagesResponseBodyImagesImage {
	s.ImageVersion = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetIsCopied(v bool) *DescribeImagesResponseBodyImagesImage {
	s.IsCopied = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetIsPublic(v bool) *DescribeImagesResponseBodyImagesImage {
	s.IsPublic = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetIsSelfShared(v string) *DescribeImagesResponseBodyImagesImage {
	s.IsSelfShared = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetIsSubscribed(v bool) *DescribeImagesResponseBodyImagesImage {
	s.IsSubscribed = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetIsSupportCloudinit(v bool) *DescribeImagesResponseBodyImagesImage {
	s.IsSupportCloudinit = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetIsSupportIoOptimized(v bool) *DescribeImagesResponseBodyImagesImage {
	s.IsSupportIoOptimized = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetLicenseType(v string) *DescribeImagesResponseBodyImagesImage {
	s.LicenseType = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetLoginAsNonRootSupported(v bool) *DescribeImagesResponseBodyImagesImage {
	s.LoginAsNonRootSupported = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetOSName(v string) *DescribeImagesResponseBodyImagesImage {
	s.OSName = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetOSNameEn(v string) *DescribeImagesResponseBodyImagesImage {
	s.OSNameEn = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetOSType(v string) *DescribeImagesResponseBodyImagesImage {
	s.OSType = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetPlatform(v string) *DescribeImagesResponseBodyImagesImage {
	s.Platform = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetProductCode(v string) *DescribeImagesResponseBodyImagesImage {
	s.ProductCode = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetProgress(v string) *DescribeImagesResponseBodyImagesImage {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetResourceGroupId(v string) *DescribeImagesResponseBodyImagesImage {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetSize(v int32) *DescribeImagesResponseBodyImagesImage {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetStatus(v string) *DescribeImagesResponseBodyImagesImage {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetSupplierName(v string) *DescribeImagesResponseBodyImagesImage {
	s.SupplierName = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetTags(v *DescribeImagesResponseBodyImagesImageTags) *DescribeImagesResponseBodyImagesImage {
	s.Tags = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) SetUsage(v string) *DescribeImagesResponseBodyImagesImage {
	s.Usage = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImage) Validate() error {
	if s.DetectionOptions != nil {
		if err := s.DetectionOptions.Validate(); err != nil {
			return err
		}
	}
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
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagesResponseBodyImagesImageDetectionOptions struct {
	Items  *DescribeImagesResponseBodyImagesImageDetectionOptionsItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Status *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImageDetectionOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageDetectionOptions) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptions) GetItems() *DescribeImagesResponseBodyImagesImageDetectionOptionsItems {
	return s.Items
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptions) GetStatus() *string {
	return s.Status
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptions) SetItems(v *DescribeImagesResponseBodyImagesImageDetectionOptionsItems) *DescribeImagesResponseBodyImagesImageDetectionOptions {
	s.Items = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptions) SetStatus(v string) *DescribeImagesResponseBodyImagesImageDetectionOptions {
	s.Status = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptions) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImagesResponseBodyImagesImageDetectionOptionsItems struct {
	Item []*DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImagesImageDetectionOptionsItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageDetectionOptionsItems) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItems) GetItem() []*DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem {
	return s.Item
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItems) SetItem(v []*DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) *DescribeImagesResponseBodyImagesImageDetectionOptionsItems {
	s.Item = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RiskCode  *string `json:"RiskCode,omitempty" xml:"RiskCode,omitempty"`
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) GetName() *string {
	return s.Name
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) GetRiskCode() *string {
	return s.RiskCode
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) GetValue() *string {
	return s.Value
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) SetName(v string) *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem {
	s.Name = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) SetRiskCode(v string) *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem {
	s.RiskCode = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) SetRiskLevel(v string) *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) SetValue(v string) *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem {
	s.Value = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDetectionOptionsItemsItem) Validate() error {
	return dara.Validate(s)
}

type DescribeImagesResponseBodyImagesImageDiskDeviceMappings struct {
	DiskDeviceMapping []*DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) GetDiskDeviceMapping() []*DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) SetDiskDeviceMapping(v []*DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) *DescribeImagesResponseBodyImagesImageDiskDeviceMappings {
	s.DiskDeviceMapping = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappings) Validate() error {
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

type DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping struct {
	Device          *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Encrypted       *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Format          *string `json:"Format,omitempty" xml:"Format,omitempty"`
	ImportOSSBucket *string `json:"ImportOSSBucket,omitempty" xml:"ImportOSSBucket,omitempty"`
	ImportOSSObject *string `json:"ImportOSSObject,omitempty" xml:"ImportOSSObject,omitempty"`
	Progress        *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RemainTime      *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	Size            *string `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId      *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetDevice() *string {
	return s.Device
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetImportOSSBucket() *string {
	return s.ImportOSSBucket
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetImportOSSObject() *string {
	return s.ImportOSSObject
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetProgress() *string {
	return s.Progress
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetRemainTime() *int32 {
	return s.RemainTime
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetSize() *string {
	return s.Size
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) GetType() *string {
	return s.Type
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetDevice(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Device = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetEncrypted(v bool) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Encrypted = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetFormat(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetImportOSSBucket(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.ImportOSSBucket = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetImportOSSObject(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.ImportOSSObject = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetProgress(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Progress = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetRemainTime(v int32) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.RemainTime = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetSize(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Size = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetSnapshotId(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) SetType(v string) *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping {
	s.Type = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageDiskDeviceMappingsDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}

type DescribeImagesResponseBodyImagesImageFeatures struct {
	CpuOnlineDowngrade    *string `json:"CpuOnlineDowngrade,omitempty" xml:"CpuOnlineDowngrade,omitempty"`
	CpuOnlineUpgrade      *string `json:"CpuOnlineUpgrade,omitempty" xml:"CpuOnlineUpgrade,omitempty"`
	ImdsSupport           *string `json:"ImdsSupport,omitempty" xml:"ImdsSupport,omitempty"`
	MemoryOnlineDowngrade *string `json:"MemoryOnlineDowngrade,omitempty" xml:"MemoryOnlineDowngrade,omitempty"`
	MemoryOnlineUpgrade   *string `json:"MemoryOnlineUpgrade,omitempty" xml:"MemoryOnlineUpgrade,omitempty"`
	NvmeSupport           *string `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImageFeatures) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageFeatures) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) GetCpuOnlineDowngrade() *string {
	return s.CpuOnlineDowngrade
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) GetCpuOnlineUpgrade() *string {
	return s.CpuOnlineUpgrade
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) GetImdsSupport() *string {
	return s.ImdsSupport
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) GetMemoryOnlineDowngrade() *string {
	return s.MemoryOnlineDowngrade
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) GetMemoryOnlineUpgrade() *string {
	return s.MemoryOnlineUpgrade
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) SetCpuOnlineDowngrade(v string) *DescribeImagesResponseBodyImagesImageFeatures {
	s.CpuOnlineDowngrade = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) SetCpuOnlineUpgrade(v string) *DescribeImagesResponseBodyImagesImageFeatures {
	s.CpuOnlineUpgrade = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) SetImdsSupport(v string) *DescribeImagesResponseBodyImagesImageFeatures {
	s.ImdsSupport = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) SetMemoryOnlineDowngrade(v string) *DescribeImagesResponseBodyImagesImageFeatures {
	s.MemoryOnlineDowngrade = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) SetMemoryOnlineUpgrade(v string) *DescribeImagesResponseBodyImagesImageFeatures {
	s.MemoryOnlineUpgrade = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) SetNvmeSupport(v string) *DescribeImagesResponseBodyImagesImageFeatures {
	s.NvmeSupport = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageFeatures) Validate() error {
	return dara.Validate(s)
}

type DescribeImagesResponseBodyImagesImageTags struct {
	Tag []*DescribeImagesResponseBodyImagesImageTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImagesResponseBodyImagesImageTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageTags) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageTags) GetTag() []*DescribeImagesResponseBodyImagesImageTagsTag {
	return s.Tag
}

func (s *DescribeImagesResponseBodyImagesImageTags) SetTag(v []*DescribeImagesResponseBodyImagesImageTagsTag) *DescribeImagesResponseBodyImagesImageTags {
	s.Tag = v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageTags) Validate() error {
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

type DescribeImagesResponseBodyImagesImageTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeImagesResponseBodyImagesImageTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponseBodyImagesImageTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponseBodyImagesImageTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeImagesResponseBodyImagesImageTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeImagesResponseBodyImagesImageTagsTag) SetTagKey(v string) *DescribeImagesResponseBodyImagesImageTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageTagsTag) SetTagValue(v string) *DescribeImagesResponseBodyImagesImageTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeImagesResponseBodyImagesImageTagsTag) Validate() error {
	return dara.Validate(s)
}
