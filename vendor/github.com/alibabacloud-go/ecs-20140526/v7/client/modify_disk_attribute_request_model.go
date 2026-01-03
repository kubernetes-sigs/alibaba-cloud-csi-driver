// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBurstingEnabled(v bool) *ModifyDiskAttributeRequest
	GetBurstingEnabled() *bool
	SetDeleteAutoSnapshot(v bool) *ModifyDiskAttributeRequest
	GetDeleteAutoSnapshot() *bool
	SetDeleteWithInstance(v bool) *ModifyDiskAttributeRequest
	GetDeleteWithInstance() *bool
	SetDescription(v string) *ModifyDiskAttributeRequest
	GetDescription() *string
	SetDiskId(v string) *ModifyDiskAttributeRequest
	GetDiskId() *string
	SetDiskIds(v []*string) *ModifyDiskAttributeRequest
	GetDiskIds() []*string
	SetDiskName(v string) *ModifyDiskAttributeRequest
	GetDiskName() *string
	SetEnableAutoSnapshot(v bool) *ModifyDiskAttributeRequest
	GetEnableAutoSnapshot() *bool
	SetOwnerAccount(v string) *ModifyDiskAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDiskAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDiskAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDiskAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDiskAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDiskAttributeRequest struct {
	// Specifies whether to enable performance burst for the disk if the disk supports performance burst. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  An error is reported if you specify this parameter for a disk that does not support performance burst.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Specifies whether to delete the automatic snapshots of the disk when the disk is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is empty by default, which indicates that the current value remains unchanged.
	//
	// example:
	//
	// false
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Specifies whether to release the disk together with the associated instance. This parameter is empty by default, which indicates that the current value remains unchanged.
	//
	// An error is returned if you set `DeleteWithInstance` to `false` in one of the following cases:
	//
	// 	- The disk is a local disk.
	//
	// 	- The disk is a basic disk and is not removable. If the Portable attribute of a disk is set to false, the disk is not removable.
	//
	// **
	//
	// **Warning*	- If you set DeleteWithInstance to false and the instance to which the disk is attached is locked for security reasons, the DeleteWithInstance attribute of the disk is ignored and the disk is released together with the instance. If "LockReason" : "security" is displayed in the response when you query information about an instance, the instance is locked for security reasons.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the disk. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the disk whose attributes you want to modify.
	//
	// >  You can specify `DiskId` or `DiskIds.N`, but not both.
	//
	// example:
	//
	// d-bp1famypsnar20bv****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The IDs of the disks whose attributes you want to modify. Valid values of N: 0 to 100.
	//
	// >  You can specify `DiskId` or `DiskIds.N`, but not both.
	//
	// example:
	//
	// d-bp1famypsnar20bv****
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// The name of the disk. The name must be 2 to 128 characters in length and can contain Unicode characters under the Decimal Number category and the categories whose names contain Letter. The name can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// MyDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to enable the automatic snapshot policy feature. Valid values:
	//
	// 	- true: enables the automatic snapshot policy feature for the cloud disk.
	//
	// 	- false: disables the automatic snapshot policy feature for the cloud disk.
	//
	// This parameter is empty by default, which indicates that the current value remains unchanged.
	//
	// >  This parameter is deprecated. By default, the automatic snapshot policy feature is enabled for cloud disks. You only need to apply an automatic snapshot policy to a cloud disk before you can use the automatic snapshot policy.
	//
	// example:
	//
	// true
	EnableAutoSnapshot *bool   `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDiskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskAttributeRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDiskAttributeRequest) GetDeleteAutoSnapshot() *bool {
	return s.DeleteAutoSnapshot
}

func (s *ModifyDiskAttributeRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *ModifyDiskAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDiskAttributeRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyDiskAttributeRequest) GetDiskIds() []*string {
	return s.DiskIds
}

func (s *ModifyDiskAttributeRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyDiskAttributeRequest) GetEnableAutoSnapshot() *bool {
	return s.EnableAutoSnapshot
}

func (s *ModifyDiskAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDiskAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDiskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiskAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDiskAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDiskAttributeRequest) SetBurstingEnabled(v bool) *ModifyDiskAttributeRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDeleteAutoSnapshot(v bool) *ModifyDiskAttributeRequest {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDeleteWithInstance(v bool) *ModifyDiskAttributeRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDescription(v string) *ModifyDiskAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDiskId(v string) *ModifyDiskAttributeRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDiskIds(v []*string) *ModifyDiskAttributeRequest {
	s.DiskIds = v
	return s
}

func (s *ModifyDiskAttributeRequest) SetDiskName(v string) *ModifyDiskAttributeRequest {
	s.DiskName = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetEnableAutoSnapshot(v bool) *ModifyDiskAttributeRequest {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetOwnerAccount(v string) *ModifyDiskAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetOwnerId(v int64) *ModifyDiskAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetRegionId(v string) *ModifyDiskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDiskAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDiskAttributeRequest) SetResourceOwnerId(v int64) *ModifyDiskAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDiskAttributeRequest) Validate() error {
	return dara.Validate(s)
}
