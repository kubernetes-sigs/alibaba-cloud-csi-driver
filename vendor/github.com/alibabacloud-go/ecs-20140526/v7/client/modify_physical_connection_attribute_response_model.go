// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhysicalConnectionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPhysicalConnectionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPhysicalConnectionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPhysicalConnectionAttributeResponseBody) *ModifyPhysicalConnectionAttributeResponse
	GetBody() *ModifyPhysicalConnectionAttributeResponseBody
}

type ModifyPhysicalConnectionAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPhysicalConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPhysicalConnectionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhysicalConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyPhysicalConnectionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPhysicalConnectionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPhysicalConnectionAttributeResponse) GetBody() *ModifyPhysicalConnectionAttributeResponseBody {
	return s.Body
}

func (s *ModifyPhysicalConnectionAttributeResponse) SetHeaders(v map[string]*string) *ModifyPhysicalConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyPhysicalConnectionAttributeResponse) SetStatusCode(v int32) *ModifyPhysicalConnectionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeResponse) SetBody(v *ModifyPhysicalConnectionAttributeResponseBody) *ModifyPhysicalConnectionAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyPhysicalConnectionAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
