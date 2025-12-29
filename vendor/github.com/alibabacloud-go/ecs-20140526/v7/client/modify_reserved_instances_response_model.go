// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyReservedInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyReservedInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyReservedInstancesResponseBody) *ModifyReservedInstancesResponse
	GetBody() *ModifyReservedInstancesResponseBody
}

type ModifyReservedInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyReservedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyReservedInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstancesResponse) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyReservedInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyReservedInstancesResponse) GetBody() *ModifyReservedInstancesResponseBody {
	return s.Body
}

func (s *ModifyReservedInstancesResponse) SetHeaders(v map[string]*string) *ModifyReservedInstancesResponse {
	s.Headers = v
	return s
}

func (s *ModifyReservedInstancesResponse) SetStatusCode(v int32) *ModifyReservedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyReservedInstancesResponse) SetBody(v *ModifyReservedInstancesResponseBody) *ModifyReservedInstancesResponse {
	s.Body = v
	return s
}

func (s *ModifyReservedInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
