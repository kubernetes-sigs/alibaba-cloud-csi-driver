// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAssistantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallCloudAssistantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallCloudAssistantResponse
	GetStatusCode() *int32
	SetBody(v *InstallCloudAssistantResponseBody) *InstallCloudAssistantResponse
	GetBody() *InstallCloudAssistantResponseBody
}

type InstallCloudAssistantResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCloudAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCloudAssistantResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAssistantResponse) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallCloudAssistantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallCloudAssistantResponse) GetBody() *InstallCloudAssistantResponseBody {
	return s.Body
}

func (s *InstallCloudAssistantResponse) SetHeaders(v map[string]*string) *InstallCloudAssistantResponse {
	s.Headers = v
	return s
}

func (s *InstallCloudAssistantResponse) SetStatusCode(v int32) *InstallCloudAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCloudAssistantResponse) SetBody(v *InstallCloudAssistantResponseBody) *InstallCloudAssistantResponse {
	s.Body = v
	return s
}

func (s *InstallCloudAssistantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
