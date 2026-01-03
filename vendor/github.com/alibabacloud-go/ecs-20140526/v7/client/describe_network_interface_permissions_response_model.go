// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkInterfacePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkInterfacePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkInterfacePermissionsResponseBody) *DescribeNetworkInterfacePermissionsResponse
	GetBody() *DescribeNetworkInterfacePermissionsResponseBody
}

type DescribeNetworkInterfacePermissionsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkInterfacePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkInterfacePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacePermissionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkInterfacePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkInterfacePermissionsResponse) GetBody() *DescribeNetworkInterfacePermissionsResponseBody {
	return s.Body
}

func (s *DescribeNetworkInterfacePermissionsResponse) SetHeaders(v map[string]*string) *DescribeNetworkInterfacePermissionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponse) SetStatusCode(v int32) *DescribeNetworkInterfacePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponse) SetBody(v *DescribeNetworkInterfacePermissionsResponseBody) *DescribeNetworkInterfacePermissionsResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkInterfacePermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
