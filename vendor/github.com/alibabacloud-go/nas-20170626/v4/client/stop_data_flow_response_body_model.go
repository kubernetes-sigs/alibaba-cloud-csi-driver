// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopDataFlowResponseBody
	GetRequestId() *string
}

type StopDataFlowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDataFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StopDataFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDataFlowResponseBody) SetRequestId(v string) *StopDataFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDataFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
