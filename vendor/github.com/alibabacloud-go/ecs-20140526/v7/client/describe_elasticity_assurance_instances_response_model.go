// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssuranceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticityAssuranceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticityAssuranceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticityAssuranceInstancesResponseBody) *DescribeElasticityAssuranceInstancesResponse
	GetBody() *DescribeElasticityAssuranceInstancesResponseBody
}

type DescribeElasticityAssuranceInstancesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticityAssuranceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticityAssuranceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticityAssuranceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticityAssuranceInstancesResponse) GetBody() *DescribeElasticityAssuranceInstancesResponseBody {
	return s.Body
}

func (s *DescribeElasticityAssuranceInstancesResponse) SetHeaders(v map[string]*string) *DescribeElasticityAssuranceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponse) SetStatusCode(v int32) *DescribeElasticityAssuranceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponse) SetBody(v *DescribeElasticityAssuranceInstancesResponseBody) *DescribeElasticityAssuranceInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
