// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedInstanceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeReservedInstanceAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeReservedInstanceAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeReservedInstanceAutoRenewAttributeRequest
	GetRegionId() *string
	SetReservedInstanceId(v []*string) *DescribeReservedInstanceAutoRenewAttributeRequest
	GetReservedInstanceId() []*string
	SetResourceOwnerAccount(v string) *DescribeReservedInstanceAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeReservedInstanceAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeReservedInstanceAutoRenewAttributeRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the reserved instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of reserved instances.
	ReservedInstanceId   []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeReservedInstanceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeReservedInstanceAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeReservedInstanceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) SetRegionId(v string) *DescribeReservedInstanceAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) SetReservedInstanceId(v []*string) *DescribeReservedInstanceAutoRenewAttributeRequest {
	s.ReservedInstanceId = v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeReservedInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeReservedInstanceAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
