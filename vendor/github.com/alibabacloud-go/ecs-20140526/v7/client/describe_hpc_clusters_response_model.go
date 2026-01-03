// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHpcClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHpcClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHpcClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHpcClustersResponseBody) *DescribeHpcClustersResponse
	GetBody() *DescribeHpcClustersResponseBody
}

type DescribeHpcClustersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHpcClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHpcClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHpcClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeHpcClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHpcClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHpcClustersResponse) GetBody() *DescribeHpcClustersResponseBody {
	return s.Body
}

func (s *DescribeHpcClustersResponse) SetHeaders(v map[string]*string) *DescribeHpcClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeHpcClustersResponse) SetStatusCode(v int32) *DescribeHpcClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHpcClustersResponse) SetBody(v *DescribeHpcClustersResponseBody) *DescribeHpcClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeHpcClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
