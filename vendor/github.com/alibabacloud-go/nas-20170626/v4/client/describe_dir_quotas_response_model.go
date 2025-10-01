// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDirQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDirQuotasResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDirQuotasResponseBody) *DescribeDirQuotasResponse
	GetBody() *DescribeDirQuotasResponseBody
}

type DescribeDirQuotasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDirQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDirQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirQuotasResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDirQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDirQuotasResponse) GetBody() *DescribeDirQuotasResponseBody {
	return s.Body
}

func (s *DescribeDirQuotasResponse) SetHeaders(v map[string]*string) *DescribeDirQuotasResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirQuotasResponse) SetStatusCode(v int32) *DescribeDirQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirQuotasResponse) SetBody(v *DescribeDirQuotasResponseBody) *DescribeDirQuotasResponse {
	s.Body = v
	return s
}

func (s *DescribeDirQuotasResponse) Validate() error {
	return dara.Validate(s)
}
