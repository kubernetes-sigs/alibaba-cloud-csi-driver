// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstanceRamRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachInstanceRamRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachInstanceRamRoleResponse
	GetStatusCode() *int32
	SetBody(v *DetachInstanceRamRoleResponseBody) *DetachInstanceRamRoleResponse
	GetBody() *DetachInstanceRamRoleResponseBody
}

type DetachInstanceRamRoleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachInstanceRamRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachInstanceRamRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceRamRoleResponse) GoString() string {
	return s.String()
}

func (s *DetachInstanceRamRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachInstanceRamRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachInstanceRamRoleResponse) GetBody() *DetachInstanceRamRoleResponseBody {
	return s.Body
}

func (s *DetachInstanceRamRoleResponse) SetHeaders(v map[string]*string) *DetachInstanceRamRoleResponse {
	s.Headers = v
	return s
}

func (s *DetachInstanceRamRoleResponse) SetStatusCode(v int32) *DetachInstanceRamRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachInstanceRamRoleResponse) SetBody(v *DetachInstanceRamRoleResponseBody) *DetachInstanceRamRoleResponse {
	s.Body = v
	return s
}

func (s *DetachInstanceRamRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
