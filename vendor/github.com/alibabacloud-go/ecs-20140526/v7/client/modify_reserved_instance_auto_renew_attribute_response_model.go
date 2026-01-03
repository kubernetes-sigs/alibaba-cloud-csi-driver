// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstanceAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReservedInstanceAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReservedInstanceAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReservedInstanceAutoRenewAttributeResponseBody) *ModifyReservedInstanceAutoRenewAttributeResponse
	GetBody() *ModifyReservedInstanceAutoRenewAttributeResponseBody
}

type ModifyReservedInstanceAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReservedInstanceAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReservedInstanceAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstanceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponse) GetBody() *ModifyReservedInstanceAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyReservedInstanceAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponse) SetStatusCode(v int32) *ModifyReservedInstanceAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponse) SetBody(v *ModifyReservedInstanceAutoRenewAttributeResponseBody) *ModifyReservedInstanceAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
