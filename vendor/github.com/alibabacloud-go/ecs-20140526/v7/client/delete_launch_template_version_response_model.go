// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaunchTemplateVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLaunchTemplateVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLaunchTemplateVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLaunchTemplateVersionResponseBody) *DeleteLaunchTemplateVersionResponse
	GetBody() *DeleteLaunchTemplateVersionResponseBody
}

type DeleteLaunchTemplateVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLaunchTemplateVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLaunchTemplateVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLaunchTemplateVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLaunchTemplateVersionResponse) GetBody() *DeleteLaunchTemplateVersionResponseBody {
	return s.Body
}

func (s *DeleteLaunchTemplateVersionResponse) SetHeaders(v map[string]*string) *DeleteLaunchTemplateVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteLaunchTemplateVersionResponse) SetStatusCode(v int32) *DeleteLaunchTemplateVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLaunchTemplateVersionResponse) SetBody(v *DeleteLaunchTemplateVersionResponseBody) *DeleteLaunchTemplateVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteLaunchTemplateVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
