// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBandwidthPackageIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *RemoveBandwidthPackageIpsRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *RemoveBandwidthPackageIpsRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *RemoveBandwidthPackageIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveBandwidthPackageIpsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveBandwidthPackageIpsRequest
	GetRegionId() *string
	SetRemovedIpAddresses(v []*string) *RemoveBandwidthPackageIpsRequest
	GetRemovedIpAddresses() []*string
	SetResourceOwnerAccount(v string) *RemoveBandwidthPackageIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveBandwidthPackageIpsRequest
	GetResourceOwnerId() *int64
}

type RemoveBandwidthPackageIpsRequest struct {
	// This parameter is required.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RemovedIpAddresses   []*string `json:"RemovedIpAddresses,omitempty" xml:"RemovedIpAddresses,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RemoveBandwidthPackageIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveBandwidthPackageIpsRequest) GoString() string {
	return s.String()
}

func (s *RemoveBandwidthPackageIpsRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *RemoveBandwidthPackageIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveBandwidthPackageIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveBandwidthPackageIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveBandwidthPackageIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveBandwidthPackageIpsRequest) GetRemovedIpAddresses() []*string {
	return s.RemovedIpAddresses
}

func (s *RemoveBandwidthPackageIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveBandwidthPackageIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveBandwidthPackageIpsRequest) SetBandwidthPackageId(v string) *RemoveBandwidthPackageIpsRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) SetClientToken(v string) *RemoveBandwidthPackageIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) SetOwnerAccount(v string) *RemoveBandwidthPackageIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) SetOwnerId(v int64) *RemoveBandwidthPackageIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) SetRegionId(v string) *RemoveBandwidthPackageIpsRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) SetRemovedIpAddresses(v []*string) *RemoveBandwidthPackageIpsRequest {
	s.RemovedIpAddresses = v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) SetResourceOwnerAccount(v string) *RemoveBandwidthPackageIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) SetResourceOwnerId(v int64) *RemoveBandwidthPackageIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveBandwidthPackageIpsRequest) Validate() error {
	return dara.Validate(s)
}
