// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticityAssuranceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticityAssuranceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticityAssuranceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticityAssuranceResponseBody) *ModifyElasticityAssuranceResponse
	GetBody() *ModifyElasticityAssuranceResponseBody
}

type ModifyElasticityAssuranceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticityAssuranceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticityAssuranceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticityAssuranceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticityAssuranceResponse) GetBody() *ModifyElasticityAssuranceResponseBody {
	return s.Body
}

func (s *ModifyElasticityAssuranceResponse) SetHeaders(v map[string]*string) *ModifyElasticityAssuranceResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticityAssuranceResponse) SetStatusCode(v int32) *ModifyElasticityAssuranceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticityAssuranceResponse) SetBody(v *ModifyElasticityAssuranceResponseBody) *ModifyElasticityAssuranceResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticityAssuranceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
