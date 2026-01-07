// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProtocolServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyProtocolServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyProtocolServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyProtocolServiceResponseBody) *ModifyProtocolServiceResponse
	GetBody() *ModifyProtocolServiceResponseBody
}

type ModifyProtocolServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProtocolServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProtocolServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyProtocolServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyProtocolServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyProtocolServiceResponse) GetBody() *ModifyProtocolServiceResponseBody {
	return s.Body
}

func (s *ModifyProtocolServiceResponse) SetHeaders(v map[string]*string) *ModifyProtocolServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyProtocolServiceResponse) SetStatusCode(v int32) *ModifyProtocolServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProtocolServiceResponse) SetBody(v *ModifyProtocolServiceResponseBody) *ModifyProtocolServiceResponse {
	s.Body = v
	return s
}

func (s *ModifyProtocolServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
