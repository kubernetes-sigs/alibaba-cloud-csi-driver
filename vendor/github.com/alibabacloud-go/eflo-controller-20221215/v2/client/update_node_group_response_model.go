// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodeGroupResponseBody) *UpdateNodeGroupResponse
	GetBody() *UpdateNodeGroupResponseBody
}

type UpdateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodeGroupResponse) GetBody() *UpdateNodeGroupResponseBody {
	return s.Body
}

func (s *UpdateNodeGroupResponse) SetHeaders(v map[string]*string) *UpdateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeGroupResponse) SetStatusCode(v int32) *UpdateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeGroupResponse) SetBody(v *UpdateNodeGroupResponseBody) *UpdateNodeGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
