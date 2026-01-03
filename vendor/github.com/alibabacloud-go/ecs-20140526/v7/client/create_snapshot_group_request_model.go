// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSnapshotGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateSnapshotGroupRequest
	GetDescription() *string
	SetDiskId(v []*string) *CreateSnapshotGroupRequest
	GetDiskId() []*string
	SetExcludeDiskId(v []*string) *CreateSnapshotGroupRequest
	GetExcludeDiskId() []*string
	SetInstanceId(v string) *CreateSnapshotGroupRequest
	GetInstanceId() *string
	SetInstantAccess(v bool) *CreateSnapshotGroupRequest
	GetInstantAccess() *bool
	SetInstantAccessRetentionDays(v int32) *CreateSnapshotGroupRequest
	GetInstantAccessRetentionDays() *int32
	SetName(v string) *CreateSnapshotGroupRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateSnapshotGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSnapshotGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateSnapshotGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSnapshotGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateSnapshotGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSnapshotGroupRequest
	GetResourceOwnerId() *int64
	SetStorageLocationArn(v string) *CreateSnapshotGroupRequest
	GetStorageLocationArn() *string
	SetTag(v []*CreateSnapshotGroupRequestTag) *CreateSnapshotGroupRequest
	GetTag() []*CreateSnapshotGroupRequestTag
}

type CreateSnapshotGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the snapshot-consistent group. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the cloud disks for which you want to create a snapshot-consistent group. You can specify the IDs of cloud disks that are attached to multiple instances within the same zone. Valid values of N: 1 to 16. A snapshot-consistent group can contain snapshots of up to 16 cloud disks whose total disk size does not exceed 32 TiB.
	//
	// Take note of the following:
	//
	// 	- You cannot specify both the DiskId.N and `ExcludeDiskId.N` parameters in the same request.
	//
	// 	- If you specify `InstanceId`, you can specify the IDs of cloud disks that are attached only to the specified instance. You cannot specify the IDs of cloud disks that are attached to multiple instances.
	DiskId []*string `json:"DiskId,omitempty" xml:"DiskId,omitempty" type:"Repeated"`
	// The IDs of the cloud disks for which you do not want to create snapshots. After you specify the IDs of cloud disks, the snapshot-consistent group that you create does not contain the snapshots of the specified cloud disks. Valid values of N: 1 to 16.
	//
	// This parameter is empty by default, which indicates that snapshots are created for all disks of the instance.
	//
	// >  This parameter cannot be set at the same time as the `DiskId.N`.
	//
	// example:
	//
	// d-j6cf7l0ewidb78lq****
	ExcludeDiskId []*string `json:"ExcludeDiskId,omitempty" xml:"ExcludeDiskId,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// i-j6ca469urv8ei629****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable the instant access feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  This parameter is no longer used. By default, new standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// false
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// The number of days for which the instant access feature is available. Unit: days. Valid values: 1 to 65535.
	//
	// This parameter takes effect only when `InstantAccess` is set to true. The instant access feature is automatically disabled when the specified duration ends.
	//
	// This parameter is left empty by default, which indicates that the instant access feature is automatically disabled when the instant access snapshots are released.
	//
	// >  This parameter is no longer used. By default, new standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// 1
	InstantAccessRetentionDays *int32 `json:"InstantAccessRetentionDays,omitempty" xml:"InstantAccessRetentionDays,omitempty"`
	// The name of the snapshot-consistent group. The name must be 2 to 128 characters in length. The name can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testName
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the snapshot-consistent group belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// > This parameter is unavailable for public use.
	//
	// example:
	//
	// null
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The list of tags.
	Tag []*CreateSnapshotGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateSnapshotGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSnapshotGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSnapshotGroupRequest) GetDiskId() []*string {
	return s.DiskId
}

func (s *CreateSnapshotGroupRequest) GetExcludeDiskId() []*string {
	return s.ExcludeDiskId
}

func (s *CreateSnapshotGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSnapshotGroupRequest) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *CreateSnapshotGroupRequest) GetInstantAccessRetentionDays() *int32 {
	return s.InstantAccessRetentionDays
}

func (s *CreateSnapshotGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateSnapshotGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSnapshotGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSnapshotGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSnapshotGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSnapshotGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSnapshotGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSnapshotGroupRequest) GetStorageLocationArn() *string {
	return s.StorageLocationArn
}

func (s *CreateSnapshotGroupRequest) GetTag() []*CreateSnapshotGroupRequestTag {
	return s.Tag
}

func (s *CreateSnapshotGroupRequest) SetClientToken(v string) *CreateSnapshotGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetDescription(v string) *CreateSnapshotGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetDiskId(v []*string) *CreateSnapshotGroupRequest {
	s.DiskId = v
	return s
}

func (s *CreateSnapshotGroupRequest) SetExcludeDiskId(v []*string) *CreateSnapshotGroupRequest {
	s.ExcludeDiskId = v
	return s
}

func (s *CreateSnapshotGroupRequest) SetInstanceId(v string) *CreateSnapshotGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetInstantAccess(v bool) *CreateSnapshotGroupRequest {
	s.InstantAccess = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetInstantAccessRetentionDays(v int32) *CreateSnapshotGroupRequest {
	s.InstantAccessRetentionDays = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetName(v string) *CreateSnapshotGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetOwnerAccount(v string) *CreateSnapshotGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetOwnerId(v int64) *CreateSnapshotGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetRegionId(v string) *CreateSnapshotGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetResourceGroupId(v string) *CreateSnapshotGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetResourceOwnerAccount(v string) *CreateSnapshotGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetResourceOwnerId(v int64) *CreateSnapshotGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetStorageLocationArn(v string) *CreateSnapshotGroupRequest {
	s.StorageLocationArn = &v
	return s
}

func (s *CreateSnapshotGroupRequest) SetTag(v []*CreateSnapshotGroupRequestTag) *CreateSnapshotGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateSnapshotGroupRequest) Validate() error {
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

type CreateSnapshotGroupRequestTag struct {
	// The key of tag N of the snapshot-consistent group. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the snapshot-consistent group. Valid values of N: 1 to 20. The tag value can be an empty string. It can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSnapshotGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSnapshotGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateSnapshotGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateSnapshotGroupRequestTag) SetKey(v string) *CreateSnapshotGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSnapshotGroupRequestTag) SetValue(v string) *CreateSnapshotGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateSnapshotGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
