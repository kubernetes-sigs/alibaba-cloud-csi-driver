// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataFlowResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataFlowResponseBody) *ModifyDataFlowResponse
	GetBody() *ModifyDataFlowResponseBody
}

type ModifyDataFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataFlowResponse) GetBody() *ModifyDataFlowResponseBody {
	return s.Body
}

func (s *ModifyDataFlowResponse) SetHeaders(v map[string]*string) *ModifyDataFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataFlowResponse) SetStatusCode(v int32) *ModifyDataFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataFlowResponse) SetBody(v *ModifyDataFlowResponseBody) *ModifyDataFlowResponse {
	s.Body = v
	return s
}

func (s *ModifyDataFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
