// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstancesResponseBody) *DeleteInstancesResponse
	GetBody() *DeleteInstancesResponseBody
}

type DeleteInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstancesResponse) GetBody() *DeleteInstancesResponseBody {
	return s.Body
}

func (s *DeleteInstancesResponse) SetHeaders(v map[string]*string) *DeleteInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstancesResponse) SetStatusCode(v int32) *DeleteInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstancesResponse) SetBody(v *DeleteInstancesResponseBody) *DeleteInstancesResponse {
	s.Body = v
	return s
}

func (s *DeleteInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
