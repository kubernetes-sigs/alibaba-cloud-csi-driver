// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployDedicatedHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedeployDedicatedHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedeployDedicatedHostResponse
	GetStatusCode() *int32
	SetBody(v *RedeployDedicatedHostResponseBody) *RedeployDedicatedHostResponse
	GetBody() *RedeployDedicatedHostResponseBody
}

type RedeployDedicatedHostResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedeployDedicatedHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedeployDedicatedHostResponse) String() string {
	return dara.Prettify(s)
}

func (s RedeployDedicatedHostResponse) GoString() string {
	return s.String()
}

func (s *RedeployDedicatedHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedeployDedicatedHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedeployDedicatedHostResponse) GetBody() *RedeployDedicatedHostResponseBody {
	return s.Body
}

func (s *RedeployDedicatedHostResponse) SetHeaders(v map[string]*string) *RedeployDedicatedHostResponse {
	s.Headers = v
	return s
}

func (s *RedeployDedicatedHostResponse) SetStatusCode(v int32) *RedeployDedicatedHostResponse {
	s.StatusCode = &v
	return s
}

func (s *RedeployDedicatedHostResponse) SetBody(v *RedeployDedicatedHostResponseBody) *RedeployDedicatedHostResponse {
	s.Body = v
	return s
}

func (s *RedeployDedicatedHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
