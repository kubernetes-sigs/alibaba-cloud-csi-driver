// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterNodesResponseBody) *ListClusterNodesResponse
	GetBody() *ListClusterNodesResponseBody
}

type ListClusterNodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterNodesResponse) GetBody() *ListClusterNodesResponseBody {
	return s.Body
}

func (s *ListClusterNodesResponse) SetHeaders(v map[string]*string) *ListClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterNodesResponse) SetStatusCode(v int32) *ListClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterNodesResponse) SetBody(v *ListClusterNodesResponseBody) *ListClusterNodesResponse {
	s.Body = v
	return s
}

func (s *ListClusterNodesResponse) Validate() error {
	return dara.Validate(s)
}
