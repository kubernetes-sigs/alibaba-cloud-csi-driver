// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse
	GetBody() *DescribeInstancesResponseBody
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstancesResponse) GetBody() *DescribeInstancesResponseBody {
	return s.Body
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
