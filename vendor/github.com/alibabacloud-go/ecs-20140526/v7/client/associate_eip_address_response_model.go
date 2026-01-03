// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *AssociateEipAddressResponseBody) *AssociateEipAddressResponse
	GetBody() *AssociateEipAddressResponseBody
}

type AssociateEipAddressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateEipAddressResponse) GetBody() *AssociateEipAddressResponseBody {
	return s.Body
}

func (s *AssociateEipAddressResponse) SetHeaders(v map[string]*string) *AssociateEipAddressResponse {
	s.Headers = v
	return s
}

func (s *AssociateEipAddressResponse) SetStatusCode(v int32) *AssociateEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateEipAddressResponse) SetBody(v *AssociateEipAddressResponseBody) *AssociateEipAddressResponse {
	s.Body = v
	return s
}

func (s *AssociateEipAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
