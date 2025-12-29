// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoProvisioningGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoProvisioningGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoProvisioningGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoProvisioningGroupResponseBody) *DeleteAutoProvisioningGroupResponse
	GetBody() *DeleteAutoProvisioningGroupResponseBody
}

type DeleteAutoProvisioningGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoProvisioningGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoProvisioningGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoProvisioningGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoProvisioningGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoProvisioningGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoProvisioningGroupResponse) GetBody() *DeleteAutoProvisioningGroupResponseBody {
	return s.Body
}

func (s *DeleteAutoProvisioningGroupResponse) SetHeaders(v map[string]*string) *DeleteAutoProvisioningGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoProvisioningGroupResponse) SetStatusCode(v int32) *DeleteAutoProvisioningGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoProvisioningGroupResponse) SetBody(v *DeleteAutoProvisioningGroupResponseBody) *DeleteAutoProvisioningGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoProvisioningGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
