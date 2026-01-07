// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDisksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDisksResponseBody) *DescribeDisksResponse
	GetBody() *DescribeDisksResponseBody
}

type DescribeDisksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDisksResponse) GetBody() *DescribeDisksResponseBody {
	return s.Body
}

func (s *DescribeDisksResponse) SetHeaders(v map[string]*string) *DescribeDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisksResponse) SetStatusCode(v int32) *DescribeDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisksResponse) SetBody(v *DescribeDisksResponseBody) *DescribeDisksResponse {
	s.Body = v
	return s
}

func (s *DescribeDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
