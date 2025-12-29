// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagePipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImagePipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImagePipelinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImagePipelinesResponseBody) *DescribeImagePipelinesResponse
	GetBody() *DescribeImagePipelinesResponseBody
}

type DescribeImagePipelinesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImagePipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImagePipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagePipelinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagePipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImagePipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImagePipelinesResponse) GetBody() *DescribeImagePipelinesResponseBody {
	return s.Body
}

func (s *DescribeImagePipelinesResponse) SetHeaders(v map[string]*string) *DescribeImagePipelinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagePipelinesResponse) SetStatusCode(v int32) *DescribeImagePipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImagePipelinesResponse) SetBody(v *DescribeImagePipelinesResponseBody) *DescribeImagePipelinesResponse {
	s.Body = v
	return s
}

func (s *DescribeImagePipelinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
