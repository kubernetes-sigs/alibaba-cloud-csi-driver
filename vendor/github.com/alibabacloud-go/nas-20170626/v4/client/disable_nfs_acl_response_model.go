// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNfsAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableNfsAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableNfsAclResponse
	GetStatusCode() *int32
	SetBody(v *DisableNfsAclResponseBody) *DisableNfsAclResponse
	GetBody() *DisableNfsAclResponseBody
}

type DisableNfsAclResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableNfsAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableNfsAclResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableNfsAclResponse) GoString() string {
	return s.String()
}

func (s *DisableNfsAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableNfsAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableNfsAclResponse) GetBody() *DisableNfsAclResponseBody {
	return s.Body
}

func (s *DisableNfsAclResponse) SetHeaders(v map[string]*string) *DisableNfsAclResponse {
	s.Headers = v
	return s
}

func (s *DisableNfsAclResponse) SetStatusCode(v int32) *DisableNfsAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableNfsAclResponse) SetBody(v *DisableNfsAclResponseBody) *DisableNfsAclResponse {
	s.Body = v
	return s
}

func (s *DisableNfsAclResponse) Validate() error {
	return dara.Validate(s)
}
