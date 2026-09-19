// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleWithOIDCResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssumeRoleWithOIDCResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssumeRoleWithOIDCResponse
	GetStatusCode() *int32
	SetBody(v *AssumeRoleWithOIDCResponseBody) *AssumeRoleWithOIDCResponse
	GetBody() *AssumeRoleWithOIDCResponseBody
}

type AssumeRoleWithOIDCResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeRoleWithOIDCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeRoleWithOIDCResponse) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithOIDCResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssumeRoleWithOIDCResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssumeRoleWithOIDCResponse) GetBody() *AssumeRoleWithOIDCResponseBody {
	return s.Body
}

func (s *AssumeRoleWithOIDCResponse) SetHeaders(v map[string]*string) *AssumeRoleWithOIDCResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleWithOIDCResponse) SetStatusCode(v int32) *AssumeRoleWithOIDCResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleWithOIDCResponse) SetBody(v *AssumeRoleWithOIDCResponseBody) *AssumeRoleWithOIDCResponse {
	s.Body = v
	return s
}

func (s *AssumeRoleWithOIDCResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
