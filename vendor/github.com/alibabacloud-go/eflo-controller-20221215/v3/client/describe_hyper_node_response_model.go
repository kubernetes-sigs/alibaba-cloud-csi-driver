// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHyperNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHyperNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHyperNodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHyperNodeResponseBody) *DescribeHyperNodeResponse
	GetBody() *DescribeHyperNodeResponseBody
}

type DescribeHyperNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHyperNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHyperNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHyperNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeHyperNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHyperNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHyperNodeResponse) GetBody() *DescribeHyperNodeResponseBody {
	return s.Body
}

func (s *DescribeHyperNodeResponse) SetHeaders(v map[string]*string) *DescribeHyperNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeHyperNodeResponse) SetStatusCode(v int32) *DescribeHyperNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHyperNodeResponse) SetBody(v *DescribeHyperNodeResponseBody) *DescribeHyperNodeResponse {
	s.Body = v
	return s
}

func (s *DescribeHyperNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
