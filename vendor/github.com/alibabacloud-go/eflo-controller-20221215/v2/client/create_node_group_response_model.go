// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodeGroupResponseBody) *CreateNodeGroupResponse
	GetBody() *CreateNodeGroupResponseBody
}

type CreateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodeGroupResponse) GetBody() *CreateNodeGroupResponseBody {
	return s.Body
}

func (s *CreateNodeGroupResponse) SetHeaders(v map[string]*string) *CreateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeGroupResponse) SetStatusCode(v int32) *CreateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeGroupResponse) SetBody(v *CreateNodeGroupResponseBody) *CreateNodeGroupResponse {
	s.Body = v
	return s
}

func (s *CreateNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
