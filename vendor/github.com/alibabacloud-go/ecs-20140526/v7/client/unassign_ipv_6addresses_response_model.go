// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignIpv6AddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassignIpv6AddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassignIpv6AddressesResponse
	GetStatusCode() *int32
	SetBody(v *UnassignIpv6AddressesResponseBody) *UnassignIpv6AddressesResponse
	GetBody() *UnassignIpv6AddressesResponseBody
}

type UnassignIpv6AddressesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassignIpv6AddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassignIpv6AddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassignIpv6AddressesResponse) GoString() string {
	return s.String()
}

func (s *UnassignIpv6AddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassignIpv6AddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassignIpv6AddressesResponse) GetBody() *UnassignIpv6AddressesResponseBody {
	return s.Body
}

func (s *UnassignIpv6AddressesResponse) SetHeaders(v map[string]*string) *UnassignIpv6AddressesResponse {
	s.Headers = v
	return s
}

func (s *UnassignIpv6AddressesResponse) SetStatusCode(v int32) *UnassignIpv6AddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassignIpv6AddressesResponse) SetBody(v *UnassignIpv6AddressesResponseBody) *UnassignIpv6AddressesResponse {
	s.Body = v
	return s
}

func (s *UnassignIpv6AddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
