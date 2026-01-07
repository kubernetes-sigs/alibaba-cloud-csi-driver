// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDedicatedHostAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDedicatedHostAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDedicatedHostAttributeResponseBody) *ModifyDedicatedHostAttributeResponse
	GetBody() *ModifyDedicatedHostAttributeResponseBody
}

type ModifyDedicatedHostAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedHostAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedHostAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDedicatedHostAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDedicatedHostAttributeResponse) GetBody() *ModifyDedicatedHostAttributeResponseBody {
	return s.Body
}

func (s *ModifyDedicatedHostAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedHostAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostAttributeResponse) SetBody(v *ModifyDedicatedHostAttributeResponseBody) *ModifyDedicatedHostAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyDedicatedHostAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
