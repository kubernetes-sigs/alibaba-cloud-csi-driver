// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHpcClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteHpcClusterRequest
	GetClientToken() *string
	SetHpcClusterId(v string) *DeleteHpcClusterRequest
	GetHpcClusterId() *string
	SetOwnerAccount(v string) *DeleteHpcClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteHpcClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteHpcClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteHpcClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteHpcClusterRequest
	GetResourceOwnerId() *int64
}

type DeleteHpcClusterRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.*	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the HPC cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// hpc-cxvr5uzy54j0ya****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// example:
	//
	// EcsforCloud@Alibaba.com
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// RAM用户的虚拟账号ID。
	//
	// example:
	//
	// 155780923770
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the HPC cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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
	// EcsforCloud
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// 资源主账号的ID，亦即UID。
	//
	// example:
	//
	// 155780923770
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteHpcClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHpcClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteHpcClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteHpcClusterRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *DeleteHpcClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteHpcClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteHpcClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHpcClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteHpcClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteHpcClusterRequest) SetClientToken(v string) *DeleteHpcClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteHpcClusterRequest) SetHpcClusterId(v string) *DeleteHpcClusterRequest {
	s.HpcClusterId = &v
	return s
}

func (s *DeleteHpcClusterRequest) SetOwnerAccount(v string) *DeleteHpcClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteHpcClusterRequest) SetOwnerId(v int64) *DeleteHpcClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteHpcClusterRequest) SetRegionId(v string) *DeleteHpcClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHpcClusterRequest) SetResourceOwnerAccount(v string) *DeleteHpcClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteHpcClusterRequest) SetResourceOwnerId(v int64) *DeleteHpcClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteHpcClusterRequest) Validate() error {
	return dara.Validate(s)
}
