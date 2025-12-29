// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceVpcAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceVpcAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceVpcAttributeResponseBody) *ModifyInstanceVpcAttributeResponse
	GetBody() *ModifyInstanceVpcAttributeResponseBody
}

type ModifyInstanceVpcAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceVpcAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceVpcAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceVpcAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceVpcAttributeResponse) GetBody() *ModifyInstanceVpcAttributeResponseBody {
	return s.Body
}

func (s *ModifyInstanceVpcAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceVpcAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVpcAttributeResponse) SetStatusCode(v int32) *ModifyInstanceVpcAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceVpcAttributeResponse) SetBody(v *ModifyInstanceVpcAttributeResponseBody) *ModifyInstanceVpcAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceVpcAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
