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
	// Details about the automatic snapshot policies.
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
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-bp67acfmxazb4ph****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy.
	//
	// example:
	//
	// testAutoSnapshotPolicyName
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 0
	CopiedSnapshotsRetentionDays *int32 `json:"CopiedSnapshotsRetentionDays,omitempty" xml:"CopiedSnapshotsRetentionDays,omitempty"`
	// Encryption configurations for cross-region snapshot replication.
	CopyEncryptionConfiguration *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyCopyEncryptionConfiguration `json:"CopyEncryptionConfiguration,omitempty" xml:"CopyEncryptionConfiguration,omitempty" type:"Struct"`
	// The time when the automatic snapshot policy was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-12-10T16:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The number of disks to which the automatic snapshot policy is applied.
	//
	// example:
	//
	// 1
	DiskNums *int32 `json:"DiskNums,omitempty" xml:"DiskNums,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	EnableCrossRegionCopy *bool `json:"EnableCrossRegionCopy,omitempty" xml:"EnableCrossRegionCopy,omitempty"`
	// The region ID of the automatic snapshot policy.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The days of the week on which to create automatic snapshots. Valid values: 1 to 7, which correspond to the days of the week. For example, 1 indicates Monday. One or more days can be specified.
	//
	// example:
	//
	// ["6"]
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2kkmhmhs****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The retention period of the automatic snapshots. Unit: days. Valid values:
	//
	// 	- \\-1: Automatic snapshots are retained until they are deleted.
	//
	// 	- 1 to 65536: Auto snapshots are retained for the specified number of days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	//
	// example:
	//
	// 7
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The status of the automatic snapshot policy. Valid values:
	//
	// 	- Normal: The automatic snapshot policy is normal.
	//
	// 	- Expire: The automatic snapshot policy cannot be used because your account has overdue payments.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the automatic snapshot policy.
	Tags *DescribeAutoSnapshotPolicyExResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// test
	TargetCopyRegions *string `json:"TargetCopyRegions,omitempty" xml:"TargetCopyRegions,omitempty"`
	// The points in time of the day at which to create automatic snapshots.
	//
	// The time is displayed in UTC+8. Unit: hours. Valid values: 0 to 23, which correspond to the 24 points in time on the hour from 00:00:00 to 23:00:00. For example, 1 indicates 01:00:00. Multiple points in time can be specified.
	//
	// The parameter value is a JSON array that contains up to 24 points in time separated by commas (,). Example: `["0", "1", ... "23"]`.
	//
	// example:
	//
	// ["1"]
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
	// The type of the automatic snapshot policy. Valid values:
	//
	// 	- Custom: user-defined snapshot policy.
	//
	// 	- System: system-defined snapshot policy.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The number of extended volumes to which the automatic snapshot policy is applied.
	//
	// example:
	//
	// 2
	VolumeNums *int32 `json:"VolumeNums,omitempty" xml:"VolumeNums,omitempty"`
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
	// Whether to enable encryption for cross-region snapshot replication. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key used to encrypt snapshots in cross-region snapshot replication.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
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
	// The tag key of the automatic snapshot policy.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the automatic snapshot policy.
	//
	// example:
	//
	// TestValue
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
