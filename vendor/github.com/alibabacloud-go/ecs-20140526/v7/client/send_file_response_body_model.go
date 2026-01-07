// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *SendFileResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *SendFileResponseBody
	GetRequestId() *string
}

type SendFileResponseBody struct {
	// The ID of the command task.
	//
	// example:
	//
	// f-7d2a745b412b46****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendFileResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *SendFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendFileResponseBody) SetInvokeId(v string) *SendFileResponseBody {
	s.InvokeId = &v
	return s
}

func (s *SendFileResponseBody) SetRequestId(v string) *SendFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendFileResponseBody) Validate() error {
	return dara.Validate(s)
}
