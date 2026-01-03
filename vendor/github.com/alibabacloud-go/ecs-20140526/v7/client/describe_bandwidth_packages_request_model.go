// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DescribeBandwidthPackagesRequest
	GetBandwidthPackageId() *string
	SetNatGatewayId(v string) *DescribeBandwidthPackagesRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *DescribeBandwidthPackagesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBandwidthPackagesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBandwidthPackagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBandwidthPackagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBandwidthPackagesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeBandwidthPackagesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBandwidthPackagesRequest
	GetResourceOwnerId() *int64
}

type DescribeBandwidthPackagesRequest struct {
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	NatGatewayId       *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBandwidthPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagesRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeBandwidthPackagesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeBandwidthPackagesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBandwidthPackagesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBandwidthPackagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBandwidthPackagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBandwidthPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBandwidthPackagesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBandwidthPackagesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBandwidthPackagesRequest) SetBandwidthPackageId(v string) *DescribeBandwidthPackagesRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetNatGatewayId(v string) *DescribeBandwidthPackagesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetOwnerAccount(v string) *DescribeBandwidthPackagesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetOwnerId(v int64) *DescribeBandwidthPackagesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetPageNumber(v int32) *DescribeBandwidthPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetPageSize(v int32) *DescribeBandwidthPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetRegionId(v string) *DescribeBandwidthPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetResourceOwnerAccount(v string) *DescribeBandwidthPackagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) SetResourceOwnerId(v int64) *DescribeBandwidthPackagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBandwidthPackagesRequest) Validate() error {
	return dara.Validate(s)
}
