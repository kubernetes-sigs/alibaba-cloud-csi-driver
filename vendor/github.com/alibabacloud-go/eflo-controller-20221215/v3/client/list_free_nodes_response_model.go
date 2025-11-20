// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreeNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFreeNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFreeNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListFreeNodesResponseBody) *ListFreeNodesResponse
	GetBody() *ListFreeNodesResponseBody
}

type ListFreeNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFreeNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFreeNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesResponse) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFreeNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFreeNodesResponse) GetBody() *ListFreeNodesResponseBody {
	return s.Body
}

func (s *ListFreeNodesResponse) SetHeaders(v map[string]*string) *ListFreeNodesResponse {
	s.Headers = v
	return s
}

func (s *ListFreeNodesResponse) SetStatusCode(v int32) *ListFreeNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFreeNodesResponse) SetBody(v *ListFreeNodesResponseBody) *ListFreeNodesResponse {
	s.Body = v
	return s
}

func (s *ListFreeNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
