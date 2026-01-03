// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVncUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceVncUrlRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceVncUrlRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceVncUrlRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceVncUrlRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceVncUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceVncUrlRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceVncUrlRequest struct {
	// The ID of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1hzoinajzkh91h****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
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

func (s DescribeInstanceVncUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVncUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceVncUrlRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceVncUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceVncUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceVncUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceVncUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceVncUrlRequest) SetInstanceId(v string) *DescribeInstanceVncUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetOwnerAccount(v string) *DescribeInstanceVncUrlRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetOwnerId(v int64) *DescribeInstanceVncUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetRegionId(v string) *DescribeInstanceVncUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetResourceOwnerAccount(v string) *DescribeInstanceVncUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetResourceOwnerId(v int64) *DescribeInstanceVncUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) Validate() error {
	return dara.Validate(s)
}
