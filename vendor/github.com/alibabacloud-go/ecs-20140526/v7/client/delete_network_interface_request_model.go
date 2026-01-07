// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *DeleteNetworkInterfaceRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *DeleteNetworkInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNetworkInterfaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNetworkInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNetworkInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNetworkInterfaceRequest
	GetResourceOwnerId() *int64
}

type DeleteNetworkInterfaceRequest struct {
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

func (s DeleteNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfaceRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DeleteNetworkInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNetworkInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNetworkInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNetworkInterfaceRequest) SetNetworkInterfaceId(v string) *DeleteNetworkInterfaceRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DeleteNetworkInterfaceRequest) SetOwnerAccount(v string) *DeleteNetworkInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNetworkInterfaceRequest) SetOwnerId(v int64) *DeleteNetworkInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNetworkInterfaceRequest) SetRegionId(v string) *DeleteNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkInterfaceRequest) SetResourceOwnerAccount(v string) *DeleteNetworkInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNetworkInterfaceRequest) SetResourceOwnerId(v int64) *DeleteNetworkInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
