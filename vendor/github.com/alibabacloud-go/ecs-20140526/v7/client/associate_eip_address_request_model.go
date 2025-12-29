// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *AssociateEipAddressRequest
	GetAllocationId() *string
	SetInstanceId(v string) *AssociateEipAddressRequest
	GetInstanceId() *string
	SetInstanceType(v string) *AssociateEipAddressRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *AssociateEipAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateEipAddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateEipAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateEipAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateEipAddressRequest
	GetResourceOwnerId() *int64
}

type AssociateEipAddressRequest struct {
	// This parameter is required.
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AssociateEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *AssociateEipAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateEipAddressRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AssociateEipAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateEipAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateEipAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateEipAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateEipAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateEipAddressRequest) SetAllocationId(v string) *AssociateEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetInstanceId(v string) *AssociateEipAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetInstanceType(v string) *AssociateEipAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *AssociateEipAddressRequest) SetOwnerAccount(v string) *AssociateEipAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateEipAddressRequest) SetOwnerId(v int64) *AssociateEipAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetRegionId(v string) *AssociateEipAddressRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateEipAddressRequest) SetResourceOwnerAccount(v string) *AssociateEipAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateEipAddressRequest) SetResourceOwnerId(v int64) *AssociateEipAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
