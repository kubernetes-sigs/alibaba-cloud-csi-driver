// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNetworkInterfaceQoSRequest interface {
  dara.Model
  String() string
  GoString() string
  SetNetworkInterfaceId(v string) *EnableNetworkInterfaceQoSRequest
  GetNetworkInterfaceId() *string 
  SetOwnerAccount(v string) *EnableNetworkInterfaceQoSRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableNetworkInterfaceQoSRequest
  GetOwnerId() *int64 
  SetQoS(v *EnableNetworkInterfaceQoSRequestQoS) *EnableNetworkInterfaceQoSRequest
  GetQoS() *EnableNetworkInterfaceQoSRequestQoS 
  SetRegionId(v string) *EnableNetworkInterfaceQoSRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableNetworkInterfaceQoSRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableNetworkInterfaceQoSRequest
  GetResourceOwnerId() *int64 
}

type EnableNetworkInterfaceQoSRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // eni-2zeh9atclduxvf1z****
  NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  QoS *EnableNetworkInterfaceQoSRequestQoS `json:"QoS,omitempty" xml:"QoS,omitempty" type:"Struct"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableNetworkInterfaceQoSRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableNetworkInterfaceQoSRequest) GoString() string {
  return s.String()
}

func (s *EnableNetworkInterfaceQoSRequest) GetNetworkInterfaceId() *string  {
  return s.NetworkInterfaceId
}

func (s *EnableNetworkInterfaceQoSRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableNetworkInterfaceQoSRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableNetworkInterfaceQoSRequest) GetQoS() *EnableNetworkInterfaceQoSRequestQoS  {
  return s.QoS
}

func (s *EnableNetworkInterfaceQoSRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableNetworkInterfaceQoSRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableNetworkInterfaceQoSRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableNetworkInterfaceQoSRequest) SetNetworkInterfaceId(v string) *EnableNetworkInterfaceQoSRequest {
  s.NetworkInterfaceId = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequest) SetOwnerAccount(v string) *EnableNetworkInterfaceQoSRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequest) SetOwnerId(v int64) *EnableNetworkInterfaceQoSRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequest) SetQoS(v *EnableNetworkInterfaceQoSRequestQoS) *EnableNetworkInterfaceQoSRequest {
  s.QoS = v
  return s
}

func (s *EnableNetworkInterfaceQoSRequest) SetRegionId(v string) *EnableNetworkInterfaceQoSRequest {
  s.RegionId = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequest) SetResourceOwnerAccount(v string) *EnableNetworkInterfaceQoSRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequest) SetResourceOwnerId(v int64) *EnableNetworkInterfaceQoSRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequest) Validate() error {
  if s.QoS != nil {
    if err := s.QoS.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableNetworkInterfaceQoSRequestQoS struct {
  // example:
  // 
  // 50000
  BandwidthRx *int64 `json:"BandwidthRx,omitempty" xml:"BandwidthRx,omitempty"`
  // example:
  // 
  // 50000
  BandwidthTx *int64 `json:"BandwidthTx,omitempty" xml:"BandwidthTx,omitempty"`
  // example:
  // 
  // 50000
  ConcurrentConnections *int64 `json:"ConcurrentConnections,omitempty" xml:"ConcurrentConnections,omitempty"`
  // example:
  // 
  // 50000
  PpsRx *int64 `json:"PpsRx,omitempty" xml:"PpsRx,omitempty"`
  // example:
  // 
  // 50000
  PpsTx *int64 `json:"PpsTx,omitempty" xml:"PpsTx,omitempty"`
}

func (s EnableNetworkInterfaceQoSRequestQoS) String() string {
  return dara.Prettify(s)
}

func (s EnableNetworkInterfaceQoSRequestQoS) GoString() string {
  return s.String()
}

func (s *EnableNetworkInterfaceQoSRequestQoS) GetBandwidthRx() *int64  {
  return s.BandwidthRx
}

func (s *EnableNetworkInterfaceQoSRequestQoS) GetBandwidthTx() *int64  {
  return s.BandwidthTx
}

func (s *EnableNetworkInterfaceQoSRequestQoS) GetConcurrentConnections() *int64  {
  return s.ConcurrentConnections
}

func (s *EnableNetworkInterfaceQoSRequestQoS) GetPpsRx() *int64  {
  return s.PpsRx
}

func (s *EnableNetworkInterfaceQoSRequestQoS) GetPpsTx() *int64  {
  return s.PpsTx
}

func (s *EnableNetworkInterfaceQoSRequestQoS) SetBandwidthRx(v int64) *EnableNetworkInterfaceQoSRequestQoS {
  s.BandwidthRx = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequestQoS) SetBandwidthTx(v int64) *EnableNetworkInterfaceQoSRequestQoS {
  s.BandwidthTx = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequestQoS) SetConcurrentConnections(v int64) *EnableNetworkInterfaceQoSRequestQoS {
  s.ConcurrentConnections = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequestQoS) SetPpsRx(v int64) *EnableNetworkInterfaceQoSRequestQoS {
  s.PpsRx = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequestQoS) SetPpsTx(v int64) *EnableNetworkInterfaceQoSRequestQoS {
  s.PpsTx = &v
  return s
}

func (s *EnableNetworkInterfaceQoSRequestQoS) Validate() error {
  return dara.Validate(s)
}

