// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataFlowResponseBody
	GetRequestId() *string
}

type DeleteDataFlowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataFlowResponseBody) SetRequestId(v string) *DeleteDataFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
