// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataFlowAutoRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDataFlowAutoRefreshResponseBody
	GetRequestId() *string
}

type ModifyDataFlowAutoRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataFlowAutoRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataFlowAutoRefreshResponseBody) SetRequestId(v string) *ModifyDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
