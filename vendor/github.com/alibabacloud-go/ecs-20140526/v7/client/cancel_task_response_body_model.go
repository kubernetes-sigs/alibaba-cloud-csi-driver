// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelTaskResponseBody
	GetRequestId() *string
}

type CancelTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
