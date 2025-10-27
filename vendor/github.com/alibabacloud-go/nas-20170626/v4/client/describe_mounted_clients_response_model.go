// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountedClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMountedClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMountedClientsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMountedClientsResponseBody) *DescribeMountedClientsResponse
	GetBody() *DescribeMountedClientsResponseBody
}

type DescribeMountedClientsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMountedClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMountedClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountedClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMountedClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMountedClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMountedClientsResponse) GetBody() *DescribeMountedClientsResponseBody {
	return s.Body
}

func (s *DescribeMountedClientsResponse) SetHeaders(v map[string]*string) *DescribeMountedClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMountedClientsResponse) SetStatusCode(v int32) *DescribeMountedClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMountedClientsResponse) SetBody(v *DescribeMountedClientsResponseBody) *DescribeMountedClientsResponse {
	s.Body = v
	return s
}

func (s *DescribeMountedClientsResponse) Validate() error {
	return dara.Validate(s)
}
