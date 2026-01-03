// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseEipAddressResponseBody) *ReleaseEipAddressResponse
	GetBody() *ReleaseEipAddressResponseBody
}

type ReleaseEipAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseEipAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseEipAddressResponse) GetBody() *ReleaseEipAddressResponseBody {
	return s.Body
}

func (s *ReleaseEipAddressResponse) SetHeaders(v map[string]*string) *ReleaseEipAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseEipAddressResponse) SetStatusCode(v int32) *ReleaseEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseEipAddressResponse) SetBody(v *ReleaseEipAddressResponseBody) *ReleaseEipAddressResponse {
	s.Body = v
	return s
}

func (s *ReleaseEipAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
