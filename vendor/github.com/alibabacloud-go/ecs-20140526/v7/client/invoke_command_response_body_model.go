// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *InvokeCommandResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *InvokeCommandResponseBody
	GetRequestId() *string
}

type InvokeCommandResponseBody struct {
	// The ID of the command task.
	//
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeCommandResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeCommandResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *InvokeCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeCommandResponseBody) SetInvokeId(v string) *InvokeCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *InvokeCommandResponseBody) SetRequestId(v string) *InvokeCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
