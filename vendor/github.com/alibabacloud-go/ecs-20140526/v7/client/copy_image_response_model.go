// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyImageResponse
	GetStatusCode() *int32
	SetBody(v *CopyImageResponseBody) *CopyImageResponse
	GetBody() *CopyImageResponseBody
}

type CopyImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyImageResponse) GoString() string {
	return s.String()
}

func (s *CopyImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyImageResponse) GetBody() *CopyImageResponseBody {
	return s.Body
}

func (s *CopyImageResponse) SetHeaders(v map[string]*string) *CopyImageResponse {
	s.Headers = v
	return s
}

func (s *CopyImageResponse) SetStatusCode(v int32) *CopyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyImageResponse) SetBody(v *CopyImageResponseBody) *CopyImageResponse {
	s.Body = v
	return s
}

func (s *CopyImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
