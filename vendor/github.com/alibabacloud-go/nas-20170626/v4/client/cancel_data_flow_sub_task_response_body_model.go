// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelDataFlowSubTaskResponseBody
	GetRequestId() *string
}

type CancelDataFlowSubTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDataFlowSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDataFlowSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDataFlowSubTaskResponseBody) SetRequestId(v string) *CancelDataFlowSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDataFlowSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
