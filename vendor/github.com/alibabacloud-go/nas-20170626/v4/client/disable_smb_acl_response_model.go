// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmbAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSmbAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSmbAclResponse
	GetStatusCode() *int32
	SetBody(v *DisableSmbAclResponseBody) *DisableSmbAclResponse
	GetBody() *DisableSmbAclResponseBody
}

type DisableSmbAclResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSmbAclResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSmbAclResponse) GoString() string {
	return s.String()
}

func (s *DisableSmbAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSmbAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSmbAclResponse) GetBody() *DisableSmbAclResponseBody {
	return s.Body
}

func (s *DisableSmbAclResponse) SetHeaders(v map[string]*string) *DisableSmbAclResponse {
	s.Headers = v
	return s
}

func (s *DisableSmbAclResponse) SetStatusCode(v int32) *DisableSmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSmbAclResponse) SetBody(v *DisableSmbAclResponseBody) *DisableSmbAclResponse {
	s.Body = v
	return s
}

func (s *DisableSmbAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
