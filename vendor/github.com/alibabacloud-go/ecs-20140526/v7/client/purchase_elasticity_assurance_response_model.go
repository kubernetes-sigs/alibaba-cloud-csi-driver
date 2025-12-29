// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseElasticityAssuranceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurchaseElasticityAssuranceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurchaseElasticityAssuranceResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseElasticityAssuranceResponseBody) *PurchaseElasticityAssuranceResponse
	GetBody() *PurchaseElasticityAssuranceResponseBody
}

type PurchaseElasticityAssuranceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseElasticityAssuranceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseElasticityAssuranceResponse) String() string {
	return dara.Prettify(s)
}

func (s PurchaseElasticityAssuranceResponse) GoString() string {
	return s.String()
}

func (s *PurchaseElasticityAssuranceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurchaseElasticityAssuranceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurchaseElasticityAssuranceResponse) GetBody() *PurchaseElasticityAssuranceResponseBody {
	return s.Body
}

func (s *PurchaseElasticityAssuranceResponse) SetHeaders(v map[string]*string) *PurchaseElasticityAssuranceResponse {
	s.Headers = v
	return s
}

func (s *PurchaseElasticityAssuranceResponse) SetStatusCode(v int32) *PurchaseElasticityAssuranceResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseElasticityAssuranceResponse) SetBody(v *PurchaseElasticityAssuranceResponseBody) *PurchaseElasticityAssuranceResponse {
	s.Body = v
	return s
}

func (s *PurchaseElasticityAssuranceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
