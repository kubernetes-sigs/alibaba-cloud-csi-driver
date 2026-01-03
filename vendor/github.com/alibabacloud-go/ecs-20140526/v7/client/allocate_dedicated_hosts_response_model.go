// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDedicatedHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateDedicatedHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateDedicatedHostsResponse
	GetStatusCode() *int32
	SetBody(v *AllocateDedicatedHostsResponseBody) *AllocateDedicatedHostsResponse
	GetBody() *AllocateDedicatedHostsResponseBody
}

type AllocateDedicatedHostsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateDedicatedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateDedicatedHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateDedicatedHostsResponse) GoString() string {
	return s.String()
}

func (s *AllocateDedicatedHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateDedicatedHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateDedicatedHostsResponse) GetBody() *AllocateDedicatedHostsResponseBody {
	return s.Body
}

func (s *AllocateDedicatedHostsResponse) SetHeaders(v map[string]*string) *AllocateDedicatedHostsResponse {
	s.Headers = v
	return s
}

func (s *AllocateDedicatedHostsResponse) SetStatusCode(v int32) *AllocateDedicatedHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateDedicatedHostsResponse) SetBody(v *AllocateDedicatedHostsResponseBody) *AllocateDedicatedHostsResponse {
	s.Body = v
	return s
}

func (s *AllocateDedicatedHostsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
