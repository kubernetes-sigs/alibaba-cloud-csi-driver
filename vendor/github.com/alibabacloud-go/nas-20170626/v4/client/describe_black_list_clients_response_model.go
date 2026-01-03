// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlackListClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBlackListClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBlackListClientsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBlackListClientsResponseBody) *DescribeBlackListClientsResponse
	GetBody() *DescribeBlackListClientsResponseBody
}

type DescribeBlackListClientsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBlackListClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBlackListClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlackListClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBlackListClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBlackListClientsResponse) GetBody() *DescribeBlackListClientsResponseBody {
	return s.Body
}

func (s *DescribeBlackListClientsResponse) SetHeaders(v map[string]*string) *DescribeBlackListClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlackListClientsResponse) SetStatusCode(v int32) *DescribeBlackListClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlackListClientsResponse) SetBody(v *DescribeBlackListClientsResponseBody) *DescribeBlackListClientsResponse {
	s.Body = v
	return s
}

func (s *DescribeBlackListClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
