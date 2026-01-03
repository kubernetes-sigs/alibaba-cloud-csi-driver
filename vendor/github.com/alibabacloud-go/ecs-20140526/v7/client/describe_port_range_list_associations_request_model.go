// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribePortRangeListAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePortRangeListAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribePortRangeListAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePortRangeListAssociationsRequest
	GetOwnerId() *int64
	SetPortRangeListId(v string) *DescribePortRangeListAssociationsRequest
	GetPortRangeListId() *string
	SetRegionId(v string) *DescribePortRangeListAssociationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePortRangeListAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePortRangeListAssociationsRequest
	GetResourceOwnerId() *int64
}

type DescribePortRangeListAssociationsRequest struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh15YZPnzqF7Vs2EB6Ix327v
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the port list.
	//
	// This parameter is required.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The region ID of the port list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribePortRangeListAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListAssociationsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePortRangeListAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePortRangeListAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePortRangeListAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePortRangeListAssociationsRequest) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *DescribePortRangeListAssociationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePortRangeListAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePortRangeListAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePortRangeListAssociationsRequest) SetMaxResults(v int32) *DescribePortRangeListAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) SetNextToken(v string) *DescribePortRangeListAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) SetOwnerAccount(v string) *DescribePortRangeListAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) SetOwnerId(v int64) *DescribePortRangeListAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) SetPortRangeListId(v string) *DescribePortRangeListAssociationsRequest {
	s.PortRangeListId = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) SetRegionId(v string) *DescribePortRangeListAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) SetResourceOwnerAccount(v string) *DescribePortRangeListAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) SetResourceOwnerId(v int64) *DescribePortRangeListAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePortRangeListAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
