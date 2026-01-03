// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterManagedInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeregisterManagedInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeregisterManagedInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeregisterManagedInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeregisterManagedInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeregisterManagedInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeregisterManagedInstanceRequest
	GetResourceOwnerId() *int64
}

type DeregisterManagedInstanceRequest struct {
	// The managed instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mi-hz01axdfas****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), China (Hong Kong), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia).
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DeregisterManagedInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeregisterManagedInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeregisterManagedInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeregisterManagedInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeregisterManagedInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeregisterManagedInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeregisterManagedInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeregisterManagedInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeregisterManagedInstanceRequest) SetInstanceId(v string) *DeregisterManagedInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeregisterManagedInstanceRequest) SetOwnerAccount(v string) *DeregisterManagedInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeregisterManagedInstanceRequest) SetOwnerId(v int64) *DeregisterManagedInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeregisterManagedInstanceRequest) SetRegionId(v string) *DeregisterManagedInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeregisterManagedInstanceRequest) SetResourceOwnerAccount(v string) *DeregisterManagedInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeregisterManagedInstanceRequest) SetResourceOwnerId(v int64) *DeregisterManagedInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeregisterManagedInstanceRequest) Validate() error {
	return dara.Validate(s)
}
