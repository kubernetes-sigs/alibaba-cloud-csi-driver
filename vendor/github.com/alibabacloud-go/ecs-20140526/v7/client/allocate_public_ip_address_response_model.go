// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocatePublicIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocatePublicIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocatePublicIpAddressResponseBody) *AllocatePublicIpAddressResponse
	GetBody() *AllocatePublicIpAddressResponseBody
}

type AllocatePublicIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocatePublicIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocatePublicIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicIpAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocatePublicIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocatePublicIpAddressResponse) GetBody() *AllocatePublicIpAddressResponseBody {
	return s.Body
}

func (s *AllocatePublicIpAddressResponse) SetHeaders(v map[string]*string) *AllocatePublicIpAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicIpAddressResponse) SetStatusCode(v int32) *AllocatePublicIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePublicIpAddressResponse) SetBody(v *AllocatePublicIpAddressResponseBody) *AllocatePublicIpAddressResponse {
	s.Body = v
	return s
}

func (s *AllocatePublicIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
