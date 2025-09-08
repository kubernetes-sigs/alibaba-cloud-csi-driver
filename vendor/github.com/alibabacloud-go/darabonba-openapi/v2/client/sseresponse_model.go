// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

// Description:
//
// This is for OpenApi SDK
type iSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int) *SSEResponse
	GetStatusCode() *int
	SetEvent(v *dara.SSEEvent) *SSEResponse
	GetEvent() *dara.SSEEvent
}

type SSEResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	// HTTP Status Code
	StatusCode *int           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Event      *dara.SSEEvent `json:"event,omitempty" xml:"event,omitempty" require:"true"`
}

func (s SSEResponse) String() string {
	return dara.Prettify(s)
}

func (s SSEResponse) GoString() string {
	return s.String()
}

func (s *SSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SSEResponse) GetStatusCode() *int {
	return s.StatusCode
}

func (s *SSEResponse) GetEvent() *dara.SSEEvent {
	return s.Event
}

func (s *SSEResponse) SetHeaders(v map[string]*string) *SSEResponse {
	s.Headers = v
	return s
}

func (s *SSEResponse) SetStatusCode(v int) *SSEResponse {
	s.StatusCode = &v
	return s
}

func (s *SSEResponse) SetEvent(v *dara.SSEEvent) *SSEResponse {
	s.Event = v
	return s
}

func (s *SSEResponse) Validate() error {
	return dara.Validate(s)
}
