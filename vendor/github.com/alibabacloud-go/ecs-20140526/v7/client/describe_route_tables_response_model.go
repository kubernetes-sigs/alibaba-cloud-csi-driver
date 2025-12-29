// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouteTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouteTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouteTablesResponseBody) *DescribeRouteTablesResponse
	GetBody() *DescribeRouteTablesResponseBody
}

type DescribeRouteTablesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouteTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouteTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouteTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouteTablesResponse) GetBody() *DescribeRouteTablesResponseBody {
	return s.Body
}

func (s *DescribeRouteTablesResponse) SetHeaders(v map[string]*string) *DescribeRouteTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteTablesResponse) SetStatusCode(v int32) *DescribeRouteTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouteTablesResponse) SetBody(v *DescribeRouteTablesResponseBody) *DescribeRouteTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeRouteTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
