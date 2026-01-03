// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *JoinResourceGroupResponseBody) *JoinResourceGroupResponse
	GetBody() *JoinResourceGroupResponseBody
}

type JoinResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinResourceGroupResponse) GetBody() *JoinResourceGroupResponseBody {
	return s.Body
}

func (s *JoinResourceGroupResponse) SetHeaders(v map[string]*string) *JoinResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinResourceGroupResponse) SetStatusCode(v int32) *JoinResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinResourceGroupResponse) SetBody(v *JoinResourceGroupResponseBody) *JoinResourceGroupResponse {
	s.Body = v
	return s
}

func (s *JoinResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
