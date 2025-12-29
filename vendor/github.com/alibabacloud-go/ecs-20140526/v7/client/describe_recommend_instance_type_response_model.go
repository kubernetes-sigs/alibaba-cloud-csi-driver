// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecommendInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecommendInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecommendInstanceTypeResponseBody) *DescribeRecommendInstanceTypeResponse
	GetBody() *DescribeRecommendInstanceTypeResponseBody
}

type DescribeRecommendInstanceTypeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecommendInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecommendInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecommendInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecommendInstanceTypeResponse) GetBody() *DescribeRecommendInstanceTypeResponseBody {
	return s.Body
}

func (s *DescribeRecommendInstanceTypeResponse) SetHeaders(v map[string]*string) *DescribeRecommendInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponse) SetStatusCode(v int32) *DescribeRecommendInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecommendInstanceTypeResponse) SetBody(v *DescribeRecommendInstanceTypeResponseBody) *DescribeRecommendInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeRecommendInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
