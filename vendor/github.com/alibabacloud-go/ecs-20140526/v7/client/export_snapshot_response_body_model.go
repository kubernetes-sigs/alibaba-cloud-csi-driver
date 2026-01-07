// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportSnapshotResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExportSnapshotResponseBody
  GetRequestId() *string 
  SetTaskId(v string) *ExportSnapshotResponseBody
  GetTaskId() *string 
}

type ExportSnapshotResponseBody struct {
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExportSnapshotResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportSnapshotResponseBody) GoString() string {
  return s.String()
}

func (s *ExportSnapshotResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportSnapshotResponseBody) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportSnapshotResponseBody) SetRequestId(v string) *ExportSnapshotResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportSnapshotResponseBody) SetTaskId(v string) *ExportSnapshotResponseBody {
  s.TaskId = &v
  return s
}

func (s *ExportSnapshotResponseBody) Validate() error {
  return dara.Validate(s)
}

