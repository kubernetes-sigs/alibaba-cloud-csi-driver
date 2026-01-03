// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBandwidthPackageSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBandwidthPackageSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBandwidthPackageSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBandwidthPackageSpecResponseBody) *ModifyBandwidthPackageSpecResponse
	GetBody() *ModifyBandwidthPackageSpecResponseBody
}

type ModifyBandwidthPackageSpecResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBandwidthPackageSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBandwidthPackageSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBandwidthPackageSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyBandwidthPackageSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBandwidthPackageSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBandwidthPackageSpecResponse) GetBody() *ModifyBandwidthPackageSpecResponseBody {
	return s.Body
}

func (s *ModifyBandwidthPackageSpecResponse) SetHeaders(v map[string]*string) *ModifyBandwidthPackageSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyBandwidthPackageSpecResponse) SetStatusCode(v int32) *ModifyBandwidthPackageSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBandwidthPackageSpecResponse) SetBody(v *ModifyBandwidthPackageSpecResponseBody) *ModifyBandwidthPackageSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyBandwidthPackageSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
