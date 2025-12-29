// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImageResponseBody) *DeleteImageResponse
	GetBody() *DeleteImageResponseBody
}

type DeleteImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImageResponse) GetBody() *DeleteImageResponseBody {
	return s.Body
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetStatusCode(v int32) *DeleteImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

func (s *DeleteImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
