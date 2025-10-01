// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessRuleResponseBody) *DeleteAccessRuleResponse
	GetBody() *DeleteAccessRuleResponseBody
}

type DeleteAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessRuleResponse) GetBody() *DeleteAccessRuleResponseBody {
	return s.Body
}

func (s *DeleteAccessRuleResponse) SetHeaders(v map[string]*string) *DeleteAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessRuleResponse) SetStatusCode(v int32) *DeleteAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessRuleResponse) SetBody(v *DeleteAccessRuleResponseBody) *DeleteAccessRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessRuleResponse) Validate() error {
	return dara.Validate(s)
}
