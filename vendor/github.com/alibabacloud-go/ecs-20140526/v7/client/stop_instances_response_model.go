// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopInstancesResponse
	GetStatusCode() *int32
	SetBody(v *StopInstancesResponseBody) *StopInstancesResponse
	GetBody() *StopInstancesResponseBody
}

type StopInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesResponse) GoString() string {
	return s.String()
}

func (s *StopInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopInstancesResponse) GetBody() *StopInstancesResponseBody {
	return s.Body
}

func (s *StopInstancesResponse) SetHeaders(v map[string]*string) *StopInstancesResponse {
	s.Headers = v
	return s
}

func (s *StopInstancesResponse) SetStatusCode(v int32) *StopInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstancesResponse) SetBody(v *StopInstancesResponseBody) *StopInstancesResponse {
	s.Body = v
	return s
}

func (s *StopInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
