// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmbAclResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSmbAclResponseBody
  GetRequestId() *string 
}

type EnableSmbAclResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 98696EF0-1607-4E9D-B01D-F20930B6****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableSmbAclResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSmbAclResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSmbAclResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSmbAclResponseBody) SetRequestId(v string) *EnableSmbAclResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSmbAclResponseBody) Validate() error {
  return dara.Validate(s)
}

