// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityGroupResponseBody) *DeleteSecurityGroupResponse
	GetBody() *DeleteSecurityGroupResponseBody
}

type DeleteSecurityGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityGroupResponse) GetBody() *DeleteSecurityGroupResponseBody {
	return s.Body
}

func (s *DeleteSecurityGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityGroupResponse) SetStatusCode(v int32) *DeleteSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityGroupResponse) SetBody(v *DeleteSecurityGroupResponseBody) *DeleteSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
