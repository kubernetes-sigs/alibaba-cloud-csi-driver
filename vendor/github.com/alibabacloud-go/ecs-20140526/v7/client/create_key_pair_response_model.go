// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKeyPairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKeyPairResponse
	GetStatusCode() *int32
	SetBody(v *CreateKeyPairResponseBody) *CreateKeyPairResponse
	GetBody() *CreateKeyPairResponseBody
}

type CreateKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyPairResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKeyPairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKeyPairResponse) GetBody() *CreateKeyPairResponseBody {
	return s.Body
}

func (s *CreateKeyPairResponse) SetHeaders(v map[string]*string) *CreateKeyPairResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyPairResponse) SetStatusCode(v int32) *CreateKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyPairResponse) SetBody(v *CreateKeyPairResponseBody) *CreateKeyPairResponse {
	s.Body = v
	return s
}

func (s *CreateKeyPairResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
