// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDedicatedHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewDedicatedHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewDedicatedHostsResponse
	GetStatusCode() *int32
	SetBody(v *RenewDedicatedHostsResponseBody) *RenewDedicatedHostsResponse
	GetBody() *RenewDedicatedHostsResponseBody
}

type RenewDedicatedHostsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewDedicatedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewDedicatedHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewDedicatedHostsResponse) GoString() string {
	return s.String()
}

func (s *RenewDedicatedHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewDedicatedHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewDedicatedHostsResponse) GetBody() *RenewDedicatedHostsResponseBody {
	return s.Body
}

func (s *RenewDedicatedHostsResponse) SetHeaders(v map[string]*string) *RenewDedicatedHostsResponse {
	s.Headers = v
	return s
}

func (s *RenewDedicatedHostsResponse) SetStatusCode(v int32) *RenewDedicatedHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewDedicatedHostsResponse) SetBody(v *RenewDedicatedHostsResponseBody) *RenewDedicatedHostsResponse {
	s.Body = v
	return s
}

func (s *RenewDedicatedHostsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
