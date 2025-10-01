// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDataFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDataFlowResponse
	GetStatusCode() *int32
	SetBody(v *StartDataFlowResponseBody) *StartDataFlowResponse
	GetBody() *StartDataFlowResponseBody
}

type StartDataFlowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDataFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDataFlowResponse) GoString() string {
	return s.String()
}

func (s *StartDataFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDataFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDataFlowResponse) GetBody() *StartDataFlowResponseBody {
	return s.Body
}

func (s *StartDataFlowResponse) SetHeaders(v map[string]*string) *StartDataFlowResponse {
	s.Headers = v
	return s
}

func (s *StartDataFlowResponse) SetStatusCode(v int32) *StartDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDataFlowResponse) SetBody(v *StartDataFlowResponseBody) *StartDataFlowResponse {
	s.Body = v
	return s
}

func (s *StartDataFlowResponse) Validate() error {
	return dara.Validate(s)
}
