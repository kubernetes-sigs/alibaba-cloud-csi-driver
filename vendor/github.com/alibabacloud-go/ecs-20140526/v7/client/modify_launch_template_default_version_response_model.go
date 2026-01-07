// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLaunchTemplateDefaultVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLaunchTemplateDefaultVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLaunchTemplateDefaultVersionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLaunchTemplateDefaultVersionResponseBody) *ModifyLaunchTemplateDefaultVersionResponse
	GetBody() *ModifyLaunchTemplateDefaultVersionResponseBody
}

type ModifyLaunchTemplateDefaultVersionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLaunchTemplateDefaultVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLaunchTemplateDefaultVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLaunchTemplateDefaultVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyLaunchTemplateDefaultVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLaunchTemplateDefaultVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLaunchTemplateDefaultVersionResponse) GetBody() *ModifyLaunchTemplateDefaultVersionResponseBody {
	return s.Body
}

func (s *ModifyLaunchTemplateDefaultVersionResponse) SetHeaders(v map[string]*string) *ModifyLaunchTemplateDefaultVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionResponse) SetStatusCode(v int32) *ModifyLaunchTemplateDefaultVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionResponse) SetBody(v *ModifyLaunchTemplateDefaultVersionResponseBody) *ModifyLaunchTemplateDefaultVersionResponse {
	s.Body = v
	return s
}

func (s *ModifyLaunchTemplateDefaultVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
