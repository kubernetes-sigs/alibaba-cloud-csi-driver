// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesFullStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstancesFullStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstancesFullStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstancesFullStatusResponseBody) *DescribeInstancesFullStatusResponse
	GetBody() *DescribeInstancesFullStatusResponseBody
}

type DescribeInstancesFullStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesFullStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancesFullStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesFullStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesFullStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstancesFullStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstancesFullStatusResponse) GetBody() *DescribeInstancesFullStatusResponseBody {
	return s.Body
}

func (s *DescribeInstancesFullStatusResponse) SetHeaders(v map[string]*string) *DescribeInstancesFullStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesFullStatusResponse) SetStatusCode(v int32) *DescribeInstancesFullStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesFullStatusResponse) SetBody(v *DescribeInstancesFullStatusResponseBody) *DescribeInstancesFullStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeInstancesFullStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
