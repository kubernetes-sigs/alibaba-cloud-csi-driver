// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnablePhysicalConnectionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnablePhysicalConnectionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnablePhysicalConnectionResponse
  GetStatusCode() *int32 
  SetBody(v *EnablePhysicalConnectionResponseBody) *EnablePhysicalConnectionResponse
  GetBody() *EnablePhysicalConnectionResponseBody 
}

type EnablePhysicalConnectionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnablePhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnablePhysicalConnectionResponse) String() string {
  return dara.Prettify(s)
}

func (s EnablePhysicalConnectionResponse) GoString() string {
  return s.String()
}

func (s *EnablePhysicalConnectionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnablePhysicalConnectionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnablePhysicalConnectionResponse) GetBody() *EnablePhysicalConnectionResponseBody  {
  return s.Body
}

func (s *EnablePhysicalConnectionResponse) SetHeaders(v map[string]*string) *EnablePhysicalConnectionResponse {
  s.Headers = v
  return s
}

func (s *EnablePhysicalConnectionResponse) SetStatusCode(v int32) *EnablePhysicalConnectionResponse {
  s.StatusCode = &v
  return s
}

func (s *EnablePhysicalConnectionResponse) SetBody(v *EnablePhysicalConnectionResponseBody) *EnablePhysicalConnectionResponse {
  s.Body = v
  return s
}

func (s *EnablePhysicalConnectionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

