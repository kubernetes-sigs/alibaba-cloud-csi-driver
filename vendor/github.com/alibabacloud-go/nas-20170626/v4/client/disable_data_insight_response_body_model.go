// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDataInsightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDataInsightResponseBody
	GetRequestId() *string
}

type DisableDataInsightResponseBody struct {
	// example:
	//
	// 5B4511A7-C99E-4071-AA8C-32E2529D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDataInsightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDataInsightResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDataInsightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDataInsightResponseBody) SetRequestId(v string) *DisableDataInsightResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDataInsightResponseBody) Validate() error {
	return dara.Validate(s)
}
