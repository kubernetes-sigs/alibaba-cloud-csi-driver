// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBandwidthPackageIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *AddBandwidthPackageIpsRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *AddBandwidthPackageIpsRequest
	GetClientToken() *string
	SetIpCount(v string) *AddBandwidthPackageIpsRequest
	GetIpCount() *string
	SetOwnerAccount(v string) *AddBandwidthPackageIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddBandwidthPackageIpsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddBandwidthPackageIpsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddBandwidthPackageIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddBandwidthPackageIpsRequest
	GetResourceOwnerId() *int64
}

type AddBandwidthPackageIpsRequest struct {
	// This parameter is required.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	IpCount      *string `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddBandwidthPackageIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddBandwidthPackageIpsRequest) GoString() string {
	return s.String()
}

func (s *AddBandwidthPackageIpsRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *AddBandwidthPackageIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddBandwidthPackageIpsRequest) GetIpCount() *string {
	return s.IpCount
}

func (s *AddBandwidthPackageIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddBandwidthPackageIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddBandwidthPackageIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddBandwidthPackageIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddBandwidthPackageIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddBandwidthPackageIpsRequest) SetBandwidthPackageId(v string) *AddBandwidthPackageIpsRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) SetClientToken(v string) *AddBandwidthPackageIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) SetIpCount(v string) *AddBandwidthPackageIpsRequest {
	s.IpCount = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) SetOwnerAccount(v string) *AddBandwidthPackageIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) SetOwnerId(v int64) *AddBandwidthPackageIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) SetRegionId(v string) *AddBandwidthPackageIpsRequest {
	s.RegionId = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) SetResourceOwnerAccount(v string) *AddBandwidthPackageIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) SetResourceOwnerId(v int64) *AddBandwidthPackageIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddBandwidthPackageIpsRequest) Validate() error {
	return dara.Validate(s)
}
