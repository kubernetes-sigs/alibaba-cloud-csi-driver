// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterManagedInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeregisterManagedInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeregisterManagedInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeregisterManagedInstanceResponseBody) *DeregisterManagedInstanceResponse
	GetBody() *DeregisterManagedInstanceResponseBody
}

type DeregisterManagedInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterManagedInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterManagedInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeregisterManagedInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeregisterManagedInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeregisterManagedInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeregisterManagedInstanceResponse) GetBody() *DeregisterManagedInstanceResponseBody {
	return s.Body
}

func (s *DeregisterManagedInstanceResponse) SetHeaders(v map[string]*string) *DeregisterManagedInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeregisterManagedInstanceResponse) SetStatusCode(v int32) *DeregisterManagedInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterManagedInstanceResponse) SetBody(v *DeregisterManagedInstanceResponseBody) *DeregisterManagedInstanceResponse {
	s.Body = v
	return s
}

func (s *DeregisterManagedInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
