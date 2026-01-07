// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminatePhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *TerminatePhysicalConnectionRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *TerminatePhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TerminatePhysicalConnectionRequest
	GetOwnerId() *int64
	SetPhysicalConnectionId(v string) *TerminatePhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *TerminatePhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *TerminatePhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TerminatePhysicalConnectionRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *TerminatePhysicalConnectionRequest
	GetUserCidr() *string
}

type TerminatePhysicalConnectionRequest struct {
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

func (s TerminatePhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminatePhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *TerminatePhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TerminatePhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TerminatePhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TerminatePhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *TerminatePhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TerminatePhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TerminatePhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TerminatePhysicalConnectionRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *TerminatePhysicalConnectionRequest) SetClientToken(v string) *TerminatePhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetOwnerAccount(v string) *TerminatePhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetOwnerId(v int64) *TerminatePhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetPhysicalConnectionId(v string) *TerminatePhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetRegionId(v string) *TerminatePhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *TerminatePhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetResourceOwnerId(v int64) *TerminatePhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) SetUserCidr(v string) *TerminatePhysicalConnectionRequest {
	s.UserCidr = &v
	return s
}

func (s *TerminatePhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
