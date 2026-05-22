// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyExResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicies(v *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPolicyExResponseBody
	GetAutoSnapshotPolicies() *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies
	SetPageNumber(v int32) *DescribeAutoSnapshotPolicyExResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoSnapshotPolicyExResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAutoSnapshotPolicyExResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAutoSnapshotPolicyExResponseBody
	GetTotalCount() *int32
}

type DescribeAutoSnapshotPolicyExResponseBody struct {
	AutoSnapshotPolicies *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies `json:"AutoSnapshotPolicies,omitempty" xml:"AutoSnapshotPolicies,omitempty" type:"Struct"`
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of automatic snapshot policies
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoSnapshotPolicyExResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) GetAutoSnapshotPolicies() *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies {
	return s.AutoSnapshotPolicies
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) SetAutoSnapshotPolicies(v *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPolicyExResponseBody {
	s.AutoSnapshotPolicies = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotPolicyExResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotPolicyExResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPolicyExResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotPolicyExResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBody) Validate() error {
	if s.AutoSnapshotPolicies != nil {
		if err := s.AutoSnapshotPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies struct {
	AutoSnapshotPolicy []*DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy `json:"AutoSnapshotPolicy,omitempty" xml:"AutoSnapshotPolicy,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies) GetAutoSnapshotPolicy() []*DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	return s.AutoSnapshotPolicy
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies) SetAutoSnapshotPolicy(v []*DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies {
	s.AutoSnapshotPolicy = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPolicies) Validate() error {
	if s.AutoSnapshotPolicy != nil {
		for _, item := range s.AutoSnapshotPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy struct {
	AutoSnapshotPolicyId         *string                                                                                                    `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	AutoSnapshotPolicyName       *string                                                                                                    `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	CopiedSnapshotsRetentionDays *int32                                                                                                     `json:"CopiedSnapshotsRetentionDays,omitempty" xml:"CopiedSnapshotsRetentionDays,omitempty"`
	CopyEncryptionConfiguration  *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration `json:"CopyEncryptionConfiguration,omitempty" xml:"CopyEncryptionConfiguration,omitempty" type:"Struct"`
	CreationTime                 *string                                                                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DiskNums                     *int32                                                                                                     `json:"DiskNums,omitempty" xml:"DiskNums,omitempty"`
	EnableCrossRegionCopy        *bool                                                                                                      `json:"EnableCrossRegionCopy,omitempty" xml:"EnableCrossRegionCopy,omitempty"`
	RegionId                     *string                                                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepeatWeekdays               *string                                                                                                    `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	ResourceGroupId              *string                                                                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RetentionDays                *int32                                                                                                     `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	Status                       *string                                                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                         *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags                        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TargetCopyRegions            *string                                                                                                    `json:"TargetCopyRegions,omitempty" xml:"TargetCopyRegions,omitempty"`
	TimePoints                   *string                                                                                                    `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
	Type                         *string                                                                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	VolumeNums                   *int32                                                                                                     `json:"VolumeNums,omitempty" xml:"VolumeNums,omitempty"`
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetCopiedSnapshotsRetentionDays() *int32 {
	return s.CopiedSnapshotsRetentionDays
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetCopyEncryptionConfiguration() *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration {
	return s.CopyEncryptionConfiguration
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetDiskNums() *int32 {
	return s.DiskNums
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetEnableCrossRegionCopy() *bool {
	return s.EnableCrossRegionCopy
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetTags() *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags {
	return s.Tags
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetTargetCopyRegions() *string {
	return s.TargetCopyRegions
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetTimePoints() *string {
	return s.TimePoints
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetType() *string {
	return s.Type
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetVolumeNums() *int32 {
	return s.VolumeNums
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetCopiedSnapshotsRetentionDays(v int32) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.CopiedSnapshotsRetentionDays = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetCopyEncryptionConfiguration(v *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.CopyEncryptionConfiguration = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetCreationTime(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.CreationTime = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetDiskNums(v int32) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.DiskNums = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetEnableCrossRegionCopy(v bool) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.EnableCrossRegionCopy = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRegionId(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRepeatWeekdays(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RepeatWeekdays = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetResourceGroupId(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRetentionDays(v int32) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RetentionDays = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetStatus(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.Status = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetTags(v *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.Tags = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetTargetCopyRegions(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.TargetCopyRegions = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetTimePoints(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.TimePoints = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetType(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.Type = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetVolumeNums(v int32) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.VolumeNums = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) Validate() error {
	if s.CopyEncryptionConfiguration != nil {
		if err := s.CopyEncryptionConfiguration.Validate(); err != nil {
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

type DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration struct {
	Encrypted *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId  *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) SetEncrypted(v bool) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration {
	s.Encrypted = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) SetKMSKeyId(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags struct {
	Tag []*DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags) GetTag() []*DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag {
	return s.Tag
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags) SetTag(v []*DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags {
	s.Tag = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags) Validate() error {
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

type DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) SetTagKey(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) SetTagValue(v string) *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTagsTag) Validate() error {
	return dara.Validate(s)
}
