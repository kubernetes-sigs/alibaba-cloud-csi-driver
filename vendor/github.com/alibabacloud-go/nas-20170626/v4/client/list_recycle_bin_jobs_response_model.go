// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecycleBinJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecycleBinJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecycleBinJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecycleBinJobsResponseBody) *ListRecycleBinJobsResponse
	GetBody() *ListRecycleBinJobsResponseBody
}

type ListRecycleBinJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecycleBinJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecycleBinJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecycleBinJobsResponse) GoString() string {
	return s.String()
}

func (s *ListRecycleBinJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecycleBinJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecycleBinJobsResponse) GetBody() *ListRecycleBinJobsResponseBody {
	return s.Body
}

func (s *ListRecycleBinJobsResponse) SetHeaders(v map[string]*string) *ListRecycleBinJobsResponse {
	s.Headers = v
	return s
}

func (s *ListRecycleBinJobsResponse) SetStatusCode(v int32) *ListRecycleBinJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecycleBinJobsResponse) SetBody(v *ListRecycleBinJobsResponseBody) *ListRecycleBinJobsResponse {
	s.Body = v
	return s
}

func (s *ListRecycleBinJobsResponse) Validate() error {
	return dara.Validate(s)
}
