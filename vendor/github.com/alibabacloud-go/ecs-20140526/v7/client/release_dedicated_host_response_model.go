// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseDedicatedHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseDedicatedHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseDedicatedHostResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseDedicatedHostResponseBody) *ReleaseDedicatedHostResponse
	GetBody() *ReleaseDedicatedHostResponseBody
}

type ReleaseDedicatedHostResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseDedicatedHostResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseDedicatedHostResponse) GoString() string {
	return s.String()
}

func (s *ReleaseDedicatedHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseDedicatedHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseDedicatedHostResponse) GetBody() *ReleaseDedicatedHostResponseBody {
	return s.Body
}

func (s *ReleaseDedicatedHostResponse) SetHeaders(v map[string]*string) *ReleaseDedicatedHostResponse {
	s.Headers = v
	return s
}

func (s *ReleaseDedicatedHostResponse) SetStatusCode(v int32) *ReleaseDedicatedHostResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseDedicatedHostResponse) SetBody(v *ReleaseDedicatedHostResponseBody) *ReleaseDedicatedHostResponse {
	s.Body = v
	return s
}

func (s *ReleaseDedicatedHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
