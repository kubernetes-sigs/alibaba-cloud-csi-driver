// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartDataFlowResponseBody
	GetRequestId() *string
}

type StartDataFlowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDataFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StartDataFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDataFlowResponseBody) SetRequestId(v string) *StartDataFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDataFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
