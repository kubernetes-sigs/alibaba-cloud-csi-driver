// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateHaVipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociateHaVipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociateHaVipResponse
	GetStatusCode() *int32
	SetBody(v *UnassociateHaVipResponseBody) *UnassociateHaVipResponse
	GetBody() *UnassociateHaVipResponseBody
}

type UnassociateHaVipResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociateHaVipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociateHaVipResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociateHaVipResponse) GoString() string {
	return s.String()
}

func (s *UnassociateHaVipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociateHaVipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociateHaVipResponse) GetBody() *UnassociateHaVipResponseBody {
	return s.Body
}

func (s *UnassociateHaVipResponse) SetHeaders(v map[string]*string) *UnassociateHaVipResponse {
	s.Headers = v
	return s
}

func (s *UnassociateHaVipResponse) SetStatusCode(v int32) *UnassociateHaVipResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociateHaVipResponse) SetBody(v *UnassociateHaVipResponseBody) *UnassociateHaVipResponse {
	s.Body = v
	return s
}

func (s *UnassociateHaVipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
