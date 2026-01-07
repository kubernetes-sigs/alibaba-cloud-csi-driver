// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeletePhysicalConnectionRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeletePhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePhysicalConnectionRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *DeletePhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *DeletePhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletePhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePhysicalConnectionRequest
	GetResourceOwnerId() *int64
}

type DeletePhysicalConnectionRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeletePhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeletePhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeletePhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DeletePhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePhysicalConnectionRequest) SetClientToken(v string) *DeletePhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletePhysicalConnectionRequest) SetOwnerAccount(v string) *DeletePhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePhysicalConnectionRequest) SetOwnerId(v int64) *DeletePhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePhysicalConnectionRequest) SetPhysicalConnectionId(v string) *DeletePhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DeletePhysicalConnectionRequest) SetRegionId(v string) *DeletePhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *DeletePhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePhysicalConnectionRequest) SetResourceOwnerId(v int64) *DeletePhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
