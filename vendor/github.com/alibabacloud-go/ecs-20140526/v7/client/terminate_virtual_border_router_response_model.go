// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateVirtualBorderRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateVirtualBorderRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateVirtualBorderRouterResponse
	GetStatusCode() *int32
	SetBody(v *TerminateVirtualBorderRouterResponseBody) *TerminateVirtualBorderRouterResponse
	GetBody() *TerminateVirtualBorderRouterResponseBody
}

type TerminateVirtualBorderRouterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateVirtualBorderRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateVirtualBorderRouterResponse) GoString() string {
	return s.String()
}

func (s *TerminateVirtualBorderRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateVirtualBorderRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateVirtualBorderRouterResponse) GetBody() *TerminateVirtualBorderRouterResponseBody {
	return s.Body
}

func (s *TerminateVirtualBorderRouterResponse) SetHeaders(v map[string]*string) *TerminateVirtualBorderRouterResponse {
	s.Headers = v
	return s
}

func (s *TerminateVirtualBorderRouterResponse) SetStatusCode(v int32) *TerminateVirtualBorderRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateVirtualBorderRouterResponse) SetBody(v *TerminateVirtualBorderRouterResponseBody) *TerminateVirtualBorderRouterResponse {
	s.Body = v
	return s
}

func (s *TerminateVirtualBorderRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
