// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHpcClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyHpcClusterAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyHpcClusterAttributeRequest
	GetDescription() *string
	SetHpcClusterId(v string) *ModifyHpcClusterAttributeRequest
	GetHpcClusterId() *string
	SetName(v string) *ModifyHpcClusterAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyHpcClusterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyHpcClusterAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyHpcClusterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyHpcClusterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyHpcClusterAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyHpcClusterAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.*	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the HPC cluster. The description must be 2 to 256 characters in length, and cannot start with http:// or https://.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the HPC cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// hpc-b8bq705cvx1****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// The name of the HPC cluster. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// EcsforCloud@Alibaba.com
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// RAM用户的虚拟账号ID。
	//
	// example:
	//
	// 1234567890
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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
	// 1234567890
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyHpcClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHpcClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyHpcClusterAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyHpcClusterAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyHpcClusterAttributeRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *ModifyHpcClusterAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyHpcClusterAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyHpcClusterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyHpcClusterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHpcClusterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyHpcClusterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyHpcClusterAttributeRequest) SetClientToken(v string) *ModifyHpcClusterAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetDescription(v string) *ModifyHpcClusterAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetHpcClusterId(v string) *ModifyHpcClusterAttributeRequest {
	s.HpcClusterId = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetName(v string) *ModifyHpcClusterAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetOwnerAccount(v string) *ModifyHpcClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetOwnerId(v int64) *ModifyHpcClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetRegionId(v string) *ModifyHpcClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetResourceOwnerAccount(v string) *ModifyHpcClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) SetResourceOwnerId(v int64) *ModifyHpcClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHpcClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
