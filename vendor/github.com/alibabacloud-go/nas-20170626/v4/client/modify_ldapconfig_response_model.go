// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLDAPConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLDAPConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLDAPConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLDAPConfigResponseBody) *ModifyLDAPConfigResponse
	GetBody() *ModifyLDAPConfigResponseBody
}

type ModifyLDAPConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLDAPConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLDAPConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLDAPConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLDAPConfigResponse) GetBody() *ModifyLDAPConfigResponseBody {
	return s.Body
}

func (s *ModifyLDAPConfigResponse) SetHeaders(v map[string]*string) *ModifyLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLDAPConfigResponse) SetStatusCode(v int32) *ModifyLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLDAPConfigResponse) SetBody(v *ModifyLDAPConfigResponseBody) *ModifyLDAPConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyLDAPConfigResponse) Validate() error {
	return dara.Validate(s)
}
