// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataFlowResponseBody) *CreateDataFlowResponse
	GetBody() *CreateDataFlowResponseBody
}

type CreateDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateDataFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataFlowResponse) GetBody() *CreateDataFlowResponseBody {
	return s.Body
}

func (s *CreateDataFlowResponse) SetHeaders(v map[string]*string) *CreateDataFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateDataFlowResponse) SetStatusCode(v int32) *CreateDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataFlowResponse) SetBody(v *CreateDataFlowResponseBody) *CreateDataFlowResponse {
	s.Body = v
	return s
}

func (s *CreateDataFlowResponse) Validate() error {
	return dara.Validate(s)
}
