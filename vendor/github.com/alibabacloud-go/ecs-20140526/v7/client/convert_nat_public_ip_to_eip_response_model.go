// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertNatPublicIpToEipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertNatPublicIpToEipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertNatPublicIpToEipResponse
	GetStatusCode() *int32
	SetBody(v *ConvertNatPublicIpToEipResponseBody) *ConvertNatPublicIpToEipResponse
	GetBody() *ConvertNatPublicIpToEipResponseBody
}

type ConvertNatPublicIpToEipResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertNatPublicIpToEipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertNatPublicIpToEipResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertNatPublicIpToEipResponse) GoString() string {
	return s.String()
}

func (s *ConvertNatPublicIpToEipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertNatPublicIpToEipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertNatPublicIpToEipResponse) GetBody() *ConvertNatPublicIpToEipResponseBody {
	return s.Body
}

func (s *ConvertNatPublicIpToEipResponse) SetHeaders(v map[string]*string) *ConvertNatPublicIpToEipResponse {
	s.Headers = v
	return s
}

func (s *ConvertNatPublicIpToEipResponse) SetStatusCode(v int32) *ConvertNatPublicIpToEipResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertNatPublicIpToEipResponse) SetBody(v *ConvertNatPublicIpToEipResponseBody) *ConvertNatPublicIpToEipResponse {
	s.Body = v
	return s
}

func (s *ConvertNatPublicIpToEipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
