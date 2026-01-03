// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse
	GetBody() *DescribeTasksResponseBody
}

type DescribeTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTasksResponse) GetBody() *DescribeTasksResponseBody {
	return s.Body
}

func (s *DescribeTasksResponse) SetHeaders(v map[string]*string) *DescribeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeTasksResponse) SetStatusCode(v int32) *DescribeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
