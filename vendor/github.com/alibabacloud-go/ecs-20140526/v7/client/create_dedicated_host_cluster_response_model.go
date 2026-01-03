// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedHostClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDedicatedHostClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDedicatedHostClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateDedicatedHostClusterResponseBody) *CreateDedicatedHostClusterResponse
	GetBody() *CreateDedicatedHostClusterResponseBody
}

type CreateDedicatedHostClusterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDedicatedHostClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDedicatedHostClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedHostClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedHostClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDedicatedHostClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDedicatedHostClusterResponse) GetBody() *CreateDedicatedHostClusterResponseBody {
	return s.Body
}

func (s *CreateDedicatedHostClusterResponse) SetHeaders(v map[string]*string) *CreateDedicatedHostClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedHostClusterResponse) SetStatusCode(v int32) *CreateDedicatedHostClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedHostClusterResponse) SetBody(v *CreateDedicatedHostClusterResponseBody) *CreateDedicatedHostClusterResponse {
	s.Body = v
	return s
}

func (s *CreateDedicatedHostClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
