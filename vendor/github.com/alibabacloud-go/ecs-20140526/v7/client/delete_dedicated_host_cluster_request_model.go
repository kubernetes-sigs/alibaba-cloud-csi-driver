// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDedicatedHostClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusterId(v string) *DeleteDedicatedHostClusterRequest
	GetDedicatedHostClusterId() *string
	SetOwnerAccount(v string) *DeleteDedicatedHostClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDedicatedHostClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDedicatedHostClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteDedicatedHostClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDedicatedHostClusterRequest
	GetResourceOwnerId() *int64
}

type DeleteDedicatedHostClusterRequest struct {
	// The ID of the host group.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-bp12wlf6am0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the host group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DeleteDedicatedHostClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDedicatedHostClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostClusterRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *DeleteDedicatedHostClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDedicatedHostClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDedicatedHostClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDedicatedHostClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDedicatedHostClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDedicatedHostClusterRequest) SetDedicatedHostClusterId(v string) *DeleteDedicatedHostClusterRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *DeleteDedicatedHostClusterRequest) SetOwnerAccount(v string) *DeleteDedicatedHostClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDedicatedHostClusterRequest) SetOwnerId(v int64) *DeleteDedicatedHostClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDedicatedHostClusterRequest) SetRegionId(v string) *DeleteDedicatedHostClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDedicatedHostClusterRequest) SetResourceOwnerAccount(v string) *DeleteDedicatedHostClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDedicatedHostClusterRequest) SetResourceOwnerId(v int64) *DeleteDedicatedHostClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDedicatedHostClusterRequest) Validate() error {
	return dara.Validate(s)
}
