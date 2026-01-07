// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedFeatures(v string) *CreateDiskRequest
	GetAdvancedFeatures() *string
	SetArn(v []*CreateDiskRequestArn) *CreateDiskRequest
	GetArn() []*CreateDiskRequestArn
	SetBurstingEnabled(v bool) *CreateDiskRequest
	GetBurstingEnabled() *bool
	SetClientToken(v string) *CreateDiskRequest
	GetClientToken() *string
	SetDescription(v string) *CreateDiskRequest
	GetDescription() *string
	SetDiskCategory(v string) *CreateDiskRequest
	GetDiskCategory() *string
	SetDiskName(v string) *CreateDiskRequest
	GetDiskName() *string
	SetEncryptAlgorithm(v string) *CreateDiskRequest
	GetEncryptAlgorithm() *string
	SetEncrypted(v bool) *CreateDiskRequest
	GetEncrypted() *bool
	SetInstanceId(v string) *CreateDiskRequest
	GetInstanceId() *string
	SetKMSKeyId(v string) *CreateDiskRequest
	GetKMSKeyId() *string
	SetMultiAttach(v string) *CreateDiskRequest
	GetMultiAttach() *string
	SetOwnerAccount(v string) *CreateDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDiskRequest
	GetOwnerId() *int64
	SetPerformanceLevel(v string) *CreateDiskRequest
	GetPerformanceLevel() *string
	SetProvisionedIops(v int64) *CreateDiskRequest
	GetProvisionedIops() *int64
	SetRegionId(v string) *CreateDiskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDiskRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDiskRequest
	GetResourceOwnerId() *int64
	SetSize(v int32) *CreateDiskRequest
	GetSize() *int32
	SetSnapshotId(v string) *CreateDiskRequest
	GetSnapshotId() *string
	SetStorageClusterId(v string) *CreateDiskRequest
	GetStorageClusterId() *string
	SetStorageSetId(v string) *CreateDiskRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *CreateDiskRequest
	GetStorageSetPartitionNumber() *int32
	SetTag(v []*CreateDiskRequestTag) *CreateDiskRequest
	GetTag() []*CreateDiskRequestTag
	SetZoneId(v string) *CreateDiskRequest
	GetZoneId() *string
}

type CreateDiskRequest struct {
	// This parameter is not publicly available.
	//
	// example:
	//
	// hide
	AdvancedFeatures *string `json:"AdvancedFeatures,omitempty" xml:"AdvancedFeatures,omitempty"`
	// >  This parameter is not publicly available.
	Arn []*CreateDiskRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is available only if you set `DiskCategory` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The category of the data disk. Valid values for different disk categories:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: utra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: ESSD
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// 	- cloud_essd_entry: ESSD Entry disk
	//
	// 	- cloud_regional_disk_auto: Regional ESSD
	//
	// 	- elastic_ephemeral_disk_standard: standard elastic ephemeral disk
	//
	// 	- elastic_ephemeral_disk_premium: premium elastic ephemeral disk
	//
	// Default value: cloud.
	//
	// Enumerated values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: SSD
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// 	- cloud_regional_disk_auto: Regional ESSD
	//
	// 	- cloud_essd: ESSD
	//
	// 	- elastic_ephemeral_disk_standard: standard elastic ephemeral disk.
	//
	// 	- cloud_essd: ESSD
	//
	// 	- elastic_ephemeral_disk_premium: premium elastic ephemeral disk
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). The name must start with a letter.
	//
	// This parameter is left empty by default.
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// hide
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the disk. Valid values:
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
	// The ID of the subscription instance to which to attach the subscription disk.
	//
	// 	- If you specify an instance ID, the following parameters are ignored: ResourceGroupId, Tag.N.Key, Tag.N.Value, ClientToken, and KMSKeyId.
	//
	// 	- You cannot specify both ZoneId and InstanceId in a request.
	//
	// This parameter is empty by default, which indicates that a pay-as-you-go disk is created in the region and zone specified by RegionId and ZoneId.
	//
	// example:
	//
	// i-bp18pnlg1ds9rky4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used for the disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40826X
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// Specifies whether to enable the multi-attach feature for the disk. Valid values:
	//
	// 	- Disabled
	//
	// 	- Enabled Set the value to `Enabled` only for ESSDs.
	//
	// Default value: Disabled.
	//
	// >  Disks for which the multi-attach feature is enabled support only the pay-as-you-go billing method. When `MultiAttach` is set to Enabled, you cannot specify `InstanceId`. You can call the [AttachDisk](https://help.aliyun.com/document_detail/25515.html) operation to attach disks to instances after the disks are created. Disks for which the multi-attach feature is enabled can be attached only as data disks.
	//
	// example:
	//
	// Disabled
	MultiAttach  *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The performance level of the disk if the disk is an ESSD. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// For information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Specifies the read/write IOPS that is provisioned for the ESSD AutoPL disk. Valid value:
	//
	// 	- Capacity (GiB) <= 3: not configurable
	//
	// 	- Capacity (GiB) >= 4: [0, min{(1,000
	//
	// IOPS/GiB × Capacity - Baseline IOPS), 50,000}]
	//
	// Baseline performance: max{min{1,800 + 50 × Capacity, 50,000}, 3,000}
	//
	// >  This parameter is available only if you set `DiskCategory` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The ID of the region in which to create the disk. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to add the disk.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The size of the data disk. Unit: GiB. This parameter is required. Valid values for different disk categories:
	//
	// 	- Valid values when DiskCategory is set to cloud: 5 to 2000.
	//
	// 	- Valid values when DiskCategory is set to cloud_efficiency: 20 to 32768.
	//
	// 	- Valid values when DiskCategory is set to cloud_ssd: 20 to 32768.
	//
	// 	- Valid values when DiskCategory is set to cloud_essd: vary based on the `PerformanceLevel` value.
	//
	//     	- Valid values when PerformanceLevel is set to PL0: 1 to 65536.
	//
	//     	- Valid values when PerformanceLevel is set to PL1: 20 to 65536.
	//
	//     	- Valid values when PerformanceLevel is set to PL2: 461 to 65536.
	//
	//     	- Valid values when PerformanceLevel is set to PL3: 1261 to 65536.
	//
	// 	- Valid values when DiskCategory is set to cloud_auto: 1 to 65536.
	//
	// 	- Valid values when DiskCategory is set to cloud_essd_entry: 10 to 32768.
	//
	// 	- Valid values when DiskCategory is set to cloud_regional_disk_auto: 10 to 65536.
	//
	// 	- Valid values when DiskCategory is set to elastic_ephemeral_disk_standard: 64 to 8192.
	//
	// 	- Valid values when DiskCategory is set to elastic_ephemeral_disk_premium: 64 to 8192.
	//
	// If you specify `SnapshotId`, the following limits apply to `SnapshotId` and `Size`:
	//
	// 	- If the size of the snapshot specified by `SnapshotId` is larger than the value of `Size`, the size of the created disk is equal to the size of the snapshot.
	//
	// 	- If the size of the snapshot specified by `SnapshotId` is smaller than the value of `Size`, the size of the created disk is equal to the value of `Size`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot to use to create the disk. Snapshots that were created on or before July 15, 2013 cannot be used to create disks.
	//
	// The following limits apply to `SnapshotId` and `Size`:
	//
	// 	- If the size of the snapshot specified by `SnapshotId` is larger than the value of `Size`, the size of the created disk is equal to the specified snapshot size.
	//
	// 	- If the size of the snapshot specified by `SnapshotId` is smaller than the value of `Size`, the size of the created disk is equal to the value of `Size`.
	//
	// 	- You cannot create elastic ephemeral disks from snapshots.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster in which to create the disk. To create a disk in a specific dedicated block storage cluster, you must specify this parameter.
	//
	// >  You can specify the storage set-related parameters (`StorageSetId` and `StorageSetPartitionNumber`) or the dedicated block storage cluster-related parameter (`StorageClusterId`), but not both. If you specify a storage set-related parameter and a dedicated block storage cluster-related parameter in a request, the request fails.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The ID of the storage set.
	//
	// > You cannot specify storage set-related parameters (`StorageSetId` and `StorageSetPartitionNumber`) and the dedicated block storage cluster-related parameter (`StorageClusterId`) at the same time. Otherwise, the operation cannot be called.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The number of partitions in the storage set. The value must be greater than or equal to 2 but cannot exceed the quota obtained by calling the [DescribeAccountAttributes](https://help.aliyun.com/document_detail/73772.html)operation.
	//
	// Default value: 2.
	//
	// example:
	//
	// 3
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags to add to the disk.
	Tag []*CreateDiskRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the zone in which to create the pay-as-you-go disk.
	//
	// 	- If you do not specify InstanceId, you must specify ZoneId.
	//
	// 	- You cannot specify both ZoneId and InstanceId in the same request.
	//
	// >  You do not need to specify this parameter if you set DiskCategory to `cloud_regional_disk_auto` to create a Regional ESSD.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskRequest) GetAdvancedFeatures() *string {
	return s.AdvancedFeatures
}

func (s *CreateDiskRequest) GetArn() []*CreateDiskRequestArn {
	return s.Arn
}

func (s *CreateDiskRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateDiskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDiskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDiskRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CreateDiskRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateDiskRequest) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateDiskRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDiskRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateDiskRequest) GetMultiAttach() *string {
	return s.MultiAttach
}

func (s *CreateDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDiskRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateDiskRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDiskRequest) GetSize() *int32 {
	return s.Size
}

func (s *CreateDiskRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateDiskRequest) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateDiskRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateDiskRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *CreateDiskRequest) GetTag() []*CreateDiskRequestTag {
	return s.Tag
}

func (s *CreateDiskRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDiskRequest) SetAdvancedFeatures(v string) *CreateDiskRequest {
	s.AdvancedFeatures = &v
	return s
}

func (s *CreateDiskRequest) SetArn(v []*CreateDiskRequestArn) *CreateDiskRequest {
	s.Arn = v
	return s
}

func (s *CreateDiskRequest) SetBurstingEnabled(v bool) *CreateDiskRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateDiskRequest) SetClientToken(v string) *CreateDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskRequest) SetDescription(v string) *CreateDiskRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskRequest) SetDiskCategory(v string) *CreateDiskRequest {
	s.DiskCategory = &v
	return s
}

func (s *CreateDiskRequest) SetDiskName(v string) *CreateDiskRequest {
	s.DiskName = &v
	return s
}

func (s *CreateDiskRequest) SetEncryptAlgorithm(v string) *CreateDiskRequest {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateDiskRequest) SetEncrypted(v bool) *CreateDiskRequest {
	s.Encrypted = &v
	return s
}

func (s *CreateDiskRequest) SetInstanceId(v string) *CreateDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDiskRequest) SetKMSKeyId(v string) *CreateDiskRequest {
	s.KMSKeyId = &v
	return s
}

func (s *CreateDiskRequest) SetMultiAttach(v string) *CreateDiskRequest {
	s.MultiAttach = &v
	return s
}

func (s *CreateDiskRequest) SetOwnerAccount(v string) *CreateDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDiskRequest) SetOwnerId(v int64) *CreateDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDiskRequest) SetPerformanceLevel(v string) *CreateDiskRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateDiskRequest) SetProvisionedIops(v int64) *CreateDiskRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateDiskRequest) SetRegionId(v string) *CreateDiskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskRequest) SetResourceGroupId(v string) *CreateDiskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDiskRequest) SetResourceOwnerAccount(v string) *CreateDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDiskRequest) SetResourceOwnerId(v int64) *CreateDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDiskRequest) SetSize(v int32) *CreateDiskRequest {
	s.Size = &v
	return s
}

func (s *CreateDiskRequest) SetSnapshotId(v string) *CreateDiskRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateDiskRequest) SetStorageClusterId(v string) *CreateDiskRequest {
	s.StorageClusterId = &v
	return s
}

func (s *CreateDiskRequest) SetStorageSetId(v string) *CreateDiskRequest {
	s.StorageSetId = &v
	return s
}

func (s *CreateDiskRequest) SetStorageSetPartitionNumber(v int32) *CreateDiskRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *CreateDiskRequest) SetTag(v []*CreateDiskRequestTag) *CreateDiskRequest {
	s.Tag = v
	return s
}

func (s *CreateDiskRequest) SetZoneId(v string) *CreateDiskRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDiskRequest) Validate() error {
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateDiskRequestArn struct {
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

func (s CreateDiskRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequestArn) GoString() string {
	return s.String()
}

func (s *CreateDiskRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateDiskRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateDiskRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateDiskRequestArn) SetAssumeRoleFor(v int64) *CreateDiskRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateDiskRequestArn) SetRoleType(v string) *CreateDiskRequestArn {
	s.RoleType = &v
	return s
}

func (s *CreateDiskRequestArn) SetRolearn(v string) *CreateDiskRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CreateDiskRequestArn) Validate() error {
	return dara.Validate(s)
}

type CreateDiskRequestTag struct {
	// The key of tag N to add to the disk. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the disk. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDiskRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDiskRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDiskRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDiskRequestTag) SetKey(v string) *CreateDiskRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDiskRequestTag) SetValue(v string) *CreateDiskRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDiskRequestTag) Validate() error {
	return dara.Validate(s)
}
