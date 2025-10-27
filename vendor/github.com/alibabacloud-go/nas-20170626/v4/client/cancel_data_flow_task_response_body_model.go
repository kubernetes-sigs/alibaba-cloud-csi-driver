// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelDataFlowTaskResponseBody
	GetRequestId() *string
}

type CancelDataFlowTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF518948****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDataFlowTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDataFlowTaskResponseBody) SetRequestId(v string) *CancelDataFlowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDataFlowTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
