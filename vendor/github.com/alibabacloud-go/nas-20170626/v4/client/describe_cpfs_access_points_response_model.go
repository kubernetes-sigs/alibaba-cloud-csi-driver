// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCpfsAccessPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCpfsAccessPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCpfsAccessPointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCpfsAccessPointsResponseBody) *DescribeCpfsAccessPointsResponse
	GetBody() *DescribeCpfsAccessPointsResponseBody
}

type DescribeCpfsAccessPointsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCpfsAccessPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCpfsAccessPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCpfsAccessPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCpfsAccessPointsResponse) GetBody() *DescribeCpfsAccessPointsResponseBody {
	return s.Body
}

func (s *DescribeCpfsAccessPointsResponse) SetHeaders(v map[string]*string) *DescribeCpfsAccessPointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCpfsAccessPointsResponse) SetStatusCode(v int32) *DescribeCpfsAccessPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCpfsAccessPointsResponse) SetBody(v *DescribeCpfsAccessPointsResponseBody) *DescribeCpfsAccessPointsResponse {
	s.Body = v
	return s
}

func (s *DescribeCpfsAccessPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
