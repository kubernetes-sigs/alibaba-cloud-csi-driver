// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLaunchTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateSets(v *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets) *DescribeLaunchTemplatesResponseBody
	GetLaunchTemplateSets() *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets
	SetPageNumber(v int32) *DescribeLaunchTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLaunchTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLaunchTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLaunchTemplatesResponseBody
	GetTotalCount() *int32
}

type DescribeLaunchTemplatesResponseBody struct {
	// The queried launch templates.
	LaunchTemplateSets *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets `json:"LaunchTemplateSets,omitempty" xml:"LaunchTemplateSets,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE12CBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of launch templates.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLaunchTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesResponseBody) GetLaunchTemplateSets() *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets {
	return s.LaunchTemplateSets
}

func (s *DescribeLaunchTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLaunchTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLaunchTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLaunchTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLaunchTemplatesResponseBody) SetLaunchTemplateSets(v *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets) *DescribeLaunchTemplatesResponseBody {
	s.LaunchTemplateSets = v
	return s
}

func (s *DescribeLaunchTemplatesResponseBody) SetPageNumber(v int32) *DescribeLaunchTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBody) SetPageSize(v int32) *DescribeLaunchTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBody) SetRequestId(v string) *DescribeLaunchTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBody) SetTotalCount(v int32) *DescribeLaunchTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBody) Validate() error {
	if s.LaunchTemplateSets != nil {
		if err := s.LaunchTemplateSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLaunchTemplatesResponseBodyLaunchTemplateSets struct {
	LaunchTemplateSet []*DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet `json:"LaunchTemplateSet,omitempty" xml:"LaunchTemplateSet,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSets) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets) GetLaunchTemplateSet() []*DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	return s.LaunchTemplateSet
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets) SetLaunchTemplateSet(v []*DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets {
	s.LaunchTemplateSet = v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSets) Validate() error {
	if s.LaunchTemplateSet != nil {
		for _, item := range s.LaunchTemplateSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet struct {
	// The time when the launch template was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-05-14T14:18:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Alibaba Cloud account that created the launch template.
	//
	// example:
	//
	// 1234567890
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The default version number of the launch template.
	//
	// example:
	//
	// 1
	DefaultVersionNumber *int64 `json:"DefaultVersionNumber,omitempty" xml:"DefaultVersionNumber,omitempty"`
	// The latest version number of the launch template.
	//
	// example:
	//
	// 1
	LatestVersionNumber *int64 `json:"LatestVersionNumber,omitempty" xml:"LatestVersionNumber,omitempty"`
	// The ID of the launch template.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7q****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the launch template.
	//
	// example:
	//
	// wd-152630748****
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The time when a version was added to or deleted from the launch template.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-05-14T14:18:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the resource group to which the launch template belongs.
	//
	// example:
	//
	// rg-acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the launch template.
	//
	// >  You can only call API operations to add tags to and query the tags of a launch template. You cannot add tags to or view the tags of a launch template in the ECS console.
	Tags *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetDefaultVersionNumber() *int64 {
	return s.DefaultVersionNumber
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetLatestVersionNumber() *int64 {
	return s.LatestVersionNumber
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) GetTags() *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags {
	return s.Tags
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetCreateTime(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.CreateTime = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetCreatedBy(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.CreatedBy = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetDefaultVersionNumber(v int64) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.DefaultVersionNumber = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetLatestVersionNumber(v int64) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.LatestVersionNumber = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetLaunchTemplateId(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetLaunchTemplateName(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.LaunchTemplateName = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetModifiedTime(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetResourceGroupId(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) SetTags(v *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet {
	s.Tags = v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSet) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags struct {
	Tag []*DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags) GetTag() []*DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag {
	return s.Tag
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags) SetTag(v []*DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags {
	s.Tag = v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTags) Validate() error {
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

type DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag struct {
	// The tag value of the launch template.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag key of the launch template.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) SetTagKey(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) SetTagValue(v string) *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeLaunchTemplatesResponseBodyLaunchTemplateSetsLaunchTemplateSetTagsTag) Validate() error {
	return dara.Validate(s)
}
