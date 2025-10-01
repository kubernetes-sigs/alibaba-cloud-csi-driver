// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNfsAclResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableNfsAclResponseBody
  GetRequestId() *string 
}

type EnableNfsAclResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 29F4F360-A6A8-561A-A45B-A0F6882969BA
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableNfsAclResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableNfsAclResponseBody) GoString() string {
  return s.String()
}

func (s *EnableNfsAclResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableNfsAclResponseBody) SetRequestId(v string) *EnableNfsAclResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableNfsAclResponseBody) Validate() error {
  return dara.Validate(s)
}

