// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiskResponseBody) *CreateDiskResponse
	GetBody() *CreateDiskResponseBody
}

type CreateDiskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiskResponse) GetBody() *CreateDiskResponseBody {
	return s.Body
}

func (s *CreateDiskResponse) SetHeaders(v map[string]*string) *CreateDiskResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskResponse) SetStatusCode(v int32) *CreateDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskResponse) SetBody(v *CreateDiskResponseBody) *CreateDiskResponse {
	s.Body = v
	return s
}

func (s *CreateDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
