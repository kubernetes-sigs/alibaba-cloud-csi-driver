// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvocationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvocationResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvocationResultsResponseBody) *DescribeInvocationResultsResponse
	GetBody() *DescribeInvocationResultsResponseBody
}

type DescribeInvocationResultsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvocationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvocationResultsResponse) GetBody() *DescribeInvocationResultsResponseBody {
	return s.Body
}

func (s *DescribeInvocationResultsResponse) SetHeaders(v map[string]*string) *DescribeInvocationResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationResultsResponse) SetStatusCode(v int32) *DescribeInvocationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationResultsResponse) SetBody(v *DescribeInvocationResultsResponseBody) *DescribeInvocationResultsResponse {
	s.Body = v
	return s
}

func (s *DescribeInvocationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
