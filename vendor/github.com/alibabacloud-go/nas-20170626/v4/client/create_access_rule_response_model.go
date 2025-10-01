// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessRuleResponseBody) *CreateAccessRuleResponse
	GetBody() *CreateAccessRuleResponseBody
}

type CreateAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessRuleResponse) GetBody() *CreateAccessRuleResponseBody {
	return s.Body
}

func (s *CreateAccessRuleResponse) SetHeaders(v map[string]*string) *CreateAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessRuleResponse) SetStatusCode(v int32) *CreateAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessRuleResponse) SetBody(v *CreateAccessRuleResponseBody) *CreateAccessRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAccessRuleResponse) Validate() error {
	return dara.Validate(s)
}
