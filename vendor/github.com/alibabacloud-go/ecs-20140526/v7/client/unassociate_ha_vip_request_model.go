// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnassociateHaVipRequest
	GetClientToken() *string
	SetForce(v string) *UnassociateHaVipRequest
	GetForce() *string
	SetHaVipId(v string) *UnassociateHaVipRequest
	GetHaVipId() *string
	SetInstanceId(v string) *UnassociateHaVipRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UnassociateHaVipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateHaVipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassociateHaVipRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociateHaVipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateHaVipRequest
	GetResourceOwnerId() *int64
}

type UnassociateHaVipRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Force       *string `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnassociateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateHaVipRequest) GoString() string {
	return s.String()
}

func (s *UnassociateHaVipRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassociateHaVipRequest) GetForce() *string {
	return s.Force
}

func (s *UnassociateHaVipRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *UnassociateHaVipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnassociateHaVipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateHaVipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateHaVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateHaVipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateHaVipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateHaVipRequest) SetClientToken(v string) *UnassociateHaVipRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociateHaVipRequest) SetForce(v string) *UnassociateHaVipRequest {
	s.Force = &v
	return s
}

func (s *UnassociateHaVipRequest) SetHaVipId(v string) *UnassociateHaVipRequest {
	s.HaVipId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetInstanceId(v string) *UnassociateHaVipRequest {
	s.InstanceId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetOwnerAccount(v string) *UnassociateHaVipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateHaVipRequest) SetOwnerId(v int64) *UnassociateHaVipRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetRegionId(v string) *UnassociateHaVipRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateHaVipRequest) SetResourceOwnerAccount(v string) *UnassociateHaVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateHaVipRequest) SetResourceOwnerId(v int64) *UnassociateHaVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateHaVipRequest) Validate() error {
	return dara.Validate(s)
}
