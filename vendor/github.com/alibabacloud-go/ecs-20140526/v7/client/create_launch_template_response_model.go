// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLaunchTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLaunchTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateLaunchTemplateResponseBody) *CreateLaunchTemplateResponse
	GetBody() *CreateLaunchTemplateResponseBody
}

type CreateLaunchTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLaunchTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLaunchTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLaunchTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLaunchTemplateResponse) GetBody() *CreateLaunchTemplateResponseBody {
	return s.Body
}

func (s *CreateLaunchTemplateResponse) SetHeaders(v map[string]*string) *CreateLaunchTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateLaunchTemplateResponse) SetStatusCode(v int32) *CreateLaunchTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLaunchTemplateResponse) SetBody(v *CreateLaunchTemplateResponseBody) *CreateLaunchTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateLaunchTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
