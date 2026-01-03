// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopyImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCopyImageResponseBody
	GetRequestId() *string
}

type CancelCopyImageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCopyImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCopyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCopyImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCopyImageResponseBody) SetRequestId(v string) *CancelCopyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCopyImageResponseBody) Validate() error {
	return dara.Validate(s)
}
