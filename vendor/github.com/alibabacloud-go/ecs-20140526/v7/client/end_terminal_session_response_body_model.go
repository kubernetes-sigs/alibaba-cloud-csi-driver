// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndTerminalSessionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EndTerminalSessionResponseBody
  GetRequestId() *string 
}

type EndTerminalSessionResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndTerminalSessionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndTerminalSessionResponseBody) GoString() string {
  return s.String()
}

func (s *EndTerminalSessionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndTerminalSessionResponseBody) SetRequestId(v string) *EndTerminalSessionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndTerminalSessionResponseBody) Validate() error {
  return dara.Validate(s)
}

