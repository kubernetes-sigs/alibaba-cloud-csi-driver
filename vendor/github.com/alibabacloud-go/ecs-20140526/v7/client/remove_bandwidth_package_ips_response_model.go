// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBandwidthPackageIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveBandwidthPackageIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveBandwidthPackageIpsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveBandwidthPackageIpsResponseBody) *RemoveBandwidthPackageIpsResponse
	GetBody() *RemoveBandwidthPackageIpsResponseBody
}

type RemoveBandwidthPackageIpsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveBandwidthPackageIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveBandwidthPackageIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveBandwidthPackageIpsResponse) GoString() string {
	return s.String()
}

func (s *RemoveBandwidthPackageIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveBandwidthPackageIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveBandwidthPackageIpsResponse) GetBody() *RemoveBandwidthPackageIpsResponseBody {
	return s.Body
}

func (s *RemoveBandwidthPackageIpsResponse) SetHeaders(v map[string]*string) *RemoveBandwidthPackageIpsResponse {
	s.Headers = v
	return s
}

func (s *RemoveBandwidthPackageIpsResponse) SetStatusCode(v int32) *RemoveBandwidthPackageIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBandwidthPackageIpsResponse) SetBody(v *RemoveBandwidthPackageIpsResponseBody) *RemoveBandwidthPackageIpsResponse {
	s.Body = v
	return s
}

func (s *RemoveBandwidthPackageIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
