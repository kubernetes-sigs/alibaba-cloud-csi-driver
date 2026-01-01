// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRecycleBinResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableRecycleBinResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableRecycleBinResponse
  GetStatusCode() *int32 
  SetBody(v *EnableRecycleBinResponseBody) *EnableRecycleBinResponse
  GetBody() *EnableRecycleBinResponseBody 
}

type EnableRecycleBinResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableRecycleBinResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableRecycleBinResponse) GoString() string {
  return s.String()
}

func (s *EnableRecycleBinResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableRecycleBinResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableRecycleBinResponse) GetBody() *EnableRecycleBinResponseBody  {
  return s.Body
}

func (s *EnableRecycleBinResponse) SetHeaders(v map[string]*string) *EnableRecycleBinResponse {
  s.Headers = v
  return s
}

func (s *EnableRecycleBinResponse) SetStatusCode(v int32) *EnableRecycleBinResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableRecycleBinResponse) SetBody(v *EnableRecycleBinResponseBody) *EnableRecycleBinResponse {
  s.Body = v
  return s
}

func (s *EnableRecycleBinResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

