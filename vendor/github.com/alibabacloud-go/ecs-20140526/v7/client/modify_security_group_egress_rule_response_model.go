// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupEgressRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityGroupEgressRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityGroupEgressRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityGroupEgressRuleResponseBody) *ModifySecurityGroupEgressRuleResponse
	GetBody() *ModifySecurityGroupEgressRuleResponseBody
}

type ModifySecurityGroupEgressRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityGroupEgressRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityGroupEgressRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupEgressRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupEgressRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityGroupEgressRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityGroupEgressRuleResponse) GetBody() *ModifySecurityGroupEgressRuleResponseBody {
	return s.Body
}

func (s *ModifySecurityGroupEgressRuleResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupEgressRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupEgressRuleResponse) SetStatusCode(v int32) *ModifySecurityGroupEgressRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleResponse) SetBody(v *ModifySecurityGroupEgressRuleResponseBody) *ModifySecurityGroupEgressRuleResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityGroupEgressRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
