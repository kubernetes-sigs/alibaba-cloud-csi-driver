// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouterInterfacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouterInterfacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouterInterfacesResponseBody) *DescribeRouterInterfacesResponse
	GetBody() *DescribeRouterInterfacesResponseBody
}

type DescribeRouterInterfacesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouterInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouterInterfacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouterInterfacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouterInterfacesResponse) GetBody() *DescribeRouterInterfacesResponseBody {
	return s.Body
}

func (s *DescribeRouterInterfacesResponse) SetHeaders(v map[string]*string) *DescribeRouterInterfacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouterInterfacesResponse) SetStatusCode(v int32) *DescribeRouterInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouterInterfacesResponse) SetBody(v *DescribeRouterInterfacesResponseBody) *DescribeRouterInterfacesResponse {
	s.Body = v
	return s
}

func (s *DescribeRouterInterfacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
