// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v int64) *AllocateEipAddressRequest
	GetActivityId() *int64
	SetBandwidth(v string) *AllocateEipAddressRequest
	GetBandwidth() *string
	SetClientToken(v string) *AllocateEipAddressRequest
	GetClientToken() *string
	SetISP(v string) *AllocateEipAddressRequest
	GetISP() *string
	SetInternetChargeType(v string) *AllocateEipAddressRequest
	GetInternetChargeType() *string
	SetOwnerAccount(v string) *AllocateEipAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateEipAddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AllocateEipAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AllocateEipAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateEipAddressRequest
	GetResourceOwnerId() *int64
}

type AllocateEipAddressRequest struct {
	ActivityId         *int64  `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	Bandwidth          *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ISP                *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressRequest) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *AllocateEipAddressRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *AllocateEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateEipAddressRequest) GetISP() *string {
	return s.ISP
}

func (s *AllocateEipAddressRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *AllocateEipAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateEipAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateEipAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateEipAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateEipAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateEipAddressRequest) SetActivityId(v int64) *AllocateEipAddressRequest {
	s.ActivityId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetBandwidth(v string) *AllocateEipAddressRequest {
	s.Bandwidth = &v
	return s
}

func (s *AllocateEipAddressRequest) SetClientToken(v string) *AllocateEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateEipAddressRequest) SetISP(v string) *AllocateEipAddressRequest {
	s.ISP = &v
	return s
}

func (s *AllocateEipAddressRequest) SetInternetChargeType(v string) *AllocateEipAddressRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AllocateEipAddressRequest) SetOwnerAccount(v string) *AllocateEipAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateEipAddressRequest) SetOwnerId(v int64) *AllocateEipAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetRegionId(v string) *AllocateEipAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateEipAddressRequest) SetResourceOwnerAccount(v string) *AllocateEipAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateEipAddressRequest) SetResourceOwnerId(v int64) *AllocateEipAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
