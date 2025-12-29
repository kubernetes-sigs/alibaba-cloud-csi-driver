// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DescribeInstanceTopologyRequest
	GetInstanceIds() *string
	SetOwnerId(v int64) *DescribeInstanceTopologyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceTopologyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceTopologyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceTopologyRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceTopologyRequest struct {
	// The IDs of one or more ECS instances. You can specify a maximum of 100 instance IDs.
	//
	// example:
	//
	// ["i-bp67acfmxazb4p****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ECS instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeInstanceTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstanceTopologyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceTopologyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceTopologyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceTopologyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceTopologyRequest) SetInstanceIds(v string) *DescribeInstanceTopologyRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstanceTopologyRequest) SetOwnerId(v int64) *DescribeInstanceTopologyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceTopologyRequest) SetRegionId(v string) *DescribeInstanceTopologyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceTopologyRequest) SetResourceOwnerAccount(v string) *DescribeInstanceTopologyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceTopologyRequest) SetResourceOwnerId(v int64) *DescribeInstanceTopologyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceTopologyRequest) Validate() error {
	return dara.Validate(s)
}
