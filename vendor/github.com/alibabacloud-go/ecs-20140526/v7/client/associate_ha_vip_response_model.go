// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateHaVipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateHaVipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateHaVipResponse
	GetStatusCode() *int32
	SetBody(v *AssociateHaVipResponseBody) *AssociateHaVipResponse
	GetBody() *AssociateHaVipResponseBody
}

type AssociateHaVipResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateHaVipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateHaVipResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateHaVipResponse) GoString() string {
	return s.String()
}

func (s *AssociateHaVipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateHaVipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateHaVipResponse) GetBody() *AssociateHaVipResponseBody {
	return s.Body
}

func (s *AssociateHaVipResponse) SetHeaders(v map[string]*string) *AssociateHaVipResponse {
	s.Headers = v
	return s
}

func (s *AssociateHaVipResponse) SetStatusCode(v int32) *AssociateHaVipResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateHaVipResponse) SetBody(v *AssociateHaVipResponseBody) *AssociateHaVipResponse {
	s.Body = v
	return s
}

func (s *AssociateHaVipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
