// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStorageSetDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStorageSetDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStorageSetDetailsResponseBody) *DescribeStorageSetDetailsResponse
	GetBody() *DescribeStorageSetDetailsResponseBody
}

type DescribeStorageSetDetailsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageSetDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageSetDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStorageSetDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStorageSetDetailsResponse) GetBody() *DescribeStorageSetDetailsResponseBody {
	return s.Body
}

func (s *DescribeStorageSetDetailsResponse) SetHeaders(v map[string]*string) *DescribeStorageSetDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageSetDetailsResponse) SetStatusCode(v int32) *DescribeStorageSetDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageSetDetailsResponse) SetBody(v *DescribeStorageSetDetailsResponseBody) *DescribeStorageSetDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeStorageSetDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
