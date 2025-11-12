// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendClusterResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExtendClusterResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExtendClusterResponse
  GetStatusCode() *int32 
  SetBody(v *ExtendClusterResponseBody) *ExtendClusterResponse
  GetBody() *ExtendClusterResponseBody 
}

type ExtendClusterResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExtendClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExtendClusterResponse) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterResponse) GoString() string {
  return s.String()
}

func (s *ExtendClusterResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExtendClusterResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExtendClusterResponse) GetBody() *ExtendClusterResponseBody  {
  return s.Body
}

func (s *ExtendClusterResponse) SetHeaders(v map[string]*string) *ExtendClusterResponse {
  s.Headers = v
  return s
}

func (s *ExtendClusterResponse) SetStatusCode(v int32) *ExtendClusterResponse {
  s.StatusCode = &v
  return s
}

func (s *ExtendClusterResponse) SetBody(v *ExtendClusterResponseBody) *ExtendClusterResponse {
  s.Body = v
  return s
}

func (s *ExtendClusterResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

