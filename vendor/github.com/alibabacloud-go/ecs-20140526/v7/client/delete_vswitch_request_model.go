// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteVSwitchRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVSwitchRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVSwitchRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVSwitchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVSwitchRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *DeleteVSwitchRequest
	GetVSwitchId() *string
}

type DeleteVSwitchRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DeleteVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchRequest) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVSwitchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVSwitchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVSwitchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVSwitchRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeleteVSwitchRequest) SetOwnerAccount(v string) *DeleteVSwitchRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVSwitchRequest) SetOwnerId(v int64) *DeleteVSwitchRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVSwitchRequest) SetRegionId(v string) *DeleteVSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVSwitchRequest) SetResourceOwnerAccount(v string) *DeleteVSwitchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVSwitchRequest) SetResourceOwnerId(v int64) *DeleteVSwitchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVSwitchRequest) SetVSwitchId(v string) *DeleteVSwitchRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeleteVSwitchRequest) Validate() error {
	return dara.Validate(s)
}
