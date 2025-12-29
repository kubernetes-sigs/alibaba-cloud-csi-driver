// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNetworkInterfaceQoSResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableNetworkInterfaceQoSResponseBody
  GetRequestId() *string 
}

type EnableNetworkInterfaceQoSResponseBody struct {
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableNetworkInterfaceQoSResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableNetworkInterfaceQoSResponseBody) GoString() string {
  return s.String()
}

func (s *EnableNetworkInterfaceQoSResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableNetworkInterfaceQoSResponseBody) SetRequestId(v string) *EnableNetworkInterfaceQoSResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableNetworkInterfaceQoSResponseBody) Validate() error {
  return dara.Validate(s)
}

