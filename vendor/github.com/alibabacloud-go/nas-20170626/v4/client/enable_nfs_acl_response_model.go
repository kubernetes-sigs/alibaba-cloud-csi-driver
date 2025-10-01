// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNfsAclResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableNfsAclResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableNfsAclResponse
  GetStatusCode() *int32 
  SetBody(v *EnableNfsAclResponseBody) *EnableNfsAclResponse
  GetBody() *EnableNfsAclResponseBody 
}

type EnableNfsAclResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableNfsAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableNfsAclResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableNfsAclResponse) GoString() string {
  return s.String()
}

func (s *EnableNfsAclResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableNfsAclResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableNfsAclResponse) GetBody() *EnableNfsAclResponseBody  {
  return s.Body
}

func (s *EnableNfsAclResponse) SetHeaders(v map[string]*string) *EnableNfsAclResponse {
  s.Headers = v
  return s
}

func (s *EnableNfsAclResponse) SetStatusCode(v int32) *EnableNfsAclResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableNfsAclResponse) SetBody(v *EnableNfsAclResponseBody) *EnableNfsAclResponse {
  s.Body = v
  return s
}

func (s *EnableNfsAclResponse) Validate() error {
  return dara.Validate(s)
}

