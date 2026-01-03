// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypeFamiliesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGeneration(v string) *DescribeInstanceTypeFamiliesRequest
	GetGeneration() *string
	SetOwnerAccount(v string) *DescribeInstanceTypeFamiliesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceTypeFamiliesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceTypeFamiliesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceTypeFamiliesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceTypeFamiliesRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceTypeFamiliesRequest struct {
	// The series of the instance family. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html). Valid values:
	//
	// 	- ecs-1: Series I, which consists of the earliest and cost-effective instance types
	//
	// 	- ecs-2: Series II, which provides upgraded software and hardware and higher performance than Series I
	//
	// 	- ecs-3: Series III, which consists of high-performance instance families and is suitable for different business scenarios
	//
	// 	- ecs-4: Series IV, which consists of enterprise-level instance families (such as g5, c5, and r5), ECS Bare Metal Instance families (such as ebmc5s, ebmg5s, and ebmr5s), and burstable instance families (such as t5) and can meet a wide variety of business requirements with lower latency
	//
	// 	- ecs-5: Series V, which consists of enterprise-level instance families (such as g6, c6, and r6), ECS Bare Metal Instance families (such as ebmg6, ebmg6e, and ebmc6), and storage-enhanced instance families (such as g6e) and delivers quick response times and higher performance
	//
	// 	- ecs-6: Series VI, which consists of enterprise-level instance families (such as hfc7, hfg7, and hfr7) and ECS Bare Metal Instance families (such as ebmhfg7)
	//
	// example:
	//
	// ecs-5
	Generation   *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance family. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeInstanceTypeFamiliesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeFamiliesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeFamiliesRequest) GetGeneration() *string {
	return s.Generation
}

func (s *DescribeInstanceTypeFamiliesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceTypeFamiliesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceTypeFamiliesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceTypeFamiliesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceTypeFamiliesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceTypeFamiliesRequest) SetGeneration(v string) *DescribeInstanceTypeFamiliesRequest {
	s.Generation = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesRequest) SetOwnerAccount(v string) *DescribeInstanceTypeFamiliesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesRequest) SetOwnerId(v int64) *DescribeInstanceTypeFamiliesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesRequest) SetRegionId(v string) *DescribeInstanceTypeFamiliesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesRequest) SetResourceOwnerAccount(v string) *DescribeInstanceTypeFamiliesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesRequest) SetResourceOwnerId(v int64) *DescribeInstanceTypeFamiliesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesRequest) Validate() error {
	return dara.Validate(s)
}
