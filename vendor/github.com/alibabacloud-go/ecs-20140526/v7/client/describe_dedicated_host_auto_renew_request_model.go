// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostAutoRenewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostIds(v string) *DescribeDedicatedHostAutoRenewRequest
	GetDedicatedHostIds() *string
	SetOwnerAccount(v string) *DescribeDedicatedHostAutoRenewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDedicatedHostAutoRenewRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDedicatedHostAutoRenewRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedHostAutoRenewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedHostAutoRenewRequest
	GetResourceOwnerId() *int64
}

type DescribeDedicatedHostAutoRenewRequest struct {
	// The ID of the dedicated host. You can specify up to 100 subscription dedicated host IDs. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****,dh-bp1f9vxmno****
	DedicatedHostIds *string `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the dedicated host resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDedicatedHostAutoRenewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostAutoRenewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAutoRenewRequest) GetDedicatedHostIds() *string {
	return s.DedicatedHostIds
}

func (s *DescribeDedicatedHostAutoRenewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDedicatedHostAutoRenewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedHostAutoRenewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostAutoRenewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedHostAutoRenewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedHostAutoRenewRequest) SetDedicatedHostIds(v string) *DescribeDedicatedHostAutoRenewRequest {
	s.DedicatedHostIds = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewRequest) SetOwnerAccount(v string) *DescribeDedicatedHostAutoRenewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewRequest) SetOwnerId(v int64) *DescribeDedicatedHostAutoRenewRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewRequest) SetRegionId(v string) *DescribeDedicatedHostAutoRenewRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostAutoRenewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostAutoRenewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewRequest) Validate() error {
	return dara.Validate(s)
}
