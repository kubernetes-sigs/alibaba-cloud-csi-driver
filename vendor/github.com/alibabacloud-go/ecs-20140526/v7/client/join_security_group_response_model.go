// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *JoinSecurityGroupResponseBody) *JoinSecurityGroupResponse
	GetBody() *JoinSecurityGroupResponseBody
}

type JoinSecurityGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinSecurityGroupResponse) GetBody() *JoinSecurityGroupResponseBody {
	return s.Body
}

func (s *JoinSecurityGroupResponse) SetHeaders(v map[string]*string) *JoinSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinSecurityGroupResponse) SetStatusCode(v int32) *JoinSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinSecurityGroupResponse) SetBody(v *JoinSecurityGroupResponseBody) *JoinSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *JoinSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
