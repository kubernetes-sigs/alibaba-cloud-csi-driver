// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeSecurityGroupResponseBody) *AuthorizeSecurityGroupResponse
	GetBody() *AuthorizeSecurityGroupResponseBody
}

type AuthorizeSecurityGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeSecurityGroupResponse) GetBody() *AuthorizeSecurityGroupResponseBody {
	return s.Body
}

func (s *AuthorizeSecurityGroupResponse) SetHeaders(v map[string]*string) *AuthorizeSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeSecurityGroupResponse) SetStatusCode(v int32) *AuthorizeSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeSecurityGroupResponse) SetBody(v *AuthorizeSecurityGroupResponseBody) *AuthorizeSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *AuthorizeSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
