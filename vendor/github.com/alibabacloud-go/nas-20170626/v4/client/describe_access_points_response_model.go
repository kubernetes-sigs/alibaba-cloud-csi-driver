// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessPointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessPointsResponseBody) *DescribeAccessPointsResponse
	GetBody() *DescribeAccessPointsResponseBody
}

type DescribeAccessPointsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessPointsResponse) GetBody() *DescribeAccessPointsResponseBody {
	return s.Body
}

func (s *DescribeAccessPointsResponse) SetHeaders(v map[string]*string) *DescribeAccessPointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessPointsResponse) SetStatusCode(v int32) *DescribeAccessPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessPointsResponse) SetBody(v *DescribeAccessPointsResponseBody) *DescribeAccessPointsResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessPointsResponse) Validate() error {
	return dara.Validate(s)
}
