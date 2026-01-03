// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssuranceAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticityAssuranceAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticityAssuranceAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticityAssuranceAutoRenewAttributeResponseBody) *DescribeElasticityAssuranceAutoRenewAttributeResponse
	GetBody() *DescribeElasticityAssuranceAutoRenewAttributeResponseBody
}

type DescribeElasticityAssuranceAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticityAssuranceAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponse) GetBody() *DescribeElasticityAssuranceAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeElasticityAssuranceAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeElasticityAssuranceAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponse) SetBody(v *DescribeElasticityAssuranceAutoRenewAttributeResponseBody) *DescribeElasticityAssuranceAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
