// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyReservedInstanceAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyReservedInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyReservedInstanceAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyReservedInstanceAttributeRequest
	GetRegionId() *string
	SetReservedInstanceId(v string) *ModifyReservedInstanceAttributeRequest
	GetReservedInstanceId() *string
	SetReservedInstanceName(v string) *ModifyReservedInstanceAttributeRequest
	GetReservedInstanceName() *string
	SetResourceOwnerAccount(v string) *ModifyReservedInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyReservedInstanceAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyReservedInstanceAttributeRequest struct {
	// The error code.
	//
	// example:
	//
	// ri-example
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the reserved instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The new name of the reserved instance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecsri-uf61hdhue4kcorqsk****
	ReservedInstanceId *string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty"`
	// The new description of the reserved instance. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// testReservedInstanceName
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" xml:"ReservedInstanceName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyReservedInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstanceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyReservedInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyReservedInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyReservedInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyReservedInstanceAttributeRequest) GetReservedInstanceId() *string {
	return s.ReservedInstanceId
}

func (s *ModifyReservedInstanceAttributeRequest) GetReservedInstanceName() *string {
	return s.ReservedInstanceName
}

func (s *ModifyReservedInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyReservedInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyReservedInstanceAttributeRequest) SetDescription(v string) *ModifyReservedInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) SetOwnerAccount(v string) *ModifyReservedInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) SetOwnerId(v int64) *ModifyReservedInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) SetRegionId(v string) *ModifyReservedInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) SetReservedInstanceId(v string) *ModifyReservedInstanceAttributeRequest {
	s.ReservedInstanceId = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) SetReservedInstanceName(v string) *ModifyReservedInstanceAttributeRequest {
	s.ReservedInstanceName = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyReservedInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyReservedInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyReservedInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
