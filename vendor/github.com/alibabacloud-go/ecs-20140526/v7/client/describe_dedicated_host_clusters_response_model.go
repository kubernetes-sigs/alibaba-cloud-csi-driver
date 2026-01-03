// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedHostClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedHostClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedHostClustersResponseBody) *DescribeDedicatedHostClustersResponse
	GetBody() *DescribeDedicatedHostClustersResponseBody
}

type DescribeDedicatedHostClustersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedHostClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedHostClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedHostClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedHostClustersResponse) GetBody() *DescribeDedicatedHostClustersResponseBody {
	return s.Body
}

func (s *DescribeDedicatedHostClustersResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostClustersResponse) SetStatusCode(v int32) *DescribeDedicatedHostClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostClustersResponse) SetBody(v *DescribeDedicatedHostClustersResponseBody) *DescribeDedicatedHostClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedHostClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
