// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRecycleBinResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableRecycleBinResponseBody
  GetRequestId() *string 
}

type EnableRecycleBinResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 9E15E394-38A6-457A-A62A-D9797C9A****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableRecycleBinResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableRecycleBinResponseBody) GoString() string {
  return s.String()
}

func (s *EnableRecycleBinResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableRecycleBinResponseBody) SetRequestId(v string) *EnableRecycleBinResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableRecycleBinResponseBody) Validate() error {
  return dara.Validate(s)
}

