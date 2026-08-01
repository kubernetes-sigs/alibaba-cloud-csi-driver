// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCpfsAccessPointMountedClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCpfsAccessPointMountedClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCpfsAccessPointMountedClientsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCpfsAccessPointMountedClientsResponseBody) *DescribeCpfsAccessPointMountedClientsResponse
	GetBody() *DescribeCpfsAccessPointMountedClientsResponseBody
}

type DescribeCpfsAccessPointMountedClientsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCpfsAccessPointMountedClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCpfsAccessPointMountedClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCpfsAccessPointMountedClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCpfsAccessPointMountedClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCpfsAccessPointMountedClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCpfsAccessPointMountedClientsResponse) GetBody() *DescribeCpfsAccessPointMountedClientsResponseBody {
	return s.Body
}

func (s *DescribeCpfsAccessPointMountedClientsResponse) SetHeaders(v map[string]*string) *DescribeCpfsAccessPointMountedClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponse) SetStatusCode(v int32) *DescribeCpfsAccessPointMountedClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponse) SetBody(v *DescribeCpfsAccessPointMountedClientsResponseBody) *DescribeCpfsAccessPointMountedClientsResponse {
	s.Body = v
	return s
}

func (s *DescribeCpfsAccessPointMountedClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
