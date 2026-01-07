// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse
	GetBody() *DescribeTagsResponseBody
}

type DescribeTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagsResponse) GetBody() *DescribeTagsResponseBody {
	return s.Body
}

func (s *DescribeTagsResponse) SetHeaders(v map[string]*string) *DescribeTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsResponse) SetStatusCode(v int32) *DescribeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsResponse) SetBody(v *DescribeTagsResponseBody) *DescribeTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
