// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeUserDataRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DescribeUserDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUserDataRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUserDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUserDataRequest
	GetResourceOwnerId() *int64
}

type DescribeUserDataRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp14bnftyqhxg9ij****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeUserDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUserDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserDataRequest) SetInstanceId(v string) *DescribeUserDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserDataRequest) SetOwnerId(v int64) *DescribeUserDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserDataRequest) SetRegionId(v string) *DescribeUserDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserDataRequest) SetResourceOwnerAccount(v string) *DescribeUserDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserDataRequest) SetResourceOwnerId(v int64) *DescribeUserDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserDataRequest) Validate() error {
	return dara.Validate(s)
}
