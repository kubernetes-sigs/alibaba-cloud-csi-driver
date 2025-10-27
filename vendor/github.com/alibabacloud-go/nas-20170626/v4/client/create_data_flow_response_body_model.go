// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataFlowId(v string) *CreateDataFlowResponseBody
	GetDataFlowId() *string
	SetRequestId(v string) *CreateDataFlowResponseBody
	GetRequestId() *string
}

type CreateDataFlowResponseBody struct {
	// The dataflow ID.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0D****3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataFlowResponseBody) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *CreateDataFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataFlowResponseBody) SetDataFlowId(v string) *CreateDataFlowResponseBody {
	s.DataFlowId = &v
	return s
}

func (s *CreateDataFlowResponseBody) SetRequestId(v string) *CreateDataFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
