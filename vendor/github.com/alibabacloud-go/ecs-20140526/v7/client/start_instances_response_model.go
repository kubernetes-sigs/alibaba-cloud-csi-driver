// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartInstancesResponse
	GetStatusCode() *int32
	SetBody(v *StartInstancesResponseBody) *StartInstancesResponse
	GetBody() *StartInstancesResponseBody
}

type StartInstancesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesResponse) GoString() string {
	return s.String()
}

func (s *StartInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartInstancesResponse) GetBody() *StartInstancesResponseBody {
	return s.Body
}

func (s *StartInstancesResponse) SetHeaders(v map[string]*string) *StartInstancesResponse {
	s.Headers = v
	return s
}

func (s *StartInstancesResponse) SetStatusCode(v int32) *StartInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstancesResponse) SetBody(v *StartInstancesResponseBody) *StartInstancesResponse {
	s.Body = v
	return s
}

func (s *StartInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
