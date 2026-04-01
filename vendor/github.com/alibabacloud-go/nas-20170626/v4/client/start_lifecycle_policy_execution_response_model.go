// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLifecyclePolicyExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartLifecyclePolicyExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartLifecyclePolicyExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StartLifecyclePolicyExecutionResponseBody) *StartLifecyclePolicyExecutionResponse
	GetBody() *StartLifecyclePolicyExecutionResponseBody
}

type StartLifecyclePolicyExecutionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLifecyclePolicyExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLifecyclePolicyExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartLifecyclePolicyExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartLifecyclePolicyExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartLifecyclePolicyExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartLifecyclePolicyExecutionResponse) GetBody() *StartLifecyclePolicyExecutionResponseBody {
	return s.Body
}

func (s *StartLifecyclePolicyExecutionResponse) SetHeaders(v map[string]*string) *StartLifecyclePolicyExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartLifecyclePolicyExecutionResponse) SetStatusCode(v int32) *StartLifecyclePolicyExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLifecyclePolicyExecutionResponse) SetBody(v *StartLifecyclePolicyExecutionResponseBody) *StartLifecyclePolicyExecutionResponse {
	s.Body = v
	return s
}

func (s *StartLifecyclePolicyExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
