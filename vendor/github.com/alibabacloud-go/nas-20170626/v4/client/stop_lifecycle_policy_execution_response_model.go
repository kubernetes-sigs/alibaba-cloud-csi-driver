// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLifecyclePolicyExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLifecyclePolicyExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLifecyclePolicyExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StopLifecyclePolicyExecutionResponseBody) *StopLifecyclePolicyExecutionResponse
	GetBody() *StopLifecyclePolicyExecutionResponseBody
}

type StopLifecyclePolicyExecutionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLifecyclePolicyExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLifecyclePolicyExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLifecyclePolicyExecutionResponse) GoString() string {
	return s.String()
}

func (s *StopLifecyclePolicyExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLifecyclePolicyExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLifecyclePolicyExecutionResponse) GetBody() *StopLifecyclePolicyExecutionResponseBody {
	return s.Body
}

func (s *StopLifecyclePolicyExecutionResponse) SetHeaders(v map[string]*string) *StopLifecyclePolicyExecutionResponse {
	s.Headers = v
	return s
}

func (s *StopLifecyclePolicyExecutionResponse) SetStatusCode(v int32) *StopLifecyclePolicyExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLifecyclePolicyExecutionResponse) SetBody(v *StopLifecyclePolicyExecutionResponseBody) *StopLifecyclePolicyExecutionResponse {
	s.Body = v
	return s
}

func (s *StopLifecyclePolicyExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
