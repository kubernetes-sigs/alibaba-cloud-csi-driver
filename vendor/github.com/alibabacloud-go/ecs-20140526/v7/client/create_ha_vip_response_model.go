// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHaVipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHaVipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHaVipResponse
	GetStatusCode() *int32
	SetBody(v *CreateHaVipResponseBody) *CreateHaVipResponse
	GetBody() *CreateHaVipResponseBody
}

type CreateHaVipResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHaVipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHaVipResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipResponse) GoString() string {
	return s.String()
}

func (s *CreateHaVipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHaVipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHaVipResponse) GetBody() *CreateHaVipResponseBody {
	return s.Body
}

func (s *CreateHaVipResponse) SetHeaders(v map[string]*string) *CreateHaVipResponse {
	s.Headers = v
	return s
}

func (s *CreateHaVipResponse) SetStatusCode(v int32) *CreateHaVipResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHaVipResponse) SetBody(v *CreateHaVipResponseBody) *CreateHaVipResponse {
	s.Body = v
	return s
}

func (s *CreateHaVipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
