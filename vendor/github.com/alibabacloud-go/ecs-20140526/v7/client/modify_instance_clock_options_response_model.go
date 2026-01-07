// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceClockOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceClockOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceClockOptionsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceClockOptionsResponseBody) *ModifyInstanceClockOptionsResponse
	GetBody() *ModifyInstanceClockOptionsResponseBody
}

type ModifyInstanceClockOptionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceClockOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceClockOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceClockOptionsResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceClockOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceClockOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceClockOptionsResponse) GetBody() *ModifyInstanceClockOptionsResponseBody {
	return s.Body
}

func (s *ModifyInstanceClockOptionsResponse) SetHeaders(v map[string]*string) *ModifyInstanceClockOptionsResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceClockOptionsResponse) SetStatusCode(v int32) *ModifyInstanceClockOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceClockOptionsResponse) SetBody(v *ModifyInstanceClockOptionsResponseBody) *ModifyInstanceClockOptionsResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceClockOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
