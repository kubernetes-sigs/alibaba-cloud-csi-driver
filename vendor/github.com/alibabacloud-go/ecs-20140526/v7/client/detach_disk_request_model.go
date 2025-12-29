// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteWithInstance(v bool) *DetachDiskRequest
	GetDeleteWithInstance() *bool
	SetDiskId(v string) *DetachDiskRequest
	GetDiskId() *string
	SetInstanceId(v string) *DetachDiskRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DetachDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachDiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DetachDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachDiskRequest
	GetResourceOwnerId() *int64
}

type DetachDiskRequest struct {
	// Specifies whether to release the system disk or data disk when the instance from which you want to detach the disk is released. Valid values:
	//
	// 	- true: releases the disk when the instance is released.
	//
	// 	- false: does not release the disk when the instance is released. The disk is retained as a pay-as-you-go data disk.
	//
	// Default value: true.
	//
	// Take note of the following items:
	//
	// 	- You cannot specify this parameter for disks for which the multi-attach feature is enabled.
	//
	// 	- If a data disk is to be detached, the default value is `false`.
	//
	// 	- If you want to detach an `elastic ephemeral disk`, you must set `DeleteWithInstance` to `true`.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The ID of the disk that you want to detach.
	//
	// 	- The disk that you want to detach must be attached to an ECS instance and in the In Use (`In_use`) state.
	//
	// 	- The instance from which you want to detach a data disk must be in the `Running` or `Stopped` state.
	//
	// 	- The instance from which you want to detach the system disk must be in the `Stopped` state.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the ECS instance from which you want to detach the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DetachDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDiskRequest) GoString() string {
	return s.String()
}

func (s *DetachDiskRequest) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DetachDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DetachDiskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachDiskRequest) SetDeleteWithInstance(v bool) *DetachDiskRequest {
	s.DeleteWithInstance = &v
	return s
}

func (s *DetachDiskRequest) SetDiskId(v string) *DetachDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DetachDiskRequest) SetInstanceId(v string) *DetachDiskRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachDiskRequest) SetOwnerAccount(v string) *DetachDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachDiskRequest) SetOwnerId(v int64) *DetachDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachDiskRequest) SetResourceOwnerAccount(v string) *DetachDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachDiskRequest) SetResourceOwnerId(v int64) *DetachDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachDiskRequest) Validate() error {
	return dara.Validate(s)
}
