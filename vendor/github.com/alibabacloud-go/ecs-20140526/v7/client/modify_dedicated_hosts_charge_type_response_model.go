// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostsChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDedicatedHostsChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDedicatedHostsChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDedicatedHostsChargeTypeResponseBody) *ModifyDedicatedHostsChargeTypeResponse
	GetBody() *ModifyDedicatedHostsChargeTypeResponseBody
}

type ModifyDedicatedHostsChargeTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedHostsChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedHostsChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostsChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostsChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDedicatedHostsChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDedicatedHostsChargeTypeResponse) GetBody() *ModifyDedicatedHostsChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyDedicatedHostsChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostsChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponse) SetStatusCode(v int32) *ModifyDedicatedHostsChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponse) SetBody(v *ModifyDedicatedHostsChargeTypeResponseBody) *ModifyDedicatedHostsChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
