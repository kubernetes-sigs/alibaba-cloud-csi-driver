// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySecurityGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySecurityGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifySecurityGroupAttributeResponseBody) *ModifySecurityGroupAttributeResponse
	GetBody() *ModifySecurityGroupAttributeResponseBody
}

type ModifySecurityGroupAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecurityGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecurityGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySecurityGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySecurityGroupAttributeResponse) GetBody() *ModifySecurityGroupAttributeResponseBody {
	return s.Body
}

func (s *ModifySecurityGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifySecurityGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityGroupAttributeResponse) SetStatusCode(v int32) *ModifySecurityGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityGroupAttributeResponse) SetBody(v *ModifySecurityGroupAttributeResponseBody) *ModifySecurityGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifySecurityGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
