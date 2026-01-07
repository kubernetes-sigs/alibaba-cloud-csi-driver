// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedHostTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedHostTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedHostTypesResponseBody) *DescribeDedicatedHostTypesResponse
	GetBody() *DescribeDedicatedHostTypesResponseBody
}

type DescribeDedicatedHostTypesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedHostTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedHostTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedHostTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedHostTypesResponse) GetBody() *DescribeDedicatedHostTypesResponseBody {
	return s.Body
}

func (s *DescribeDedicatedHostTypesResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostTypesResponse) SetStatusCode(v int32) *DescribeDedicatedHostTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostTypesResponse) SetBody(v *DescribeDedicatedHostTypesResponseBody) *DescribeDedicatedHostTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedHostTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
