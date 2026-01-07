// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImagesResponseBody) *DescribeImagesResponse
	GetBody() *DescribeImagesResponseBody
}

type DescribeImagesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImagesResponse) GetBody() *DescribeImagesResponseBody {
	return s.Body
}

func (s *DescribeImagesResponse) SetHeaders(v map[string]*string) *DescribeImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagesResponse) SetStatusCode(v int32) *DescribeImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImagesResponse) SetBody(v *DescribeImagesResponseBody) *DescribeImagesResponse {
	s.Body = v
	return s
}

func (s *DescribeImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
