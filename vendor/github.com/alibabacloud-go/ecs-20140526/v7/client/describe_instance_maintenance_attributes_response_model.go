// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMaintenanceAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceMaintenanceAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceMaintenanceAttributesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceMaintenanceAttributesResponseBody) *DescribeInstanceMaintenanceAttributesResponse
	GetBody() *DescribeInstanceMaintenanceAttributesResponseBody
}

type DescribeInstanceMaintenanceAttributesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceMaintenanceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceMaintenanceAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceMaintenanceAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceMaintenanceAttributesResponse) GetBody() *DescribeInstanceMaintenanceAttributesResponseBody {
	return s.Body
}

func (s *DescribeInstanceMaintenanceAttributesResponse) SetHeaders(v map[string]*string) *DescribeInstanceMaintenanceAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponse) SetStatusCode(v int32) *DescribeInstanceMaintenanceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponse) SetBody(v *DescribeInstanceMaintenanceAttributesResponseBody) *DescribeInstanceMaintenanceAttributesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
