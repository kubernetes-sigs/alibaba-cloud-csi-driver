// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptInquiredSystemEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptInquiredSystemEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptInquiredSystemEventResponse
	GetStatusCode() *int32
	SetBody(v *AcceptInquiredSystemEventResponseBody) *AcceptInquiredSystemEventResponse
	GetBody() *AcceptInquiredSystemEventResponseBody
}

type AcceptInquiredSystemEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptInquiredSystemEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptInquiredSystemEventResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptInquiredSystemEventResponse) GoString() string {
	return s.String()
}

func (s *AcceptInquiredSystemEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptInquiredSystemEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptInquiredSystemEventResponse) GetBody() *AcceptInquiredSystemEventResponseBody {
	return s.Body
}

func (s *AcceptInquiredSystemEventResponse) SetHeaders(v map[string]*string) *AcceptInquiredSystemEventResponse {
	s.Headers = v
	return s
}

func (s *AcceptInquiredSystemEventResponse) SetStatusCode(v int32) *AcceptInquiredSystemEventResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptInquiredSystemEventResponse) SetBody(v *AcceptInquiredSystemEventResponseBody) *AcceptInquiredSystemEventResponse {
	s.Body = v
	return s
}

func (s *AcceptInquiredSystemEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
