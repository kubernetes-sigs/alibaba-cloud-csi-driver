// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetTestTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetTestTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetTestTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetTestTaskResponseBody) *CreateNetTestTaskResponse
	GetBody() *CreateNetTestTaskResponseBody
}

type CreateNetTestTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetTestTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetTestTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetTestTaskResponse) GetBody() *CreateNetTestTaskResponseBody {
	return s.Body
}

func (s *CreateNetTestTaskResponse) SetHeaders(v map[string]*string) *CreateNetTestTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateNetTestTaskResponse) SetStatusCode(v int32) *CreateNetTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetTestTaskResponse) SetBody(v *CreateNetTestTaskResponseBody) *CreateNetTestTaskResponse {
	s.Body = v
	return s
}

func (s *CreateNetTestTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
