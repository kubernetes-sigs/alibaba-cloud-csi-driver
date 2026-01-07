// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *AttachKeyPairResponseBody) *AttachKeyPairResponse
	GetBody() *AttachKeyPairResponseBody
}

type AttachKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponse) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachKeyPairResponse) GetBody() *AttachKeyPairResponseBody {
	return s.Body
}

func (s *AttachKeyPairResponse) SetHeaders(v map[string]*string) *AttachKeyPairResponse {
	s.Headers = v
	return s
}

func (s *AttachKeyPairResponse) SetStatusCode(v int32) *AttachKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachKeyPairResponse) SetBody(v *AttachKeyPairResponseBody) *AttachKeyPairResponse {
	s.Body = v
	return s
}

func (s *AttachKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
