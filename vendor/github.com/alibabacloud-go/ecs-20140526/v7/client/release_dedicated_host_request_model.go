// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDedicatedHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostId(v string) *ReleaseDedicatedHostRequest
	GetDedicatedHostId() *string
	SetOwnerAccount(v string) *ReleaseDedicatedHostRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseDedicatedHostRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReleaseDedicatedHostRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReleaseDedicatedHostRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseDedicatedHostRequest
	GetResourceOwnerId() *int64
	SetTerminateSubscription(v bool) *ReleaseDedicatedHostRequest
	GetTerminateSubscription() *bool
}

type ReleaseDedicatedHostRequest struct {
	// The ID of the dedicated host.
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp199lyny9b3****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the dedicated host. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The expiration time of the subscription dedicated host.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	TerminateSubscription *bool `json:"TerminateSubscription,omitempty" xml:"TerminateSubscription,omitempty"`
}

func (s ReleaseDedicatedHostRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *ReleaseDedicatedHostRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *ReleaseDedicatedHostRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseDedicatedHostRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseDedicatedHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseDedicatedHostRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseDedicatedHostRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseDedicatedHostRequest) GetTerminateSubscription() *bool {
	return s.TerminateSubscription
}

func (s *ReleaseDedicatedHostRequest) SetDedicatedHostId(v string) *ReleaseDedicatedHostRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ReleaseDedicatedHostRequest) SetOwnerAccount(v string) *ReleaseDedicatedHostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseDedicatedHostRequest) SetOwnerId(v int64) *ReleaseDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseDedicatedHostRequest) SetRegionId(v string) *ReleaseDedicatedHostRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseDedicatedHostRequest) SetResourceOwnerAccount(v string) *ReleaseDedicatedHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseDedicatedHostRequest) SetResourceOwnerId(v int64) *ReleaseDedicatedHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseDedicatedHostRequest) SetTerminateSubscription(v bool) *ReleaseDedicatedHostRequest {
	s.TerminateSubscription = &v
	return s
}

func (s *ReleaseDedicatedHostRequest) Validate() error {
	return dara.Validate(s)
}
