// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendClusterResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExtendClusterResponseBody
  GetRequestId() *string 
  SetTaskId(v string) *ExtendClusterResponseBody
  GetTaskId() *string 
}

type ExtendClusterResponseBody struct {
  // Request ID
  // 
  // example:
  // 
  // 03668372-18FF-5959-98D9-6B36A4643C7A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Task ID
  // 
  // example:
  // 
  // i158475611663639202234
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExtendClusterResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterResponseBody) GoString() string {
  return s.String()
}

func (s *ExtendClusterResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExtendClusterResponseBody) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExtendClusterResponseBody) SetRequestId(v string) *ExtendClusterResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExtendClusterResponseBody) SetTaskId(v string) *ExtendClusterResponseBody {
  s.TaskId = &v
  return s
}

func (s *ExtendClusterResponseBody) Validate() error {
  return dara.Validate(s)
}

