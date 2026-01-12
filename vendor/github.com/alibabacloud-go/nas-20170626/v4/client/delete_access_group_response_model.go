// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessGroupResponseBody) *DeleteAccessGroupResponse
	GetBody() *DeleteAccessGroupResponseBody
}

type DeleteAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessGroupResponse) GetBody() *DeleteAccessGroupResponseBody {
	return s.Body
}

func (s *DeleteAccessGroupResponse) SetHeaders(v map[string]*string) *DeleteAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessGroupResponse) SetStatusCode(v int32) *DeleteAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessGroupResponse) SetBody(v *DeleteAccessGroupResponseBody) *DeleteAccessGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
