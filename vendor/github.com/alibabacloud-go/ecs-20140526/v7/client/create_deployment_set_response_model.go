// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeploymentSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeploymentSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeploymentSetResponseBody) *CreateDeploymentSetResponse
	GetBody() *CreateDeploymentSetResponseBody
}

type CreateDeploymentSetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentSetResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeploymentSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeploymentSetResponse) GetBody() *CreateDeploymentSetResponseBody {
	return s.Body
}

func (s *CreateDeploymentSetResponse) SetHeaders(v map[string]*string) *CreateDeploymentSetResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentSetResponse) SetStatusCode(v int32) *CreateDeploymentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentSetResponse) SetBody(v *CreateDeploymentSetResponseBody) *CreateDeploymentSetResponse {
	s.Body = v
	return s
}

func (s *CreateDeploymentSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
