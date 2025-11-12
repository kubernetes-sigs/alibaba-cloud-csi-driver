// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHyperNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHyperNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHyperNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListHyperNodesResponseBody) *ListHyperNodesResponse
	GetBody() *ListHyperNodesResponseBody
}

type ListHyperNodesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHyperNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHyperNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHyperNodesResponse) GoString() string {
	return s.String()
}

func (s *ListHyperNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHyperNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHyperNodesResponse) GetBody() *ListHyperNodesResponseBody {
	return s.Body
}

func (s *ListHyperNodesResponse) SetHeaders(v map[string]*string) *ListHyperNodesResponse {
	s.Headers = v
	return s
}

func (s *ListHyperNodesResponse) SetStatusCode(v int32) *ListHyperNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHyperNodesResponse) SetBody(v *ListHyperNodesResponseBody) *ListHyperNodesResponse {
	s.Body = v
	return s
}

func (s *ListHyperNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
