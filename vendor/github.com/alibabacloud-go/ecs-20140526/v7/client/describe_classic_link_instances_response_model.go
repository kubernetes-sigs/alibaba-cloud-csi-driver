// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassicLinkInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClassicLinkInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClassicLinkInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClassicLinkInstancesResponseBody) *DescribeClassicLinkInstancesResponse
	GetBody() *DescribeClassicLinkInstancesResponseBody
}

type DescribeClassicLinkInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClassicLinkInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClassicLinkInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassicLinkInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClassicLinkInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClassicLinkInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClassicLinkInstancesResponse) GetBody() *DescribeClassicLinkInstancesResponseBody {
	return s.Body
}

func (s *DescribeClassicLinkInstancesResponse) SetHeaders(v map[string]*string) *DescribeClassicLinkInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeClassicLinkInstancesResponse) SetStatusCode(v int32) *DescribeClassicLinkInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClassicLinkInstancesResponse) SetBody(v *DescribeClassicLinkInstancesResponseBody) *DescribeClassicLinkInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeClassicLinkInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
