// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminatePhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminatePhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminatePhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *TerminatePhysicalConnectionResponseBody) *TerminatePhysicalConnectionResponse
	GetBody() *TerminatePhysicalConnectionResponseBody
}

type TerminatePhysicalConnectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminatePhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminatePhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminatePhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *TerminatePhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminatePhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminatePhysicalConnectionResponse) GetBody() *TerminatePhysicalConnectionResponseBody {
	return s.Body
}

func (s *TerminatePhysicalConnectionResponse) SetHeaders(v map[string]*string) *TerminatePhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *TerminatePhysicalConnectionResponse) SetStatusCode(v int32) *TerminatePhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminatePhysicalConnectionResponse) SetBody(v *TerminatePhysicalConnectionResponseBody) *TerminatePhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *TerminatePhysicalConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
