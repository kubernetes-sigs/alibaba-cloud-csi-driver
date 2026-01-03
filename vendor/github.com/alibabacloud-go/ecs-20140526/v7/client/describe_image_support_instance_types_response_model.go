// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSupportInstanceTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageSupportInstanceTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageSupportInstanceTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageSupportInstanceTypesResponseBody) *DescribeImageSupportInstanceTypesResponse
	GetBody() *DescribeImageSupportInstanceTypesResponseBody
}

type DescribeImageSupportInstanceTypesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageSupportInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageSupportInstanceTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSupportInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageSupportInstanceTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageSupportInstanceTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageSupportInstanceTypesResponse) GetBody() *DescribeImageSupportInstanceTypesResponseBody {
	return s.Body
}

func (s *DescribeImageSupportInstanceTypesResponse) SetHeaders(v map[string]*string) *DescribeImageSupportInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponse) SetStatusCode(v int32) *DescribeImageSupportInstanceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponse) SetBody(v *DescribeImageSupportInstanceTypesResponseBody) *DescribeImageSupportInstanceTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
