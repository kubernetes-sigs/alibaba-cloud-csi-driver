// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmbAclRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFileSystemId(v string) *EnableSmbAclRequest
  GetFileSystemId() *string 
  SetKeytab(v string) *EnableSmbAclRequest
  GetKeytab() *string 
  SetKeytabMd5(v string) *EnableSmbAclRequest
  GetKeytabMd5() *string 
}

type EnableSmbAclRequest struct {
  // The ID of the file system.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 31a8e4****
  FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
  // The string that is generated after the system encodes the keytab file by using Base64.
  // 
  // example:
  // 
  // BQIAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAQAIqIx6v7p11oUAAABHAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAAwAIqIx6v7p11oUAAABPAAIADUFMSUFEVEVTVC5DT00ABGNpZnMAGXNtYnNlcnZlcjI0LmFsaWFkdGVzdC5jb20AAAABAAAAAAEAFwAQnQZWB3RAPHU7PMIJyBWePAAAAF8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQASACAGJ7F0s+bcBjf6jD5HlvlRLmPSOW+qDZe0Qk0lQcf8WwAAAE8AAgANQUxJQURURVNULkNPTQAEY2lmcwAZc21ic2VydmVyMjQuYWxpYWR0ZXN0LmNvbQAAAAEAAAAAAQARABDdFmanrSIatnDDhxxxxx
  Keytab *string `json:"Keytab,omitempty" xml:"Keytab,omitempty"`
  // The string that is generated after the system encodes the keytab file by using MD5.
  // 
  // example:
  // 
  // E3CCF7E2416DF04FA958AA4513EAxxxx
  KeytabMd5 *string `json:"KeytabMd5,omitempty" xml:"KeytabMd5,omitempty"`
}

func (s EnableSmbAclRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSmbAclRequest) GoString() string {
  return s.String()
}

func (s *EnableSmbAclRequest) GetFileSystemId() *string  {
  return s.FileSystemId
}

func (s *EnableSmbAclRequest) GetKeytab() *string  {
  return s.Keytab
}

func (s *EnableSmbAclRequest) GetKeytabMd5() *string  {
  return s.KeytabMd5
}

func (s *EnableSmbAclRequest) SetFileSystemId(v string) *EnableSmbAclRequest {
  s.FileSystemId = &v
  return s
}

func (s *EnableSmbAclRequest) SetKeytab(v string) *EnableSmbAclRequest {
  s.Keytab = &v
  return s
}

func (s *EnableSmbAclRequest) SetKeytabMd5(v string) *EnableSmbAclRequest {
  s.KeytabMd5 = &v
  return s
}

func (s *EnableSmbAclRequest) Validate() error {
  return dara.Validate(s)
}

