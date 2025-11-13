// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterHyperNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterHyperNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterHyperNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterHyperNodesResponseBody) *ListClusterHyperNodesResponse
	GetBody() *ListClusterHyperNodesResponseBody
}

type ListClusterHyperNodesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterHyperNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterHyperNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHyperNodesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterHyperNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterHyperNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterHyperNodesResponse) GetBody() *ListClusterHyperNodesResponseBody {
	return s.Body
}

func (s *ListClusterHyperNodesResponse) SetHeaders(v map[string]*string) *ListClusterHyperNodesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterHyperNodesResponse) SetStatusCode(v int32) *ListClusterHyperNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterHyperNodesResponse) SetBody(v *ListClusterHyperNodesResponseBody) *ListClusterHyperNodesResponse {
	s.Body = v
	return s
}

func (s *ListClusterHyperNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
