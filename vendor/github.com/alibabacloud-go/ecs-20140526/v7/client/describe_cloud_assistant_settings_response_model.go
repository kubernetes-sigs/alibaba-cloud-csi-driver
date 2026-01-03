// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudAssistantSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudAssistantSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudAssistantSettingsResponseBody) *DescribeCloudAssistantSettingsResponse
	GetBody() *DescribeCloudAssistantSettingsResponseBody
}

type DescribeCloudAssistantSettingsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudAssistantSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudAssistantSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudAssistantSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudAssistantSettingsResponse) GetBody() *DescribeCloudAssistantSettingsResponseBody {
	return s.Body
}

func (s *DescribeCloudAssistantSettingsResponse) SetHeaders(v map[string]*string) *DescribeCloudAssistantSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponse) SetStatusCode(v int32) *DescribeCloudAssistantSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudAssistantSettingsResponse) SetBody(v *DescribeCloudAssistantSettingsResponseBody) *DescribeCloudAssistantSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudAssistantSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
