// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *CreateVpcRequest
	GetCidrBlock() *string
	SetClientToken(v string) *CreateVpcRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVpcRequest
	GetDescription() *string
	SetOwnerAccount(v string) *CreateVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVpcRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpcRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *CreateVpcRequest
	GetUserCidr() *string
	SetVpcName(v string) *CreateVpcRequest
	GetVpcName() *string
}

type CreateVpcRequest struct {
	CidrBlock    *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserCidr             *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	VpcName              *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s CreateVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *CreateVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpcRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *CreateVpcRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *CreateVpcRequest) SetCidrBlock(v string) *CreateVpcRequest {
	s.CidrBlock = &v
	return s
}

func (s *CreateVpcRequest) SetClientToken(v string) *CreateVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcRequest) SetDescription(v string) *CreateVpcRequest {
	s.Description = &v
	return s
}

func (s *CreateVpcRequest) SetOwnerAccount(v string) *CreateVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpcRequest) SetOwnerId(v int64) *CreateVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpcRequest) SetRegionId(v string) *CreateVpcRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcRequest) SetResourceOwnerAccount(v string) *CreateVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpcRequest) SetResourceOwnerId(v int64) *CreateVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpcRequest) SetUserCidr(v string) *CreateVpcRequest {
	s.UserCidr = &v
	return s
}

func (s *CreateVpcRequest) SetVpcName(v string) *CreateVpcRequest {
	s.VpcName = &v
	return s
}

func (s *CreateVpcRequest) Validate() error {
	return dara.Validate(s)
}
