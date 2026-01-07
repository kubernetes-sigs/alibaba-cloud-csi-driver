// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentSetSupportedInstanceTypeFamilyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest
	GetResourceOwnerId() *int64
	SetStrategy(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest
	GetStrategy() *string
}

type DescribeDeploymentSetSupportedInstanceTypeFamilyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The deployment strategy. Valid values:
	//
	// 	- Availability: high availability strategy
	//
	// 	- AvailabilityGroup: high availability group strategy
	//
	// 	- LowLatency: low latency strategy
	//
	// Default value: Availability.
	//
	// example:
	//
	// Availability
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) SetOwnerAccount(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) SetOwnerId(v int64) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) SetRegionId(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) SetResourceOwnerAccount(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) SetResourceOwnerId(v int64) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) SetStrategy(v string) *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest {
	s.Strategy = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyRequest) Validate() error {
	return dara.Validate(s)
}
