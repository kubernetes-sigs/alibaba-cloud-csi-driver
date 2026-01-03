// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssuranceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions) *DescribeElasticityAssuranceInstancesRequest
	GetPrivatePoolOptions() *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions
	SetMaxResults(v int32) *DescribeElasticityAssuranceInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeElasticityAssuranceInstancesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeElasticityAssuranceInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeElasticityAssuranceInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeElasticityAssuranceInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeElasticityAssuranceInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeElasticityAssuranceInstancesRequest
	GetResourceOwnerId() *int64
}

type DescribeElasticityAssuranceInstancesRequest struct {
	PrivatePoolOptions *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The number of entries to return on each page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the query. Set the value to the NextToken value obtained from the response to the preceding request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the elasticity assurance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeElasticityAssuranceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetPrivatePoolOptions() *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeElasticityAssuranceInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetPrivatePoolOptions(v *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions) *DescribeElasticityAssuranceInstancesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetMaxResults(v int32) *DescribeElasticityAssuranceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetNextToken(v string) *DescribeElasticityAssuranceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetOwnerAccount(v string) *DescribeElasticityAssuranceInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetOwnerId(v int64) *DescribeElasticityAssuranceInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetRegionId(v string) *DescribeElasticityAssuranceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetResourceOwnerAccount(v string) *DescribeElasticityAssuranceInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) SetResourceOwnerId(v int64) *DescribeElasticityAssuranceInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions struct {
	// The ID of the elasticity assurance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions) SetId(v string) *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
