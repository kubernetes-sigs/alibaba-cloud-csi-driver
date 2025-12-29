// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewElasticityAssurancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewElasticityAssurancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewElasticityAssurancesResponse
	GetStatusCode() *int32
	SetBody(v *RenewElasticityAssurancesResponseBody) *RenewElasticityAssurancesResponse
	GetBody() *RenewElasticityAssurancesResponseBody
}

type RenewElasticityAssurancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewElasticityAssurancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewElasticityAssurancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewElasticityAssurancesResponse) GoString() string {
	return s.String()
}

func (s *RenewElasticityAssurancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewElasticityAssurancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewElasticityAssurancesResponse) GetBody() *RenewElasticityAssurancesResponseBody {
	return s.Body
}

func (s *RenewElasticityAssurancesResponse) SetHeaders(v map[string]*string) *RenewElasticityAssurancesResponse {
	s.Headers = v
	return s
}

func (s *RenewElasticityAssurancesResponse) SetStatusCode(v int32) *RenewElasticityAssurancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewElasticityAssurancesResponse) SetBody(v *RenewElasticityAssurancesResponseBody) *RenewElasticityAssurancesResponse {
	s.Body = v
	return s
}

func (s *RenewElasticityAssurancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
