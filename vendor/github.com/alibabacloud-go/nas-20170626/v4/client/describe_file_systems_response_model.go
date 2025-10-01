// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileSystemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFileSystemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFileSystemsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFileSystemsResponseBody) *DescribeFileSystemsResponse
	GetBody() *DescribeFileSystemsResponseBody
}

type DescribeFileSystemsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileSystemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileSystemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFileSystemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFileSystemsResponse) GetBody() *DescribeFileSystemsResponseBody {
	return s.Body
}

func (s *DescribeFileSystemsResponse) SetHeaders(v map[string]*string) *DescribeFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileSystemsResponse) SetStatusCode(v int32) *DescribeFileSystemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileSystemsResponse) SetBody(v *DescribeFileSystemsResponseBody) *DescribeFileSystemsResponse {
	s.Body = v
	return s
}

func (s *DescribeFileSystemsResponse) Validate() error {
	return dara.Validate(s)
}
