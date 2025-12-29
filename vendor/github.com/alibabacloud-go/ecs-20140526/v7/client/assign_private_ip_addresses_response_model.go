// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignPrivateIpAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignPrivateIpAddressesResponse
	GetStatusCode() *int32
	SetBody(v *AssignPrivateIpAddressesResponseBody) *AssignPrivateIpAddressesResponse
	GetBody() *AssignPrivateIpAddressesResponseBody
}

type AssignPrivateIpAddressesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignPrivateIpAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignPrivateIpAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressesResponse) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignPrivateIpAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignPrivateIpAddressesResponse) GetBody() *AssignPrivateIpAddressesResponseBody {
	return s.Body
}

func (s *AssignPrivateIpAddressesResponse) SetHeaders(v map[string]*string) *AssignPrivateIpAddressesResponse {
	s.Headers = v
	return s
}

func (s *AssignPrivateIpAddressesResponse) SetStatusCode(v int32) *AssignPrivateIpAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignPrivateIpAddressesResponse) SetBody(v *AssignPrivateIpAddressesResponseBody) *AssignPrivateIpAddressesResponse {
	s.Body = v
	return s
}

func (s *AssignPrivateIpAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
