// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortRangeListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortRangeListsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortRangeListsResponseBody) *DescribePortRangeListsResponse
	GetBody() *DescribePortRangeListsResponseBody
}

type DescribePortRangeListsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortRangeListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortRangeListsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsResponse) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortRangeListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortRangeListsResponse) GetBody() *DescribePortRangeListsResponseBody {
	return s.Body
}

func (s *DescribePortRangeListsResponse) SetHeaders(v map[string]*string) *DescribePortRangeListsResponse {
	s.Headers = v
	return s
}

func (s *DescribePortRangeListsResponse) SetStatusCode(v int32) *DescribePortRangeListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortRangeListsResponse) SetBody(v *DescribePortRangeListsResponseBody) *DescribePortRangeListsResponse {
	s.Body = v
	return s
}

func (s *DescribePortRangeListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
