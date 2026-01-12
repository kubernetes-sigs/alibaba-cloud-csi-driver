// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataFlowSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataFlowSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataFlowSubTaskResponseBody) *CreateDataFlowSubTaskResponse
	GetBody() *CreateDataFlowSubTaskResponseBody
}

type CreateDataFlowSubTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataFlowSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataFlowSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowSubTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDataFlowSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataFlowSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataFlowSubTaskResponse) GetBody() *CreateDataFlowSubTaskResponseBody {
	return s.Body
}

func (s *CreateDataFlowSubTaskResponse) SetHeaders(v map[string]*string) *CreateDataFlowSubTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDataFlowSubTaskResponse) SetStatusCode(v int32) *CreateDataFlowSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataFlowSubTaskResponse) SetBody(v *CreateDataFlowSubTaskResponseBody) *CreateDataFlowSubTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDataFlowSubTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
