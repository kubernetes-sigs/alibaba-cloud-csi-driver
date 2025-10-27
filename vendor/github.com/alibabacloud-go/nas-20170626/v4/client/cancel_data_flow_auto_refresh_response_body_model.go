// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowAutoRefreshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelDataFlowAutoRefreshResponseBody
	GetRequestId() *string
}

type CancelDataFlowAutoRefreshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDataFlowAutoRefreshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowAutoRefreshResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDataFlowAutoRefreshResponseBody) SetRequestId(v string) *CancelDataFlowAutoRefreshResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshResponseBody) Validate() error {
	return dara.Validate(s)
}
