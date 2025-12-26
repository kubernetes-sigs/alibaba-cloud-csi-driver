// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFilesetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFilesetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFilesetsResponseBody) *DescribeFilesetsResponse
	GetBody() *DescribeFilesetsResponseBody
}

type DescribeFilesetsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFilesetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFilesetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFilesetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFilesetsResponse) GetBody() *DescribeFilesetsResponseBody {
	return s.Body
}

func (s *DescribeFilesetsResponse) SetHeaders(v map[string]*string) *DescribeFilesetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFilesetsResponse) SetStatusCode(v int32) *DescribeFilesetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFilesetsResponse) SetBody(v *DescribeFilesetsResponseBody) *DescribeFilesetsResponse {
	s.Body = v
	return s
}

func (s *DescribeFilesetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
