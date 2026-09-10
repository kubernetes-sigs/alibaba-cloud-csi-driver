// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDataInsightRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFileSystemId(v string) *EnableDataInsightRequest
  GetFileSystemId() *string 
}

type EnableDataInsightRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // bmcpfs-290w65p03ok64y*****
  FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s EnableDataInsightRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDataInsightRequest) GoString() string {
  return s.String()
}

func (s *EnableDataInsightRequest) GetFileSystemId() *string  {
  return s.FileSystemId
}

func (s *EnableDataInsightRequest) SetFileSystemId(v string) *EnableDataInsightRequest {
  s.FileSystemId = &v
  return s
}

func (s *EnableDataInsightRequest) Validate() error {
  return dara.Validate(s)
}

