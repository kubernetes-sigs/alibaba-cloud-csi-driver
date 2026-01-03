// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupEgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeSecurityGroupEgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeSecurityGroupEgressResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeSecurityGroupEgressResponseBody) *AuthorizeSecurityGroupEgressResponse
	GetBody() *AuthorizeSecurityGroupEgressResponseBody
}

type AuthorizeSecurityGroupEgressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeSecurityGroupEgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeSecurityGroupEgressResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeSecurityGroupEgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeSecurityGroupEgressResponse) GetBody() *AuthorizeSecurityGroupEgressResponseBody {
	return s.Body
}

func (s *AuthorizeSecurityGroupEgressResponse) SetHeaders(v map[string]*string) *AuthorizeSecurityGroupEgressResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeSecurityGroupEgressResponse) SetStatusCode(v int32) *AuthorizeSecurityGroupEgressResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressResponse) SetBody(v *AuthorizeSecurityGroupEgressResponseBody) *AuthorizeSecurityGroupEgressResponse {
	s.Body = v
	return s
}

func (s *AuthorizeSecurityGroupEgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
