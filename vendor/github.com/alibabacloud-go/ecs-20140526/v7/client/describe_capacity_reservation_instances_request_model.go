// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *DescribeCapacityReservationInstancesRequestPrivatePoolOptions) *DescribeCapacityReservationInstancesRequest
	GetPrivatePoolOptions() *DescribeCapacityReservationInstancesRequestPrivatePoolOptions
	SetMaxResults(v int32) *DescribeCapacityReservationInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCapacityReservationInstancesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeCapacityReservationInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCapacityReservationInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCapacityReservationInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCapacityReservationInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCapacityReservationInstancesRequest
	GetResourceOwnerId() *int64
}

type DescribeCapacityReservationInstancesRequest struct {
	PrivatePoolOptions *DescribeCapacityReservationInstancesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The number of entries per page.
	//
	// 	- Valid values: 1 to 100.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of the NextToken parameter.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the capacity reservation. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeCapacityReservationInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationInstancesRequest) GetPrivatePoolOptions() *DescribeCapacityReservationInstancesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeCapacityReservationInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCapacityReservationInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCapacityReservationInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCapacityReservationInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCapacityReservationInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCapacityReservationInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCapacityReservationInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCapacityReservationInstancesRequest) SetPrivatePoolOptions(v *DescribeCapacityReservationInstancesRequestPrivatePoolOptions) *DescribeCapacityReservationInstancesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) SetMaxResults(v int32) *DescribeCapacityReservationInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) SetNextToken(v string) *DescribeCapacityReservationInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) SetOwnerAccount(v string) *DescribeCapacityReservationInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) SetOwnerId(v int64) *DescribeCapacityReservationInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) SetRegionId(v string) *DescribeCapacityReservationInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) SetResourceOwnerAccount(v string) *DescribeCapacityReservationInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) SetResourceOwnerId(v int64) *DescribeCapacityReservationInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCapacityReservationInstancesRequestPrivatePoolOptions struct {
	// The ID of the capacity reservation.
	//
	// This parameter is required.
	//
	// example:
	//
	// crp-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeCapacityReservationInstancesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationInstancesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationInstancesRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *DescribeCapacityReservationInstancesRequestPrivatePoolOptions) SetId(v string) *DescribeCapacityReservationInstancesRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *DescribeCapacityReservationInstancesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
