// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBorderRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirtualBorderRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirtualBorderRouterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirtualBorderRouterResponseBody) *DeleteVirtualBorderRouterResponse
	GetBody() *DeleteVirtualBorderRouterResponseBody
}

type DeleteVirtualBorderRouterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualBorderRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBorderRouterResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBorderRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirtualBorderRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirtualBorderRouterResponse) GetBody() *DeleteVirtualBorderRouterResponseBody {
	return s.Body
}

func (s *DeleteVirtualBorderRouterResponse) SetHeaders(v map[string]*string) *DeleteVirtualBorderRouterResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualBorderRouterResponse) SetStatusCode(v int32) *DeleteVirtualBorderRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualBorderRouterResponse) SetBody(v *DeleteVirtualBorderRouterResponseBody) *DeleteVirtualBorderRouterResponse {
	s.Body = v
	return s
}

func (s *DeleteVirtualBorderRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
