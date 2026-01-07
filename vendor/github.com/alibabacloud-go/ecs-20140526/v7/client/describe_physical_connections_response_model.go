// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhysicalConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhysicalConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhysicalConnectionsResponseBody) *DescribePhysicalConnectionsResponse
	GetBody() *DescribePhysicalConnectionsResponseBody
}

type DescribePhysicalConnectionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhysicalConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhysicalConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhysicalConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhysicalConnectionsResponse) GetBody() *DescribePhysicalConnectionsResponseBody {
	return s.Body
}

func (s *DescribePhysicalConnectionsResponse) SetHeaders(v map[string]*string) *DescribePhysicalConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePhysicalConnectionsResponse) SetStatusCode(v int32) *DescribePhysicalConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhysicalConnectionsResponse) SetBody(v *DescribePhysicalConnectionsResponseBody) *DescribePhysicalConnectionsResponse {
	s.Body = v
	return s
}

func (s *DescribePhysicalConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
