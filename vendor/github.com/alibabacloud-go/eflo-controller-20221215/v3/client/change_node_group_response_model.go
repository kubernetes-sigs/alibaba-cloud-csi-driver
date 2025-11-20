// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeNodeGroupResponseBody) *ChangeNodeGroupResponse
	GetBody() *ChangeNodeGroupResponseBody
}

type ChangeNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeNodeGroupResponse) GetBody() *ChangeNodeGroupResponseBody {
	return s.Body
}

func (s *ChangeNodeGroupResponse) SetHeaders(v map[string]*string) *ChangeNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeNodeGroupResponse) SetStatusCode(v int32) *ChangeNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeNodeGroupResponse) SetBody(v *ChangeNodeGroupResponseBody) *ChangeNodeGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
