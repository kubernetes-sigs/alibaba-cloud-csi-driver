// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RunInstancesResponseBody) *RunInstancesResponse
	GetBody() *RunInstancesResponseBody
}

type RunInstancesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesResponse) GoString() string {
	return s.String()
}

func (s *RunInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunInstancesResponse) GetBody() *RunInstancesResponseBody {
	return s.Body
}

func (s *RunInstancesResponse) SetHeaders(v map[string]*string) *RunInstancesResponse {
	s.Headers = v
	return s
}

func (s *RunInstancesResponse) SetStatusCode(v int32) *RunInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunInstancesResponse) SetBody(v *RunInstancesResponseBody) *RunInstancesResponse {
	s.Body = v
	return s
}

func (s *RunInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
