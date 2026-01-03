// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDiskEncryptionByDefaultResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EnableDiskEncryptionByDefaultResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EnableDiskEncryptionByDefaultResponse
  GetStatusCode() *int32 
  SetBody(v *EnableDiskEncryptionByDefaultResponseBody) *EnableDiskEncryptionByDefaultResponse
  GetBody() *EnableDiskEncryptionByDefaultResponseBody 
}

type EnableDiskEncryptionByDefaultResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EnableDiskEncryptionByDefaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDiskEncryptionByDefaultResponse) String() string {
  return dara.Prettify(s)
}

func (s EnableDiskEncryptionByDefaultResponse) GoString() string {
  return s.String()
}

func (s *EnableDiskEncryptionByDefaultResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EnableDiskEncryptionByDefaultResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EnableDiskEncryptionByDefaultResponse) GetBody() *EnableDiskEncryptionByDefaultResponseBody  {
  return s.Body
}

func (s *EnableDiskEncryptionByDefaultResponse) SetHeaders(v map[string]*string) *EnableDiskEncryptionByDefaultResponse {
  s.Headers = v
  return s
}

func (s *EnableDiskEncryptionByDefaultResponse) SetStatusCode(v int32) *EnableDiskEncryptionByDefaultResponse {
  s.StatusCode = &v
  return s
}

func (s *EnableDiskEncryptionByDefaultResponse) SetBody(v *EnableDiskEncryptionByDefaultResponseBody) *EnableDiskEncryptionByDefaultResponse {
  s.Body = v
  return s
}

func (s *EnableDiskEncryptionByDefaultResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

