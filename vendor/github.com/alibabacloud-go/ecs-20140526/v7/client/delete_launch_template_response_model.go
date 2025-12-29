// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaunchTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLaunchTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLaunchTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLaunchTemplateResponseBody) *DeleteLaunchTemplateResponse
	GetBody() *DeleteLaunchTemplateResponseBody
}

type DeleteLaunchTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLaunchTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLaunchTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLaunchTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLaunchTemplateResponse) GetBody() *DeleteLaunchTemplateResponseBody {
	return s.Body
}

func (s *DeleteLaunchTemplateResponse) SetHeaders(v map[string]*string) *DeleteLaunchTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteLaunchTemplateResponse) SetStatusCode(v int32) *DeleteLaunchTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLaunchTemplateResponse) SetBody(v *DeleteLaunchTemplateResponseBody) *DeleteLaunchTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteLaunchTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
