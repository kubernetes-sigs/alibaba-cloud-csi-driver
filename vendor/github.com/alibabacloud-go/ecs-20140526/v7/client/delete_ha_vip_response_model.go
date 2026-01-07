// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHaVipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHaVipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHaVipResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHaVipResponseBody) *DeleteHaVipResponse
	GetBody() *DeleteHaVipResponseBody
}

type DeleteHaVipResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHaVipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHaVipResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHaVipResponse) GoString() string {
	return s.String()
}

func (s *DeleteHaVipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHaVipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHaVipResponse) GetBody() *DeleteHaVipResponseBody {
	return s.Body
}

func (s *DeleteHaVipResponse) SetHeaders(v map[string]*string) *DeleteHaVipResponse {
	s.Headers = v
	return s
}

func (s *DeleteHaVipResponse) SetStatusCode(v int32) *DeleteHaVipResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHaVipResponse) SetBody(v *DeleteHaVipResponseBody) *DeleteHaVipResponse {
	s.Body = v
	return s
}

func (s *DeleteHaVipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
