// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLifecycleRetrieveJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLifecycleRetrieveJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLifecycleRetrieveJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListLifecycleRetrieveJobsResponseBody) *ListLifecycleRetrieveJobsResponse
	GetBody() *ListLifecycleRetrieveJobsResponseBody
}

type ListLifecycleRetrieveJobsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLifecycleRetrieveJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLifecycleRetrieveJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLifecycleRetrieveJobsResponse) GoString() string {
	return s.String()
}

func (s *ListLifecycleRetrieveJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLifecycleRetrieveJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLifecycleRetrieveJobsResponse) GetBody() *ListLifecycleRetrieveJobsResponseBody {
	return s.Body
}

func (s *ListLifecycleRetrieveJobsResponse) SetHeaders(v map[string]*string) *ListLifecycleRetrieveJobsResponse {
	s.Headers = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponse) SetStatusCode(v int32) *ListLifecycleRetrieveJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLifecycleRetrieveJobsResponse) SetBody(v *ListLifecycleRetrieveJobsResponseBody) *ListLifecycleRetrieveJobsResponse {
	s.Body = v
	return s
}

func (s *ListLifecycleRetrieveJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
