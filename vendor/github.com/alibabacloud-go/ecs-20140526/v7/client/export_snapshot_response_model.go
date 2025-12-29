// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportSnapshotResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExportSnapshotResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExportSnapshotResponse
  GetStatusCode() *int32 
  SetBody(v *ExportSnapshotResponseBody) *ExportSnapshotResponse
  GetBody() *ExportSnapshotResponseBody 
}

type ExportSnapshotResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExportSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportSnapshotResponse) String() string {
  return dara.Prettify(s)
}

func (s ExportSnapshotResponse) GoString() string {
  return s.String()
}

func (s *ExportSnapshotResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExportSnapshotResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExportSnapshotResponse) GetBody() *ExportSnapshotResponseBody  {
  return s.Body
}

func (s *ExportSnapshotResponse) SetHeaders(v map[string]*string) *ExportSnapshotResponse {
  s.Headers = v
  return s
}

func (s *ExportSnapshotResponse) SetStatusCode(v int32) *ExportSnapshotResponse {
  s.StatusCode = &v
  return s
}

func (s *ExportSnapshotResponse) SetBody(v *ExportSnapshotResponseBody) *ExportSnapshotResponse {
  s.Body = v
  return s
}

func (s *ExportSnapshotResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

