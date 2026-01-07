// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusterId(v string) *ModifyDedicatedHostClusterAttributeRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostClusterName(v string) *ModifyDedicatedHostClusterAttributeRequest
	GetDedicatedHostClusterName() *string
	SetDescription(v string) *ModifyDedicatedHostClusterAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyDedicatedHostClusterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDedicatedHostClusterAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDedicatedHostClusterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDedicatedHostClusterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDedicatedHostClusterAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDedicatedHostClusterAttributeRequest struct {
	// The ID of the host group.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-bp12wlf6am0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The name of the host group. It must be 2 to 128 characters in length and start with a letter. It can contain letters, digits, periods (.), underscores (_), and hyphens (-), and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// newClusterName
	DedicatedHostClusterName *string `json:"DedicatedHostClusterName,omitempty" xml:"DedicatedHostClusterName,omitempty"`
	// The description of the host group. It must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newClusterDescription
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s ModifyDedicatedHostClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetDedicatedHostClusterName() *string {
	return s.DedicatedHostClusterName
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDedicatedHostClusterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetDedicatedHostClusterId(v string) *ModifyDedicatedHostClusterAttributeRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetDedicatedHostClusterName(v string) *ModifyDedicatedHostClusterAttributeRequest {
	s.DedicatedHostClusterName = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetDescription(v string) *ModifyDedicatedHostClusterAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetOwnerAccount(v string) *ModifyDedicatedHostClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetOwnerId(v int64) *ModifyDedicatedHostClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetRegionId(v string) *ModifyDedicatedHostClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
