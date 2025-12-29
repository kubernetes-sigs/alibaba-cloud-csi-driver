// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribePrefixListAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePrefixListAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribePrefixListAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePrefixListAssociationsRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *DescribePrefixListAssociationsRequest
	GetPrefixListId() *string
	SetRegionId(v string) *DescribePrefixListAssociationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePrefixListAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePrefixListAssociationsRequest
	GetResourceOwnerId() *int64
}

type DescribePrefixListAssociationsRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the `NextToken` value returned in the previous call to the DescribePrefixListAssociations operation. Leave this parameter empty the first time you call this operation.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// The region ID of the prefix list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribePrefixListAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListAssociationsRequest) GoString() string {
	return s.String()
}

func (s *DescribePrefixListAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePrefixListAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePrefixListAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePrefixListAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePrefixListAssociationsRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DescribePrefixListAssociationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePrefixListAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePrefixListAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePrefixListAssociationsRequest) SetMaxResults(v int32) *DescribePrefixListAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) SetNextToken(v string) *DescribePrefixListAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) SetOwnerAccount(v string) *DescribePrefixListAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) SetOwnerId(v int64) *DescribePrefixListAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) SetPrefixListId(v string) *DescribePrefixListAssociationsRequest {
	s.PrefixListId = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) SetRegionId(v string) *DescribePrefixListAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) SetResourceOwnerAccount(v string) *DescribePrefixListAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) SetResourceOwnerId(v int64) *DescribePrefixListAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePrefixListAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
