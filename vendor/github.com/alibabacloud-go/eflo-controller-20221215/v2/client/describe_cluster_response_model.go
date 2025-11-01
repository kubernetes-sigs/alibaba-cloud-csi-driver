// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse
	GetBody() *DescribeClusterResponseBody
}

type DescribeClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterResponse) GetBody() *DescribeClusterResponseBody {
	return s.Body
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetStatusCode(v int32) *DescribeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
