// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessGroupsResponseBody) *DescribeAccessGroupsResponse
	GetBody() *DescribeAccessGroupsResponseBody
}

type DescribeAccessGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessGroupsResponse) GetBody() *DescribeAccessGroupsResponseBody {
	return s.Body
}

func (s *DescribeAccessGroupsResponse) SetHeaders(v map[string]*string) *DescribeAccessGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessGroupsResponse) SetStatusCode(v int32) *DescribeAccessGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessGroupsResponse) SetBody(v *DescribeAccessGroupsResponseBody) *DescribeAccessGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessGroupsResponse) Validate() error {
	return dara.Validate(s)
}
