// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDiskDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDiskDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDiskDeploymentResponseBody) *ModifyDiskDeploymentResponse
	GetBody() *ModifyDiskDeploymentResponseBody
}

type ModifyDiskDeploymentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDiskDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDiskDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskDeploymentResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDiskDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDiskDeploymentResponse) GetBody() *ModifyDiskDeploymentResponseBody {
	return s.Body
}

func (s *ModifyDiskDeploymentResponse) SetHeaders(v map[string]*string) *ModifyDiskDeploymentResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskDeploymentResponse) SetStatusCode(v int32) *ModifyDiskDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskDeploymentResponse) SetBody(v *ModifyDiskDeploymentResponseBody) *ModifyDiskDeploymentResponse {
	s.Body = v
	return s
}

func (s *ModifyDiskDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
