// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootNodesResponse
	GetStatusCode() *int32
	SetBody(v *RebootNodesResponseBody) *RebootNodesResponse
	GetBody() *RebootNodesResponseBody
}

type RebootNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootNodesResponse) GoString() string {
	return s.String()
}

func (s *RebootNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootNodesResponse) GetBody() *RebootNodesResponseBody {
	return s.Body
}

func (s *RebootNodesResponse) SetHeaders(v map[string]*string) *RebootNodesResponse {
	s.Headers = v
	return s
}

func (s *RebootNodesResponse) SetStatusCode(v int32) *RebootNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootNodesResponse) SetBody(v *RebootNodesResponseBody) *RebootNodesResponse {
	s.Body = v
	return s
}

func (s *RebootNodesResponse) Validate() error {
	return dara.Validate(s)
}
