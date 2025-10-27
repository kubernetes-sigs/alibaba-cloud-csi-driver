// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowSubTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataFlowSubTaskId(v string) *CreateDataFlowSubTaskResponseBody
	GetDataFlowSubTaskId() *string
	SetRequestId(v string) *CreateDataFlowSubTaskResponseBody
	GetRequestId() *string
}

type CreateDataFlowSubTaskResponseBody struct {
	// The ID of the data streaming task.
	//
	// example:
	//
	// subTaskId-370kyfmyknxcyzw****
	DataFlowSubTaskId *string `json:"DataFlowSubTaskId,omitempty" xml:"DataFlowSubTaskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A70BEE5D-76D3-49FB-B58F-1F398211A5C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataFlowSubTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowSubTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataFlowSubTaskResponseBody) GetDataFlowSubTaskId() *string {
	return s.DataFlowSubTaskId
}

func (s *CreateDataFlowSubTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataFlowSubTaskResponseBody) SetDataFlowSubTaskId(v string) *CreateDataFlowSubTaskResponseBody {
	s.DataFlowSubTaskId = &v
	return s
}

func (s *CreateDataFlowSubTaskResponseBody) SetRequestId(v string) *CreateDataFlowSubTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataFlowSubTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
