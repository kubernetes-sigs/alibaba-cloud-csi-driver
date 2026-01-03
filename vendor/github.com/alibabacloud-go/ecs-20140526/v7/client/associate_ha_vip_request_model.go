// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateHaVipRequest
	GetClientToken() *string
	SetHaVipId(v string) *AssociateHaVipRequest
	GetHaVipId() *string
	SetInstanceId(v string) *AssociateHaVipRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *AssociateHaVipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateHaVipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateHaVipRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateHaVipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateHaVipRequest
	GetResourceOwnerId() *int64
}

type AssociateHaVipRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s AssociateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateHaVipRequest) GoString() string {
	return s.String()
}

func (s *AssociateHaVipRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateHaVipRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *AssociateHaVipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateHaVipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateHaVipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateHaVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateHaVipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateHaVipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateHaVipRequest) SetClientToken(v string) *AssociateHaVipRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateHaVipRequest) SetHaVipId(v string) *AssociateHaVipRequest {
	s.HaVipId = &v
	return s
}

func (s *AssociateHaVipRequest) SetInstanceId(v string) *AssociateHaVipRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateHaVipRequest) SetOwnerAccount(v string) *AssociateHaVipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateHaVipRequest) SetOwnerId(v int64) *AssociateHaVipRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateHaVipRequest) SetRegionId(v string) *AssociateHaVipRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateHaVipRequest) SetResourceOwnerAccount(v string) *AssociateHaVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateHaVipRequest) SetResourceOwnerId(v int64) *AssociateHaVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateHaVipRequest) Validate() error {
	return dara.Validate(s)
}
