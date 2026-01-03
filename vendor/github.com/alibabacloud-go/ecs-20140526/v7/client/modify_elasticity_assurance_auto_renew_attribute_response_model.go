// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticityAssuranceAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticityAssuranceAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticityAssuranceAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticityAssuranceAutoRenewAttributeResponseBody) *ModifyElasticityAssuranceAutoRenewAttributeResponse
	GetBody() *ModifyElasticityAssuranceAutoRenewAttributeResponseBody
}

type ModifyElasticityAssuranceAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticityAssuranceAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticityAssuranceAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponse) GetBody() *ModifyElasticityAssuranceAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyElasticityAssuranceAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponse) SetStatusCode(v int32) *ModifyElasticityAssuranceAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponse) SetBody(v *ModifyElasticityAssuranceAutoRenewAttributeResponseBody) *ModifyElasticityAssuranceAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticityAssuranceAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
