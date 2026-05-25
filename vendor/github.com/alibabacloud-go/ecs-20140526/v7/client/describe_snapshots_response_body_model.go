// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSnapshotsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeSnapshotsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v *DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody
	GetSnapshots() *DescribeSnapshotsResponseBodySnapshots
	SetTotalCount(v int32) *DescribeSnapshotsResponseBody
	GetTotalCount() *int32
}

type DescribeSnapshotsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
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
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots *DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The total number of snapshots.
	//
	// > When using the `MaxResults` and `NextToken` parameters for a paginated query, the returned `TotalCount` parameter value is invalid.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotsResponseBody) GetSnapshots() *DescribeSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeSnapshotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnapshotsResponseBody) SetNextToken(v string) *DescribeSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetPageNumber(v int32) *DescribeSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetPageSize(v int32) *DescribeSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v *DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetTotalCount(v int32) *DescribeSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) Validate() error {
	if s.Snapshots != nil {
		if err := s.Snapshots.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotsResponseBodySnapshots struct {
	Snapshot []*DescribeSnapshotsResponseBodySnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshot() []*DescribeSnapshotsResponseBodySnapshotsSnapshot {
	return s.Snapshot
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshot(v []*DescribeSnapshotsResponseBodySnapshotsSnapshot) *DescribeSnapshotsResponseBodySnapshots {
	s.Snapshot = v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) Validate() error {
	if s.Snapshot != nil {
		for _, item := range s.Snapshot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnapshotsResponseBodySnapshotsSnapshot struct {
	Available    *bool   `json:"Available,omitempty" xml:"Available,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Encrypted    *bool   `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// example:
	//
	// 65535
	FullSnapshotSizeInBytes    *int64  `json:"FullSnapshotSizeInBytes,omitempty" xml:"FullSnapshotSizeInBytes,omitempty"`
	InstantAccess              *bool   `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	InstantAccessRetentionDays *int32  `json:"InstantAccessRetentionDays,omitempty" xml:"InstantAccessRetentionDays,omitempty"`
	KMSKeyId                   *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	LastModifiedTime           *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	ProductCode                *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Progress                   *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemainTime                 *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	ResourceGroupId            *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RetentionDays              *int32  `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SnapshotId                 *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotLinkId             *string `json:"SnapshotLinkId,omitempty" xml:"SnapshotLinkId,omitempty"`
	SnapshotName               *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotSN                 *string `json:"SnapshotSN,omitempty" xml:"SnapshotSN,omitempty"`
	SnapshotType               *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SourceDiskId               *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	SourceDiskSize             *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	SourceDiskType             *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// example:
	//
	// cn-hangzhou
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// example:
	//
	// s-bp67acfmxazb4p****
	SourceSnapshotId  *string                                             `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	SourceStorageType *string                                             `json:"SourceStorageType,omitempty" xml:"SourceStorageType,omitempty"`
	Status            *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags              *DescribeSnapshotsResponseBodySnapshotsSnapshotTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Usage             *string                                             `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetAvailable() *bool {
	return s.Available
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetCategory() *string {
	return s.Category
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetDescription() *string {
	return s.Description
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetFullSnapshotSizeInBytes() *int64 {
	return s.FullSnapshotSizeInBytes
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetInstantAccessRetentionDays() *int32 {
	return s.InstantAccessRetentionDays
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetRemainTime() *int32 {
	return s.RemainTime
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotLinkId() *string {
	return s.SnapshotLinkId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotSN() *string {
	return s.SnapshotSN
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceDiskSize() *string {
	return s.SourceDiskSize
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceSnapshotId() *string {
	return s.SourceSnapshotId
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetSourceStorageType() *string {
	return s.SourceStorageType
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetTags() *DescribeSnapshotsResponseBodySnapshotsSnapshotTags {
	return s.Tags
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) GetUsage() *string {
	return s.Usage
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetAvailable(v bool) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Available = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetCategory(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Category = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetCreationTime(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetEncrypted(v bool) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Encrypted = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetFullSnapshotSizeInBytes(v int64) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.FullSnapshotSizeInBytes = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetInstantAccess(v bool) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.InstantAccess = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetInstantAccessRetentionDays(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.InstantAccessRetentionDays = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetKMSKeyId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetLastModifiedTime(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetProductCode(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.ProductCode = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRegionId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetResourceGroupId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetRetentionDays(v int32) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.RetentionDays = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotLinkId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotLinkId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotSN(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotSN = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotType(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceDiskId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceDiskSize(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceDiskType(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceRegionId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceRegionId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceSnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetSourceStorageType(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceStorageType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetTags(v *DescribeSnapshotsResponseBodySnapshotsSnapshotTags) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Tags = v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) SetUsage(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshot {
	s.Usage = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshot) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotsResponseBodySnapshotsSnapshotTags struct {
	Tag []*DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshotTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshotTags) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTags) GetTag() []*DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag {
	return s.Tag
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTags) SetTag(v []*DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) *DescribeSnapshotsResponseBodySnapshotsSnapshotTags {
	s.Tag = v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTags) Validate() error {
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

type DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) SetTagKey(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) SetTagValue(v string) *DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshotsSnapshotTagsTag) Validate() error {
	return dara.Validate(s)
}
