// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyPairsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKeyPairsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKeyPairsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKeyPairsResponseBody) *DeleteKeyPairsResponse
	GetBody() *DeleteKeyPairsResponseBody
}

type DeleteKeyPairsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeyPairsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKeyPairsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKeyPairsResponse) GetBody() *DeleteKeyPairsResponseBody {
	return s.Body
}

func (s *DeleteKeyPairsResponse) SetHeaders(v map[string]*string) *DeleteKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyPairsResponse) SetStatusCode(v int32) *DeleteKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeyPairsResponse) SetBody(v *DeleteKeyPairsResponseBody) *DeleteKeyPairsResponse {
	s.Body = v
	return s
}

func (s *DeleteKeyPairsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
