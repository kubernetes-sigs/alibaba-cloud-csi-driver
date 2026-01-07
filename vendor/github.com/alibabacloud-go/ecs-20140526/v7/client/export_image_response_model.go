// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportImageResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportImageResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportImageResponse
  GetStatusCode() *int32 
  SetBody(v *ExportImageResponseBody) *ExportImageResponse
  GetBody() *ExportImageResponseBody 
}

type ExportImageResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportImageResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportImageResponse) GoString() string {
  return s.String()
}

func (s *ExportImageResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportImageResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportImageResponse) GetBody() *ExportImageResponseBody  {
  return s.Body
}

func (s *ExportImageResponse) SetHeaders(v map[string]*string) *ExportImageResponse {
  s.Headers = v
  return s
}

func (s *ExportImageResponse) SetStatusCode(v int32) *ExportImageResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportImageResponse) SetBody(v *ExportImageResponseBody) *ExportImageResponse {
  s.Body = v
  return s
}

func (s *ExportImageResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

