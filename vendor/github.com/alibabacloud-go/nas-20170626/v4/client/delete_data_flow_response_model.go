// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataFlowResponseBody) *DeleteDataFlowResponse
	GetBody() *DeleteDataFlowResponseBody
}

type DeleteDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataFlowResponse) GetBody() *DeleteDataFlowResponseBody {
	return s.Body
}

func (s *DeleteDataFlowResponse) SetHeaders(v map[string]*string) *DeleteDataFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataFlowResponse) SetStatusCode(v int32) *DeleteDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataFlowResponse) SetBody(v *DeleteDataFlowResponseBody) *DeleteDataFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteDataFlowResponse) Validate() error {
	return dara.Validate(s)
}
