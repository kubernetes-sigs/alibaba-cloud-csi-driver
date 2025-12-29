// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssuranceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) *DescribeElasticityAssuranceAutoRenewAttributeRequest
	GetPrivatePoolOptions() *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions
	SetOwnerAccount(v string) *DescribeElasticityAssuranceAutoRenewAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeElasticityAssuranceAutoRenewAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeElasticityAssuranceAutoRenewAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeElasticityAssuranceAutoRenewAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeElasticityAssuranceAutoRenewAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeElasticityAssuranceAutoRenewAttributeRequest struct {
	PrivatePoolOptions *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	OwnerAccount       *string                                                                 `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64                                                                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DescribeElasticityAssuranceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) GetPrivatePoolOptions() *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) SetPrivatePoolOptions(v *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) *DescribeElasticityAssuranceAutoRenewAttributeRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeElasticityAssuranceAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeElasticityAssuranceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) SetRegionId(v string) *DescribeElasticityAssuranceAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeElasticityAssuranceAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeElasticityAssuranceAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions struct {
	// The IDs of elasticity assurances.
	//
	// **Limits**: You can specify up to 50 elasticity assurance IDs in a single request.
	Id []*string `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) GetId() []*string {
	return s.Id
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) SetId(v []*string) *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions {
	s.Id = v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
