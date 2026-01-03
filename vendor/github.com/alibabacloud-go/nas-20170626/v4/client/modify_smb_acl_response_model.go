// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmbAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySmbAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySmbAclResponse
	GetStatusCode() *int32
	SetBody(v *ModifySmbAclResponseBody) *ModifySmbAclResponse
	GetBody() *ModifySmbAclResponseBody
}

type ModifySmbAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySmbAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySmbAclResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySmbAclResponse) GoString() string {
	return s.String()
}

func (s *ModifySmbAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySmbAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySmbAclResponse) GetBody() *ModifySmbAclResponseBody {
	return s.Body
}

func (s *ModifySmbAclResponse) SetHeaders(v map[string]*string) *ModifySmbAclResponse {
	s.Headers = v
	return s
}

func (s *ModifySmbAclResponse) SetStatusCode(v int32) *ModifySmbAclResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySmbAclResponse) SetBody(v *ModifySmbAclResponseBody) *ModifySmbAclResponse {
	s.Body = v
	return s
}

func (s *ModifySmbAclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
