// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCopiedSnapshotsRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest
	GetCopiedSnapshotsRetentionDays() *int32
	SetCopyEncryptionConfiguration(v *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) *CreateAutoSnapshotPolicyRequest
	GetCopyEncryptionConfiguration() *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration
	SetEnableCrossRegionCopy(v bool) *CreateAutoSnapshotPolicyRequest
	GetEnableCrossRegionCopy() *bool
	SetOwnerId(v int64) *CreateAutoSnapshotPolicyRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateAutoSnapshotPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateAutoSnapshotPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAutoSnapshotPolicyRequest
	GetResourceOwnerId() *int64
	SetStorageLocationArn(v string) *CreateAutoSnapshotPolicyRequest
	GetStorageLocationArn() *string
	SetTag(v []*CreateAutoSnapshotPolicyRequestTag) *CreateAutoSnapshotPolicyRequest
	GetTag() []*CreateAutoSnapshotPolicyRequestTag
	SetTargetCopyRegions(v string) *CreateAutoSnapshotPolicyRequest
	GetTargetCopyRegions() *string
	SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyName() *string
	SetRegionId(v string) *CreateAutoSnapshotPolicyRequest
	GetRegionId() *string
	SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest
	GetRepeatWeekdays() *string
	SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest
	GetRetentionDays() *int32
	SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest
	GetTimePoints() *string
}

type CreateAutoSnapshotPolicyRequest struct {
	// The retention period of the snapshot copy in the destination region. Unit: days. Valid values:
	//
	// 	- \\-1: The snapshot copy is retained until it is deleted.
	//
	// 	- 1 to 65535: The snapshot copy is retained for the specified number of days. After the retention period of the snapshot copy expires, the snapshot copy is automatically deleted.
	//
	// Default value: -1.
	//
	// example:
	//
	// 30
	CopiedSnapshotsRetentionDays *int32 `json:"CopiedSnapshotsRetentionDays,omitempty" xml:"CopiedSnapshotsRetentionDays,omitempty"`
	// The encryption parameters for cross-region snapshot replication.
	CopyEncryptionConfiguration *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration `json:"CopyEncryptionConfiguration,omitempty" xml:"CopyEncryptionConfiguration,omitempty" type:"Struct"`
	// Specifies whether to enable cross-region replication for snapshots.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableCrossRegionCopy *bool  `json:"EnableCrossRegionCopy,omitempty" xml:"EnableCrossRegionCopy,omitempty"`
	OwnerId               *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2kkmhmhs****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The tags to add to the automatic snapshot policy.
	Tag []*CreateAutoSnapshotPolicyRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The destination region to which to copy the snapshot. You can specify only a single destination region.
	//
	// example:
	//
	// ["cn-hangzhou"]
	TargetCopyRegions *string `json:"TargetCopyRegions,omitempty" xml:"TargetCopyRegions,omitempty"`
	// The name of the automatic snapshot policy. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// TestName
	AutoSnapshotPolicyName *string `json:"autoSnapshotPolicyName,omitempty" xml:"autoSnapshotPolicyName,omitempty"`
	// The ID of the region in which to create the automatic snapshot policy. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The days of the week on which to create automatic snapshots. Valid values: 1 to 7, which correspond to Monday to Sunday. 1 indicates Monday. Format description:
	//
	// 	- Set this parameter to a JSON-formatted array. For example, a value of ["1"] specifies automatic snapshots to be created every Monday.
	//
	// 	- To schedule multiple automatic snapshots to be created in a week, you can specify multiple values. Separate the values with commas (,). You can specify a maximum of seven days. For example, a value of ["1","3","5"] specifies automatic snapshots to be created every Monday, Wednesday, and Friday.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["1","2"]
	RepeatWeekdays *string `json:"repeatWeekdays,omitempty" xml:"repeatWeekdays,omitempty"`
	// The retention period of the automatic snapshot. Unit: days. Valid values:
	//
	// 	- \\-1: The automatic snapshot is retained until it is deleted.
	//
	// 	- 1 to 65535: The automatic snapshot is retained for the specified number of days. After the retention period of the automatic snapshot expires, the automatic snapshot is automatically deleted.
	//
	// Default value: -1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The points in time of the day at which to create automatic snapshots. The time must be in UTC+8. Unit: hours. Valid values: 0 to 23, which correspond to the 24 on-the-hour points in time from 00:00:00 to 23:00:00. For example, 1 indicates 01:00:00. Format description:
	//
	// 	- Set this parameter to a JSON-formatted array. For example, a value of ["1"] specifies automatic snapshots to be created at 01:00:00.
	//
	// 	- To schedule multiple automatic snapshots to be created in a day, you can specify multiple values. Separate the values with commas (,). You can specify up to 24 points in time. For example, a value of ["1","3","5"] specifies automatic snapshots to be created at 01:00:00, 03:00:00, and 05:00:00.
	//
	// >  If an automatic snapshot is being created when the time scheduled for creating another automatic snapshot is due, the new snapshot task is skipped. This may occur when a disk contains a large volume of data. For example, you scheduled snapshots to be automatically created at 09:00, 10:00, 11:00, and 12:00. The system starts to create a snapshot for the disk at 09:00:00. The process takes 80 minutes to complete because the disk contains a large volume of data and ends at 10:20:00. The system skips the automatic snapshot task scheduled for 10:00:00 and creates the next automatic snapshot for the disk at 11:00:00.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["0", "1", … "23"]
	TimePoints *string `json:"timePoints,omitempty" xml:"timePoints,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) GetCopiedSnapshotsRetentionDays() *int32 {
	return s.CopiedSnapshotsRetentionDays
}

func (s *CreateAutoSnapshotPolicyRequest) GetCopyEncryptionConfiguration() *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	return s.CopyEncryptionConfiguration
}

func (s *CreateAutoSnapshotPolicyRequest) GetEnableCrossRegionCopy() *bool {
	return s.EnableCrossRegionCopy
}

func (s *CreateAutoSnapshotPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoSnapshotPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAutoSnapshotPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAutoSnapshotPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAutoSnapshotPolicyRequest) GetStorageLocationArn() *string {
	return s.StorageLocationArn
}

func (s *CreateAutoSnapshotPolicyRequest) GetTag() []*CreateAutoSnapshotPolicyRequestTag {
	return s.Tag
}

func (s *CreateAutoSnapshotPolicyRequest) GetTargetCopyRegions() *string {
	return s.TargetCopyRegions
}

func (s *CreateAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *CreateAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAutoSnapshotPolicyRequest) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *CreateAutoSnapshotPolicyRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateAutoSnapshotPolicyRequest) GetTimePoints() *string {
	return s.TimePoints
}

func (s *CreateAutoSnapshotPolicyRequest) SetCopiedSnapshotsRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.CopiedSnapshotsRetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetCopyEncryptionConfiguration(v *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) *CreateAutoSnapshotPolicyRequest {
	s.CopyEncryptionConfiguration = v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetEnableCrossRegionCopy(v bool) *CreateAutoSnapshotPolicyRequest {
	s.EnableCrossRegionCopy = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetOwnerId(v int64) *CreateAutoSnapshotPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetResourceGroupId(v string) *CreateAutoSnapshotPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetResourceOwnerAccount(v string) *CreateAutoSnapshotPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetResourceOwnerId(v int64) *CreateAutoSnapshotPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetStorageLocationArn(v string) *CreateAutoSnapshotPolicyRequest {
	s.StorageLocationArn = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTag(v []*CreateAutoSnapshotPolicyRequestTag) *CreateAutoSnapshotPolicyRequest {
	s.Tag = v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTargetCopyRegions(v string) *CreateAutoSnapshotPolicyRequest {
	s.TargetCopyRegions = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRegionId(v string) *CreateAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) Validate() error {
	if s.CopyEncryptionConfiguration != nil {
		if err := s.CopyEncryptionConfiguration.Validate(); err != nil {
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

type CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration struct {
	// >  This parameter is not publicly available.
	Arn []*CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable cross-region snapshot replication and encryption. Valid values:
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
	// The ID of the Key Management Service (KMS) key used in cross-region snapshot replication and encryption.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40826X
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GetArn() []*CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	return s.Arn
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) SetArn(v []*CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	s.Arn = v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) SetEncrypted(v bool) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	s.Encrypted = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) SetKMSKeyId(v string) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration {
	s.KMSKeyId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfiguration) Validate() error {
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 1000000000
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// hide
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// hide
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) SetAssumeRoleFor(v int64) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) SetRoleType(v string) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	s.RoleType = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) SetRolearn(v string) *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn {
	s.Rolearn = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestCopyEncryptionConfigurationArn) Validate() error {
	return dara.Validate(s)
}

type CreateAutoSnapshotPolicyRequestTag struct {
	// The key of tag N to add to the automatic snapshot policy. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the automatic snapshot policy. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with acs:.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAutoSnapshotPolicyRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAutoSnapshotPolicyRequestTag) SetKey(v string) *CreateAutoSnapshotPolicyRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestTag) SetValue(v string) *CreateAutoSnapshotPolicyRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequestTag) Validate() error {
	return dara.Validate(s)
}
