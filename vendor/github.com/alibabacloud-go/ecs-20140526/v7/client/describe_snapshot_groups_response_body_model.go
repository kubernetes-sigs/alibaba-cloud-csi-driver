// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSnapshotGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSnapshotGroupsResponseBody
	GetRequestId() *string
	SetSnapshotGroups(v *DescribeSnapshotGroupsResponseBodySnapshotGroups) *DescribeSnapshotGroupsResponseBody
	GetSnapshotGroups() *DescribeSnapshotGroupsResponseBodySnapshotGroups
}

type DescribeSnapshotGroupsResponseBody struct {
	// The token used to start the next query.
	//
	// > If the return value is empty, no more data exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3F9A4CC4-362F-469A-B9EF-B3204EF8AA3A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the snapshot-consistent groups.
	SnapshotGroups *DescribeSnapshotGroupsResponseBodySnapshotGroups `json:"SnapshotGroups,omitempty" xml:"SnapshotGroups,omitempty" type:"Struct"`
}

func (s DescribeSnapshotGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotGroupsResponseBody) GetSnapshotGroups() *DescribeSnapshotGroupsResponseBodySnapshotGroups {
	return s.SnapshotGroups
}

func (s *DescribeSnapshotGroupsResponseBody) SetNextToken(v string) *DescribeSnapshotGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBody) SetRequestId(v string) *DescribeSnapshotGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBody) SetSnapshotGroups(v *DescribeSnapshotGroupsResponseBodySnapshotGroups) *DescribeSnapshotGroupsResponseBody {
	s.SnapshotGroups = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBody) Validate() error {
	if s.SnapshotGroups != nil {
		if err := s.SnapshotGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotGroupsResponseBodySnapshotGroups struct {
	SnapshotGroup []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup `json:"SnapshotGroup,omitempty" xml:"SnapshotGroup,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroups) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroups) GetSnapshotGroup() []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	return s.SnapshotGroup
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroups) SetSnapshotGroup(v []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) *DescribeSnapshotGroupsResponseBodySnapshotGroups {
	s.SnapshotGroup = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroups) Validate() error {
	if s.SnapshotGroup != nil {
		for _, item := range s.SnapshotGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup struct {
	// The creation time. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-23T10:58:48Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the snapshot-consistent group.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance to which the snapshot-consistent group belongs. This parameter has a value only when all disk snapshots in the snapshot-consistent group belong to the same instance. If disk snapshots in the snapshot-consistent group belong to different instances, you can check the response parameters that start with `Snapshots.Snapshot.Tags.` to determine the ID of the instance to which each snapshot in the snapshot-consistent group belongs.
	//
	// example:
	//
	// i-j6ca469urv8ei629****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the snapshot-consistent group.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	ProgressStatus *string `json:"ProgressStatus,omitempty" xml:"ProgressStatus,omitempty"`
	// The ID of the resource group to which the snapshot-consistent group belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the snapshot-consistent group.
	//
	// example:
	//
	// ssg-j6ciyh3k52qp7ovm****
	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" xml:"SnapshotGroupId,omitempty"`
	// The information about the snapshots in the snapshot-consistent group.
	Snapshots *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The state of the snapshot-consistent group. Valid values:
	//
	// 	- progressing: The snapshot-consistent group was being created.
	//
	// 	- accomplished: The snapshot-consistent group was created.
	//
	// 	- failed: The snapshot-consistent group failed to be created.
	//
	// example:
	//
	// accomplished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the snapshot-consistent group.
	Tags *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetDescription() *string {
	return s.Description
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetName() *string {
	return s.Name
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetProgressStatus() *string {
	return s.ProgressStatus
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetSnapshotGroupId() *string {
	return s.SnapshotGroupId
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetSnapshots() *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots {
	return s.Snapshots
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) GetTags() *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags {
	return s.Tags
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetCreationTime(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetDescription(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetInstanceId(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetName(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.Name = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetProgressStatus(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.ProgressStatus = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetResourceGroupId(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetSnapshotGroupId(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.SnapshotGroupId = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetSnapshots(v *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetStatus(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) SetTags(v *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup {
	s.Tags = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroup) Validate() error {
	if s.Snapshots != nil {
		if err := s.Snapshots.Validate(); err != nil {
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

type DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots struct {
	Snapshot []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots) GetSnapshot() []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	return s.Snapshot
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots) SetSnapshot(v []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots {
	s.Snapshot = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshots) Validate() error {
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

type DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot struct {
	// Indicates whether the snapshot can be shared and be used to create or roll back a disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Available *bool `json:"Available,omitempty" xml:"Available,omitempty"`
	// Indicates whether the instant access feature is enabled. Valid values:
	//
	// 	- true: The instant access feature is enabled. By default, the instant access feature is enabled for ESSDs.
	//
	// 	- false: The instant access feature is disabled. The snapshot is a standard snapshot for which the instant access feature is disabled.
	//
	// >  This parameter is no longer used. By default, standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// true
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// The validity period of the instant access feature. When the validity period ends, the instant access snapshot is automatically released.
	//
	// >  This parameter is no longer used. By default, standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// 3
	InstantAccessRetentionDays *int32 `json:"InstantAccessRetentionDays,omitempty" xml:"InstantAccessRetentionDays,omitempty"`
	// The progress of the snapshot creation task. Unit: percent (%).
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-j6cbzmrlbf09w72q****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the source disk. This parameter is retained even after the source disk of the snapshot is released.
	//
	// example:
	//
	// d-j6c3ogynmvpi6wy7****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The type of the source disk. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// system
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The tags of the snapshot. The default values contain snapshot source information.
	Tags *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetAvailable() *bool {
	return s.Available
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetInstantAccessRetentionDays() *int32 {
	return s.InstantAccessRetentionDays
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) GetTags() *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags {
	return s.Tags
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetAvailable(v bool) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.Available = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetInstantAccess(v bool) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.InstantAccess = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetInstantAccessRetentionDays(v int32) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.InstantAccessRetentionDays = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetProgress(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetSnapshotId(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetSourceDiskId(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetSourceDiskType(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) SetTags(v *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot {
	s.Tags = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshot) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags struct {
	Tag []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags) GetTag() []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag {
	return s.Tag
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags) SetTag(v []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags {
	s.Tag = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTags) Validate() error {
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

type DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag struct {
	// The tag key of the snapshot. The default values of Key and Value contain snapshot source information.
	//
	// example:
	//
	// acs:ecs:createFrom
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the snapshot. The default values of Key and Value contain snapshot source information.
	//
	// example:
	//
	// i-bp11qm0o3dk4iuc****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) SetKey(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) SetValue(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupSnapshotsSnapshotTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags struct {
	Tag []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags) GetTag() []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag {
	return s.Tag
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags) SetTag(v []*DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags {
	s.Tag = v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTags) Validate() error {
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

type DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag struct {
	// The tag key of the snapshot-consistent group.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the snapshot-consistent group.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) SetKey(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) SetValue(v string) *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeSnapshotGroupsResponseBodySnapshotGroupsSnapshotGroupTagsTag) Validate() error {
	return dara.Validate(s)
}
