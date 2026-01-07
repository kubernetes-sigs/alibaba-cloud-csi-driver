// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *RevokeSecurityGroupResponseBody) *RevokeSecurityGroupResponse
	GetBody() *RevokeSecurityGroupResponseBody
}

type RevokeSecurityGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeSecurityGroupResponse) GetBody() *RevokeSecurityGroupResponseBody {
	return s.Body
}

func (s *RevokeSecurityGroupResponse) SetHeaders(v map[string]*string) *RevokeSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *RevokeSecurityGroupResponse) SetStatusCode(v int32) *RevokeSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeSecurityGroupResponse) SetBody(v *RevokeSecurityGroupResponseBody) *RevokeSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *RevokeSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
