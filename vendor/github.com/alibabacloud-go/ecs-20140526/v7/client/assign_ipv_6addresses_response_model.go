// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignIpv6AddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignIpv6AddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignIpv6AddressesResponse
	GetStatusCode() *int32
	SetBody(v *AssignIpv6AddressesResponseBody) *AssignIpv6AddressesResponse
	GetBody() *AssignIpv6AddressesResponseBody
}

type AssignIpv6AddressesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignIpv6AddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignIpv6AddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignIpv6AddressesResponse) GoString() string {
	return s.String()
}

func (s *AssignIpv6AddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignIpv6AddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignIpv6AddressesResponse) GetBody() *AssignIpv6AddressesResponseBody {
	return s.Body
}

func (s *AssignIpv6AddressesResponse) SetHeaders(v map[string]*string) *AssignIpv6AddressesResponse {
	s.Headers = v
	return s
}

func (s *AssignIpv6AddressesResponse) SetStatusCode(v int32) *AssignIpv6AddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignIpv6AddressesResponse) SetBody(v *AssignIpv6AddressesResponseBody) *AssignIpv6AddressesResponse {
	s.Body = v
	return s
}

func (s *AssignIpv6AddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
