// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserDataResponseBody) *DescribeUserDataResponse
	GetBody() *DescribeUserDataResponseBody
}

type DescribeUserDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserDataResponse) GetBody() *DescribeUserDataResponseBody {
	return s.Body
}

func (s *DescribeUserDataResponse) SetHeaders(v map[string]*string) *DescribeUserDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDataResponse) SetStatusCode(v int32) *DescribeUserDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserDataResponse) SetBody(v *DescribeUserDataResponseBody) *DescribeUserDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUserDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
