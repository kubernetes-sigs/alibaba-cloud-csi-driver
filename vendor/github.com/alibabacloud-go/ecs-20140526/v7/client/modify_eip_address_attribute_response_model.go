// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEipAddressAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEipAddressAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEipAddressAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEipAddressAttributeResponseBody) *ModifyEipAddressAttributeResponse
	GetBody() *ModifyEipAddressAttributeResponseBody
}

type ModifyEipAddressAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEipAddressAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEipAddressAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEipAddressAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyEipAddressAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEipAddressAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEipAddressAttributeResponse) GetBody() *ModifyEipAddressAttributeResponseBody {
	return s.Body
}

func (s *ModifyEipAddressAttributeResponse) SetHeaders(v map[string]*string) *ModifyEipAddressAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyEipAddressAttributeResponse) SetStatusCode(v int32) *ModifyEipAddressAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEipAddressAttributeResponse) SetBody(v *ModifyEipAddressAttributeResponseBody) *ModifyEipAddressAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyEipAddressAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
