// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataFlowTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataFlowTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataFlowTaskResponseBody) *CreateDataFlowTaskResponse
	GetBody() *CreateDataFlowTaskResponseBody
}

type CreateDataFlowTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataFlowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataFlowTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDataFlowTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataFlowTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataFlowTaskResponse) GetBody() *CreateDataFlowTaskResponseBody {
	return s.Body
}

func (s *CreateDataFlowTaskResponse) SetHeaders(v map[string]*string) *CreateDataFlowTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDataFlowTaskResponse) SetStatusCode(v int32) *CreateDataFlowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataFlowTaskResponse) SetBody(v *CreateDataFlowTaskResponseBody) *CreateDataFlowTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDataFlowTaskResponse) Validate() error {
	return dara.Validate(s)
}
