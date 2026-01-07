// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHpcClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHpcClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHpcClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateHpcClusterResponseBody) *CreateHpcClusterResponse
	GetBody() *CreateHpcClusterResponseBody
}

type CreateHpcClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHpcClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHpcClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHpcClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateHpcClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHpcClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHpcClusterResponse) GetBody() *CreateHpcClusterResponseBody {
	return s.Body
}

func (s *CreateHpcClusterResponse) SetHeaders(v map[string]*string) *CreateHpcClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateHpcClusterResponse) SetStatusCode(v int32) *CreateHpcClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHpcClusterResponse) SetBody(v *CreateHpcClusterResponseBody) *CreateHpcClusterResponse {
	s.Body = v
	return s
}

func (s *CreateHpcClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
