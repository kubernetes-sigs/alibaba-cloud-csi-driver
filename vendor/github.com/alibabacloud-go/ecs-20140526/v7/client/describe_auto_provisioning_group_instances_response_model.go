// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoProvisioningGroupInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoProvisioningGroupInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoProvisioningGroupInstancesResponseBody) *DescribeAutoProvisioningGroupInstancesResponse
	GetBody() *DescribeAutoProvisioningGroupInstancesResponseBody
}

type DescribeAutoProvisioningGroupInstancesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoProvisioningGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoProvisioningGroupInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoProvisioningGroupInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoProvisioningGroupInstancesResponse) GetBody() *DescribeAutoProvisioningGroupInstancesResponseBody {
	return s.Body
}

func (s *DescribeAutoProvisioningGroupInstancesResponse) SetHeaders(v map[string]*string) *DescribeAutoProvisioningGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponse) SetStatusCode(v int32) *DescribeAutoProvisioningGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponse) SetBody(v *DescribeAutoProvisioningGroupInstancesResponseBody) *DescribeAutoProvisioningGroupInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
