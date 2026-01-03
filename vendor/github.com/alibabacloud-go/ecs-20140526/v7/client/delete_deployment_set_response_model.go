// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeploymentSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeploymentSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeploymentSetResponseBody) *DeleteDeploymentSetResponse
	GetBody() *DeleteDeploymentSetResponseBody
}

type DeleteDeploymentSetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeploymentSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeploymentSetResponse) GetBody() *DeleteDeploymentSetResponseBody {
	return s.Body
}

func (s *DeleteDeploymentSetResponse) SetHeaders(v map[string]*string) *DeleteDeploymentSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentSetResponse) SetStatusCode(v int32) *DeleteDeploymentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentSetResponse) SetBody(v *DeleteDeploymentSetResponseBody) *DeleteDeploymentSetResponse {
	s.Body = v
	return s
}

func (s *DeleteDeploymentSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
