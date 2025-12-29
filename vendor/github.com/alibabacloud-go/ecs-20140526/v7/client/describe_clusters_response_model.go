// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClustersResponseBody) *DescribeClustersResponse
	GetBody() *DescribeClustersResponseBody
}

type DescribeClustersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClustersResponse) GetBody() *DescribeClustersResponseBody {
	return s.Body
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersResponse) SetStatusCode(v int32) *DescribeClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClustersResponse) SetBody(v *DescribeClustersResponseBody) *DescribeClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
