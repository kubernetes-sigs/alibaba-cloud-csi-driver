// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceStatusResponseBody) *DescribeInstanceStatusResponse
	GetBody() *DescribeInstanceStatusResponseBody
}

type DescribeInstanceStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceStatusResponse) GetBody() *DescribeInstanceStatusResponseBody {
	return s.Body
}

func (s *DescribeInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStatusResponse) SetStatusCode(v int32) *DescribeInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceStatusResponse) SetBody(v *DescribeInstanceStatusResponseBody) *DescribeInstanceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
