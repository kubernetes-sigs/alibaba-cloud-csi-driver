// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupEgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeSecurityGroupEgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeSecurityGroupEgressResponse
	GetStatusCode() *int32
	SetBody(v *RevokeSecurityGroupEgressResponseBody) *RevokeSecurityGroupEgressResponse
	GetBody() *RevokeSecurityGroupEgressResponseBody
}

type RevokeSecurityGroupEgressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeSecurityGroupEgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeSecurityGroupEgressResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupEgressResponse) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupEgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeSecurityGroupEgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeSecurityGroupEgressResponse) GetBody() *RevokeSecurityGroupEgressResponseBody {
	return s.Body
}

func (s *RevokeSecurityGroupEgressResponse) SetHeaders(v map[string]*string) *RevokeSecurityGroupEgressResponse {
	s.Headers = v
	return s
}

func (s *RevokeSecurityGroupEgressResponse) SetStatusCode(v int32) *RevokeSecurityGroupEgressResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeSecurityGroupEgressResponse) SetBody(v *RevokeSecurityGroupEgressResponseBody) *RevokeSecurityGroupEgressResponse {
	s.Body = v
	return s
}

func (s *RevokeSecurityGroupEgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
