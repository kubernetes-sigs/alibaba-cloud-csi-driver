// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBorderRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirtualBorderRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirtualBorderRouterResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirtualBorderRouterResponseBody) *CreateVirtualBorderRouterResponse
	GetBody() *CreateVirtualBorderRouterResponseBody
}

type CreateVirtualBorderRouterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualBorderRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBorderRouterResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualBorderRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirtualBorderRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirtualBorderRouterResponse) GetBody() *CreateVirtualBorderRouterResponseBody {
	return s.Body
}

func (s *CreateVirtualBorderRouterResponse) SetHeaders(v map[string]*string) *CreateVirtualBorderRouterResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualBorderRouterResponse) SetStatusCode(v int32) *CreateVirtualBorderRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualBorderRouterResponse) SetBody(v *CreateVirtualBorderRouterResponseBody) *CreateVirtualBorderRouterResponse {
	s.Body = v
	return s
}

func (s *CreateVirtualBorderRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
