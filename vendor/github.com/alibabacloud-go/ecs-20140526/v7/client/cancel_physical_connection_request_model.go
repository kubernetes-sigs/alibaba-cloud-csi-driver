// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelPhysicalConnectionRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CancelPhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelPhysicalConnectionRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *CancelPhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *CancelPhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelPhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelPhysicalConnectionRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *CancelPhysicalConnectionRequest
	GetUserCidr() *string
}

type CancelPhysicalConnectionRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserCidr             *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s CancelPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *CancelPhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelPhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelPhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelPhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CancelPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelPhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelPhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelPhysicalConnectionRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *CancelPhysicalConnectionRequest) SetClientToken(v string) *CancelPhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetOwnerAccount(v string) *CancelPhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetOwnerId(v int64) *CancelPhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetPhysicalConnectionId(v string) *CancelPhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetRegionId(v string) *CancelPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetResourceOwnerAccount(v string) *CancelPhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetResourceOwnerId(v int64) *CancelPhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) SetUserCidr(v string) *CancelPhysicalConnectionRequest {
	s.UserCidr = &v
	return s
}

func (s *CancelPhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
