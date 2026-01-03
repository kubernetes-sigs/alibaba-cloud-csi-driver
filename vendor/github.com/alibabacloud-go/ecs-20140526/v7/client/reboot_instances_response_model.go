// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RebootInstancesResponseBody) *RebootInstancesResponse
	GetBody() *RebootInstancesResponseBody
}

type RebootInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesResponse) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootInstancesResponse) GetBody() *RebootInstancesResponseBody {
	return s.Body
}

func (s *RebootInstancesResponse) SetHeaders(v map[string]*string) *RebootInstancesResponse {
	s.Headers = v
	return s
}

func (s *RebootInstancesResponse) SetStatusCode(v int32) *RebootInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootInstancesResponse) SetBody(v *RebootInstancesResponseBody) *RebootInstancesResponse {
	s.Body = v
	return s
}

func (s *RebootInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
