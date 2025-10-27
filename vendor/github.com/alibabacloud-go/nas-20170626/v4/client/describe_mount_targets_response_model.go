// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMountTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMountTargetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMountTargetsResponseBody) *DescribeMountTargetsResponse
	GetBody() *DescribeMountTargetsResponseBody
}

type DescribeMountTargetsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMountTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMountTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMountTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMountTargetsResponse) GetBody() *DescribeMountTargetsResponseBody {
	return s.Body
}

func (s *DescribeMountTargetsResponse) SetHeaders(v map[string]*string) *DescribeMountTargetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMountTargetsResponse) SetStatusCode(v int32) *DescribeMountTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMountTargetsResponse) SetBody(v *DescribeMountTargetsResponseBody) *DescribeMountTargetsResponse {
	s.Body = v
	return s
}

func (s *DescribeMountTargetsResponse) Validate() error {
	return dara.Validate(s)
}
