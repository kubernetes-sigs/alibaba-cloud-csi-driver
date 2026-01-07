// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoProvisioningGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoProvisioningGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoProvisioningGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoProvisioningGroupResponseBody) *ModifyAutoProvisioningGroupResponse
	GetBody() *ModifyAutoProvisioningGroupResponseBody
}

type ModifyAutoProvisioningGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoProvisioningGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoProvisioningGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoProvisioningGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoProvisioningGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoProvisioningGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoProvisioningGroupResponse) GetBody() *ModifyAutoProvisioningGroupResponseBody {
	return s.Body
}

func (s *ModifyAutoProvisioningGroupResponse) SetHeaders(v map[string]*string) *ModifyAutoProvisioningGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoProvisioningGroupResponse) SetStatusCode(v int32) *ModifyAutoProvisioningGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoProvisioningGroupResponse) SetBody(v *ModifyAutoProvisioningGroupResponseBody) *ModifyAutoProvisioningGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoProvisioningGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
