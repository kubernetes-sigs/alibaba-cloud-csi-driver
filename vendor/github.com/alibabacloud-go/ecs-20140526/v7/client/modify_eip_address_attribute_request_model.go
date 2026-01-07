// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEipAddressAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *ModifyEipAddressAttributeRequest
	GetAllocationId() *string
	SetBandwidth(v string) *ModifyEipAddressAttributeRequest
	GetBandwidth() *string
	SetOwnerAccount(v string) *ModifyEipAddressAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyEipAddressAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyEipAddressAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyEipAddressAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyEipAddressAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyEipAddressAttributeRequest struct {
	// This parameter is required.
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// This parameter is required.
	Bandwidth            *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyEipAddressAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEipAddressAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyEipAddressAttributeRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *ModifyEipAddressAttributeRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyEipAddressAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyEipAddressAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyEipAddressAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEipAddressAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyEipAddressAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyEipAddressAttributeRequest) SetAllocationId(v string) *ModifyEipAddressAttributeRequest {
	s.AllocationId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetBandwidth(v string) *ModifyEipAddressAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetOwnerAccount(v string) *ModifyEipAddressAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetOwnerId(v int64) *ModifyEipAddressAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetRegionId(v string) *ModifyEipAddressAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetResourceOwnerAccount(v string) *ModifyEipAddressAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) SetResourceOwnerId(v int64) *ModifyEipAddressAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyEipAddressAttributeRequest) Validate() error {
	return dara.Validate(s)
}
