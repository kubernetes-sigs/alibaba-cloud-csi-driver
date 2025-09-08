// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSendFileResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSendFileResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSendFileResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSendFileResultsResponseBody) *DescribeSendFileResultsResponse
	GetBody() *DescribeSendFileResultsResponseBody
}

type DescribeSendFileResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSendFileResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSendFileResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSendFileResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSendFileResultsResponse) GetBody() *DescribeSendFileResultsResponseBody {
	return s.Body
}

func (s *DescribeSendFileResultsResponse) SetHeaders(v map[string]*string) *DescribeSendFileResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSendFileResultsResponse) SetStatusCode(v int32) *DescribeSendFileResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSendFileResultsResponse) SetBody(v *DescribeSendFileResultsResponseBody) *DescribeSendFileResultsResponse {
	s.Body = v
	return s
}

func (s *DescribeSendFileResultsResponse) Validate() error {
	return dara.Validate(s)
}
