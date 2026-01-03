// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignPrivateIpAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassignPrivateIpAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassignPrivateIpAddressesResponse
	GetStatusCode() *int32
	SetBody(v *UnassignPrivateIpAddressesResponseBody) *UnassignPrivateIpAddressesResponse
	GetBody() *UnassignPrivateIpAddressesResponseBody
}

type UnassignPrivateIpAddressesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassignPrivateIpAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassignPrivateIpAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassignPrivateIpAddressesResponse) GoString() string {
	return s.String()
}

func (s *UnassignPrivateIpAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassignPrivateIpAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassignPrivateIpAddressesResponse) GetBody() *UnassignPrivateIpAddressesResponseBody {
	return s.Body
}

func (s *UnassignPrivateIpAddressesResponse) SetHeaders(v map[string]*string) *UnassignPrivateIpAddressesResponse {
	s.Headers = v
	return s
}

func (s *UnassignPrivateIpAddressesResponse) SetStatusCode(v int32) *UnassignPrivateIpAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassignPrivateIpAddressesResponse) SetBody(v *UnassignPrivateIpAddressesResponseBody) *UnassignPrivateIpAddressesResponse {
	s.Body = v
	return s
}

func (s *UnassignPrivateIpAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
