// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoReleaseTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceAutoReleaseTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceAutoReleaseTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceAutoReleaseTimeResponseBody) *ModifyInstanceAutoReleaseTimeResponse
	GetBody() *ModifyInstanceAutoReleaseTimeResponseBody
}

type ModifyInstanceAutoReleaseTimeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAutoReleaseTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceAutoReleaseTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoReleaseTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoReleaseTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceAutoReleaseTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceAutoReleaseTimeResponse) GetBody() *ModifyInstanceAutoReleaseTimeResponseBody {
	return s.Body
}

func (s *ModifyInstanceAutoReleaseTimeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAutoReleaseTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeResponse) SetStatusCode(v int32) *ModifyInstanceAutoReleaseTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeResponse) SetBody(v *ModifyInstanceAutoReleaseTimeResponseBody) *ModifyInstanceAutoReleaseTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
