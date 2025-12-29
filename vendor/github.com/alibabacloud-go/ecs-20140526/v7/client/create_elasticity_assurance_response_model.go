// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticityAssuranceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateElasticityAssuranceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateElasticityAssuranceResponse
	GetStatusCode() *int32
	SetBody(v *CreateElasticityAssuranceResponseBody) *CreateElasticityAssuranceResponse
	GetBody() *CreateElasticityAssuranceResponseBody
}

type CreateElasticityAssuranceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateElasticityAssuranceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateElasticityAssuranceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceResponse) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateElasticityAssuranceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateElasticityAssuranceResponse) GetBody() *CreateElasticityAssuranceResponseBody {
	return s.Body
}

func (s *CreateElasticityAssuranceResponse) SetHeaders(v map[string]*string) *CreateElasticityAssuranceResponse {
	s.Headers = v
	return s
}

func (s *CreateElasticityAssuranceResponse) SetStatusCode(v int32) *CreateElasticityAssuranceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateElasticityAssuranceResponse) SetBody(v *CreateElasticityAssuranceResponseBody) *CreateElasticityAssuranceResponse {
	s.Body = v
	return s
}

func (s *CreateElasticityAssuranceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
