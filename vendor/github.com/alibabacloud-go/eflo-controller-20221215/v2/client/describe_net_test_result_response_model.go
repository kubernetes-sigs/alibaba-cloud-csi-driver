// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetTestResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetTestResultResponseBody) *DescribeNetTestResultResponse
	GetBody() *DescribeNetTestResultResponseBody
}

type DescribeNetTestResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetTestResultResponse) GetBody() *DescribeNetTestResultResponseBody {
	return s.Body
}

func (s *DescribeNetTestResultResponse) SetHeaders(v map[string]*string) *DescribeNetTestResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetTestResultResponse) SetStatusCode(v int32) *DescribeNetTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetTestResultResponse) SetBody(v *DescribeNetTestResultResponseBody) *DescribeNetTestResultResponse {
	s.Body = v
	return s
}

func (s *DescribeNetTestResultResponse) Validate() error {
	return dara.Validate(s)
}
