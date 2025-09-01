// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVscResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVscResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVscResponseBody) *DescribeVscResponse
	GetBody() *DescribeVscResponseBody
}

type DescribeVscResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVscResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVscResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVscResponse) GetBody() *DescribeVscResponseBody {
	return s.Body
}

func (s *DescribeVscResponse) SetHeaders(v map[string]*string) *DescribeVscResponse {
	s.Headers = v
	return s
}

func (s *DescribeVscResponse) SetStatusCode(v int32) *DescribeVscResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVscResponse) SetBody(v *DescribeVscResponseBody) *DescribeVscResponse {
	s.Body = v
	return s
}

func (s *DescribeVscResponse) Validate() error {
	return dara.Validate(s)
}
