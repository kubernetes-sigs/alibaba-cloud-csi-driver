// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageResponseBody) *CreateImageResponse
	GetBody() *CreateImageResponseBody
}

type CreateImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageResponse) GetBody() *CreateImageResponseBody {
	return s.Body
}

func (s *CreateImageResponse) SetHeaders(v map[string]*string) *CreateImageResponse {
	s.Headers = v
	return s
}

func (s *CreateImageResponse) SetStatusCode(v int32) *CreateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
	return s
}

func (s *CreateImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
