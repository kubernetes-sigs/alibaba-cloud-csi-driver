// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndTerminalSessionResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EndTerminalSessionResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EndTerminalSessionResponse
  GetStatusCode() *int32 
  SetBody(v *EndTerminalSessionResponseBody) *EndTerminalSessionResponse
  GetBody() *EndTerminalSessionResponseBody 
}

type EndTerminalSessionResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EndTerminalSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EndTerminalSessionResponse) String() string {
  return dara.Prettify(s)
}

func (s EndTerminalSessionResponse) GoString() string {
  return s.String()
}

func (s *EndTerminalSessionResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EndTerminalSessionResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EndTerminalSessionResponse) GetBody() *EndTerminalSessionResponseBody  {
  return s.Body
}

func (s *EndTerminalSessionResponse) SetHeaders(v map[string]*string) *EndTerminalSessionResponse {
  s.Headers = v
  return s
}

func (s *EndTerminalSessionResponse) SetStatusCode(v int32) *EndTerminalSessionResponse {
  s.StatusCode = &v
  return s
}

func (s *EndTerminalSessionResponse) SetBody(v *EndTerminalSessionResponseBody) *EndTerminalSessionResponse {
  s.Body = v
  return s
}

func (s *EndTerminalSessionResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

