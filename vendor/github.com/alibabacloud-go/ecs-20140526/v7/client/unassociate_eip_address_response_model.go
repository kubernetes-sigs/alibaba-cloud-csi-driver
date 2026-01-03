// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateEipAddressResponseBody) *UnassociateEipAddressResponse
	GetBody() *UnassociateEipAddressResponseBody
}

type UnassociateEipAddressResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateEipAddressResponse) GoString() string {
	return s.String()
}

func (s *UnassociateEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateEipAddressResponse) GetBody() *UnassociateEipAddressResponseBody {
	return s.Body
}

func (s *UnassociateEipAddressResponse) SetHeaders(v map[string]*string) *UnassociateEipAddressResponse {
	s.Headers = v
	return s
}

func (s *UnassociateEipAddressResponse) SetStatusCode(v int32) *UnassociateEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateEipAddressResponse) SetBody(v *UnassociateEipAddressResponseBody) *UnassociateEipAddressResponse {
	s.Body = v
	return s
}

func (s *UnassociateEipAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
