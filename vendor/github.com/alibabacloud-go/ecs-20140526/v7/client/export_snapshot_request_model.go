// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportSnapshotRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOssBucket(v string) *ExportSnapshotRequest
  GetOssBucket() *string 
  SetOwnerId(v int64) *ExportSnapshotRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *ExportSnapshotRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *ExportSnapshotRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *ExportSnapshotRequest
  GetResourceOwnerId() *int64 
  SetRoleName(v string) *ExportSnapshotRequest
  GetRoleName() *string 
  SetSnapshotId(v string) *ExportSnapshotRequest
  GetSnapshotId() *string 
}

type ExportSnapshotRequest struct {
  // This parameter is required.
  OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // This parameter is required.
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
  // This parameter is required.
  SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ExportSnapshotRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportSnapshotRequest) GoString() string {
  return s.String()
}

func (s *ExportSnapshotRequest) GetOssBucket() *string  {
  return s.OssBucket
}

func (s *ExportSnapshotRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ExportSnapshotRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportSnapshotRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *ExportSnapshotRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *ExportSnapshotRequest) GetRoleName() *string  {
  return s.RoleName
}

func (s *ExportSnapshotRequest) GetSnapshotId() *string  {
  return s.SnapshotId
}

func (s *ExportSnapshotRequest) SetOssBucket(v string) *ExportSnapshotRequest {
  s.OssBucket = &v
  return s
}

func (s *ExportSnapshotRequest) SetOwnerId(v int64) *ExportSnapshotRequest {
  s.OwnerId = &v
  return s
}

func (s *ExportSnapshotRequest) SetRegionId(v string) *ExportSnapshotRequest {
  s.RegionId = &v
  return s
}

func (s *ExportSnapshotRequest) SetResourceOwnerAccount(v string) *ExportSnapshotRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *ExportSnapshotRequest) SetResourceOwnerId(v int64) *ExportSnapshotRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *ExportSnapshotRequest) SetRoleName(v string) *ExportSnapshotRequest {
  s.RoleName = &v
  return s
}

func (s *ExportSnapshotRequest) SetSnapshotId(v string) *ExportSnapshotRequest {
  s.SnapshotId = &v
  return s
}

func (s *ExportSnapshotRequest) Validate() error {
  return dara.Validate(s)
}

