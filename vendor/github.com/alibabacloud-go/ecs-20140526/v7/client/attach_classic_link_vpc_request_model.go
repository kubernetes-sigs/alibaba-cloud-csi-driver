// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachClassicLinkVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AttachClassicLinkVpcRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *AttachClassicLinkVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachClassicLinkVpcRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachClassicLinkVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachClassicLinkVpcRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *AttachClassicLinkVpcRequest
	GetVpcId() *string
}

type AttachClassicLinkVpcRequest struct {
	// The ID of the instance that is deployed in the classic network. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/25506.html) operation to query available instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1gtjxuuvwj17zr****
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
	// The ID of the VPC for which the ClassicLink feature is enabled. You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query available VPCs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1j4z1sr8zxu4l8u****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AttachClassicLinkVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachClassicLinkVpcRequest) GoString() string {
	return s.String()
}

func (s *AttachClassicLinkVpcRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachClassicLinkVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachClassicLinkVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachClassicLinkVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachClassicLinkVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachClassicLinkVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AttachClassicLinkVpcRequest) SetInstanceId(v string) *AttachClassicLinkVpcRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachClassicLinkVpcRequest) SetOwnerId(v int64) *AttachClassicLinkVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachClassicLinkVpcRequest) SetRegionId(v string) *AttachClassicLinkVpcRequest {
	s.RegionId = &v
	return s
}

func (s *AttachClassicLinkVpcRequest) SetResourceOwnerAccount(v string) *AttachClassicLinkVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachClassicLinkVpcRequest) SetResourceOwnerId(v int64) *AttachClassicLinkVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachClassicLinkVpcRequest) SetVpcId(v string) *AttachClassicLinkVpcRequest {
	s.VpcId = &v
	return s
}

func (s *AttachClassicLinkVpcRequest) Validate() error {
	return dara.Validate(s)
}
