// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachClassicLinkVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DetachClassicLinkVpcRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DetachClassicLinkVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachClassicLinkVpcRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachClassicLinkVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachClassicLinkVpcRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DetachClassicLinkVpcRequest
	GetVpcId() *string
}

type DetachClassicLinkVpcRequest struct {
	// The ID of the instance that resides in the classic network.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// RAM用户的虚拟账号ID。
	//
	// example:
	//
	// 155780923770
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源主账号的账号名称。
	//
	// example:
	//
	// ECSforCloud
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// 资源主账号的ID，亦即UID。
	//
	// example:
	//
	// 155780923770
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC to which the instance is connected.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp67acfmxazb4p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DetachClassicLinkVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachClassicLinkVpcRequest) GoString() string {
	return s.String()
}

func (s *DetachClassicLinkVpcRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachClassicLinkVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachClassicLinkVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachClassicLinkVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachClassicLinkVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachClassicLinkVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DetachClassicLinkVpcRequest) SetInstanceId(v string) *DetachClassicLinkVpcRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachClassicLinkVpcRequest) SetOwnerId(v int64) *DetachClassicLinkVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachClassicLinkVpcRequest) SetRegionId(v string) *DetachClassicLinkVpcRequest {
	s.RegionId = &v
	return s
}

func (s *DetachClassicLinkVpcRequest) SetResourceOwnerAccount(v string) *DetachClassicLinkVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachClassicLinkVpcRequest) SetResourceOwnerId(v int64) *DetachClassicLinkVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachClassicLinkVpcRequest) SetVpcId(v string) *DetachClassicLinkVpcRequest {
	s.VpcId = &v
	return s
}

func (s *DetachClassicLinkVpcRequest) Validate() error {
	return dara.Validate(s)
}
