// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportImageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *ExportImageResponseBody
  GetRegionId() *string 
  SetRequestId(v string) *ExportImageResponseBody
  GetRequestId() *string 
  SetTaskId(v string) *ExportImageResponseBody
  GetTaskId() *string 
}

type ExportImageResponseBody struct {
  // The region ID.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // C8B26B44-0189-443E-9816-D951F596****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The ID of the task that is used to export the custom image.
  // 
  // example:
  // 
  // tsk-bp67acfmxazb4p****
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExportImageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExportImageResponseBody) GoString() string {
  return s.String()
}

func (s *ExportImageResponseBody) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportImageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExportImageResponseBody) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportImageResponseBody) SetRegionId(v string) *ExportImageResponseBody {
  s.RegionId = &v
  return s
}

func (s *ExportImageResponseBody) SetRequestId(v string) *ExportImageResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExportImageResponseBody) SetTaskId(v string) *ExportImageResponseBody {
  s.TaskId = &v
  return s
}

func (s *ExportImageResponseBody) Validate() error {
  return dara.Validate(s)
}

