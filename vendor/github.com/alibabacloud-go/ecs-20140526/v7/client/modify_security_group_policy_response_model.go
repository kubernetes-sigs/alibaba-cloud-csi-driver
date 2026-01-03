// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityGroupPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityGroupPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityGroupPolicyResponseBody) *ModifySecurityGroupPolicyResponse
	GetBody() *ModifySecurityGroupPolicyResponseBody
}

type ModifySecurityGroupPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityGroupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityGroupPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityGroupPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityGroupPolicyResponse) GetBody() *ModifySecurityGroupPolicyResponseBody {
	return s.Body
}

func (s *ModifySecurityGroupPolicyResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupPolicyResponse) SetStatusCode(v int32) *ModifySecurityGroupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityGroupPolicyResponse) SetBody(v *ModifySecurityGroupPolicyResponseBody) *ModifySecurityGroupPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityGroupPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
