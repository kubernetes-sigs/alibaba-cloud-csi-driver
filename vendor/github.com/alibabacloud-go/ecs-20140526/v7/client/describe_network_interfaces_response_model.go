// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkInterfacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkInterfacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkInterfacesResponseBody) *DescribeNetworkInterfacesResponse
	GetBody() *DescribeNetworkInterfacesResponseBody
}

type DescribeNetworkInterfacesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkInterfacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkInterfacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkInterfacesResponse) GetBody() *DescribeNetworkInterfacesResponseBody {
	return s.Body
}

func (s *DescribeNetworkInterfacesResponse) SetHeaders(v map[string]*string) *DescribeNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkInterfacesResponse) SetStatusCode(v int32) *DescribeNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkInterfacesResponse) SetBody(v *DescribeNetworkInterfacesResponseBody) *DescribeNetworkInterfacesResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkInterfacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
