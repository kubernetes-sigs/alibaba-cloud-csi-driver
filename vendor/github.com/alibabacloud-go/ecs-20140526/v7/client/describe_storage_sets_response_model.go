// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStorageSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStorageSetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStorageSetsResponseBody) *DescribeStorageSetsResponse
	GetBody() *DescribeStorageSetsResponseBody
}

type DescribeStorageSetsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStorageSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStorageSetsResponse) GetBody() *DescribeStorageSetsResponseBody {
	return s.Body
}

func (s *DescribeStorageSetsResponse) SetHeaders(v map[string]*string) *DescribeStorageSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageSetsResponse) SetStatusCode(v int32) *DescribeStorageSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageSetsResponse) SetBody(v *DescribeStorageSetsResponseBody) *DescribeStorageSetsResponse {
	s.Body = v
	return s
}

func (s *DescribeStorageSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
