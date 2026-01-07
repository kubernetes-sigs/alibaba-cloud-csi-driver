// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManagedInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeManagedInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeManagedInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeManagedInstancesResponseBody) *DescribeManagedInstancesResponse
	GetBody() *DescribeManagedInstancesResponseBody
}

type DescribeManagedInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeManagedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeManagedInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeManagedInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeManagedInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeManagedInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeManagedInstancesResponse) GetBody() *DescribeManagedInstancesResponseBody {
	return s.Body
}

func (s *DescribeManagedInstancesResponse) SetHeaders(v map[string]*string) *DescribeManagedInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeManagedInstancesResponse) SetStatusCode(v int32) *DescribeManagedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeManagedInstancesResponse) SetBody(v *DescribeManagedInstancesResponseBody) *DescribeManagedInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeManagedInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
