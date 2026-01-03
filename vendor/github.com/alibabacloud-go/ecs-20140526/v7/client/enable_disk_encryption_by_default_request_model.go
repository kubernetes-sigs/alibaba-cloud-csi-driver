// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDiskEncryptionByDefaultRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOwnerAccount(v string) *EnableDiskEncryptionByDefaultRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableDiskEncryptionByDefaultRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EnableDiskEncryptionByDefaultRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *EnableDiskEncryptionByDefaultRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableDiskEncryptionByDefaultRequest
  GetResourceOwnerId() *int64 
}

type EnableDiskEncryptionByDefaultRequest struct {
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableDiskEncryptionByDefaultRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDiskEncryptionByDefaultRequest) GoString() string {
  return s.String()
}

func (s *EnableDiskEncryptionByDefaultRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableDiskEncryptionByDefaultRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableDiskEncryptionByDefaultRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableDiskEncryptionByDefaultRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableDiskEncryptionByDefaultRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableDiskEncryptionByDefaultRequest) SetOwnerAccount(v string) *EnableDiskEncryptionByDefaultRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableDiskEncryptionByDefaultRequest) SetOwnerId(v int64) *EnableDiskEncryptionByDefaultRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableDiskEncryptionByDefaultRequest) SetRegionId(v string) *EnableDiskEncryptionByDefaultRequest {
  s.RegionId = &v
  return s
}

func (s *EnableDiskEncryptionByDefaultRequest) SetResourceOwnerAccount(v string) *EnableDiskEncryptionByDefaultRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableDiskEncryptionByDefaultRequest) SetResourceOwnerId(v int64) *EnableDiskEncryptionByDefaultRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableDiskEncryptionByDefaultRequest) Validate() error {
  return dara.Validate(s)
}

