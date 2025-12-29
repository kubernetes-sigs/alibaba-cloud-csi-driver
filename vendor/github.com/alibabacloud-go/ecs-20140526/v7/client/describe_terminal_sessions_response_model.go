// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTerminalSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTerminalSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTerminalSessionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTerminalSessionsResponseBody) *DescribeTerminalSessionsResponse
	GetBody() *DescribeTerminalSessionsResponseBody
}

type DescribeTerminalSessionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTerminalSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTerminalSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSessionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTerminalSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTerminalSessionsResponse) GetBody() *DescribeTerminalSessionsResponseBody {
	return s.Body
}

func (s *DescribeTerminalSessionsResponse) SetHeaders(v map[string]*string) *DescribeTerminalSessionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTerminalSessionsResponse) SetStatusCode(v int32) *DescribeTerminalSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTerminalSessionsResponse) SetBody(v *DescribeTerminalSessionsResponseBody) *DescribeTerminalSessionsResponse {
	s.Body = v
	return s
}

func (s *DescribeTerminalSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
