// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNfsAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNfsAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNfsAclResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNfsAclResponseBody) *DescribeNfsAclResponse
	GetBody() *DescribeNfsAclResponseBody
}

type DescribeNfsAclResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNfsAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNfsAclResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNfsAclResponse) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNfsAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNfsAclResponse) GetBody() *DescribeNfsAclResponseBody {
	return s.Body
}

func (s *DescribeNfsAclResponse) SetHeaders(v map[string]*string) *DescribeNfsAclResponse {
	s.Headers = v
	return s
}

func (s *DescribeNfsAclResponse) SetStatusCode(v int32) *DescribeNfsAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNfsAclResponse) SetBody(v *DescribeNfsAclResponseBody) *DescribeNfsAclResponse {
	s.Body = v
	return s
}

func (s *DescribeNfsAclResponse) Validate() error {
	return dara.Validate(s)
}
