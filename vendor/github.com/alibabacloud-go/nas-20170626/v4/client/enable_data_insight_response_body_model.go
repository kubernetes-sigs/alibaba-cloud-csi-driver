// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDataInsightResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDataInsightResponseBody
  GetRequestId() *string 
}

type EnableDataInsightResponseBody struct {
  // example:
  // 
  // 2D69A58F-345C-4FDE-88E4-BF518948****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDataInsightResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDataInsightResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDataInsightResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDataInsightResponseBody) SetRequestId(v string) *EnableDataInsightResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDataInsightResponseBody) Validate() error {
  return dara.Validate(s)
}

