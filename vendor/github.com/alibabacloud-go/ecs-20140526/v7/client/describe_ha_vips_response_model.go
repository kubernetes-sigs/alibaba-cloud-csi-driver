// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHaVipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHaVipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHaVipsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHaVipsResponseBody) *DescribeHaVipsResponse
	GetBody() *DescribeHaVipsResponseBody
}

type DescribeHaVipsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHaVipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHaVipsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHaVipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHaVipsResponse) GetBody() *DescribeHaVipsResponseBody {
	return s.Body
}

func (s *DescribeHaVipsResponse) SetHeaders(v map[string]*string) *DescribeHaVipsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHaVipsResponse) SetStatusCode(v int32) *DescribeHaVipsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHaVipsResponse) SetBody(v *DescribeHaVipsResponseBody) *DescribeHaVipsResponse {
	s.Body = v
	return s
}

func (s *DescribeHaVipsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
