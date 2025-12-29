// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFromFamilyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *DescribeImageFromFamilyResponseBodyImage) *DescribeImageFromFamilyResponseBody
	GetImage() *DescribeImageFromFamilyResponseBodyImage
	SetRequestId(v string) *DescribeImageFromFamilyResponseBody
	GetRequestId() *string
}

type DescribeImageFromFamilyResponseBody struct {
	// The image information.
	Image *DescribeImageFromFamilyResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageFromFamilyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyResponseBody) GetImage() *DescribeImageFromFamilyResponseBodyImage {
	return s.Image
}

func (s *DescribeImageFromFamilyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageFromFamilyResponseBody) SetImage(v *DescribeImageFromFamilyResponseBodyImage) *DescribeImageFromFamilyResponseBody {
	s.Image = v
	return s
}

func (s *DescribeImageFromFamilyResponseBody) SetRequestId(v string) *DescribeImageFromFamilyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBody) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageFromFamilyResponseBodyImage struct {
	// The architecture of the image. Valid values:
	//
	// 	- i386
	//
	// 	- x86_64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 2018-01-10T01:01:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the volume.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mappings between the disk and the snapshot in the image.
	DiskDeviceMappings *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings `json:"DiskDeviceMappings,omitempty" xml:"DiskDeviceMappings,omitempty" type:"Struct"`
	// The name of the image family.
	//
	// example:
	//
	// testImageFamily
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-bp1g7004ksh0oeuc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The alias of the image owner. Valid values:
	//
	// 	- system: public images provided by Alibaba Cloud
	//
	// 	- self: your custom images
	//
	// 	- others: shared images from other Alibaba Cloud accounts
	//
	// 	- marketplace: Alibaba Cloud Marketplace images
	//
	// example:
	//
	// self
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The image version.
	//
	// example:
	//
	// 2
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// Indicates whether the image is a copy of another image.
	//
	// example:
	//
	// false
	IsCopied *bool `json:"IsCopied,omitempty" xml:"IsCopied,omitempty"`
	// Indicates whether the custom image was shared to other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsSelfShared *string `json:"IsSelfShared,omitempty" xml:"IsSelfShared,omitempty"`
	// Indicates whether you have subscribed to the service terms of the image product corresponding to the image product code.
	//
	// example:
	//
	// false
	IsSubscribed *bool `json:"IsSubscribed,omitempty" xml:"IsSubscribed,omitempty"`
	// Indicates whether cloud-init is supported.
	//
	// example:
	//
	// true
	IsSupportCloudinit *bool `json:"IsSupportCloudinit,omitempty" xml:"IsSupportCloudinit,omitempty"`
	// Indicates whether the image can be used on I/O optimized instances.
	//
	// example:
	//
	// true
	IsSupportIoOptimized *bool `json:"IsSupportIoOptimized,omitempty" xml:"IsSupportIoOptimized,omitempty"`
	// The display name of the operating system in Chinese.
	//
	// example:
	//
	// Alibaba Cloud Linux 2.1903
	OSName *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- windows
	//
	// 	- linux
	//
	// example:
	//
	// linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The operating system.
	//
	// example:
	//
	// Aliyun
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The product code of the Alibaba Cloud Marketplace image.
	//
	// example:
	//
	// jxsc00****
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The image creation progress in percentage.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The size of the image. Unit: GiB.
	//
	// example:
	//
	// 80
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The state of the image. Valid values:
	//
	// 	- UnAvailable
	//
	// 	- Available
	//
	// 	- Creating
	//
	// 	- CreateFailed
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the image.
	Tags *DescribeImageFromFamilyResponseBodyImageTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// Indicates whether the image has been used to create ECS instances. Valid values:
	//
	// 	- instance: The image was used to create one or more ECS instances.
	//
	// 	- none: The image was not used to create ECS instances.
	//
	// example:
	//
	// none
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeImageFromFamilyResponseBodyImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyResponseBodyImage) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetDiskDeviceMappings() *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings {
	return s.DiskDeviceMappings
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetIsCopied() *bool {
	return s.IsCopied
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetIsSelfShared() *string {
	return s.IsSelfShared
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetIsSubscribed() *bool {
	return s.IsSubscribed
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetIsSupportCloudinit() *bool {
	return s.IsSupportCloudinit
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetIsSupportIoOptimized() *bool {
	return s.IsSupportIoOptimized
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetOSName() *string {
	return s.OSName
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetOSType() *string {
	return s.OSType
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetProgress() *string {
	return s.Progress
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetSize() *int32 {
	return s.Size
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetTags() *DescribeImageFromFamilyResponseBodyImageTags {
	return s.Tags
}

func (s *DescribeImageFromFamilyResponseBodyImage) GetUsage() *string {
	return s.Usage
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetArchitecture(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.Architecture = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetCreationTime(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.CreationTime = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetDescription(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.Description = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetDiskDeviceMappings(v *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings) *DescribeImageFromFamilyResponseBodyImage {
	s.DiskDeviceMappings = v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetImageFamily(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.ImageFamily = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetImageId(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetImageName(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.ImageName = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetImageOwnerAlias(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetImageVersion(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.ImageVersion = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetIsCopied(v bool) *DescribeImageFromFamilyResponseBodyImage {
	s.IsCopied = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetIsSelfShared(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.IsSelfShared = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetIsSubscribed(v bool) *DescribeImageFromFamilyResponseBodyImage {
	s.IsSubscribed = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetIsSupportCloudinit(v bool) *DescribeImageFromFamilyResponseBodyImage {
	s.IsSupportCloudinit = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetIsSupportIoOptimized(v bool) *DescribeImageFromFamilyResponseBodyImage {
	s.IsSupportIoOptimized = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetOSName(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.OSName = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetOSType(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.OSType = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetPlatform(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.Platform = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetProductCode(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.ProductCode = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetProgress(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.Progress = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetSize(v int32) *DescribeImageFromFamilyResponseBodyImage {
	s.Size = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetStatus(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.Status = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetTags(v *DescribeImageFromFamilyResponseBodyImageTags) *DescribeImageFromFamilyResponseBodyImage {
	s.Tags = v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) SetUsage(v string) *DescribeImageFromFamilyResponseBodyImage {
	s.Usage = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImage) Validate() error {
	if s.DiskDeviceMappings != nil {
		if err := s.DiskDeviceMappings.Validate(); err != nil {
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

type DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings struct {
	DiskDeviceMapping []*DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping `json:"DiskDeviceMapping,omitempty" xml:"DiskDeviceMapping,omitempty" type:"Repeated"`
}

func (s DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings) GetDiskDeviceMapping() []*DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	return s.DiskDeviceMapping
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings) SetDiskDeviceMapping(v []*DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings {
	s.DiskDeviceMapping = v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappings) Validate() error {
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

type DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping struct {
	// The device name of the disk. Example: /dev/xvdb.
	//
	// >  This parameter will be removed in the future. To ensure compatibility, we recommend that you use other parameters.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The image format.
	//
	// example:
	//
	// qcow2
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The OSS bucket that contains the imported image file.
	//
	// example:
	//
	// testEcsImport
	ImportOSSBucket *string `json:"ImportOSSBucket,omitempty" xml:"ImportOSSBucket,omitempty"`
	// The OSS object to which the imported image belongs.
	//
	// example:
	//
	// imageImport
	ImportOSSObject *string `json:"ImportOSSObject,omitempty" xml:"ImportOSSObject,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 80
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-bp17ot2q7x72ggtw****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The image type.
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GetDevice() *string {
	return s.Device
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GetFormat() *string {
	return s.Format
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GetImportOSSBucket() *string {
	return s.ImportOSSBucket
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GetImportOSSObject() *string {
	return s.ImportOSSObject
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GetSize() *string {
	return s.Size
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) GetType() *string {
	return s.Type
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) SetDevice(v string) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	s.Device = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) SetFormat(v string) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	s.Format = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) SetImportOSSBucket(v string) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	s.ImportOSSBucket = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) SetImportOSSObject(v string) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	s.ImportOSSObject = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) SetSize(v string) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	s.Size = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) SetSnapshotId(v string) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) SetType(v string) *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping {
	s.Type = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageDiskDeviceMappingsDiskDeviceMapping) Validate() error {
	return dara.Validate(s)
}

type DescribeImageFromFamilyResponseBodyImageTags struct {
	Tag []*DescribeImageFromFamilyResponseBodyImageTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImageFromFamilyResponseBodyImageTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyResponseBodyImageTags) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyResponseBodyImageTags) GetTag() []*DescribeImageFromFamilyResponseBodyImageTagsTag {
	return s.Tag
}

func (s *DescribeImageFromFamilyResponseBodyImageTags) SetTag(v []*DescribeImageFromFamilyResponseBodyImageTagsTag) *DescribeImageFromFamilyResponseBodyImageTags {
	s.Tag = v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageTags) Validate() error {
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

type DescribeImageFromFamilyResponseBodyImageTagsTag struct {
	// The tag key of the custom image.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the custom image.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeImageFromFamilyResponseBodyImageTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyResponseBodyImageTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyResponseBodyImageTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeImageFromFamilyResponseBodyImageTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeImageFromFamilyResponseBodyImageTagsTag) SetTagKey(v string) *DescribeImageFromFamilyResponseBodyImageTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageTagsTag) SetTagValue(v string) *DescribeImageFromFamilyResponseBodyImageTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeImageFromFamilyResponseBodyImageTagsTag) Validate() error {
	return dara.Validate(s)
}
