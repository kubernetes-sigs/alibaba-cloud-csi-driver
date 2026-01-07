// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachDiskResponseBody
	GetRequestId() *string
}

type DetachDiskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachDiskResponseBody) SetRequestId(v string) *DetachDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
