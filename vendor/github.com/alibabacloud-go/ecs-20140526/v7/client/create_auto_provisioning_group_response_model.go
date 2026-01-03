// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoProvisioningGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoProvisioningGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoProvisioningGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoProvisioningGroupResponseBody) *CreateAutoProvisioningGroupResponse
	GetBody() *CreateAutoProvisioningGroupResponseBody
}

type CreateAutoProvisioningGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoProvisioningGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoProvisioningGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoProvisioningGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoProvisioningGroupResponse) GetBody() *CreateAutoProvisioningGroupResponseBody {
	return s.Body
}

func (s *CreateAutoProvisioningGroupResponse) SetHeaders(v map[string]*string) *CreateAutoProvisioningGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoProvisioningGroupResponse) SetStatusCode(v int32) *CreateAutoProvisioningGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponse) SetBody(v *CreateAutoProvisioningGroupResponseBody) *CreateAutoProvisioningGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAutoProvisioningGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
