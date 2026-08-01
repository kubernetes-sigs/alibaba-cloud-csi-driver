// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateFileSystemRequest
	GetBandwidth() *int64
	SetCapacity(v int64) *CreateFileSystemRequest
	GetCapacity() *int64
	SetChargeType(v string) *CreateFileSystemRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateFileSystemRequest
	GetClientToken() *string
	SetDescription(v string) *CreateFileSystemRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateFileSystemRequest
	GetDryRun() *bool
	SetDuration(v int32) *CreateFileSystemRequest
	GetDuration() *int32
	SetEncryptType(v int32) *CreateFileSystemRequest
	GetEncryptType() *int32
	SetFileSystemType(v string) *CreateFileSystemRequest
	GetFileSystemType() *string
	SetKmsKeyId(v string) *CreateFileSystemRequest
	GetKmsKeyId() *string
	SetProtocolType(v string) *CreateFileSystemRequest
	GetProtocolType() *string
	SetRedundancyType(v string) *CreateFileSystemRequest
	GetRedundancyType() *string
	SetRedundancyVSwitchIds(v []*string) *CreateFileSystemRequest
	GetRedundancyVSwitchIds() []*string
	SetResourceGroupId(v string) *CreateFileSystemRequest
	GetResourceGroupId() *string
	SetSnapshotId(v string) *CreateFileSystemRequest
	GetSnapshotId() *string
	SetStorageType(v string) *CreateFileSystemRequest
	GetStorageType() *string
	SetTag(v []*CreateFileSystemRequestTag) *CreateFileSystemRequest
	GetTag() []*CreateFileSystemRequestTag
	SetVSwitchId(v string) *CreateFileSystemRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateFileSystemRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateFileSystemRequest
	GetZoneId() *string
}

type CreateFileSystemRequest struct {
	// The maximum throughput of the file system.
	//
	// Unit: MB/s.
	//
	// For available specification values, see the actual specifications on the buy page.
	//
	// <props="china">
	//
	// [Parallel file system CPFS pay-as-you-go buy page](https://common-buy.aliyun.com/?commodityCode=nas_cpfs_post#/buy)
	//
	//
	//
	// <props="intl">
	//
	// [Parallel file system CPFS pay-as-you-go buy page](https://common-buy-intl.alibabacloud.com/?spm=5176.nas_overview.0.0.7ea01dbft0dTui&commodityCode=nas_cpfspost_public_intl#/buy)
	//
	// example:
	//
	// 150
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The capacity of the file system. Unit: GiB.
	//
	// This parameter is required and takes effect only when FileSystemType is set to extreme, cpfs, or cpfsse.
	//
	// For available values, see the actual specifications on the buy page:
	//
	// <props="china">
	//
	// -  [Extreme NAS pay-as-you-go buy page](https://common-buy.aliyun.com/?commodityCode=nas_extreme_post#/buy)
	//
	// - [Parallel file system CPFS pay-as-you-go buy page](https://common-buy.aliyun.com/?commodityCode=nas_cpfs_post#/buy)
	//
	//
	//
	// <props="intl">
	//
	// - [Extreme NAS pay-as-you-go buy page](https://common-buy-intl.alibabacloud.com/?commodityCode=nas_extpost_public_intl#/buy)
	//
	// - [Parallel file system CPFS pay-as-you-go buy page](https://common-buy-intl.alibabacloud.com/?spm=5176.nas_overview.0.0.7ea01dbft0dTui&commodityCode=nas_cpfspost_public_intl#/buy)
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// - PayAsYouGo (default): Pay-as-you-go.
	//
	// - Subscription: Subscription.
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Ensures the idempotence of the request. Generate a unique parameter value from your client. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the file system.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// 此文件系统的描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating the instance. The check items include required parameters, request format, service limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request. After the check succeeds, the instance is created.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration.
	//
	// Unit: months. This parameter is required and takes effect only when ChargeType is set to Subscription.
	//
	// If a subscription instance is not renewed upon expiration, the instance is automatically released.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to encrypt the file system.
	//
	// Uses keys managed by Key Management Service (KMS) to encrypt data stored on the file system. No decryption is required when reading or writing encrypted data.
	//
	// Valid values:
	//
	// - 0 (default): Not encrypted.
	//
	// - 1: NAS-managed key. Supported when FileSystemType is set to standard or extreme.
	//
	// - 2: Custom Key (KMS). Supported when FileSystemType is set to standard or extreme.
	//
	// > - Extreme NAS: The Custom Key (KMS) feature is supported in all regions except China (Hangzhou) Finance Cloud.
	//
	// > - General-purpose NAS: The Custom Key (KMS) feature is supported in all regions.
	//
	// example:
	//
	// 1
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// - standard (default): General-purpose NAS file system.
	//
	// - extreme: Extreme NAS file system.
	//
	// - cpfs: Cloud Parallel File Storage (CPFS) (locally redundant).
	//
	// - cpfsse: Cloud Parallel File Storage (CPFS) SE (zone-redundant).
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The KMS key ID.
	//
	// This parameter is required only when EncryptType is set to 2.
	//
	// example:
	//
	// fcbd****-62**-4a**-b605-c58cc1d5****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The file transfer protocol type.
	//
	// - If FileSystemType is set to standard, valid values: NFS and SMB.
	//
	// - If FileSystemType is set to extreme, valid values: NFS.
	//
	// - If FileSystemType is set to cpfs, valid values: cpfs.
	//
	// - If FileSystemType is set to cpfsse, valid values: cpfs.
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage redundancy type. This parameter takes effect only for CPFS SE.
	//
	// Valid values: ZRS.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ZRS
	RedundancyType *string `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	// The list of zone-redundant vSwitch IDs.
	//
	// If RedundancyType is set to ZRS, this parameter is required. You must specify three vSwitch IDs, each from a different zone.
	//
	// if can be null:
	// true
	RedundancyVSwitchIds []*string `json:"RedundancyVSwitchIds,omitempty" xml:"RedundancyVSwitchIds,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// You can view resource group IDs in the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups?).
	//
	// example:
	//
	// rg-acfmwavnfdf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is supported only for Extreme NAS file systems with the Advanced storage type.
	//
	// > A file system created from a snapshot has the same version as the source file system of the snapshot. For example, if the source file system version is 1 and you want to create a version 2 file system, first create file system A from the snapshot, then create file system B that meets the version 2 configuration. Copy the data from file system A to file system B, and migrate your workloads to file system B after the copy is complete.
	//
	// example:
	//
	// s-extreme-snapsho****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The storage type.
	//
	// - If FileSystemType is set to standard, valid values: Performance, Capacity, and Premium.
	//
	// - If FileSystemType is set to extreme, valid values: standard and advance.
	//
	// - If FileSystemType is set to cpfs, valid values: advance_100 (100 MB/s/TiB baseline), advance_200 (200 MB/s/TiB baseline), and economic.
	//
	// - If FileSystemType is set to cpfsse, valid values: advance_100 (100 MB/s/TiB baseline).
	//
	// This parameter is required.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tags.
	//
	// Array length: 1 to 20. If the array contains multiple tag objects, the tag key (Key) must be unique.
	Tag []*CreateFileSystemRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// - If FileSystemType is set to cpfs, this parameter is required.
	//
	// - If FileSystemType is not set to cpfs, this parameter is reserved and does not take effect. You do not need to configure it.
	//
	// example:
	//
	// vsw-bp131dkqilvw5pnlt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// - If FileSystemType is set to cpfs or cpfsse, this parameter is required.
	//
	// - If FileSystemType is set to standard or extreme, this parameter is reserved and does not take effect. You do not need to configure it.
	//
	// example:
	//
	// vpc-bp18cx9a7zoh0h9b4****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// A zone is an independent physical area within a region that has its own power supply and network.
	//
	// If FileSystemType is set to standard, this parameter is optional. By default, an active zone that matches the conditional ProtocolType and StorageType is randomly selected.
	//
	// If FileSystemType is set to extreme or cpfs, this parameter is required.
	//
	// >  - File systems and Elastic Computing Service (ECS) instances in different zones of the same region can communicate with each other.
	//
	// >  - Place the file system and the ECS server in the same zone to avoid cross-zone latency.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateFileSystemRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateFileSystemRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateFileSystemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFileSystemRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFileSystemRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateFileSystemRequest) GetEncryptType() *int32 {
	return s.EncryptType
}

func (s *CreateFileSystemRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *CreateFileSystemRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateFileSystemRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *CreateFileSystemRequest) GetRedundancyType() *string {
	return s.RedundancyType
}

func (s *CreateFileSystemRequest) GetRedundancyVSwitchIds() []*string {
	return s.RedundancyVSwitchIds
}

func (s *CreateFileSystemRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateFileSystemRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateFileSystemRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateFileSystemRequest) GetTag() []*CreateFileSystemRequestTag {
	return s.Tag
}

func (s *CreateFileSystemRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateFileSystemRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateFileSystemRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateFileSystemRequest) SetBandwidth(v int64) *CreateFileSystemRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateFileSystemRequest) SetCapacity(v int64) *CreateFileSystemRequest {
	s.Capacity = &v
	return s
}

func (s *CreateFileSystemRequest) SetChargeType(v string) *CreateFileSystemRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateFileSystemRequest) SetClientToken(v string) *CreateFileSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFileSystemRequest) SetDescription(v string) *CreateFileSystemRequest {
	s.Description = &v
	return s
}

func (s *CreateFileSystemRequest) SetDryRun(v bool) *CreateFileSystemRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFileSystemRequest) SetDuration(v int32) *CreateFileSystemRequest {
	s.Duration = &v
	return s
}

func (s *CreateFileSystemRequest) SetEncryptType(v int32) *CreateFileSystemRequest {
	s.EncryptType = &v
	return s
}

func (s *CreateFileSystemRequest) SetFileSystemType(v string) *CreateFileSystemRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateFileSystemRequest) SetKmsKeyId(v string) *CreateFileSystemRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateFileSystemRequest) SetProtocolType(v string) *CreateFileSystemRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequest) SetRedundancyType(v string) *CreateFileSystemRequest {
	s.RedundancyType = &v
	return s
}

func (s *CreateFileSystemRequest) SetRedundancyVSwitchIds(v []*string) *CreateFileSystemRequest {
	s.RedundancyVSwitchIds = v
	return s
}

func (s *CreateFileSystemRequest) SetResourceGroupId(v string) *CreateFileSystemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFileSystemRequest) SetSnapshotId(v string) *CreateFileSystemRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateFileSystemRequest) SetStorageType(v string) *CreateFileSystemRequest {
	s.StorageType = &v
	return s
}

func (s *CreateFileSystemRequest) SetTag(v []*CreateFileSystemRequestTag) *CreateFileSystemRequest {
	s.Tag = v
	return s
}

func (s *CreateFileSystemRequest) SetVSwitchId(v string) *CreateFileSystemRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateFileSystemRequest) SetVpcId(v string) *CreateFileSystemRequest {
	s.VpcId = &v
	return s
}

func (s *CreateFileSystemRequest) SetZoneId(v string) *CreateFileSystemRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateFileSystemRequest) Validate() error {
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

type CreateFileSystemRequestTag struct {
	// The tag key.
	//
	// Limits:
	//
	// - The tag key cannot be empty.
	//
	// - The tag key can be up to 128 characters in length.
	//
	// - The tag key cannot start with `aliyun` or `acs:`.
	//
	// - The tag key cannot contain `http://` or `https://`.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// nastest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// - The tag value cannot be empty.
	//
	// - The tag value can be up to 128 characters in length.
	//
	// - The tag value cannot contain `http://` or `https://`.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFileSystemRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFileSystemRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFileSystemRequestTag) SetKey(v string) *CreateFileSystemRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFileSystemRequestTag) SetValue(v string) *CreateFileSystemRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFileSystemRequestTag) Validate() error {
	return dara.Validate(s)
}
