// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyVSwitchAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyVSwitchAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVSwitchAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVSwitchAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVSwitchAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVSwitchAttributeRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchId() *string
	SetVSwitchName(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchName() *string
}

type ModifyVSwitchAttributeRequest struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s ModifyVSwitchAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVSwitchAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVSwitchAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVSwitchAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVSwitchAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVSwitchAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *ModifyVSwitchAttributeRequest) SetDescription(v string) *ModifyVSwitchAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetOwnerAccount(v string) *ModifyVSwitchAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetOwnerId(v int64) *ModifyVSwitchAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetRegionId(v string) *ModifyVSwitchAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVSwitchAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetResourceOwnerId(v int64) *ModifyVSwitchAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchId(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchName(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchName = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) Validate() error {
	return dara.Validate(s)
}
