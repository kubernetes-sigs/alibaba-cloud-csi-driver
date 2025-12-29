// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopyImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelCopyImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelCopyImageResponse
	GetStatusCode() *int32
	SetBody(v *CancelCopyImageResponseBody) *CancelCopyImageResponse
	GetBody() *CancelCopyImageResponseBody
}

type CancelCopyImageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelCopyImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelCopyImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelCopyImageResponse) GoString() string {
	return s.String()
}

func (s *CancelCopyImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelCopyImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelCopyImageResponse) GetBody() *CancelCopyImageResponseBody {
	return s.Body
}

func (s *CancelCopyImageResponse) SetHeaders(v map[string]*string) *CancelCopyImageResponse {
	s.Headers = v
	return s
}

func (s *CancelCopyImageResponse) SetStatusCode(v int32) *CancelCopyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelCopyImageResponse) SetBody(v *CancelCopyImageResponseBody) *CancelCopyImageResponse {
	s.Body = v
	return s
}

func (s *CancelCopyImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
