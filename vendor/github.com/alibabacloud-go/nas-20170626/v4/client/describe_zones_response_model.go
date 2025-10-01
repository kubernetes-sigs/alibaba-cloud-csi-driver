// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse
	GetBody() *DescribeZonesResponseBody
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeZonesResponse) GetBody() *DescribeZonesResponseBody {
	return s.Body
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeZonesResponse) Validate() error {
	return dara.Validate(s)
}
