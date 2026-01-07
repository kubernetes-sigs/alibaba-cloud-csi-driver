// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribePrefixListAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePrefixListAttributesRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *DescribePrefixListAttributesRequest
	GetPrefixListId() *string
	SetRegionId(v string) *DescribePrefixListAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePrefixListAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePrefixListAttributesRequest
	GetResourceOwnerId() *int64
}

type DescribePrefixListAttributesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the prefix list.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePrefixListAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePrefixListAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePrefixListAttributesRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DescribePrefixListAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePrefixListAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePrefixListAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePrefixListAttributesRequest) SetOwnerAccount(v string) *DescribePrefixListAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePrefixListAttributesRequest) SetOwnerId(v int64) *DescribePrefixListAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePrefixListAttributesRequest) SetPrefixListId(v string) *DescribePrefixListAttributesRequest {
	s.PrefixListId = &v
	return s
}

func (s *DescribePrefixListAttributesRequest) SetRegionId(v string) *DescribePrefixListAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePrefixListAttributesRequest) SetResourceOwnerAccount(v string) *DescribePrefixListAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePrefixListAttributesRequest) SetResourceOwnerId(v int64) *DescribePrefixListAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePrefixListAttributesRequest) Validate() error {
	return dara.Validate(s)
}
