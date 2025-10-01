// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDataFlowSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDataFlowSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelDataFlowSubTaskResponseBody) *CancelDataFlowSubTaskResponse
	GetBody() *CancelDataFlowSubTaskResponseBody
}

type CancelDataFlowSubTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDataFlowSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDataFlowSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowSubTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelDataFlowSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDataFlowSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDataFlowSubTaskResponse) GetBody() *CancelDataFlowSubTaskResponseBody {
	return s.Body
}

func (s *CancelDataFlowSubTaskResponse) SetHeaders(v map[string]*string) *CancelDataFlowSubTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelDataFlowSubTaskResponse) SetStatusCode(v int32) *CancelDataFlowSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDataFlowSubTaskResponse) SetBody(v *CancelDataFlowSubTaskResponseBody) *CancelDataFlowSubTaskResponse {
	s.Body = v
	return s
}

func (s *CancelDataFlowSubTaskResponse) Validate() error {
	return dara.Validate(s)
}
