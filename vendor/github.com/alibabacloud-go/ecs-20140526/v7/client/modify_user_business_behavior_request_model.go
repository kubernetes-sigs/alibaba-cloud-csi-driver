// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserBusinessBehaviorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifyUserBusinessBehaviorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyUserBusinessBehaviorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyUserBusinessBehaviorRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyUserBusinessBehaviorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyUserBusinessBehaviorRequest
	GetResourceOwnerId() *int64
	SetStatusKey(v string) *ModifyUserBusinessBehaviorRequest
	GetStatusKey() *string
	SetStatusValue(v string) *ModifyUserBusinessBehaviorRequest
	GetStatusValue() *string
}

type ModifyUserBusinessBehaviorRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StatusKey *string `json:"statusKey,omitempty" xml:"statusKey,omitempty"`
	// This parameter is required.
	StatusValue *string `json:"statusValue,omitempty" xml:"statusValue,omitempty"`
}

func (s ModifyUserBusinessBehaviorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserBusinessBehaviorRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserBusinessBehaviorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyUserBusinessBehaviorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyUserBusinessBehaviorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserBusinessBehaviorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyUserBusinessBehaviorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyUserBusinessBehaviorRequest) GetStatusKey() *string {
	return s.StatusKey
}

func (s *ModifyUserBusinessBehaviorRequest) GetStatusValue() *string {
	return s.StatusValue
}

func (s *ModifyUserBusinessBehaviorRequest) SetOwnerAccount(v string) *ModifyUserBusinessBehaviorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyUserBusinessBehaviorRequest) SetOwnerId(v int64) *ModifyUserBusinessBehaviorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUserBusinessBehaviorRequest) SetRegionId(v string) *ModifyUserBusinessBehaviorRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserBusinessBehaviorRequest) SetResourceOwnerAccount(v string) *ModifyUserBusinessBehaviorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUserBusinessBehaviorRequest) SetResourceOwnerId(v int64) *ModifyUserBusinessBehaviorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUserBusinessBehaviorRequest) SetStatusKey(v string) *ModifyUserBusinessBehaviorRequest {
	s.StatusKey = &v
	return s
}

func (s *ModifyUserBusinessBehaviorRequest) SetStatusValue(v string) *ModifyUserBusinessBehaviorRequest {
	s.StatusValue = &v
	return s
}

func (s *ModifyUserBusinessBehaviorRequest) Validate() error {
	return dara.Validate(s)
}
