// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudAssistantSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudAssistantSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudAssistantSettingsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudAssistantSettingsResponseBody) *ModifyCloudAssistantSettingsResponse
	GetBody() *ModifyCloudAssistantSettingsResponseBody
}

type ModifyCloudAssistantSettingsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudAssistantSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudAssistantSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudAssistantSettingsResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudAssistantSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudAssistantSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudAssistantSettingsResponse) GetBody() *ModifyCloudAssistantSettingsResponseBody {
	return s.Body
}

func (s *ModifyCloudAssistantSettingsResponse) SetHeaders(v map[string]*string) *ModifyCloudAssistantSettingsResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudAssistantSettingsResponse) SetStatusCode(v int32) *ModifyCloudAssistantSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudAssistantSettingsResponse) SetBody(v *ModifyCloudAssistantSettingsResponseBody) *ModifyCloudAssistantSettingsResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudAssistantSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
