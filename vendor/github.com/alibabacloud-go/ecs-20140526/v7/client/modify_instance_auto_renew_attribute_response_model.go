// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceAutoRenewAttributeResponseBody) *ModifyInstanceAutoRenewAttributeResponse
	GetBody() *ModifyInstanceAutoRenewAttributeResponseBody
}

type ModifyInstanceAutoRenewAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceAutoRenewAttributeResponse) GetBody() *ModifyInstanceAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *ModifyInstanceAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeResponse) SetStatusCode(v int32) *ModifyInstanceAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeResponse) SetBody(v *ModifyInstanceAutoRenewAttributeResponseBody) *ModifyInstanceAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
