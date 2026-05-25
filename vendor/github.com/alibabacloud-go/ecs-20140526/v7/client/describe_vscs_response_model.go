// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVscsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVscsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVscsResponseBody) *DescribeVscsResponse
	GetBody() *DescribeVscsResponseBody
}

type DescribeVscsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVscsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVscsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVscsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVscsResponse) GetBody() *DescribeVscsResponseBody {
	return s.Body
}

func (s *DescribeVscsResponse) SetHeaders(v map[string]*string) *DescribeVscsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVscsResponse) SetStatusCode(v int32) *DescribeVscsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVscsResponse) SetBody(v *DescribeVscsResponseBody) *DescribeVscsResponse {
	s.Body = v
	return s
}

func (s *DescribeVscsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
