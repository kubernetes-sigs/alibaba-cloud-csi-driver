// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcAttributeResponseBody) *ModifyVpcAttributeResponse
	GetBody() *ModifyVpcAttributeResponseBody
}

type ModifyVpcAttributeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcAttributeResponse) GetBody() *ModifyVpcAttributeResponseBody {
	return s.Body
}

func (s *ModifyVpcAttributeResponse) SetHeaders(v map[string]*string) *ModifyVpcAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcAttributeResponse) SetStatusCode(v int32) *ModifyVpcAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcAttributeResponse) SetBody(v *ModifyVpcAttributeResponseBody) *ModifyVpcAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
