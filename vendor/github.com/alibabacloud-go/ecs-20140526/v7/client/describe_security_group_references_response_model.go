// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupReferencesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityGroupReferencesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityGroupReferencesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityGroupReferencesResponseBody) *DescribeSecurityGroupReferencesResponse
	GetBody() *DescribeSecurityGroupReferencesResponseBody
}

type DescribeSecurityGroupReferencesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityGroupReferencesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityGroupReferencesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupReferencesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupReferencesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityGroupReferencesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityGroupReferencesResponse) GetBody() *DescribeSecurityGroupReferencesResponseBody {
	return s.Body
}

func (s *DescribeSecurityGroupReferencesResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupReferencesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupReferencesResponse) SetStatusCode(v int32) *DescribeSecurityGroupReferencesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityGroupReferencesResponse) SetBody(v *DescribeSecurityGroupReferencesResponseBody) *DescribeSecurityGroupReferencesResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityGroupReferencesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
