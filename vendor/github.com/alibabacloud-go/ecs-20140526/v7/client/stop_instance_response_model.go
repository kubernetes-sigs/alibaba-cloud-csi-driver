// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopInstanceResponseBody) *StopInstanceResponse
	GetBody() *StopInstanceResponseBody
}

type StopInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopInstanceResponse) GetBody() *StopInstanceResponseBody {
	return s.Body
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

func (s *StopInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
