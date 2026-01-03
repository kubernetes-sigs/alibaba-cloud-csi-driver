// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachDiskResponseBody
	GetRequestId() *string
}

type AttachDiskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDiskResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDiskResponseBody) SetRequestId(v string) *AttachDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
