// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimulatedSystemEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSimulatedSystemEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSimulatedSystemEventsResponse
	GetStatusCode() *int32
	SetBody(v *CreateSimulatedSystemEventsResponseBody) *CreateSimulatedSystemEventsResponse
	GetBody() *CreateSimulatedSystemEventsResponseBody
}

type CreateSimulatedSystemEventsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimulatedSystemEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimulatedSystemEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSimulatedSystemEventsResponse) GoString() string {
	return s.String()
}

func (s *CreateSimulatedSystemEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSimulatedSystemEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSimulatedSystemEventsResponse) GetBody() *CreateSimulatedSystemEventsResponseBody {
	return s.Body
}

func (s *CreateSimulatedSystemEventsResponse) SetHeaders(v map[string]*string) *CreateSimulatedSystemEventsResponse {
	s.Headers = v
	return s
}

func (s *CreateSimulatedSystemEventsResponse) SetStatusCode(v int32) *CreateSimulatedSystemEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimulatedSystemEventsResponse) SetBody(v *CreateSimulatedSystemEventsResponseBody) *CreateSimulatedSystemEventsResponse {
	s.Body = v
	return s
}

func (s *CreateSimulatedSystemEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
