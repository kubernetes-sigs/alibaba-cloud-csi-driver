// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReservedInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReservedInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReservedInstanceAttributeResponseBody) *ModifyReservedInstanceAttributeResponse
	GetBody() *ModifyReservedInstanceAttributeResponseBody
}

type ModifyReservedInstanceAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReservedInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReservedInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReservedInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReservedInstanceAttributeResponse) GetBody() *ModifyReservedInstanceAttributeResponseBody {
	return s.Body
}

func (s *ModifyReservedInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyReservedInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyReservedInstanceAttributeResponse) SetStatusCode(v int32) *ModifyReservedInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReservedInstanceAttributeResponse) SetBody(v *ModifyReservedInstanceAttributeResponseBody) *ModifyReservedInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyReservedInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
