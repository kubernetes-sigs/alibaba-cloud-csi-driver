// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHpcClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateHpcClusterRequest
	GetClientToken() *string
	SetDescription(v string) *CreateHpcClusterRequest
	GetDescription() *string
	SetName(v string) *CreateHpcClusterRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateHpcClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateHpcClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateHpcClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateHpcClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateHpcClusterRequest
	GetResourceOwnerId() *int64
}

type CreateHpcClusterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a client token. Make sure that a unique client token is used for each request. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the HPC cluster. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// testHPCDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the HPC cluster. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// hpc-Cluster-01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// RAM用户的账号登录名称。
	//
	// example:
	//
	// ECSforCloud@Alibaba.com
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// RAM用户的虚拟账号ID。
	//
	// example:
	//
	// 1234567890
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the HPC cluster. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
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
	// 1234567890
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateHpcClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHpcClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateHpcClusterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateHpcClusterRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHpcClusterRequest) GetName() *string {
	return s.Name
}

func (s *CreateHpcClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateHpcClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateHpcClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHpcClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateHpcClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateHpcClusterRequest) SetClientToken(v string) *CreateHpcClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHpcClusterRequest) SetDescription(v string) *CreateHpcClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateHpcClusterRequest) SetName(v string) *CreateHpcClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateHpcClusterRequest) SetOwnerAccount(v string) *CreateHpcClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateHpcClusterRequest) SetOwnerId(v int64) *CreateHpcClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHpcClusterRequest) SetRegionId(v string) *CreateHpcClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHpcClusterRequest) SetResourceOwnerAccount(v string) *CreateHpcClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHpcClusterRequest) SetResourceOwnerId(v int64) *CreateHpcClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHpcClusterRequest) Validate() error {
	return dara.Validate(s)
}
