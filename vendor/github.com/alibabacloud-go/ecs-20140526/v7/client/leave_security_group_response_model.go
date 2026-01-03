// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LeaveSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LeaveSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *LeaveSecurityGroupResponseBody) *LeaveSecurityGroupResponse
	GetBody() *LeaveSecurityGroupResponseBody
}

type LeaveSecurityGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LeaveSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LeaveSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s LeaveSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *LeaveSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LeaveSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LeaveSecurityGroupResponse) GetBody() *LeaveSecurityGroupResponseBody {
	return s.Body
}

func (s *LeaveSecurityGroupResponse) SetHeaders(v map[string]*string) *LeaveSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *LeaveSecurityGroupResponse) SetStatusCode(v int32) *LeaveSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *LeaveSecurityGroupResponse) SetBody(v *LeaveSecurityGroupResponseBody) *LeaveSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *LeaveSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
