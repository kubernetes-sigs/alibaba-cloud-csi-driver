// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRecycleBinRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFileSystemId(v string) *EnableRecycleBinRequest
  GetFileSystemId() *string 
  SetReservedDays(v int64) *EnableRecycleBinRequest
  GetReservedDays() *int64 
}

type EnableRecycleBinRequest struct {
  // The ID of the file system for which you want to enable the recycle bin feature.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1ca404****
  FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
  // The retention period of the files in the recycle bin. Unit: days.
  // 
  // Valid values: 1 to 180.
  // 
  // Default value: 3.
  // 
  // example:
  // 
  // 3
  ReservedDays *int64 `json:"ReservedDays,omitempty" xml:"ReservedDays,omitempty"`
}

func (s EnableRecycleBinRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableRecycleBinRequest) GoString() string {
  return s.String()
}

func (s *EnableRecycleBinRequest) GetFileSystemId() *string  {
  return s.FileSystemId
}

func (s *EnableRecycleBinRequest) GetReservedDays() *int64  {
  return s.ReservedDays
}

func (s *EnableRecycleBinRequest) SetFileSystemId(v string) *EnableRecycleBinRequest {
  s.FileSystemId = &v
  return s
}

func (s *EnableRecycleBinRequest) SetReservedDays(v int64) *EnableRecycleBinRequest {
  s.ReservedDays = &v
  return s
}

func (s *EnableRecycleBinRequest) Validate() error {
  return dara.Validate(s)
}

