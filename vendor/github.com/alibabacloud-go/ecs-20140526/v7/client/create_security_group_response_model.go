// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityGroupResponseBody) *CreateSecurityGroupResponse
	GetBody() *CreateSecurityGroupResponseBody
}

type CreateSecurityGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityGroupResponse) GetBody() *CreateSecurityGroupResponseBody {
	return s.Body
}

func (s *CreateSecurityGroupResponse) SetHeaders(v map[string]*string) *CreateSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityGroupResponse) SetStatusCode(v int32) *CreateSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityGroupResponse) SetBody(v *CreateSecurityGroupResponseBody) *CreateSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
