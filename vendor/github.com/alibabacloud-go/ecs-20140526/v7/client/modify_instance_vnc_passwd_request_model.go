// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVncPasswdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceVncPasswdRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceVncPasswdRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceVncPasswdRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceVncPasswdRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceVncPasswdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceVncPasswdRequest
	GetResourceOwnerId() *int64
	SetVncPassword(v string) *ModifyInstanceVncPasswdRequest
	GetVncPassword() *string
}

type ModifyInstanceVncPasswdRequest struct {
	// The ID of the ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new VNC password of the ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs123
	VncPassword *string `json:"VncPassword,omitempty" xml:"VncPassword,omitempty"`
}

func (s ModifyInstanceVncPasswdRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVncPasswdRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceVncPasswdRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceVncPasswdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceVncPasswdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceVncPasswdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceVncPasswdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceVncPasswdRequest) GetVncPassword() *string {
	return s.VncPassword
}

func (s *ModifyInstanceVncPasswdRequest) SetInstanceId(v string) *ModifyInstanceVncPasswdRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVncPasswdRequest) SetOwnerAccount(v string) *ModifyInstanceVncPasswdRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceVncPasswdRequest) SetOwnerId(v int64) *ModifyInstanceVncPasswdRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceVncPasswdRequest) SetRegionId(v string) *ModifyInstanceVncPasswdRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceVncPasswdRequest) SetResourceOwnerAccount(v string) *ModifyInstanceVncPasswdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceVncPasswdRequest) SetResourceOwnerId(v int64) *ModifyInstanceVncPasswdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceVncPasswdRequest) SetVncPassword(v string) *ModifyInstanceVncPasswdRequest {
	s.VncPassword = &v
	return s
}

func (s *ModifyInstanceVncPasswdRequest) Validate() error {
	return dara.Validate(s)
}
