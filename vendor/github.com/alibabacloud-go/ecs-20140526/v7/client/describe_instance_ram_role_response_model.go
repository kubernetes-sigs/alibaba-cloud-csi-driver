// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRamRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceRamRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceRamRoleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceRamRoleResponseBody) *DescribeInstanceRamRoleResponse
	GetBody() *DescribeInstanceRamRoleResponseBody
}

type DescribeInstanceRamRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRamRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRamRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRamRoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRamRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceRamRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceRamRoleResponse) GetBody() *DescribeInstanceRamRoleResponseBody {
	return s.Body
}

func (s *DescribeInstanceRamRoleResponse) SetHeaders(v map[string]*string) *DescribeInstanceRamRoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRamRoleResponse) SetStatusCode(v int32) *DescribeInstanceRamRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRamRoleResponse) SetBody(v *DescribeInstanceRamRoleResponseBody) *DescribeInstanceRamRoleResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceRamRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
