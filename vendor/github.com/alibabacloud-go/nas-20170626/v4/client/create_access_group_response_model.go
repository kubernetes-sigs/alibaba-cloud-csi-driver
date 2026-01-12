// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessGroupResponseBody) *CreateAccessGroupResponse
	GetBody() *CreateAccessGroupResponseBody
}

type CreateAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessGroupResponse) GetBody() *CreateAccessGroupResponseBody {
	return s.Body
}

func (s *CreateAccessGroupResponse) SetHeaders(v map[string]*string) *CreateAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessGroupResponse) SetStatusCode(v int32) *CreateAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessGroupResponse) SetBody(v *CreateAccessGroupResponseBody) *CreateAccessGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAccessGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
