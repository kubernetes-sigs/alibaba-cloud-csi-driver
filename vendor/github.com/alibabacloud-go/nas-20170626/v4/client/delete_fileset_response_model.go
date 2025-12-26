// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFilesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFilesetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFilesetResponseBody) *DeleteFilesetResponse
	GetBody() *DeleteFilesetResponseBody
}

type DeleteFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFilesetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesetResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFilesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFilesetResponse) GetBody() *DeleteFilesetResponseBody {
	return s.Body
}

func (s *DeleteFilesetResponse) SetHeaders(v map[string]*string) *DeleteFilesetResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilesetResponse) SetStatusCode(v int32) *DeleteFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilesetResponse) SetBody(v *DeleteFilesetResponseBody) *DeleteFilesetResponse {
	s.Body = v
	return s
}

func (s *DeleteFilesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
