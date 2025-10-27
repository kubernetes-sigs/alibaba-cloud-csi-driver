// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSmbAclResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableSmbAclResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableSmbAclResponse
  GetStatusCode() *int32 
  SetBody(v *EnableSmbAclResponseBody) *EnableSmbAclResponse
  GetBody() *EnableSmbAclResponseBody 
}

type EnableSmbAclResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableSmbAclResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableSmbAclResponse) GoString() string {
  return s.String()
}

func (s *EnableSmbAclResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableSmbAclResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableSmbAclResponse) GetBody() *EnableSmbAclResponseBody  {
  return s.Body
}

func (s *EnableSmbAclResponse) SetHeaders(v map[string]*string) *EnableSmbAclResponse {
  s.Headers = v
  return s
}

func (s *EnableSmbAclResponse) SetStatusCode(v int32) *EnableSmbAclResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableSmbAclResponse) SetBody(v *EnableSmbAclResponseBody) *EnableSmbAclResponse {
  s.Body = v
  return s
}

func (s *EnableSmbAclResponse) Validate() error {
  return dara.Validate(s)
}

