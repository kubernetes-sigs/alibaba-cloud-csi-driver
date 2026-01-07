// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyPairsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKeyPairsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKeyPairsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKeyPairsResponseBody) *DescribeKeyPairsResponse
	GetBody() *DescribeKeyPairsResponseBody
}

type DescribeKeyPairsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeyPairsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKeyPairsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKeyPairsResponse) GetBody() *DescribeKeyPairsResponseBody {
	return s.Body
}

func (s *DescribeKeyPairsResponse) SetHeaders(v map[string]*string) *DescribeKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeyPairsResponse) SetStatusCode(v int32) *DescribeKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeyPairsResponse) SetBody(v *DescribeKeyPairsResponseBody) *DescribeKeyPairsResponse {
	s.Body = v
	return s
}

func (s *DescribeKeyPairsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
