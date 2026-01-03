// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBandwidthPackageSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *ModifyBandwidthPackageSpecRequest
	GetBandwidth() *string
	SetBandwidthPackageId(v string) *ModifyBandwidthPackageSpecRequest
	GetBandwidthPackageId() *string
	SetOwnerAccount(v string) *ModifyBandwidthPackageSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBandwidthPackageSpecRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyBandwidthPackageSpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyBandwidthPackageSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBandwidthPackageSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyBandwidthPackageSpecRequest struct {
	// This parameter is required.
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// This parameter is required.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBandwidthPackageSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyBandwidthPackageSpecRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyBandwidthPackageSpecRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ModifyBandwidthPackageSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBandwidthPackageSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBandwidthPackageSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBandwidthPackageSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBandwidthPackageSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBandwidthPackageSpecRequest) SetBandwidth(v string) *ModifyBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyBandwidthPackageSpecRequest) SetBandwidthPackageId(v string) *ModifyBandwidthPackageSpecRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ModifyBandwidthPackageSpecRequest) SetOwnerAccount(v string) *ModifyBandwidthPackageSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBandwidthPackageSpecRequest) SetOwnerId(v int64) *ModifyBandwidthPackageSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBandwidthPackageSpecRequest) SetRegionId(v string) *ModifyBandwidthPackageSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBandwidthPackageSpecRequest) SetResourceOwnerAccount(v string) *ModifyBandwidthPackageSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBandwidthPackageSpecRequest) SetResourceOwnerId(v int64) *ModifyBandwidthPackageSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBandwidthPackageSpecRequest) Validate() error {
	return dara.Validate(s)
}
