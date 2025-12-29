// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceNetworkSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceNetworkSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceNetworkSpecResponseBody) *ModifyInstanceNetworkSpecResponse
	GetBody() *ModifyInstanceNetworkSpecResponseBody
}

type ModifyInstanceNetworkSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceNetworkSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceNetworkSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceNetworkSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceNetworkSpecResponse) GetBody() *ModifyInstanceNetworkSpecResponseBody {
	return s.Body
}

func (s *ModifyInstanceNetworkSpecResponse) SetHeaders(v map[string]*string) *ModifyInstanceNetworkSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNetworkSpecResponse) SetStatusCode(v int32) *ModifyInstanceNetworkSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNetworkSpecResponse) SetBody(v *ModifyInstanceNetworkSpecResponseBody) *ModifyInstanceNetworkSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceNetworkSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
