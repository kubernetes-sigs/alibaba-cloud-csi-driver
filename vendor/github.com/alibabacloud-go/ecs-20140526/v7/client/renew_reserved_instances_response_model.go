// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewReservedInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewReservedInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewReservedInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RenewReservedInstancesResponseBody) *RenewReservedInstancesResponse
	GetBody() *RenewReservedInstancesResponseBody
}

type RenewReservedInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewReservedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewReservedInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewReservedInstancesResponse) GoString() string {
	return s.String()
}

func (s *RenewReservedInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewReservedInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewReservedInstancesResponse) GetBody() *RenewReservedInstancesResponseBody {
	return s.Body
}

func (s *RenewReservedInstancesResponse) SetHeaders(v map[string]*string) *RenewReservedInstancesResponse {
	s.Headers = v
	return s
}

func (s *RenewReservedInstancesResponse) SetStatusCode(v int32) *RenewReservedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewReservedInstancesResponse) SetBody(v *RenewReservedInstancesResponseBody) *RenewReservedInstancesResponse {
	s.Body = v
	return s
}

func (s *RenewReservedInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
