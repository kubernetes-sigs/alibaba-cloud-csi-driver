// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTopologyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceTopologyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceTopologyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceTopologyResponseBody) *DescribeInstanceTopologyResponse
	GetBody() *DescribeInstanceTopologyResponseBody
}

type DescribeInstanceTopologyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceTopologyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceTopologyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceTopologyResponse) GetBody() *DescribeInstanceTopologyResponseBody {
	return s.Body
}

func (s *DescribeInstanceTopologyResponse) SetHeaders(v map[string]*string) *DescribeInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTopologyResponse) SetStatusCode(v int32) *DescribeInstanceTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTopologyResponse) SetBody(v *DescribeInstanceTopologyResponseBody) *DescribeInstanceTopologyResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceTopologyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
