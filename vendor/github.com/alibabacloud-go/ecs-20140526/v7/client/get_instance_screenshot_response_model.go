// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceScreenshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceScreenshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceScreenshotResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceScreenshotResponseBody) *GetInstanceScreenshotResponse
	GetBody() *GetInstanceScreenshotResponseBody
}

type GetInstanceScreenshotResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceScreenshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceScreenshotResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceScreenshotResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceScreenshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceScreenshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceScreenshotResponse) GetBody() *GetInstanceScreenshotResponseBody {
	return s.Body
}

func (s *GetInstanceScreenshotResponse) SetHeaders(v map[string]*string) *GetInstanceScreenshotResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceScreenshotResponse) SetStatusCode(v int32) *GetInstanceScreenshotResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceScreenshotResponse) SetBody(v *GetInstanceScreenshotResponseBody) *GetInstanceScreenshotResponse {
	s.Body = v
	return s
}

func (s *GetInstanceScreenshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
