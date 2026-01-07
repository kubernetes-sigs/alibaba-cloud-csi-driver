// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAutoReleaseTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDedicatedHostAutoReleaseTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDedicatedHostAutoReleaseTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDedicatedHostAutoReleaseTimeResponseBody) *ModifyDedicatedHostAutoReleaseTimeResponse
	GetBody() *ModifyDedicatedHostAutoReleaseTimeResponseBody
}

type ModifyDedicatedHostAutoReleaseTimeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDedicatedHostAutoReleaseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDedicatedHostAutoReleaseTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAutoReleaseTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponse) GetBody() *ModifyDedicatedHostAutoReleaseTimeResponseBody {
	return s.Body
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedHostAutoReleaseTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponse) SetStatusCode(v int32) *ModifyDedicatedHostAutoReleaseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponse) SetBody(v *ModifyDedicatedHostAutoReleaseTimeResponseBody) *ModifyDedicatedHostAutoReleaseTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
