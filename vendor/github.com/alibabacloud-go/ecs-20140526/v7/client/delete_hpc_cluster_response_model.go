// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHpcClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHpcClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHpcClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHpcClusterResponseBody) *DeleteHpcClusterResponse
	GetBody() *DeleteHpcClusterResponseBody
}

type DeleteHpcClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHpcClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHpcClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHpcClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteHpcClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHpcClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHpcClusterResponse) GetBody() *DeleteHpcClusterResponseBody {
	return s.Body
}

func (s *DeleteHpcClusterResponse) SetHeaders(v map[string]*string) *DeleteHpcClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteHpcClusterResponse) SetStatusCode(v int32) *DeleteHpcClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHpcClusterResponse) SetBody(v *DeleteHpcClusterResponseBody) *DeleteHpcClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteHpcClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
