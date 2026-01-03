// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityGroupRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityGroupRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityGroupRuleResponseBody) *ModifySecurityGroupRuleResponse
	GetBody() *ModifySecurityGroupRuleResponseBody
}

type ModifySecurityGroupRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityGroupRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityGroupRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityGroupRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityGroupRuleResponse) GetBody() *ModifySecurityGroupRuleResponseBody {
	return s.Body
}

func (s *ModifySecurityGroupRuleResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupRuleResponse) SetStatusCode(v int32) *ModifySecurityGroupRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityGroupRuleResponse) SetBody(v *ModifySecurityGroupRuleResponseBody) *ModifySecurityGroupRuleResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityGroupRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
