// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNodeGroupResponseBody) *DeleteNodeGroupResponse
	GetBody() *DeleteNodeGroupResponseBody
}

type DeleteNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNodeGroupResponse) GetBody() *DeleteNodeGroupResponseBody {
	return s.Body
}

func (s *DeleteNodeGroupResponse) SetHeaders(v map[string]*string) *DeleteNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeGroupResponse) SetStatusCode(v int32) *DeleteNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeGroupResponse) SetBody(v *DeleteNodeGroupResponseBody) *DeleteNodeGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
