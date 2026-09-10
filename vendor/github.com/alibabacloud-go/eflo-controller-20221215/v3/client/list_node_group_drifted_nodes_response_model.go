// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupDriftedNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeGroupDriftedNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeGroupDriftedNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeGroupDriftedNodesResponseBody) *ListNodeGroupDriftedNodesResponse
	GetBody() *ListNodeGroupDriftedNodesResponseBody
}

type ListNodeGroupDriftedNodesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeGroupDriftedNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeGroupDriftedNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupDriftedNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupDriftedNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeGroupDriftedNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeGroupDriftedNodesResponse) GetBody() *ListNodeGroupDriftedNodesResponseBody {
	return s.Body
}

func (s *ListNodeGroupDriftedNodesResponse) SetHeaders(v map[string]*string) *ListNodeGroupDriftedNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupDriftedNodesResponse) SetStatusCode(v int32) *ListNodeGroupDriftedNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponse) SetBody(v *ListNodeGroupDriftedNodesResponseBody) *ListNodeGroupDriftedNodesResponse {
	s.Body = v
	return s
}

func (s *ListNodeGroupDriftedNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
