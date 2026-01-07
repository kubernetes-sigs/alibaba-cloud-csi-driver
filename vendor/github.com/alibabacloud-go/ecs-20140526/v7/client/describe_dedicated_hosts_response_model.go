// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedHostsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedHostsResponseBody) *DescribeDedicatedHostsResponse
	GetBody() *DescribeDedicatedHostsResponseBody
}

type DescribeDedicatedHostsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedHostsResponse) GetBody() *DescribeDedicatedHostsResponseBody {
	return s.Body
}

func (s *DescribeDedicatedHostsResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostsResponse) SetStatusCode(v int32) *DescribeDedicatedHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostsResponse) SetBody(v *DescribeDedicatedHostsResponseBody) *DescribeDedicatedHostsResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedHostsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
