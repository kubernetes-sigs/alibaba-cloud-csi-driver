// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountAttributesResponseBody) *DescribeAccountAttributesResponse
	GetBody() *DescribeAccountAttributesResponseBody
}

type DescribeAccountAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountAttributesResponse) GetBody() *DescribeAccountAttributesResponseBody {
	return s.Body
}

func (s *DescribeAccountAttributesResponse) SetHeaders(v map[string]*string) *DescribeAccountAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountAttributesResponse) SetStatusCode(v int32) *DescribeAccountAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountAttributesResponse) SetBody(v *DescribeAccountAttributesResponseBody) *DescribeAccountAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
