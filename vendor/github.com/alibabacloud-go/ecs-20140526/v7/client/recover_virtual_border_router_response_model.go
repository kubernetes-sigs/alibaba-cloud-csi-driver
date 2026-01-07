// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverVirtualBorderRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverVirtualBorderRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverVirtualBorderRouterResponse
	GetStatusCode() *int32
	SetBody(v *RecoverVirtualBorderRouterResponseBody) *RecoverVirtualBorderRouterResponse
	GetBody() *RecoverVirtualBorderRouterResponseBody
}

type RecoverVirtualBorderRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverVirtualBorderRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverVirtualBorderRouterResponse) GoString() string {
	return s.String()
}

func (s *RecoverVirtualBorderRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverVirtualBorderRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverVirtualBorderRouterResponse) GetBody() *RecoverVirtualBorderRouterResponseBody {
	return s.Body
}

func (s *RecoverVirtualBorderRouterResponse) SetHeaders(v map[string]*string) *RecoverVirtualBorderRouterResponse {
	s.Headers = v
	return s
}

func (s *RecoverVirtualBorderRouterResponse) SetStatusCode(v int32) *RecoverVirtualBorderRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverVirtualBorderRouterResponse) SetBody(v *RecoverVirtualBorderRouterResponseBody) *RecoverVirtualBorderRouterResponse {
	s.Body = v
	return s
}

func (s *RecoverVirtualBorderRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
