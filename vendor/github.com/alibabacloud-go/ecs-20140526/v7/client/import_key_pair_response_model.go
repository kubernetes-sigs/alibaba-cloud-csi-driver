// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *ImportKeyPairResponseBody) *ImportKeyPairResponse
	GetBody() *ImportKeyPairResponseBody
}

type ImportKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportKeyPairResponse) GetBody() *ImportKeyPairResponseBody {
	return s.Body
}

func (s *ImportKeyPairResponse) SetHeaders(v map[string]*string) *ImportKeyPairResponse {
	s.Headers = v
	return s
}

func (s *ImportKeyPairResponse) SetStatusCode(v int32) *ImportKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKeyPairResponse) SetBody(v *ImportKeyPairResponseBody) *ImportKeyPairResponse {
	s.Body = v
	return s
}

func (s *ImportKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
