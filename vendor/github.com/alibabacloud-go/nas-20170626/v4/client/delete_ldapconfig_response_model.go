// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLDAPConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLDAPConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLDAPConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLDAPConfigResponseBody) *DeleteLDAPConfigResponse
	GetBody() *DeleteLDAPConfigResponseBody
}

type DeleteLDAPConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLDAPConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLDAPConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLDAPConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLDAPConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLDAPConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLDAPConfigResponse) GetBody() *DeleteLDAPConfigResponseBody {
	return s.Body
}

func (s *DeleteLDAPConfigResponse) SetHeaders(v map[string]*string) *DeleteLDAPConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLDAPConfigResponse) SetStatusCode(v int32) *DeleteLDAPConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLDAPConfigResponse) SetBody(v *DeleteLDAPConfigResponseBody) *DeleteLDAPConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLDAPConfigResponse) Validate() error {
	return dara.Validate(s)
}
