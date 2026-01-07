// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeploymentSetAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeploymentSetAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeploymentSetAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeploymentSetAttributeResponseBody) *ModifyDeploymentSetAttributeResponse
	GetBody() *ModifyDeploymentSetAttributeResponseBody
}

type ModifyDeploymentSetAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeploymentSetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeploymentSetAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeploymentSetAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeploymentSetAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeploymentSetAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeploymentSetAttributeResponse) GetBody() *ModifyDeploymentSetAttributeResponseBody {
	return s.Body
}

func (s *ModifyDeploymentSetAttributeResponse) SetHeaders(v map[string]*string) *ModifyDeploymentSetAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeploymentSetAttributeResponse) SetStatusCode(v int32) *ModifyDeploymentSetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeploymentSetAttributeResponse) SetBody(v *ModifyDeploymentSetAttributeResponseBody) *ModifyDeploymentSetAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyDeploymentSetAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
