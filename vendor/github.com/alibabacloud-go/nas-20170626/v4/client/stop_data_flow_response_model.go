// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDataFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDataFlowResponse
	GetStatusCode() *int32
	SetBody(v *StopDataFlowResponseBody) *StopDataFlowResponse
	GetBody() *StopDataFlowResponseBody
}

type StopDataFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDataFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDataFlowResponse) GoString() string {
	return s.String()
}

func (s *StopDataFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDataFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDataFlowResponse) GetBody() *StopDataFlowResponseBody {
	return s.Body
}

func (s *StopDataFlowResponse) SetHeaders(v map[string]*string) *StopDataFlowResponse {
	s.Headers = v
	return s
}

func (s *StopDataFlowResponse) SetStatusCode(v int32) *StopDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDataFlowResponse) SetBody(v *StopDataFlowResponseBody) *StopDataFlowResponse {
	s.Body = v
	return s
}

func (s *StopDataFlowResponse) Validate() error {
	return dara.Validate(s)
}
