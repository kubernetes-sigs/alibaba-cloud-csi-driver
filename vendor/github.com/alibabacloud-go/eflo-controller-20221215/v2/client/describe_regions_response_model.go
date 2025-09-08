// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse
	GetBody() *DescribeRegionsResponseBody
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionsResponse) GetBody() *DescribeRegionsResponseBody {
	return s.Body
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeRegionsResponse) Validate() error {
	return dara.Validate(s)
}
