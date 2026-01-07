// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleasePublicIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleasePublicIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleasePublicIpAddressResponseBody) *ReleasePublicIpAddressResponse
	GetBody() *ReleasePublicIpAddressResponseBody
}

type ReleasePublicIpAddressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePublicIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePublicIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicIpAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleasePublicIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleasePublicIpAddressResponse) GetBody() *ReleasePublicIpAddressResponseBody {
	return s.Body
}

func (s *ReleasePublicIpAddressResponse) SetHeaders(v map[string]*string) *ReleasePublicIpAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicIpAddressResponse) SetStatusCode(v int32) *ReleasePublicIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePublicIpAddressResponse) SetBody(v *ReleasePublicIpAddressResponseBody) *ReleasePublicIpAddressResponse {
	s.Body = v
	return s
}

func (s *ReleasePublicIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
