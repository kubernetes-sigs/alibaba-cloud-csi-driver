// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReActivateInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReActivateInstancesRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ReActivateInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReActivateInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReActivateInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReActivateInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReActivateInstancesRequest
	GetResourceOwnerId() *int64
}

type ReActivateInstancesRequest struct {
	// The ID of the instance that you want to reactivate.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReActivateInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ReActivateInstancesRequest) GoString() string {
	return s.String()
}

func (s *ReActivateInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReActivateInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReActivateInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReActivateInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReActivateInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReActivateInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReActivateInstancesRequest) SetInstanceId(v string) *ReActivateInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ReActivateInstancesRequest) SetOwnerAccount(v string) *ReActivateInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReActivateInstancesRequest) SetOwnerId(v int64) *ReActivateInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *ReActivateInstancesRequest) SetRegionId(v string) *ReActivateInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ReActivateInstancesRequest) SetResourceOwnerAccount(v string) *ReActivateInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReActivateInstancesRequest) SetResourceOwnerId(v int64) *ReActivateInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReActivateInstancesRequest) Validate() error {
	return dara.Validate(s)
}
