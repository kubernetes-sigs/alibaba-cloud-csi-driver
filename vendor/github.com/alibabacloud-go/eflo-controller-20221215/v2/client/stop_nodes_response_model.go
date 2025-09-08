// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopNodesResponse
	GetStatusCode() *int32
	SetBody(v *StopNodesResponseBody) *StopNodesResponse
	GetBody() *StopNodesResponseBody
}

type StopNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopNodesResponse) GoString() string {
	return s.String()
}

func (s *StopNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopNodesResponse) GetBody() *StopNodesResponseBody {
	return s.Body
}

func (s *StopNodesResponse) SetHeaders(v map[string]*string) *StopNodesResponse {
	s.Headers = v
	return s
}

func (s *StopNodesResponse) SetStatusCode(v int32) *StopNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopNodesResponse) SetBody(v *StopNodesResponseBody) *StopNodesResponse {
	s.Body = v
	return s
}

func (s *StopNodesResponse) Validate() error {
	return dara.Validate(s)
}
