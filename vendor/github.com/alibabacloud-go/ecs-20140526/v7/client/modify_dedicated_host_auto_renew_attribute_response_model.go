// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDedicatedHostAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDedicatedHostAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDedicatedHostAutoRenewAttributeResponseBody) *ModifyDedicatedHostAutoRenewAttributeResponse
	GetBody() *ModifyDedicatedHostAutoRenewAttributeResponseBody
}

type ModifyDedicatedHostAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedHostAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedHostAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponse) GetBody() *ModifyDedicatedHostAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedHostAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponse) SetBody(v *ModifyDedicatedHostAutoRenewAttributeResponseBody) *ModifyDedicatedHostAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
