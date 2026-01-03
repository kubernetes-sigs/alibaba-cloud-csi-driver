// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVpcRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpcRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DeleteVpcRequest
	GetVpcId() *string
}

type DeleteVpcRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteVpcRequest) SetOwnerAccount(v string) *DeleteVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpcRequest) SetOwnerId(v int64) *DeleteVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpcRequest) SetRegionId(v string) *DeleteVpcRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpcRequest) SetResourceOwnerAccount(v string) *DeleteVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpcRequest) SetResourceOwnerId(v int64) *DeleteVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpcRequest) SetVpcId(v string) *DeleteVpcRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteVpcRequest) Validate() error {
	return dara.Validate(s)
}
