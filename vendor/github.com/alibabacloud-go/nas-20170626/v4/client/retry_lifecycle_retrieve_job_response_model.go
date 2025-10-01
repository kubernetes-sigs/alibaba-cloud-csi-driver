// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryLifecycleRetrieveJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryLifecycleRetrieveJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryLifecycleRetrieveJobResponse
	GetStatusCode() *int32
	SetBody(v *RetryLifecycleRetrieveJobResponseBody) *RetryLifecycleRetrieveJobResponse
	GetBody() *RetryLifecycleRetrieveJobResponseBody
}

type RetryLifecycleRetrieveJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryLifecycleRetrieveJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *RetryLifecycleRetrieveJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryLifecycleRetrieveJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryLifecycleRetrieveJobResponse) GetBody() *RetryLifecycleRetrieveJobResponseBody {
	return s.Body
}

func (s *RetryLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *RetryLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *RetryLifecycleRetrieveJobResponse) SetStatusCode(v int32) *RetryLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryLifecycleRetrieveJobResponse) SetBody(v *RetryLifecycleRetrieveJobResponseBody) *RetryLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

func (s *RetryLifecycleRetrieveJobResponse) Validate() error {
	return dara.Validate(s)
}
