// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmbAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSmbAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSmbAclResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSmbAclResponseBody) *DescribeSmbAclResponse
	GetBody() *DescribeSmbAclResponseBody
}

type DescribeSmbAclResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSmbAclResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmbAclResponse) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSmbAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSmbAclResponse) GetBody() *DescribeSmbAclResponseBody {
	return s.Body
}

func (s *DescribeSmbAclResponse) SetHeaders(v map[string]*string) *DescribeSmbAclResponse {
	s.Headers = v
	return s
}

func (s *DescribeSmbAclResponse) SetStatusCode(v int32) *DescribeSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSmbAclResponse) SetBody(v *DescribeSmbAclResponseBody) *DescribeSmbAclResponse {
	s.Body = v
	return s
}

func (s *DescribeSmbAclResponse) Validate() error {
	return dara.Validate(s)
}
