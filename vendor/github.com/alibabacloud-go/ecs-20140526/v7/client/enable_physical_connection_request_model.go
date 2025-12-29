// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePhysicalConnectionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnablePhysicalConnectionRequest
  GetClientToken() *string 
  SetOwnerAccount(v string) *EnablePhysicalConnectionRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnablePhysicalConnectionRequest
  GetOwnerId() *int64 
  SetPhysicalConnectionId(v string) *EnablePhysicalConnectionRequest
  GetPhysicalConnectionId() *string 
  SetRegionId(v string) *EnablePhysicalConnectionRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnablePhysicalConnectionRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnablePhysicalConnectionRequest
  GetResourceOwnerId() *int64 
  SetUserCidr(v string) *EnablePhysicalConnectionRequest
  GetUserCidr() *string 
}

type EnablePhysicalConnectionRequest struct {
  // This parameter is required.
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // This parameter is required.
  PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
  // This parameter is required.
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s EnablePhysicalConnectionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnablePhysicalConnectionRequest) GoString() string {
  return s.String()
}

func (s *EnablePhysicalConnectionRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnablePhysicalConnectionRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnablePhysicalConnectionRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnablePhysicalConnectionRequest) GetPhysicalConnectionId() *string  {
  return s.PhysicalConnectionId
}

func (s *EnablePhysicalConnectionRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnablePhysicalConnectionRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnablePhysicalConnectionRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnablePhysicalConnectionRequest) GetUserCidr() *string  {
  return s.UserCidr
}

func (s *EnablePhysicalConnectionRequest) SetClientToken(v string) *EnablePhysicalConnectionRequest {
  s.ClientToken = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetOwnerAccount(v string) *EnablePhysicalConnectionRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetOwnerId(v int64) *EnablePhysicalConnectionRequest {
  s.OwnerId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetPhysicalConnectionId(v string) *EnablePhysicalConnectionRequest {
  s.PhysicalConnectionId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetRegionId(v string) *EnablePhysicalConnectionRequest {
  s.RegionId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *EnablePhysicalConnectionRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetResourceOwnerId(v int64) *EnablePhysicalConnectionRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) SetUserCidr(v string) *EnablePhysicalConnectionRequest {
  s.UserCidr = &v
  return s
}

func (s *EnablePhysicalConnectionRequest) Validate() error {
  return dara.Validate(s)
}

