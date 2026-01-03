// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DeleteBandwidthPackageRequest
	GetBandwidthPackageId() *string
	SetOwnerAccount(v string) *DeleteBandwidthPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteBandwidthPackageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteBandwidthPackageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBandwidthPackageRequest
	GetResourceOwnerId() *int64
}

type DeleteBandwidthPackageRequest struct {
	// This parameter is required.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DeleteBandwidthPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBandwidthPackageRequest) SetBandwidthPackageId(v string) *DeleteBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetOwnerAccount(v string) *DeleteBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetOwnerId(v int64) *DeleteBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetRegionId(v string) *DeleteBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetResourceOwnerAccount(v string) *DeleteBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) SetResourceOwnerId(v int64) *DeleteBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
