// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssurancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticityAssurancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticityAssurancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticityAssurancesResponseBody) *DescribeElasticityAssurancesResponse
	GetBody() *DescribeElasticityAssurancesResponseBody
}

type DescribeElasticityAssurancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticityAssurancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticityAssurancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssurancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssurancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticityAssurancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticityAssurancesResponse) GetBody() *DescribeElasticityAssurancesResponseBody {
	return s.Body
}

func (s *DescribeElasticityAssurancesResponse) SetHeaders(v map[string]*string) *DescribeElasticityAssurancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticityAssurancesResponse) SetStatusCode(v int32) *DescribeElasticityAssurancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticityAssurancesResponse) SetBody(v *DescribeElasticityAssurancesResponseBody) *DescribeElasticityAssurancesResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticityAssurancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
