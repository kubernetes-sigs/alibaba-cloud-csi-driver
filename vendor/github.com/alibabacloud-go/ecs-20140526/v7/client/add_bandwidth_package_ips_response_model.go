// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBandwidthPackageIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBandwidthPackageIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBandwidthPackageIpsResponse
	GetStatusCode() *int32
	SetBody(v *AddBandwidthPackageIpsResponseBody) *AddBandwidthPackageIpsResponse
	GetBody() *AddBandwidthPackageIpsResponseBody
}

type AddBandwidthPackageIpsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBandwidthPackageIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBandwidthPackageIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBandwidthPackageIpsResponse) GoString() string {
	return s.String()
}

func (s *AddBandwidthPackageIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBandwidthPackageIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBandwidthPackageIpsResponse) GetBody() *AddBandwidthPackageIpsResponseBody {
	return s.Body
}

func (s *AddBandwidthPackageIpsResponse) SetHeaders(v map[string]*string) *AddBandwidthPackageIpsResponse {
	s.Headers = v
	return s
}

func (s *AddBandwidthPackageIpsResponse) SetStatusCode(v int32) *AddBandwidthPackageIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBandwidthPackageIpsResponse) SetBody(v *AddBandwidthPackageIpsResponseBody) *AddBandwidthPackageIpsResponse {
	s.Body = v
	return s
}

func (s *AddBandwidthPackageIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
