// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteActivationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteActivationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteActivationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteActivationResponseBody) *DeleteActivationResponse
	GetBody() *DeleteActivationResponseBody
}

type DeleteActivationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteActivationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteActivationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteActivationResponse) GoString() string {
	return s.String()
}

func (s *DeleteActivationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteActivationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteActivationResponse) GetBody() *DeleteActivationResponseBody {
	return s.Body
}

func (s *DeleteActivationResponse) SetHeaders(v map[string]*string) *DeleteActivationResponse {
	s.Headers = v
	return s
}

func (s *DeleteActivationResponse) SetStatusCode(v int32) *DeleteActivationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteActivationResponse) SetBody(v *DeleteActivationResponseBody) *DeleteActivationResponse {
	s.Body = v
	return s
}

func (s *DeleteActivationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
