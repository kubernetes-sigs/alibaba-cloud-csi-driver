// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeDiskResponse
	GetStatusCode() *int32
	SetBody(v *ResizeDiskResponseBody) *ResizeDiskResponse
	GetBody() *ResizeDiskResponseBody
}

type ResizeDiskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeDiskResponse) GoString() string {
	return s.String()
}

func (s *ResizeDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeDiskResponse) GetBody() *ResizeDiskResponseBody {
	return s.Body
}

func (s *ResizeDiskResponse) SetHeaders(v map[string]*string) *ResizeDiskResponse {
	s.Headers = v
	return s
}

func (s *ResizeDiskResponse) SetStatusCode(v int32) *ResizeDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeDiskResponse) SetBody(v *ResizeDiskResponseBody) *ResizeDiskResponse {
	s.Body = v
	return s
}

func (s *ResizeDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
