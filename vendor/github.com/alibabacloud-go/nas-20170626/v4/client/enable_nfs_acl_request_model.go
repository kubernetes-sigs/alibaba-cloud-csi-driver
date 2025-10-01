// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNfsAclRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFileSystemId(v string) *EnableNfsAclRequest
  GetFileSystemId() *string 
}

type EnableNfsAclRequest struct {
  // The ID of the file system.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 43f264xxxx
  FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s EnableNfsAclRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableNfsAclRequest) GoString() string {
  return s.String()
}

func (s *EnableNfsAclRequest) GetFileSystemId() *string  {
  return s.FileSystemId
}

func (s *EnableNfsAclRequest) SetFileSystemId(v string) *EnableNfsAclRequest {
  s.FileSystemId = &v
  return s
}

func (s *EnableNfsAclRequest) Validate() error {
  return dara.Validate(s)
}

