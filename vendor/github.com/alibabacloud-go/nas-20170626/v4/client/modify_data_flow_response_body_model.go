// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDataFlowResponseBody
	GetRequestId() *string
}

type ModifyDataFlowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataFlowResponseBody) SetRequestId(v string) *ModifyDataFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
