// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInvocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopInvocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopInvocationResponse
	GetStatusCode() *int32
	SetBody(v *StopInvocationResponseBody) *StopInvocationResponse
	GetBody() *StopInvocationResponseBody
}

type StopInvocationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInvocationResponse) String() string {
	return dara.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopInvocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopInvocationResponse) GetBody() *StopInvocationResponseBody {
	return s.Body
}

func (s *StopInvocationResponse) SetHeaders(v map[string]*string) *StopInvocationResponse {
	s.Headers = v
	return s
}

func (s *StopInvocationResponse) SetStatusCode(v int32) *StopInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInvocationResponse) SetBody(v *StopInvocationResponseBody) *StopInvocationResponse {
	s.Body = v
	return s
}

func (s *StopInvocationResponse) Validate() error {
	return dara.Validate(s)
}
