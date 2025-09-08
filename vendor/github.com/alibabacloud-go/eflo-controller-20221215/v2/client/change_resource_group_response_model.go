// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse
	GetBody() *ChangeResourceGroupResponseBody
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeResourceGroupResponse) GetBody() *ChangeResourceGroupResponseBody {
	return s.Body
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ChangeResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
