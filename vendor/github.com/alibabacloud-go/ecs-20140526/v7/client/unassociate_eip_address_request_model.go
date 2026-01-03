// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *UnassociateEipAddressRequest
	GetAllocationId() *string
	SetInstanceId(v string) *UnassociateEipAddressRequest
	GetInstanceId() *string
	SetInstanceType(v string) *UnassociateEipAddressRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *UnassociateEipAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateEipAddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassociateEipAddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociateEipAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateEipAddressRequest
	GetResourceOwnerId() *int64
}

type UnassociateEipAddressRequest struct {
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

func (s UnassociateEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *UnassociateEipAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnassociateEipAddressRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UnassociateEipAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateEipAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateEipAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateEipAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateEipAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateEipAddressRequest) SetAllocationId(v string) *UnassociateEipAddressRequest {
	s.AllocationId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetInstanceId(v string) *UnassociateEipAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetInstanceType(v string) *UnassociateEipAddressRequest {
	s.InstanceType = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetOwnerAccount(v string) *UnassociateEipAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetOwnerId(v int64) *UnassociateEipAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetRegionId(v string) *UnassociateEipAddressRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetResourceOwnerAccount(v string) *UnassociateEipAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateEipAddressRequest) SetResourceOwnerId(v int64) *UnassociateEipAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
