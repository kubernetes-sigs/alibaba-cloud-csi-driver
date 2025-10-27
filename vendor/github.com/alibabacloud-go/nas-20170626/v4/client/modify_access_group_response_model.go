// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccessGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccessGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccessGroupResponseBody) *ModifyAccessGroupResponse
	GetBody() *ModifyAccessGroupResponseBody
}

type ModifyAccessGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccessGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccessGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccessGroupResponse) GetBody() *ModifyAccessGroupResponseBody {
	return s.Body
}

func (s *ModifyAccessGroupResponse) SetHeaders(v map[string]*string) *ModifyAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessGroupResponse) SetStatusCode(v int32) *ModifyAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessGroupResponse) SetBody(v *ModifyAccessGroupResponseBody) *ModifyAccessGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyAccessGroupResponse) Validate() error {
	return dara.Validate(s)
}
