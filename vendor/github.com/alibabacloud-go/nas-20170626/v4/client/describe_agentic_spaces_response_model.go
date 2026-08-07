// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticSpacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticSpacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticSpacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticSpacesResponseBody) *DescribeAgenticSpacesResponse
	GetBody() *DescribeAgenticSpacesResponseBody
}

type DescribeAgenticSpacesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticSpacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticSpacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticSpacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticSpacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticSpacesResponse) GetBody() *DescribeAgenticSpacesResponseBody {
	return s.Body
}

func (s *DescribeAgenticSpacesResponse) SetHeaders(v map[string]*string) *DescribeAgenticSpacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticSpacesResponse) SetStatusCode(v int32) *DescribeAgenticSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticSpacesResponse) SetBody(v *DescribeAgenticSpacesResponseBody) *DescribeAgenticSpacesResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticSpacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
