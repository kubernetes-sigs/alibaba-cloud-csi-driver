// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEipAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEipAddressesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEipAddressesResponseBody) *DescribeEipAddressesResponse
	GetBody() *DescribeEipAddressesResponseBody
}

type DescribeEipAddressesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEipAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEipAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEipAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEipAddressesResponse) GetBody() *DescribeEipAddressesResponseBody {
	return s.Body
}

func (s *DescribeEipAddressesResponse) SetHeaders(v map[string]*string) *DescribeEipAddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEipAddressesResponse) SetStatusCode(v int32) *DescribeEipAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEipAddressesResponse) SetBody(v *DescribeEipAddressesResponseBody) *DescribeEipAddressesResponse {
	s.Body = v
	return s
}

func (s *DescribeEipAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
