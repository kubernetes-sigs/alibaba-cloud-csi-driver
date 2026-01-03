// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReActivateInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReActivateInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReActivateInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ReActivateInstancesResponseBody) *ReActivateInstancesResponse
	GetBody() *ReActivateInstancesResponseBody
}

type ReActivateInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReActivateInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReActivateInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ReActivateInstancesResponse) GoString() string {
	return s.String()
}

func (s *ReActivateInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReActivateInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReActivateInstancesResponse) GetBody() *ReActivateInstancesResponseBody {
	return s.Body
}

func (s *ReActivateInstancesResponse) SetHeaders(v map[string]*string) *ReActivateInstancesResponse {
	s.Headers = v
	return s
}

func (s *ReActivateInstancesResponse) SetStatusCode(v int32) *ReActivateInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReActivateInstancesResponse) SetBody(v *ReActivateInstancesResponseBody) *ReActivateInstancesResponse {
	s.Body = v
	return s
}

func (s *ReActivateInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
