// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVscRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVscRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteVscRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVscRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVscRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVscRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVscRequest
	GetResourceOwnerId() *int64
	SetVscId(v string) *DeleteVscRequest
	GetVscId() *string
}

type DeleteVscRequest struct {
	// example:
	//
	// 123e4567-e89b-1**3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsc-bp1j8y**etwq1ow3jal
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DeleteVscRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVscRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVscRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVscRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVscRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVscRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVscRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVscRequest) GetVscId() *string {
	return s.VscId
}

func (s *DeleteVscRequest) SetClientToken(v string) *DeleteVscRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVscRequest) SetDryRun(v bool) *DeleteVscRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVscRequest) SetOwnerAccount(v string) *DeleteVscRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVscRequest) SetOwnerId(v int64) *DeleteVscRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVscRequest) SetRegionId(v string) *DeleteVscRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVscRequest) SetResourceOwnerAccount(v string) *DeleteVscRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVscRequest) SetResourceOwnerId(v int64) *DeleteVscRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVscRequest) SetVscId(v string) *DeleteVscRequest {
	s.VscId = &v
	return s
}

func (s *DeleteVscRequest) Validate() error {
	return dara.Validate(s)
}
