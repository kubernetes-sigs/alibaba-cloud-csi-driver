// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfacePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *CreateNetworkInterfacePermissionRequest
	GetAccountId() *int64
	SetNetworkInterfaceId(v string) *CreateNetworkInterfacePermissionRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *CreateNetworkInterfacePermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNetworkInterfacePermissionRequest
	GetOwnerId() *int64
	SetPermission(v string) *CreateNetworkInterfacePermissionRequest
	GetPermission() *string
	SetRegionId(v string) *CreateNetworkInterfacePermissionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNetworkInterfacePermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNetworkInterfacePermissionRequest
	GetResourceOwnerId() *int64
}

type CreateNetworkInterfacePermissionRequest struct {
	// The ID of the Alibaba Cloud partner (a certified ISV) or individual user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp14v2sdd3v8htln****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The permission on the ENI. Valid values:
	//
	// InstanceAttach: the permission to attach the ENI to an ECS instance. The ENI and the ECS instance must be in the same zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceAttach
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The region ID of the ENI. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s CreateNetworkInterfacePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfacePermissionRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfacePermissionRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *CreateNetworkInterfacePermissionRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *CreateNetworkInterfacePermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNetworkInterfacePermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNetworkInterfacePermissionRequest) GetPermission() *string {
	return s.Permission
}

func (s *CreateNetworkInterfacePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkInterfacePermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNetworkInterfacePermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNetworkInterfacePermissionRequest) SetAccountId(v int64) *CreateNetworkInterfacePermissionRequest {
	s.AccountId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) SetNetworkInterfaceId(v string) *CreateNetworkInterfacePermissionRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) SetOwnerAccount(v string) *CreateNetworkInterfacePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) SetOwnerId(v int64) *CreateNetworkInterfacePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) SetPermission(v string) *CreateNetworkInterfacePermissionRequest {
	s.Permission = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) SetRegionId(v string) *CreateNetworkInterfacePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) SetResourceOwnerAccount(v string) *CreateNetworkInterfacePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) SetResourceOwnerId(v int64) *CreateNetworkInterfacePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNetworkInterfacePermissionRequest) Validate() error {
	return dara.Validate(s)
}
