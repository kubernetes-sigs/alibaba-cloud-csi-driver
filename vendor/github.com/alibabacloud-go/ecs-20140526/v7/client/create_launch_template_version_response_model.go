// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLaunchTemplateVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLaunchTemplateVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateLaunchTemplateVersionResponseBody) *CreateLaunchTemplateVersionResponse
	GetBody() *CreateLaunchTemplateVersionResponseBody
}

type CreateLaunchTemplateVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLaunchTemplateVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLaunchTemplateVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLaunchTemplateVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLaunchTemplateVersionResponse) GetBody() *CreateLaunchTemplateVersionResponseBody {
	return s.Body
}

func (s *CreateLaunchTemplateVersionResponse) SetHeaders(v map[string]*string) *CreateLaunchTemplateVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateLaunchTemplateVersionResponse) SetStatusCode(v int32) *CreateLaunchTemplateVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLaunchTemplateVersionResponse) SetBody(v *CreateLaunchTemplateVersionResponseBody) *CreateLaunchTemplateVersionResponse {
	s.Body = v
	return s
}

func (s *CreateLaunchTemplateVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
