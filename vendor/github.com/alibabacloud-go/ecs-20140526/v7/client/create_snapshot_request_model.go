// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateSnapshotRequest
	GetCategory() *string
	SetClientToken(v string) *CreateSnapshotRequest
	GetClientToken() *string
	SetDescription(v string) *CreateSnapshotRequest
	GetDescription() *string
	SetDiskId(v string) *CreateSnapshotRequest
	GetDiskId() *string
	SetInstantAccess(v bool) *CreateSnapshotRequest
	GetInstantAccess() *bool
	SetInstantAccessRetentionDays(v int32) *CreateSnapshotRequest
	GetInstantAccessRetentionDays() *int32
	SetOwnerAccount(v string) *CreateSnapshotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSnapshotRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreateSnapshotRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateSnapshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSnapshotRequest
	GetResourceOwnerId() *int64
	SetRetentionDays(v int32) *CreateSnapshotRequest
	GetRetentionDays() *int32
	SetSnapshotName(v string) *CreateSnapshotRequest
	GetSnapshotName() *string
	SetStorageLocationArn(v string) *CreateSnapshotRequest
	GetStorageLocationArn() *string
	SetTag(v []*CreateSnapshotRequestTag) *CreateSnapshotRequest
	GetTag() []*CreateSnapshotRequestTag
}

type CreateSnapshotRequest struct {
	// The category of the snapshot. Valid values:
	//
	// 	- Standard: standard snapshot
	//
	// 	- Flash: local snapshot
	//
	// >  This parameter is no longer used. By default, new standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the snapshot. The description must be 2 to 256 characters in length and cannot start with `http:// `or `https://`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp1s5fnvk4gn2tws0****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to enable the instant access feature. Valid values:
	//
	// 	- true: enables the instant access feature. This feature can be enabled only for ESSDs.
	//
	// 	- false: does not enable the instant access feature. If InstantAccess is set to false, a standard snapshot is created.
	//
	// Default value: false.
	//
	// >  This parameter is no longer used. By default, new standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// false
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// The validity period of the instant access feature. When the validity period ends, the feature is disabled and the instant access snapshot is automatically released. This parameter takes effect only when `InstantAccess` is set to true. Unit: days. Valid values: 1 to 65535.
	//
	// By default, the value of this parameter is the same as that of `RetentionDays`.
	//
	// >  This parameter is no longer used. By default, new standard snapshots of ESSDs are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// 1
	InstantAccessRetentionDays *int32  `json:"InstantAccessRetentionDays,omitempty" xml:"InstantAccessRetentionDays,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The snapshot type. Valid values:
	//
	// 	- Standard: standard snapshot
	//
	// 	- Flash: local snapshot
	//
	// > This parameter will be removed in the future. We recommend that you use the `InstantAccess` parameter to ensure future compatibility. This parameter and the `InstantAccess` parameter cannot be specified at the same time. For more information, see the "Description" section of this topic.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the snapshot. Unit: days. Valid values: 1 to 65536. After the retention period ends, the snapshot is automatically released.
	//
	// This parameter is left empty by default, which indicates that the snapshot is not automatically released.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The name of the snapshot. The name must be 2 to 128 characters in length and start with a letter. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// >  The name cannot start with http:// or https://. The name cannot start with `auto` because the names of automatic snapshots start with auto.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// > This parameter is unavailable for public use.
	//
	// example:
	//
	// null
	StorageLocationArn *string `json:"StorageLocationArn,omitempty" xml:"StorageLocationArn,omitempty"`
	// The tags to add to the snapshot.
	Tag []*CreateSnapshotRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSnapshotRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSnapshotRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateSnapshotRequest) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *CreateSnapshotRequest) GetInstantAccessRetentionDays() *int32 {
	return s.InstantAccessRetentionDays
}

func (s *CreateSnapshotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSnapshotRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSnapshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSnapshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSnapshotRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateSnapshotRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *CreateSnapshotRequest) GetStorageLocationArn() *string {
	return s.StorageLocationArn
}

func (s *CreateSnapshotRequest) GetTag() []*CreateSnapshotRequestTag {
	return s.Tag
}

func (s *CreateSnapshotRequest) SetCategory(v string) *CreateSnapshotRequest {
	s.Category = &v
	return s
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetDiskId(v string) *CreateSnapshotRequest {
	s.DiskId = &v
	return s
}

func (s *CreateSnapshotRequest) SetInstantAccess(v bool) *CreateSnapshotRequest {
	s.InstantAccess = &v
	return s
}

func (s *CreateSnapshotRequest) SetInstantAccessRetentionDays(v int32) *CreateSnapshotRequest {
	s.InstantAccessRetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetOwnerAccount(v string) *CreateSnapshotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSnapshotRequest) SetOwnerId(v int64) *CreateSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceGroupId(v string) *CreateSnapshotRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceOwnerAccount(v string) *CreateSnapshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceOwnerId(v int64) *CreateSnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetStorageLocationArn(v string) *CreateSnapshotRequest {
	s.StorageLocationArn = &v
	return s
}

func (s *CreateSnapshotRequest) SetTag(v []*CreateSnapshotRequestTag) *CreateSnapshotRequest {
	s.Tag = v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
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

type CreateSnapshotRequestTag struct {
	// The key of tag N to add to the snapshot. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the snapshot. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSnapshotRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateSnapshotRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateSnapshotRequestTag) SetKey(v string) *CreateSnapshotRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSnapshotRequestTag) SetValue(v string) *CreateSnapshotRequestTag {
	s.Value = &v
	return s
}

func (s *CreateSnapshotRequestTag) Validate() error {
	return dara.Validate(s)
}
