// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreeHyperNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFreeHyperNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFreeHyperNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListFreeHyperNodesResponseBody) *ListFreeHyperNodesResponse
	GetBody() *ListFreeHyperNodesResponseBody
}

type ListFreeHyperNodesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFreeHyperNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFreeHyperNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFreeHyperNodesResponse) GoString() string {
	return s.String()
}

func (s *ListFreeHyperNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFreeHyperNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFreeHyperNodesResponse) GetBody() *ListFreeHyperNodesResponseBody {
	return s.Body
}

func (s *ListFreeHyperNodesResponse) SetHeaders(v map[string]*string) *ListFreeHyperNodesResponse {
	s.Headers = v
	return s
}

func (s *ListFreeHyperNodesResponse) SetStatusCode(v int32) *ListFreeHyperNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFreeHyperNodesResponse) SetBody(v *ListFreeHyperNodesResponseBody) *ListFreeHyperNodesResponse {
	s.Body = v
	return s
}

func (s *ListFreeHyperNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
