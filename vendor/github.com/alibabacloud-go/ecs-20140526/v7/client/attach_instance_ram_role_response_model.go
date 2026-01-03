// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstanceRamRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachInstanceRamRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachInstanceRamRoleResponse
	GetStatusCode() *int32
	SetBody(v *AttachInstanceRamRoleResponseBody) *AttachInstanceRamRoleResponse
	GetBody() *AttachInstanceRamRoleResponseBody
}

type AttachInstanceRamRoleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachInstanceRamRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachInstanceRamRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceRamRoleResponse) GoString() string {
	return s.String()
}

func (s *AttachInstanceRamRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachInstanceRamRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachInstanceRamRoleResponse) GetBody() *AttachInstanceRamRoleResponseBody {
	return s.Body
}

func (s *AttachInstanceRamRoleResponse) SetHeaders(v map[string]*string) *AttachInstanceRamRoleResponse {
	s.Headers = v
	return s
}

func (s *AttachInstanceRamRoleResponse) SetStatusCode(v int32) *AttachInstanceRamRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachInstanceRamRoleResponse) SetBody(v *AttachInstanceRamRoleResponseBody) *AttachInstanceRamRoleResponse {
	s.Body = v
	return s
}

func (s *AttachInstanceRamRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
