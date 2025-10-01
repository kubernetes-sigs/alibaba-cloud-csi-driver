// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataFlowAutoRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyDataFlowAutoRefreshResponseBody
	GetRequestId() *string
}

type ApplyDataFlowAutoRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyDataFlowAutoRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyDataFlowAutoRefreshResponseBody) SetRequestId(v string) *ApplyDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
