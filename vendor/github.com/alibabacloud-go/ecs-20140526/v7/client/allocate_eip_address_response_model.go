// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocateEipAddressResponseBody) *AllocateEipAddressResponse
	GetBody() *AllocateEipAddressResponseBody
}

type AllocateEipAddressResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateEipAddressResponse) GetBody() *AllocateEipAddressResponseBody {
	return s.Body
}

func (s *AllocateEipAddressResponse) SetHeaders(v map[string]*string) *AllocateEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateEipAddressResponse) SetStatusCode(v int32) *AllocateEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateEipAddressResponse) SetBody(v *AllocateEipAddressResponseBody) *AllocateEipAddressResponse {
	s.Body = v
	return s
}

func (s *AllocateEipAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
