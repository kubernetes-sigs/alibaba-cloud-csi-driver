// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccessRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccessRuleResponseBody) *ModifyAccessRuleResponse
	GetBody() *ModifyAccessRuleResponseBody
}

type ModifyAccessRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccessRuleResponse) GetBody() *ModifyAccessRuleResponseBody {
	return s.Body
}

func (s *ModifyAccessRuleResponse) SetHeaders(v map[string]*string) *ModifyAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessRuleResponse) SetStatusCode(v int32) *ModifyAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessRuleResponse) SetBody(v *ModifyAccessRuleResponseBody) *ModifyAccessRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyAccessRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
