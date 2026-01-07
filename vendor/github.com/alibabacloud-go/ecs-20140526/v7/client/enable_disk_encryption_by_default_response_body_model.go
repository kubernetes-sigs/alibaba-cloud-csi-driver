// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDiskEncryptionByDefaultResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDiskEncryptionByDefaultResponseBody
  GetRequestId() *string 
}

type EnableDiskEncryptionByDefaultResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDiskEncryptionByDefaultResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDiskEncryptionByDefaultResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDiskEncryptionByDefaultResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDiskEncryptionByDefaultResponseBody) SetRequestId(v string) *EnableDiskEncryptionByDefaultResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDiskEncryptionByDefaultResponseBody) Validate() error {
  return dara.Validate(s)
}

