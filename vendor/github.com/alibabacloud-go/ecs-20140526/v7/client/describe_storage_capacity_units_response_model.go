// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageCapacityUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStorageCapacityUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStorageCapacityUnitsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStorageCapacityUnitsResponseBody) *DescribeStorageCapacityUnitsResponse
	GetBody() *DescribeStorageCapacityUnitsResponseBody
}

type DescribeStorageCapacityUnitsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStorageCapacityUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStorageCapacityUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStorageCapacityUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStorageCapacityUnitsResponse) GetBody() *DescribeStorageCapacityUnitsResponseBody {
	return s.Body
}

func (s *DescribeStorageCapacityUnitsResponse) SetHeaders(v map[string]*string) *DescribeStorageCapacityUnitsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStorageCapacityUnitsResponse) SetStatusCode(v int32) *DescribeStorageCapacityUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStorageCapacityUnitsResponse) SetBody(v *DescribeStorageCapacityUnitsResponseBody) *DescribeStorageCapacityUnitsResponse {
	s.Body = v
	return s
}

func (s *DescribeStorageCapacityUnitsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
