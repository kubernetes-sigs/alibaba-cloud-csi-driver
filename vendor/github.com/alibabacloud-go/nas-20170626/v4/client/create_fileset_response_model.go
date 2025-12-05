// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFilesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFilesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFilesetResponse
	GetStatusCode() *int32
	SetBody(v *CreateFilesetResponseBody) *CreateFilesetResponse
	GetBody() *CreateFilesetResponseBody
}

type CreateFilesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFilesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFilesetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFilesetResponse) GoString() string {
	return s.String()
}

func (s *CreateFilesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFilesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFilesetResponse) GetBody() *CreateFilesetResponseBody {
	return s.Body
}

func (s *CreateFilesetResponse) SetHeaders(v map[string]*string) *CreateFilesetResponse {
	s.Headers = v
	return s
}

func (s *CreateFilesetResponse) SetStatusCode(v int32) *CreateFilesetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFilesetResponse) SetBody(v *CreateFilesetResponseBody) *CreateFilesetResponse {
	s.Body = v
	return s
}

func (s *CreateFilesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
