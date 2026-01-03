// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDedicatedHostClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDedicatedHostClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDedicatedHostClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDedicatedHostClusterResponseBody) *DeleteDedicatedHostClusterResponse
	GetBody() *DeleteDedicatedHostClusterResponseBody
}

type DeleteDedicatedHostClusterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDedicatedHostClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDedicatedHostClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDedicatedHostClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDedicatedHostClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDedicatedHostClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDedicatedHostClusterResponse) GetBody() *DeleteDedicatedHostClusterResponseBody {
	return s.Body
}

func (s *DeleteDedicatedHostClusterResponse) SetHeaders(v map[string]*string) *DeleteDedicatedHostClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDedicatedHostClusterResponse) SetStatusCode(v int32) *DeleteDedicatedHostClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDedicatedHostClusterResponse) SetBody(v *DeleteDedicatedHostClusterResponseBody) *DeleteDedicatedHostClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteDedicatedHostClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
