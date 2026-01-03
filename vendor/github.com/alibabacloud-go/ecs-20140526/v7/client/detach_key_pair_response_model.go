// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *DetachKeyPairResponseBody) *DetachKeyPairResponse
	GetBody() *DetachKeyPairResponseBody
}

type DetachKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponse) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachKeyPairResponse) GetBody() *DetachKeyPairResponseBody {
	return s.Body
}

func (s *DetachKeyPairResponse) SetHeaders(v map[string]*string) *DetachKeyPairResponse {
	s.Headers = v
	return s
}

func (s *DetachKeyPairResponse) SetStatusCode(v int32) *DetachKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachKeyPairResponse) SetBody(v *DetachKeyPairResponseBody) *DetachKeyPairResponse {
	s.Body = v
	return s
}

func (s *DetachKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
