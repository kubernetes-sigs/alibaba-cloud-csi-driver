// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoProvisioningGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoProvisioningGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoProvisioningGroupsResponseBody) *DescribeAutoProvisioningGroupsResponse
	GetBody() *DescribeAutoProvisioningGroupsResponseBody
}

type DescribeAutoProvisioningGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoProvisioningGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoProvisioningGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoProvisioningGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoProvisioningGroupsResponse) GetBody() *DescribeAutoProvisioningGroupsResponseBody {
	return s.Body
}

func (s *DescribeAutoProvisioningGroupsResponse) SetHeaders(v map[string]*string) *DescribeAutoProvisioningGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponse) SetStatusCode(v int32) *DescribeAutoProvisioningGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponse) SetBody(v *DescribeAutoProvisioningGroupsResponseBody) *DescribeAutoProvisioningGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoProvisioningGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
