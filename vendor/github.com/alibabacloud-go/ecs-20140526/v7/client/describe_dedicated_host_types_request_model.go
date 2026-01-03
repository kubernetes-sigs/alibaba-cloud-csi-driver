// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostType(v string) *DescribeDedicatedHostTypesRequest
	GetDedicatedHostType() *string
	SetOwnerAccount(v string) *DescribeDedicatedHostTypesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDedicatedHostTypesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDedicatedHostTypesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedHostTypesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedHostTypesRequest
	GetResourceOwnerId() *int64
	SetSupportedInstanceTypeFamily(v string) *DescribeDedicatedHostTypesRequest
	GetSupportedInstanceTypeFamily() *string
}

type DescribeDedicatedHostTypesRequest struct {
	// The dedicated host type. For more information, see [Dedicated host types](https://help.aliyun.com/document_detail/68564.html).
	//
	// example:
	//
	// ddh.sn1ne
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ECS instance family supported by the dedicated host type.
	//
	// example:
	//
	// ecs.sn1ne
	SupportedInstanceTypeFamily *string `json:"SupportedInstanceTypeFamily,omitempty" xml:"SupportedInstanceTypeFamily,omitempty"`
}

func (s DescribeDedicatedHostTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostTypesRequest) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *DescribeDedicatedHostTypesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDedicatedHostTypesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedHostTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostTypesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedHostTypesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedHostTypesRequest) GetSupportedInstanceTypeFamily() *string {
	return s.SupportedInstanceTypeFamily
}

func (s *DescribeDedicatedHostTypesRequest) SetDedicatedHostType(v string) *DescribeDedicatedHostTypesRequest {
	s.DedicatedHostType = &v
	return s
}

func (s *DescribeDedicatedHostTypesRequest) SetOwnerAccount(v string) *DescribeDedicatedHostTypesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostTypesRequest) SetOwnerId(v int64) *DescribeDedicatedHostTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostTypesRequest) SetRegionId(v string) *DescribeDedicatedHostTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostTypesRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostTypesRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostTypesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostTypesRequest) SetSupportedInstanceTypeFamily(v string) *DescribeDedicatedHostTypesRequest {
	s.SupportedInstanceTypeFamily = &v
	return s
}

func (s *DescribeDedicatedHostTypesRequest) Validate() error {
	return dara.Validate(s)
}
