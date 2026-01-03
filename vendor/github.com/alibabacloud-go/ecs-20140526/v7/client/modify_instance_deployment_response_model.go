// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceDeploymentResponseBody) *ModifyInstanceDeploymentResponse
	GetBody() *ModifyInstanceDeploymentResponseBody
}

type ModifyInstanceDeploymentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceDeploymentResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceDeploymentResponse) GetBody() *ModifyInstanceDeploymentResponseBody {
	return s.Body
}

func (s *ModifyInstanceDeploymentResponse) SetHeaders(v map[string]*string) *ModifyInstanceDeploymentResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceDeploymentResponse) SetStatusCode(v int32) *ModifyInstanceDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceDeploymentResponse) SetBody(v *ModifyInstanceDeploymentResponseBody) *ModifyInstanceDeploymentResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
