// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeploymentSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeploymentSetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeploymentSetsResponseBody) *DescribeDeploymentSetsResponse
	GetBody() *DescribeDeploymentSetsResponseBody
}

type DescribeDeploymentSetsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeploymentSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeploymentSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeploymentSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeploymentSetsResponse) GetBody() *DescribeDeploymentSetsResponseBody {
	return s.Body
}

func (s *DescribeDeploymentSetsResponse) SetHeaders(v map[string]*string) *DescribeDeploymentSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeploymentSetsResponse) SetStatusCode(v int32) *DescribeDeploymentSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeploymentSetsResponse) SetBody(v *DescribeDeploymentSetsResponseBody) *DescribeDeploymentSetsResponse {
	s.Body = v
	return s
}

func (s *DescribeDeploymentSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
