// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndTerminalSessionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOwnerAccount(v string) *EndTerminalSessionRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EndTerminalSessionRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EndTerminalSessionRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EndTerminalSessionRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EndTerminalSessionRequest
  GetResourceOwnerId() *int64 
  SetSessionId(v string) *EndTerminalSessionRequest
  GetSessionId() *string 
}

type EndTerminalSessionRequest struct {
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the session.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The session ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // s-hz023od0x9****
  SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EndTerminalSessionRequest) String() string {
  return dara.Prettify(s)
}

func (s EndTerminalSessionRequest) GoString() string {
  return s.String()
}

func (s *EndTerminalSessionRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EndTerminalSessionRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EndTerminalSessionRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EndTerminalSessionRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EndTerminalSessionRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EndTerminalSessionRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EndTerminalSessionRequest) SetOwnerAccount(v string) *EndTerminalSessionRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EndTerminalSessionRequest) SetOwnerId(v int64) *EndTerminalSessionRequest {
  s.OwnerId = &v
  return s
}

func (s *EndTerminalSessionRequest) SetRegionId(v string) *EndTerminalSessionRequest {
  s.RegionId = &v
  return s
}

func (s *EndTerminalSessionRequest) SetResourceOwnerAccount(v string) *EndTerminalSessionRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EndTerminalSessionRequest) SetResourceOwnerId(v int64) *EndTerminalSessionRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EndTerminalSessionRequest) SetSessionId(v string) *EndTerminalSessionRequest {
  s.SessionId = &v
  return s
}

func (s *EndTerminalSessionRequest) Validate() error {
  return dara.Validate(s)
}

