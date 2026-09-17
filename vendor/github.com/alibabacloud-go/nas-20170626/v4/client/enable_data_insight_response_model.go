// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDataInsightResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDataInsightResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDataInsightResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDataInsightResponseBody) *EnableDataInsightResponse
  GetBody() *EnableDataInsightResponseBody 
}

type EnableDataInsightResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDataInsightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDataInsightResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDataInsightResponse) GoString() string {
  return s.String()
}

func (s *EnableDataInsightResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDataInsightResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDataInsightResponse) GetBody() *EnableDataInsightResponseBody  {
  return s.Body
}

func (s *EnableDataInsightResponse) SetHeaders(v map[string]*string) *EnableDataInsightResponse {
  s.Headers = v
  return s
}

func (s *EnableDataInsightResponse) SetStatusCode(v int32) *EnableDataInsightResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDataInsightResponse) SetBody(v *EnableDataInsightResponseBody) *EnableDataInsightResponse {
  s.Body = v
  return s
}

func (s *EnableDataInsightResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

