// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoSnapshotPolicyExRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCopiedSnapshotsRetentionDays(v int32) *ModifyAutoSnapshotPolicyExRequest
	GetCopiedSnapshotsRetentionDays() *int32
	SetCopyEncryptionConfiguration(v *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) *ModifyAutoSnapshotPolicyExRequest
	GetCopyEncryptionConfiguration() *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration
	SetEnableCrossRegionCopy(v bool) *ModifyAutoSnapshotPolicyExRequest
	GetEnableCrossRegionCopy() *bool
	SetOwnerId(v int64) *ModifyAutoSnapshotPolicyExRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAutoSnapshotPolicyExRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAutoSnapshotPolicyExRequest
	GetResourceOwnerId() *int64
	SetTargetCopyRegions(v string) *ModifyAutoSnapshotPolicyExRequest
	GetTargetCopyRegions() *string
	SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyExRequest
	GetAutoSnapshotPolicyId() *string
	SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyExRequest
	GetAutoSnapshotPolicyName() *string
	SetRegionId(v string) *ModifyAutoSnapshotPolicyExRequest
	GetRegionId() *string
	SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyExRequest
	GetRepeatWeekdays() *string
	SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyExRequest
	GetRetentionDays() *int32
	SetTimePoints(v string) *ModifyAutoSnapshotPolicyExRequest
	GetTimePoints() *string
}

type ModifyAutoSnapshotPolicyExRequest struct {
	// The retention period of the snapshot copy in the destination region. Unit: days. Valid values:
	//
	// - -1: The snapshot copy is retained until it is deleted.
	//
	// - 1 to 65535: The snapshot copy is retained for a specified number of days.
	//
	// Default value: -1.
	//
	// example:
	//
	// 30
	CopiedSnapshotsRetentionDays *int32 `json:"CopiedSnapshotsRetentionDays,omitempty" xml:"CopiedSnapshotsRetentionDays,omitempty"`
	// The encryption configurations for cross-region snapshot replication.
	CopyEncryptionConfiguration *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration `json:"CopyEncryptionConfiguration,omitempty" xml:"CopyEncryptionConfiguration,omitempty" type:"Struct"`
	// Specifies whether to enable cross-region replication for the automatic snapshot.
	//
	// 	- true: enables cross-region replication for the automatic snapshot.
	//
	// 	- false: disables cross-region replication for the automatic snapshot.
	//
	// example:
	//
	// false
	EnableCrossRegionCopy *bool   `json:"EnableCrossRegionCopy,omitempty" xml:"EnableCrossRegionCopy,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination region to which to copy the snapshot. You can specify only a single destination region.
	//
	// example:
	//
	// ["cn-hangzhou"]
	TargetCopyRegions *string `json:"TargetCopyRegions,omitempty" xml:"TargetCopyRegions,omitempty"`
	// The ID of the automatic snapshot policy. You can call the [DescribeAutoSnapshotPolicyEx](https://help.aliyun.com/document_detail/25530.html) operation to query available automatic snapshot policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	AutoSnapshotPolicyId *string `json:"autoSnapshotPolicyId,omitempty" xml:"autoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy. If this parameter is not specified, the original name of the automatic snapshot policy is retained.
	//
	// example:
	//
	// SPTestName
	AutoSnapshotPolicyName *string `json:"autoSnapshotPolicyName,omitempty" xml:"autoSnapshotPolicyName,omitempty"`
	// The region ID of the automatic snapshot policy. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The days of the week on which to create automatic snapshots. Valid values are 1 to 7, which correspond to the days of the week. For example, a value of 1 indicates Monday.
	//
	// To schedule multiple automatic snapshots to be created in a week, you can specify multiple days.
	//
	// 	- You can specify up to seven days over a one-week period.
	//
	// 	- You must set this parameter to a JSON array such as `["1", "2" ... "7"]`. Separate the values in the array with commas (,).
	//
	// example:
	//
	// ["1", "7"]
	RepeatWeekdays *string `json:"repeatWeekdays,omitempty" xml:"repeatWeekdays,omitempty"`
	// The retention period of the automatic snapshot. Unit: days. Valid values:
	//
	// 	- \\-1: The automatic snapshot is permanently retained.
	//
	// 	- 1 to 65536: The auto snapshot is retained for the specified number of days.
	//
	// Default value: -1.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The points in time of the day at which to create automatic snapshots. The time must be in UTC+8. Unit: hours. Valid values are 0 to 23, which correspond to the 24 points in time on the hour from 00:00:00 to 23:00:00. For example, a value of 1 indicates 01:00:00.
	//
	// To schedule multiple automatic snapshots to be created in a day, you can specify multiple hours.
	//
	// 	- You can specify up to 24 points in time.
	//
	// 	- You must set this parameter to a JSON array such as `["0", "1", ... "23"]`. Separate the points in time with commas (,).
	//
	// example:
	//
	// ["0", "1"]
	TimePoints *string `json:"timePoints,omitempty" xml:"timePoints,omitempty"`
}

func (s ModifyAutoSnapshotPolicyExRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyExRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetCopiedSnapshotsRetentionDays() *int32 {
	return s.CopiedSnapshotsRetentionDays
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetCopyEncryptionConfiguration() *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration {
	return s.CopyEncryptionConfiguration
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetEnableCrossRegionCopy() *bool {
	return s.EnableCrossRegionCopy
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetTargetCopyRegions() *string {
	return s.TargetCopyRegions
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifyAutoSnapshotPolicyExRequest) GetTimePoints() *string {
	return s.TimePoints
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetCopiedSnapshotsRetentionDays(v int32) *ModifyAutoSnapshotPolicyExRequest {
	s.CopiedSnapshotsRetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetCopyEncryptionConfiguration(v *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) *ModifyAutoSnapshotPolicyExRequest {
	s.CopyEncryptionConfiguration = v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetEnableCrossRegionCopy(v bool) *ModifyAutoSnapshotPolicyExRequest {
	s.EnableCrossRegionCopy = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetOwnerId(v int64) *ModifyAutoSnapshotPolicyExRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetResourceOwnerAccount(v string) *ModifyAutoSnapshotPolicyExRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetResourceOwnerId(v int64) *ModifyAutoSnapshotPolicyExRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetTargetCopyRegions(v string) *ModifyAutoSnapshotPolicyExRequest {
	s.TargetCopyRegions = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyExRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyExRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetRegionId(v string) *ModifyAutoSnapshotPolicyExRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyExRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyExRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) SetTimePoints(v string) *ModifyAutoSnapshotPolicyExRequest {
	s.TimePoints = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequest) Validate() error {
	if s.CopyEncryptionConfiguration != nil {
		if err := s.CopyEncryptionConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration struct {
	// This parameter is not publicly available.
	Arn []*ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable encryption for cross-region snapshot replication. Valid values:
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
	// The ID of the KMS key used for encryption in cross-region snapshot replication.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40826X
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
}

func (s ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) GetArn() []*ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn {
	return s.Arn
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) SetArn(v []*ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration {
	s.Arn = v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) SetEncrypted(v bool) *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration {
	s.Encrypted = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) SetKMSKeyId(v string) *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfiguration) Validate() error {
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

type ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn struct {
	// This parameter is not publicly available.
	//
	// example:
	//
	// 1000000000
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// hide
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// hide
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) GetRoleType() *string {
	return s.RoleType
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) SetAssumeRoleFor(v int64) *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) SetRoleType(v string) *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn {
	s.RoleType = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) SetRolearn(v string) *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn {
	s.Rolearn = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyExRequestCopyEncryptionConfigurationArn) Validate() error {
	return dara.Validate(s)
}
