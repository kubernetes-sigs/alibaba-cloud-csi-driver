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
	// The ID of the execution.
	//
	// example:
	//
	// t-hz03la52z1zkvls
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
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
