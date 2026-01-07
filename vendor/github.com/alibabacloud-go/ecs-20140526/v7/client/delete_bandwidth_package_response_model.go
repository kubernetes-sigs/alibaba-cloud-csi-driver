// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBandwidthPackageResponseBody) *DeleteBandwidthPackageResponse
	GetBody() *DeleteBandwidthPackageResponseBody
}

type DeleteBandwidthPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBandwidthPackageResponse) GetBody() *DeleteBandwidthPackageResponseBody {
	return s.Body
}

func (s *DeleteBandwidthPackageResponse) SetHeaders(v map[string]*string) *DeleteBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DeleteBandwidthPackageResponse) SetStatusCode(v int32) *DeleteBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBandwidthPackageResponse) SetBody(v *DeleteBandwidthPackageResponseBody) *DeleteBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *DeleteBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
