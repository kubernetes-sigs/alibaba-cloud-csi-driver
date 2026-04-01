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
	// Specify a value based on the specifications on the buy page.
	//
	// [CPFS (Pay-as-you-go)](https://common-buy-intl.alibabacloud.com/?spm=5176.nas_overview.0.0.7ea01dbft0dTui\\&commodityCode=nas_cpfspost_public_intl#/buy)
	//
	// example:
	//
	// 150
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// Specify the capacity of the file system. Unit: GiB. This parameter is required and valid when FileSystemType is set to extreme, cpfs, or cpfsse.
	//
	// Specify a value based on the specifications on the following buy page:
	//
	// 	- [Extreme NAS (Pay-as-you-go)](https://common-buy-intl.alibabacloud.com/?commodityCode=nas_extpost_public_intl#/buy)
	//
	// 	- [CPFS (Pay-as-you-go)](https://common-buy-intl.alibabacloud.com/?spm=5176.nas_overview.0.0.7ea01dbft0dTui\\&commodityCode=nas_cpfspost_public_intl#/buy)
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PayAsYouGo (default): pay-as-you-go
	//
	// 	- Subscription
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. ClientToken only supports ASCII characters and cannot exceed 64 characters. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the file system.
	//
	// Limits:
	//
	// 	- Must be 2 to 128 characters in length.
	//
	// 	- Must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- Can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether to precheck the creation request.
	//
	// The precheck operation helps you check the validity of parameters and verify inventory. It does not actually create instances and does not incur fees.
	//
	// Valid values:
	//
	// 	- true: Checks the request without creating an instance. The system checks the required parameters, request syntax, service limits, and available NAS resources. If the request fails to pass the check, an error message is returned. If the request passes the check, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	//
	// 	- false (default): Sends the request. If the request passes the check, the instance is created.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration.
	//
	// This parameter is valid and required if ChargeType is set to Subscription. Unit: months.
	//
	// If you do not renew a subscription file system when the file system expires, the file system is automatically released.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to encrypt data in the file system.
	//
	// You can use the keys that are managed by Key Management Service (KMS) to encrypt data in a file system. When you read and write the encrypted data, the data is automatically decrypted.
	//
	// Valid values:
	//
	// 	- 0 (default): The data in the file system is not encrypted.
	//
	// 	- 1: A NAS-managed key is used to encrypt the data in the file system. This value is valid if FileSystemType is set to standard or extreme.
	//
	// 	- 2: A KMS-managed key is used to encrypt the data in the file system. This value is valid if the FileSystemType parameter is set to standard or extreme.
	//
	// >
	//
	// 	- Extreme NAS: All regions except China East 1 Finance support KMS-managed keys.
	//
	// 	- General-purpose NAS: All regions support KMS-managed keys.
	//
	// example:
	//
	// 1
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard: General-purpose NAS
	//
	// 	- extreme: Extreme NAS
	//
	// 	- cpfs: CPFS (locally redundant storage)
	//
	// 	- cpfsse: CPFS SE (zone-redundant storage)
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The ID of the KMS key.
	//
	// This parameter is required if EncryptType is set to 2.
	//
	// example:
	//
	// 3c0b3885-2adf-483d-8a65-5e280689****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// Specify the protocol type.
	//
	// 	- If FileSystemType is set to standard, set this parameter to NFS or SMB.
	//
	// 	- If FileSystemType is set to extreme, set this parameter to NFS.
	//
	// 	- If FileSystemType is set to cpfs, set this parameter to cpfs.
	//
	// 	- If FileSystemType is set to cpfsse, set this parameter to cpfs.
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// Storage redundancy type. Only available for CPFS SE.
	//
	// Valid values:
	//
	// 	- ZRS
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ZRS
	RedundancyType *string `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	// A list of IDs for the zone-redundant vSwitches. This parameter is required if RedundancyType is set to ZRS. You must enter three vSwitch IDs from three different zones.
	//
	// if can be null:
	// true
	RedundancyVSwitchIds []*string `json:"RedundancyVSwitchIds,omitempty" xml:"RedundancyVSwitchIds,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// You can log on to the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups?) to view resource group IDs.
	//
	// example:
	//
	// rg-acfmwavnfdf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is available only for advanced Extreme NAS file systems.
	//
	// > You can create a file system from a snapshot. The version of the file system is the same as that of the source file system. For example, the source file system of the snapshot uses version 1. To create a file system of version 2, create File System A from the snapshot and create File System B of version 2. Then copy the data and migrate your business from File System A to File System B.
	//
	// example:
	//
	// s-xxx
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The storage type.
	//
	// 	- If FileSystemType is set to standard, set this parameter to Performance, Capacity, or Premium.
	//
	// 	- If FileSystemType is set to extreme, set this parameter to standard or advance.
	//
	// 	- If FileSystemType is set to cpfs, set this parameter to advance_100 (100 MB/s/TiB Baseline), advance_200 (200 MB/s/TiB Baseline), or economic.
	//
	// 	- If FileSystemType is set to cpfsse, set this parameter to advance_100 (100 MB/s/TiB Baseline).
	//
	// This parameter is required.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// An array of tags.
	//
	// You can specify up to 20 tags. If you specify multiple tags, each tag key must be unique.
	Tag []*CreateFileSystemRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// 	- This parameter is required if FileSystemType is set to cpfs.
	//
	// 	- If FileSystemType is not set to cpfs, this parameter is reserved and not required.
	//
	// example:
	//
	// vsw-2ze37k6jh8ums2fw2****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// 	- This parameter is required if FileSystemType is set to cpfs or cpfsse.
	//
	// 	- This parameter is reserved and not required if FileSystemType is set to standard or extreme.
	//
	// example:
	//
	// vpc-bp1cbv1ljve4j5hlw****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone.
	//
	// Each region has multiple isolated locations known as zones. Each zone has its own independent power supply and network.
	//
	// This parameter is not required if FileSystemType is set to standard. By default, a random zone is selected based on the protocol type and storage type.
	//
	// This parameter is required if FileSystemType is set to extreme or cpfs.
	//
	// >
	//
	// 	- An Elastic Compute Service (ECS) instance and a file system that reside in different zones of the same region can access each other.
	//
	// 	- We recommend that you select the zone where the ECS instance resides. This prevents cross-zone latency between the file system and the ECS instance.
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
	// 	- Cannot be null or an empty string.
	//
	// 	- Can be up to 128 characters in length.
	//
	// 	- Cannot start with `aliyun` or `acs:`.
	//
	// 	- Cannot contain `http://` or `https://`.
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
	// 	- Cannot be null or an empty string.
	//
	// 	- Can be up to 128 characters in length.
	//
	// 	- Cannot contain `http://` or `https://`.
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
