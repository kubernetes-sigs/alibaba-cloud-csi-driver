// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPluginStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPluginStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListPluginStatusResponseBody) *ListPluginStatusResponse
	GetBody() *ListPluginStatusResponseBody
}

type ListPluginStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPluginStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPluginStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPluginStatusResponse) GoString() string {
	return s.String()
}

func (s *ListPluginStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPluginStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPluginStatusResponse) GetBody() *ListPluginStatusResponseBody {
	return s.Body
}

func (s *ListPluginStatusResponse) SetHeaders(v map[string]*string) *ListPluginStatusResponse {
	s.Headers = v
	return s
}

func (s *ListPluginStatusResponse) SetStatusCode(v int32) *ListPluginStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPluginStatusResponse) SetBody(v *ListPluginStatusResponseBody) *ListPluginStatusResponse {
	s.Body = v
	return s
}

func (s *ListPluginStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
