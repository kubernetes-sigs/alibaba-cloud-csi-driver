// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DetachNetworkInterfaceRequest
	GetInstanceId() *string
	SetNetworkInterfaceId(v string) *DetachNetworkInterfaceRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *DetachNetworkInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachNetworkInterfaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachNetworkInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachNetworkInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachNetworkInterfaceRequest
	GetResourceOwnerId() *int64
	SetTrunkNetworkInstanceId(v string) *DetachNetworkInterfaceRequest
	GetTrunkNetworkInstanceId() *string
}

type DetachNetworkInterfaceRequest struct {
	// The ID of the trunk ENI.
	//
	// >  This parameter is unavailable for use.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp67acfmxazb4p****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// eni-f8zapqwj1v1j4ia3****
	TrunkNetworkInstanceId *string `json:"TrunkNetworkInstanceId,omitempty" xml:"TrunkNetworkInstanceId,omitempty"`
}

func (s DetachNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DetachNetworkInterfaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachNetworkInterfaceRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DetachNetworkInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachNetworkInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachNetworkInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachNetworkInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachNetworkInterfaceRequest) GetTrunkNetworkInstanceId() *string {
	return s.TrunkNetworkInstanceId
}

func (s *DetachNetworkInterfaceRequest) SetInstanceId(v string) *DetachNetworkInterfaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) SetNetworkInterfaceId(v string) *DetachNetworkInterfaceRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) SetOwnerAccount(v string) *DetachNetworkInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) SetOwnerId(v int64) *DetachNetworkInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) SetRegionId(v string) *DetachNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) SetResourceOwnerAccount(v string) *DetachNetworkInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) SetResourceOwnerId(v int64) *DetachNetworkInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) SetTrunkNetworkInstanceId(v string) *DetachNetworkInterfaceRequest {
	s.TrunkNetworkInstanceId = &v
	return s
}

func (s *DetachNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
