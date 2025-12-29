// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFromFamilyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageFromFamilyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageFromFamilyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageFromFamilyResponseBody) *DescribeImageFromFamilyResponse
	GetBody() *DescribeImageFromFamilyResponseBody
}

type DescribeImageFromFamilyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageFromFamilyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageFromFamilyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFromFamilyResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageFromFamilyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageFromFamilyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageFromFamilyResponse) GetBody() *DescribeImageFromFamilyResponseBody {
	return s.Body
}

func (s *DescribeImageFromFamilyResponse) SetHeaders(v map[string]*string) *DescribeImageFromFamilyResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageFromFamilyResponse) SetStatusCode(v int32) *DescribeImageFromFamilyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageFromFamilyResponse) SetBody(v *DescribeImageFromFamilyResponseBody) *DescribeImageFromFamilyResponse {
	s.Body = v
	return s
}

func (s *DescribeImageFromFamilyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
