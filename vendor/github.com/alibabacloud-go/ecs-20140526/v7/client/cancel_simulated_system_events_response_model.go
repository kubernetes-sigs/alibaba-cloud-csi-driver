// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSimulatedSystemEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSimulatedSystemEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSimulatedSystemEventsResponse
	GetStatusCode() *int32
	SetBody(v *CancelSimulatedSystemEventsResponseBody) *CancelSimulatedSystemEventsResponse
	GetBody() *CancelSimulatedSystemEventsResponseBody
}

type CancelSimulatedSystemEventsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSimulatedSystemEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSimulatedSystemEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSimulatedSystemEventsResponse) GoString() string {
	return s.String()
}

func (s *CancelSimulatedSystemEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSimulatedSystemEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSimulatedSystemEventsResponse) GetBody() *CancelSimulatedSystemEventsResponseBody {
	return s.Body
}

func (s *CancelSimulatedSystemEventsResponse) SetHeaders(v map[string]*string) *CancelSimulatedSystemEventsResponse {
	s.Headers = v
	return s
}

func (s *CancelSimulatedSystemEventsResponse) SetStatusCode(v int32) *CancelSimulatedSystemEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSimulatedSystemEventsResponse) SetBody(v *CancelSimulatedSystemEventsResponseBody) *CancelSimulatedSystemEventsResponse {
	s.Body = v
	return s
}

func (s *CancelSimulatedSystemEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
