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
	// example:
	//
	// 150
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The capacity of the file system. Unit: GiB.
	//
	// This parameter is valid and required if the FileSystemType parameter is set to extreme.
	//
	// Specify a value based on the specifications on the following buy page:
	//
	// [Extreme NAS file system (Pay-as-you-go)](https://common-buy-intl.alibabacloud.com/?commodityCode=nas_extpost_public_intl#/buy)
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
	// 	- Subscription: subscription
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
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
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no file system is created and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, limits, and available NAS resources. If the request fails the dry run, an error message is returned. If the request passes the precheck, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a file system is created.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The subscription duration.
	//
	// This parameter is valid and required only if the ChargeType parameter is set to Subscription. Unit: months.
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
	// 	- 1: A NAS-managed key is used to encrypt the data in the file system. This value is valid only if the FileSystemType parameter is set to standard or extreme.
	//
	// 	- 2: A KMS-managed key is used to encrypt the data in the file system. This value is valid only if the FileSystemType parameter is set to standard or extreme.
	//
	// >  	- Extreme NAS file systems: All regions except China East 1 Finance support KMS-managed keys.
	//
	// > 	- General-purpose NAS file systems: All regions support KMS-managed keys.
	//
	// example:
	//
	// 1
	EncryptType *int32 `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard (default): General-purpose NAS file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// 	- cpfs: Cloud Parallel File Storage (CPFS) file system
	//
	// > CPFS file systems are available only on the China site (aliyun.com).
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The ID of the KMS key.
	//
	// This parameter is required only if the EncryptType parameter is set to 2.
	//
	// example:
	//
	// 3c0b3885-2adf-483d-8a65-5e280689****
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The protocol type.
	//
	// 	- If the FileSystemType parameter is set to standard, you can set the ProtocolType parameter to NFS or SMB.
	//
	// 	- If the FileSystemType parameter is set to extreme, you can set the ProtocolType parameter to NFS.
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
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
	// >  You can create a file system from a snapshot. In this case, the version of the file system is the same as that of the source file system. For example, the source file system of the snapshot uses version 1. To create a file system of version 2, you can create File System A from the snapshot and create File System B of version 2. You can then copy the data and migrate your business from File System A to File System B.
	//
	// example:
	//
	// s-xxx
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The storage class.
	//
	// 	- If the FileSystemType parameter is set to standard, you can set the StorageType parameter to Performance, Capacity, or Premium.
	//
	// 	- If the FileSystemType parameter is set to extreme, you can set the StorageType parameter to standard or advance.
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
	// This parameter is reserved and does not take effect. You do not need to configure this parameter.
	//
	// example:
	//
	// vsw-2ze37k6jh8ums2fw2****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// This parameter is reserved and does not take effect. You do not need to configure this parameter.
	//
	// example:
	//
	// vpc-bp1cbv1ljve4j5hlw****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// Each region has multiple isolated locations known as zones. Each zone has its own independent power supply and networks.
	//
	// This parameter is not required if the FileSystemType parameter is set to standard. By default, a random zone is selected based on the protocol type and storage type.
	//
	// This parameter is required if the FileSystemType parameter is set to extreme.
	//
	// > 	- An Elastic Compute Service (ECS) instance and a NAS file system that reside in different zones of the same region can access each other.
	//
	// >	- We recommend that you select the zone where the ECS instance resides. This prevents cross-zone latency between the file system and the ECS instance.
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
	// 	- The tag key cannot be null or an empty string.
	//
	// 	- The tag key can be up to 128 characters in length.
	//
	// 	- The tag key cannot start with `aliyun` or `acs:`.
	//
	// 	- The tag key cannot contain `http://` or `https://`.
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
	// 	- The tag value cannot be null or an empty string.
	//
	// 	- The tag value can be up to 128 characters in length.
	//
	// 	- The tag value cannot contain `http://` or `https://`.
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
