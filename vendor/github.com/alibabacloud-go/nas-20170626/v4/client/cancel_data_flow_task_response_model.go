// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDataFlowTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDataFlowTaskResponse
	GetStatusCode() *int32
	SetBody(v *CancelDataFlowTaskResponseBody) *CancelDataFlowTaskResponse
	GetBody() *CancelDataFlowTaskResponseBody
}

type CancelDataFlowTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDataFlowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDataFlowTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelDataFlowTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDataFlowTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDataFlowTaskResponse) GetBody() *CancelDataFlowTaskResponseBody {
	return s.Body
}

func (s *CancelDataFlowTaskResponse) SetHeaders(v map[string]*string) *CancelDataFlowTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelDataFlowTaskResponse) SetStatusCode(v int32) *CancelDataFlowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDataFlowTaskResponse) SetBody(v *CancelDataFlowTaskResponseBody) *CancelDataFlowTaskResponse {
	s.Body = v
	return s
}

func (s *CancelDataFlowTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
