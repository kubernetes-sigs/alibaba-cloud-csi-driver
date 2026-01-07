// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostClusterAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDedicatedHostClusterAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDedicatedHostClusterAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDedicatedHostClusterAttributeResponseBody) *ModifyDedicatedHostClusterAttributeResponse
	GetBody() *ModifyDedicatedHostClusterAttributeResponseBody
}

type ModifyDedicatedHostClusterAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedHostClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedHostClusterAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostClusterAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDedicatedHostClusterAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDedicatedHostClusterAttributeResponse) GetBody() *ModifyDedicatedHostClusterAttributeResponseBody {
	return s.Body
}

func (s *ModifyDedicatedHostClusterAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedHostClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeResponse) SetBody(v *ModifyDedicatedHostClusterAttributeResponseBody) *ModifyDedicatedHostClusterAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyDedicatedHostClusterAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
