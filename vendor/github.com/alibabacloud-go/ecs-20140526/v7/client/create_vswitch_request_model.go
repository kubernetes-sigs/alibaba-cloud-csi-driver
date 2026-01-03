// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateVSwitchRequest
	GetCidrBlock() *string
	SetClientToken(v string) *CreateVSwitchRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVSwitchRequest
	GetDescription() *string
	SetOwnerAccount(v string) *CreateVSwitchRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVSwitchRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVSwitchRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVSwitchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVSwitchRequest
	GetResourceOwnerId() *int64
	SetVSwitchName(v string) *CreateVSwitchRequest
	GetVSwitchName() *string
	SetVpcId(v string) *CreateVSwitchRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateVSwitchRequest
	GetZoneId() *string
}

type CreateVSwitchRequest struct {
	// This parameter is required.
	CidrBlock            *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VSwitchName          *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVSwitchRequest) GoString() string {
	return s.String()
}

func (s *CreateVSwitchRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateVSwitchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVSwitchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVSwitchRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVSwitchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVSwitchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVSwitchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVSwitchRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *CreateVSwitchRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVSwitchRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateVSwitchRequest) SetCidrBlock(v string) *CreateVSwitchRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateVSwitchRequest) SetClientToken(v string) *CreateVSwitchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVSwitchRequest) SetDescription(v string) *CreateVSwitchRequest {
	s.Description = &v
	return s
}

func (s *CreateVSwitchRequest) SetOwnerAccount(v string) *CreateVSwitchRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVSwitchRequest) SetOwnerId(v int64) *CreateVSwitchRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVSwitchRequest) SetRegionId(v string) *CreateVSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVSwitchRequest) SetResourceOwnerAccount(v string) *CreateVSwitchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVSwitchRequest) SetResourceOwnerId(v int64) *CreateVSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVSwitchRequest) SetVSwitchName(v string) *CreateVSwitchRequest {
	s.VSwitchName = &v
	return s
}

func (s *CreateVSwitchRequest) SetVpcId(v string) *CreateVSwitchRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVSwitchRequest) SetZoneId(v string) *CreateVSwitchRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateVSwitchRequest) Validate() error {
	return dara.Validate(s)
}
