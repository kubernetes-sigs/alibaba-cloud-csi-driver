// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLDAPConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLDAPConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLDAPConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateLDAPConfigResponseBody) *CreateLDAPConfigResponse
	GetBody() *CreateLDAPConfigResponseBody
}

type CreateLDAPConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLDAPConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLDAPConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLDAPConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLDAPConfigResponse) GetBody() *CreateLDAPConfigResponseBody {
	return s.Body
}

func (s *CreateLDAPConfigResponse) SetHeaders(v map[string]*string) *CreateLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLDAPConfigResponse) SetStatusCode(v int32) *CreateLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLDAPConfigResponse) SetBody(v *CreateLDAPConfigResponseBody) *CreateLDAPConfigResponse {
	s.Body = v
	return s
}

func (s *CreateLDAPConfigResponse) Validate() error {
	return dara.Validate(s)
}
