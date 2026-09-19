// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshNodeGroupNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshNodeGroupNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshNodeGroupNodesResponse
	GetStatusCode() *int32
	SetBody(v *RefreshNodeGroupNodesResponseBody) *RefreshNodeGroupNodesResponse
	GetBody() *RefreshNodeGroupNodesResponseBody
}

type RefreshNodeGroupNodesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshNodeGroupNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshNodeGroupNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshNodeGroupNodesResponse) GoString() string {
	return s.String()
}

func (s *RefreshNodeGroupNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshNodeGroupNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshNodeGroupNodesResponse) GetBody() *RefreshNodeGroupNodesResponseBody {
	return s.Body
}

func (s *RefreshNodeGroupNodesResponse) SetHeaders(v map[string]*string) *RefreshNodeGroupNodesResponse {
	s.Headers = v
	return s
}

func (s *RefreshNodeGroupNodesResponse) SetStatusCode(v int32) *RefreshNodeGroupNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshNodeGroupNodesResponse) SetBody(v *RefreshNodeGroupNodesResponseBody) *RefreshNodeGroupNodesResponse {
	s.Body = v
	return s
}

func (s *RefreshNodeGroupNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
