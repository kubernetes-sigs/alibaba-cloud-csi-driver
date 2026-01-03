// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceByTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceByTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceByTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceByTagsResponseBody) *DescribeResourceByTagsResponse
	GetBody() *DescribeResourceByTagsResponseBody
}

type DescribeResourceByTagsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceByTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceByTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceByTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceByTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceByTagsResponse) GetBody() *DescribeResourceByTagsResponseBody {
	return s.Body
}

func (s *DescribeResourceByTagsResponse) SetHeaders(v map[string]*string) *DescribeResourceByTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceByTagsResponse) SetStatusCode(v int32) *DescribeResourceByTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceByTagsResponse) SetBody(v *DescribeResourceByTagsResponseBody) *DescribeResourceByTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceByTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
