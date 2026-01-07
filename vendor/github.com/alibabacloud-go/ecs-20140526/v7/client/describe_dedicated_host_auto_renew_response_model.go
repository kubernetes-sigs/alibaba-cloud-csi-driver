// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostAutoRenewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedHostAutoRenewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedHostAutoRenewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedHostAutoRenewResponseBody) *DescribeDedicatedHostAutoRenewResponse
	GetBody() *DescribeDedicatedHostAutoRenewResponseBody
}

type DescribeDedicatedHostAutoRenewResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedHostAutoRenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedHostAutoRenewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostAutoRenewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAutoRenewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedHostAutoRenewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedHostAutoRenewResponse) GetBody() *DescribeDedicatedHostAutoRenewResponseBody {
	return s.Body
}

func (s *DescribeDedicatedHostAutoRenewResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostAutoRenewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponse) SetStatusCode(v int32) *DescribeDedicatedHostAutoRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponse) SetBody(v *DescribeDedicatedHostAutoRenewResponseBody) *DescribeDedicatedHostAutoRenewResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
