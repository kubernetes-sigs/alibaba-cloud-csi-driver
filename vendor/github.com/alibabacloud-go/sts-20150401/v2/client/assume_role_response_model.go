// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssumeRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssumeRoleResponse
	GetStatusCode() *int32
	SetBody(v *AssumeRoleResponseBody) *AssumeRoleResponse
	GetBody() *AssumeRoleResponseBody
}

type AssumeRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssumeRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssumeRoleResponse) GetBody() *AssumeRoleResponseBody {
	return s.Body
}

func (s *AssumeRoleResponse) SetHeaders(v map[string]*string) *AssumeRoleResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleResponse) SetStatusCode(v int32) *AssumeRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleResponse) SetBody(v *AssumeRoleResponseBody) *AssumeRoleResponse {
	s.Body = v
	return s
}

func (s *AssumeRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
