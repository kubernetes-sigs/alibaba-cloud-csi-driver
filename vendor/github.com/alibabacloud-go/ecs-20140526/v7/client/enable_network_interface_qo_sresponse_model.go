// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNetworkInterfaceQoSResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableNetworkInterfaceQoSResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableNetworkInterfaceQoSResponse
  GetStatusCode() *int32 
  SetBody(v *EnableNetworkInterfaceQoSResponseBody) *EnableNetworkInterfaceQoSResponse
  GetBody() *EnableNetworkInterfaceQoSResponseBody 
}

type EnableNetworkInterfaceQoSResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableNetworkInterfaceQoSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableNetworkInterfaceQoSResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableNetworkInterfaceQoSResponse) GoString() string {
  return s.String()
}

func (s *EnableNetworkInterfaceQoSResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableNetworkInterfaceQoSResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableNetworkInterfaceQoSResponse) GetBody() *EnableNetworkInterfaceQoSResponseBody  {
  return s.Body
}

func (s *EnableNetworkInterfaceQoSResponse) SetHeaders(v map[string]*string) *EnableNetworkInterfaceQoSResponse {
  s.Headers = v
  return s
}

func (s *EnableNetworkInterfaceQoSResponse) SetStatusCode(v int32) *EnableNetworkInterfaceQoSResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableNetworkInterfaceQoSResponse) SetBody(v *EnableNetworkInterfaceQoSResponseBody) *EnableNetworkInterfaceQoSResponse {
  s.Body = v
  return s
}

func (s *EnableNetworkInterfaceQoSResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

