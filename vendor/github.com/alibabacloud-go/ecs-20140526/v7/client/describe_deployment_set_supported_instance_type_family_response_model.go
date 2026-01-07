// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentSetSupportedInstanceTypeFamilyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse
	GetBody() *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody
}

type DescribeDeploymentSetSupportedInstanceTypeFamilyResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) GetBody() *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody {
	return s.Body
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) SetHeaders(v map[string]*string) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) SetStatusCode(v int32) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) SetBody(v *DescribeDeploymentSetSupportedInstanceTypeFamilyResponseBody) *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse {
	s.Body = v
	return s
}

func (s *DescribeDeploymentSetSupportedInstanceTypeFamilyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
